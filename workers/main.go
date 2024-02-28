package workers

func Init() {
	go checkExpiredRecords()
}
