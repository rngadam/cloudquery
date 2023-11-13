package client

import (
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func BenchmarkNoCache(b *testing.B) {
	benchmarkQueries(b, false)
}

func BenchmarkCache(b *testing.B) {
	benchmarkQueries(b, true)
}

func benchmarkQueries(b *testing.B, cache bool) {
	b.Helper()
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
		r := msg.Record
		sqlQueries[i], _ = cl.generateSQL(queries, schema.Tables{table}, r, cache)
	}
}
