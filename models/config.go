package models

type ConfigObject struct {
	Mode                            string
	IsQueryLoggingEnabled           bool
	DbHost                          string
	DbUsername                      string
	DbPassword                      string
	DbName                          string
	DbMaxOpenCons                   int64
	DbMaxIdleCons                   int64
	ListenerType                    string
	ListenerAddress                 string
	IsProcessUnstoredQueriesEnabled bool
	ProcessUnstoredQueries_Server   string
	StaticRecords                   []*Record
}
