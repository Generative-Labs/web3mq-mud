import { ComponentValue, Schema } from "@latticexyz/recs";
import { TableId } from "@latticexyz/utils";
import { Contract } from "ethers";
import { registerSchema } from "./schemas/tableSchemas";
import { registerMetadata } from "./schemas/tableMetadata";
import { decodeField } from "./schemas/decodeField";
import { TableSchema } from "./common";
import { decodeStaticField } from "./schemas/decodeStaticField";
import { DynamicSchemaType, StaticSchemaType } from "@latticexyz/schema-type";
import { decodeDynamicField } from "./schemas/decodeDynamicField";

export async function decodeStoreSetField(
  contract: Contract,
  table: TableId,
  keyTuple: string[],
  schemaIndex: number,
  data: string
): Promise<{
  schema: TableSchema;
  indexedValues: Record<number, any>;
  indexedInitialValues: Record<number, any>;
  namedValues?: Record<string, any>;
  namedInitialValues?: Record<string, any>;
}> {
  const schema = await registerSchema(contract, table);
  const indexedValues = decodeField(schema, schemaIndex, data);

  // Create an object that represents an "uninitialized" record as it would exist in Solidity
  // to help populate RECS state when using StoreSetField before StoreSetRecord.
  const defaultValues = [
    ...schema.staticFields.map((fieldType) => decodeStaticField(fieldType as StaticSchemaType, new Uint8Array(0), 0)),
    ...schema.dynamicFields.map((fieldType) => decodeDynamicField(fieldType as DynamicSchemaType, new Uint8Array(0))),
  ];
  const indexedInitialValues = Object.fromEntries(
    defaultValues.map((value, index) => [index, value])
  ) as ComponentValue;

  const metadata = await registerMetadata(contract, table);
  if (metadata) {
    const { tableName, fieldNames } = metadata;
    const namedInitialValues = Object.fromEntries(
      defaultValues.map((fieldValue, schemaIndex) => {
        return [fieldNames[schemaIndex], fieldValue];
      })
    ) as ComponentValue;
    return {
      schema,
      indexedValues,
      indexedInitialValues,
      namedValues: {
        [fieldNames[schemaIndex]]: indexedValues[schemaIndex],
      },
      namedInitialValues,
    };
  }

  console.warn(
    `Received data for ${table.toString()}, but could not find table metadata for field names. Did your contracts get autogenerated and deployed properly?`
  );
  return {
    schema,
    indexedValues,
    indexedInitialValues,
  };
}
