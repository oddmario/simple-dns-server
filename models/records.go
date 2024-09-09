package models

type Record struct {
	ID           int64
	Type         string
	Name         string
	Value        string
	TTL          int64
	SRVPriority  int64
	SRVWeight    int64
	SRVPort      int64
	SRVTarget    string
	IsDisposable bool
}
