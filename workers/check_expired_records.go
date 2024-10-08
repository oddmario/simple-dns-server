package workers

import (
	"mario/simple-dns-server/config"
	"mario/simple-dns-server/db"
	"mario/simple-dns-server/utils"
	"time"
)

func checkExpiredRecords() {
	if config.Config.Mode != "db" {
		return
	}

	t := time.NewTicker(time.Second * 10)

	defer t.Stop()

	for range t.C {
		currTimestamp := time.Now().Unix()
		currTimestampStr := utils.I64ToStr(currTimestamp)

		db.EasyExec("DELETE FROM dns_records WHERE " + currTimestampStr + " >= delete_at_timestamp AND delete_at_timestamp > 0")
	}
}
