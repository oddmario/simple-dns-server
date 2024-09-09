package records

import (
	"mario/simple-dns-server/config"
	"mario/simple-dns-server/db"
	"mario/simple-dns-server/models"
)

func GetDNSRecords(record_name_, record_type_ string) (bool, []*models.Record) {
	records := []*models.Record{}
	recordsFound := false

	if config.Config.Mode == "db" {
		res, err := db.RetriedDbQuery(10, "SELECT id, record_type, record_name, record_value, record_ttl, srv_priority, srv_weight, srv_port, srv_target, is_disposable FROM dns_records WHERE record_name = ? AND record_type = ?", record_name_, record_type_)
		if err != nil {
			// an error has occured while preparing the SQL statement

			return recordsFound, records
		}

		defer res.Close()

		for res.Next() {
			recordsFound = true

			var record_id int64
			var record_type string
			var record_name string
			var record_value string
			var record_ttl int64
			var srv_priority int64
			var srv_weight int64
			var srv_port int64
			var srv_target string
			var record_isdisposable int64

			err = res.Scan(&record_id, &record_type, &record_name, &record_value, &record_ttl, &srv_priority, &srv_weight, &srv_port, &srv_target, &record_isdisposable)
			if err != nil {
				// an error has occured

				recordsFound = false

				return recordsFound, records
			}

			isDisposable := false
			if record_isdisposable >= 1 {
				isDisposable = true
			}

			records = append(records, &models.Record{
				ID:           record_id,
				Type:         record_type,
				Name:         record_name,
				Value:        record_value,
				TTL:          record_ttl,
				SRVPriority:  srv_priority,
				SRVWeight:    srv_weight,
				SRVPort:      srv_port,
				SRVTarget:    srv_target,
				IsDisposable: isDisposable,
			})
		}
	} else {
		for _, record := range config.Config.StaticRecords {
			if record.Name == record_name_ && record.Type == record_type_ {
				recordsFound = true

				records = append(records, &models.Record{
					ID:           -1,
					Type:         record.Type,
					Name:         record.Name,
					Value:        record.Value,
					TTL:          record.TTL,
					SRVPriority:  record.SRVPriority,
					SRVWeight:    record.SRVWeight,
					SRVPort:      record.SRVPort,
					SRVTarget:    record.SRVTarget,
					IsDisposable: false,
				})
			}
		}
	}

	return recordsFound, records
}
