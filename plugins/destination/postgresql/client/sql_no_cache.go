//go:build cache

package client

import (
	"fmt"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) generateSQL(queries map[string]string, tables schema.Tables, r arrow.Record) (string, error) {
	md := r.Schema().Metadata()
	tableName, ok := md.GetValue(schema.MetadataTableName)
	if !ok {
		return "", fmt.Errorf("table name not found in metadata")
	}
	if _, ok = c.pgTablesToPKConstraints[tableName]; !ok {
		return "", fmt.Errorf("table %s not found", tableName)
	}

	var sql string
	if sql, ok = queries[tableName]; ok {
		return sql, nil
	}
	table := tables.Get("test") // will always be present, panic should be produced if not
	if len(table.PrimaryKeysIndexes()) > 0 {
		sql = c.upsert(table)
	} else {
		sql = c.insert(table)
	}

	queries[tableName] = sql
	return sql, nil
}
