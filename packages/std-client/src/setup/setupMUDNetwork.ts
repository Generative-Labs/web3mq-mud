import {
  createNetwork,
  createContracts,
  Mappings,
  createTxQueue,
  createSyncWorker,
  createSystemExecutor,
  Ack,
  InputType,
  SingletonID,
} from "@latticexyz/network";
import { BehaviorSubject, concatMap, from, Subject } from "rxjs";
import { defineComponent, Type, World } from "@latticexyz/recs";
import { computed } from "mobx";
import { keccak256 } from "@latticexyz/utils";
import { Contract, ContractInterface } from "ethers";
import { World as WorldContract } from "@latticexyz/solecs/types/ethers-contracts/World";
import WorldAbi from "@latticexyz/solecs/abi/World.sol/World.abi.json";
import { defineStringComponent } from "../components";
import keys from "lodash/keys";
import { ContractComponent, ContractComponents, NetworkComponents, SetupContractConfig } from "./types";
import {
  applyNetworkUpdates,
  createDecodeNetworkComponentUpdate,
  createEncoders,
  createSystemCallStreams,
} from "./utils";

export async function setupMUDNetwork<C extends ContractComponents, SystemTypes extends { [key: string]: Contract }>(
  networkConfig: SetupContractConfig,
  world: World,
  contractComponents: C,
  SystemAbis: { [key in keyof SystemTypes]: ContractInterface },
  options?: { initialGasPrice?: number; fetchSystemCalls?: boolean; syncThread?: "main" | "worker" }
) {
  const SystemsRegistry = findOrDefineComponent(
    contractComponents,
    defineStringComponent(world, {
      id: "SystemsRegistry",
      metadata: { contractId: "world.component.systems" },
    })
  );

  const ComponentsRegistry = findOrDefineComponent(
    contractComponents,
    defineStringComponent(world, {
      id: "ComponentsRegistry",
      metadata: { contractId: "world.component.components" },
    })
  );

  // used by SyncWorker to notify client of sync progress
  const LoadingState = findOrDefineComponent(
    contractComponents,
    defineComponent(
      world,
      {
        state: Type.Number,
        msg: Type.String,
        percentage: Type.Number,
      },
      {
        id: "LoadingState",
        metadata: { contractId: "component.LoadingState" },
      }
    )
  );

  const components: NetworkComponents<C> = {
    ...contractComponents,
    SystemsRegistry,
    ComponentsRegistry,
    LoadingState,
  };

  // Mapping from component contract id to key in components object
  const mappings: Mappings<C> = {};

  // Function to register new components in mappings object
  function registerComponent(key: string, component: ContractComponent) {
    const { contractId, tableId } = component.metadata;
    if (tableId) {
      mappings[tableId] = key;
    } else {
      mappings[keccak256(contractId)] = key;
    }
  }

  // Register initial components in mappings object
  for (const key of Object.keys(components)) {
    registerComponent(key, components[key]);
  }

  const network = await createNetwork(networkConfig);
  world.registerDisposer(network.dispose);

  // For LoadingState updates
  const singletonEntity = world.registerEntity({ id: SingletonID });
  // Register player entity
  const address = network.connectedAddress.get();
  const playerEntityId = address;
  const playerEntity = playerEntityId ? world.registerEntity({ id: playerEntityId }) : undefined;

  const signerOrProvider = computed(() => network.signer.get() || network.providers.get().json);

  const { contracts, config: contractsConfig } = await createContracts<{ World: WorldContract }>({
    config: { World: { abi: WorldAbi, address: networkConfig.worldAddress } },
    signerOrProvider,
  });

  const gasPriceInput$ = new BehaviorSubject<number>(
    // If no initial gas price is provided, check the gas price once and add a 30% buffer
    options?.initialGasPrice || Math.ceil((await signerOrProvider.get().getGasPrice()).toNumber() * 1.3)
  );

  const { txQueue, dispose: disposeTxQueue } = createTxQueue(contracts, network, gasPriceInput$, {
    devMode: networkConfig.devMode,
  });
  world.registerDisposer(disposeTxQueue);

  const { systems, registerSystem, getSystemContract } = createSystemExecutor<SystemTypes>(
    world,
    network,
    SystemsRegistry,
    SystemAbis,
    gasPriceInput$,
    {
      devMode: networkConfig.devMode,
    }
  );

  const decodeNetworkComponentUpdate = createDecodeNetworkComponentUpdate(world, components, mappings);
  const { systemCallStreams, decodeAndEmitSystemCall } = createSystemCallStreams(
    world,
    keys(SystemAbis),
    SystemsRegistry,
    getSystemContract,
    decodeNetworkComponentUpdate
  );

  // Create sync worker
  const ack$ = new Subject<Ack>();
  // Avoid passing externalProvider to sync worker (too complex to copy)
  const {
    provider: { externalProvider: _, ...providerConfig },
    ...syncWorkerConfig
  } = networkConfig;
  const { ecsEvents$, input$, dispose } = createSyncWorker<C>(ack$, { thread: options?.syncThread });
  world.registerDisposer(dispose);
  function startSync() {
    input$.next({
      type: InputType.Config,
      data: {
        ...syncWorkerConfig,
        provider: providerConfig,
        worldContract: contractsConfig.World,
        initialBlockNumber: networkConfig.initialBlockNumber ?? 0,
        // TODO: `networkConfig.devMode` is deprecated, only use `networkConfig.disableCache`
        disableCache: networkConfig.devMode || networkConfig.disableCache,
        fetchSystemCalls: options?.fetchSystemCalls,
      },
    });
  }

  const { txReduced$ } = applyNetworkUpdates(world, components, ecsEvents$, mappings, ack$, decodeAndEmitSystemCall);

  const encoders = networkConfig.encoders
    ? createEncoders(world, ComponentsRegistry, signerOrProvider)
    : new Promise((resolve) => resolve({}));

  return {
    txQueue,
    txReduced$,
    encoders,
    network,
    startSync,
    systems,
    systemCallStreams,
    gasPriceInput$,
    ecsEvent$: ecsEvents$.pipe(concatMap((updates) => from(updates))),
    mappings,
    registerComponent,
    registerSystem,
    components,
    singletonEntityId: SingletonID,
    singletonEntity,
    /* @deprecated playerEntityId is equivalent to playerEntity */
    playerEntityId,
    playerEntity,
  };
}

/**
 * Find a component in the components object by contract id, or return the component if it doesn't exist
 * @param components object of components
 * @param component component to find
 * @returns component if it exists in components object, otherwise the component passed in
 */
function findOrDefineComponent<Cs extends ContractComponents, C extends ContractComponent>(
  components: Cs,
  component: C
): C {
  const existingComponent = Object.values(components).find(
    (c) => c.metadata.contractId === component.metadata.contractId
  ) as C;

  if (existingComponent) {
    console.warn(
      "Component with contract id",
      component.metadata.contractId,
      "is defined by default in setupMUDNetwork"
    );
  }

  return existingComponent || component;
}
