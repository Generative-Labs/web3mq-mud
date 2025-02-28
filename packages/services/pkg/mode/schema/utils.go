package schema

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

/////////////////////////////////////////
// Utilities for building table names. //
/////////////////////////////////////////

func TableIdToTableName(tableId string) string {
	// Table ID comes in as a uint256 in string format comprised of two bytes16s
	// concatenated.
	// Get a byte array of the table ID.
	tableIdBytes := []byte(tableId[2:])

	b1 := tableIdBytes[:32]
	b2 := tableIdBytes[32:]
	namePart1 := string([]byte(b1))
	namePart2 := string([]byte(b2))

	p1 := string(common.FromHex("0x" + namePart1))
	p2 := string(common.FromHex("0x" + namePart2))

	return strings.Trim(p1, "\u0000") + CONNECTOR + strings.Trim(p2, "\u0000")
}

func TableNameToTableId(tableName string) string {
	// Assumes that table name was generated by TableIdToTableName.
	parts := strings.Split(tableName, CONNECTOR)
	if len(parts) != 2 {
		panic("Invalid table name: " + tableName)
	}
	return "0x" + common.Bytes2Hex(append(common.FromHex("0x"+parts[0]), common.FromHex("0x"+parts[1])...))
}

/////////////////////////////////////////
// Utilities for building field names. //
/////////////////////////////////////////

func DefaultFieldName(index int) string {
	return "field_" + fmt.Sprint(index)
}

func DefaultKeyName(index int) string {
	return "key_" + fmt.Sprint(index)
}
