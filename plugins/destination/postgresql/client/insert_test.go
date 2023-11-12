package client

import (
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func TestGenInsert(t *testing.T) {
	table := schema.TestTable("test", schema.TestSourceOptions{
		SkipMaps:      true,
		TimePrecision: time.Microsecond, // only us precision supported by time cols
	})
	writeMessages := make(message.WriteInserts, 100000)

	for i := range writeMessages {
		tg := schema.NewTestDataGenerator()

		normalRecord := tg.Generate(table, schema.GenTestDataOptions{
			MaxRows:       1,
			TimePrecision: schema.TestSourceOptions{}.TimePrecision,
		})
		writeMessages[i] = &message.WriteInsert{
			Record: normalRecord,
		}
	}
	cl := Client{}
	cl.pgTablesToPKConstraints = map[string]string{}

	cl.pgTablesToPKConstraints["test"] = "id"
	queries := make(map[string]string, 100)
	startTime := time.Now() // Start the timer
	sqlQueries := make([]string, len(writeMessages))
	for i, msg := range writeMessages {
		r := msg.Record
		sqlQueries[i], _ = cl.generateSQL(queries, schema.Tables{table}, r, true)

	}
	elapsedTime := time.Since(startTime) // Calculate elapsed time
	t.Logf("With caching enabled took %s", elapsedTime)
	queries = make(map[string]string, 100)
	startTimeNoCaching := time.Now() // Start the timer
	sqlQueriesNoCaching := make([]string, len(writeMessages))
	for i, msg := range writeMessages {
		r := msg.Record
		sqlQueriesNoCaching[i], _ = cl.generateSQL(queries, schema.Tables{table}, r, false)

	}
	elapsedTimeNoCaching := time.Since(startTimeNoCaching) // Calculate elapsed time
	t.Logf("With caching disabled took %s", elapsedTimeNoCaching)

}
