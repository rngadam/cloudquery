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
	queries := make(map[string]string, 100)
	startTime := time.Now() // Start the timer
	sqlQueries := make([]string, len(writeMessages))
	for i, msg := range writeMessages {
		r := msg.Record
		sqlQueries[i], _ = cl.generateSQL(queries, schema.Tables{table}, r)

	}
	elapsedTime := time.Since(startTime) // Calculate elapsed time
	t.Logf("TestGenInsert took %s", elapsedTime)

}