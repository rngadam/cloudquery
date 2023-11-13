package client

import (
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func BenchmarkSQL(b *testing.B) {
	table := schema.TestTable("test", schema.TestSourceOptions{
		SkipMaps:      true,
		TimePrecision: time.Microsecond, // only us precision supported by time cols
	})
	writeMessages := make(message.WriteInserts, b.N)

	tg := schema.NewTestDataGenerator()
	for i := range writeMessages {
		writeMessages[i] = &message.WriteInsert{Record: tg.Generate(table, schema.GenTestDataOptions{MaxRows: 1})}
	}

	cl := Client{}
	cl.pgTablesToPKConstraints = map[string]string{}

	cl.pgTablesToPKConstraints["test"] = "id"
	queries := make(map[string]string, 100)
	sqlQueries := make([]string, b.N)

	b.ResetTimer()

	for i, msg := range writeMessages {
		sqlQueries[i], _ = cl.generateSQL(queries, schema.Tables{table}, msg.Record)
	}
}
