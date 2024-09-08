# 🌎 Simple DNS server
A tiny DNS server that is capable of serving records configured in a MySQL table, or configured statically in a JSON file

## 🧐 Configuration documentation

- `mode`: Can be either `db` if your records are stored in a MySQL database, or `static_records` if your records are static and stored in the configuration JSON file.

- `db`: The MySQL server & database credentials. This works only if `mode` is set to `db`

- `listener`: The listening/bind settings for the DNS server (usually has to be kept binding on port 53 to be able to accept DNS requests).

- `process_unstored_dns_queries`: Should the DNS server also accept queries of records that are not stored in your database table/static records configuration? Enable this if yes.

- `static_records`: Configure your static records here, one per JSON array. This works only if `mode` is set to `static_records`

## 🛠️ Installation as a service

1. Store your configuration file at `/etc/simpledns/config.json`
   You can copy the example configuration file and change it to serve your needs.
2. If running Simple DNS server in the `db` mode, use this database structure for your records table: https://github.com/oddmario/simple-dns-server/blob/main/db_structure.sql
3. Place the binary file of Simple DNS server at `/usr/local/bin` (e.g. `/usr/local/bin/simpledns`)
4. Make the binary file executable: `chmod u+x /usr/local/bin/simpledns`
5. Create a systemd service for the application. This can be done by creating /etc/systemd/system/simpledns.service to have this content:
```
[Unit]
Description=SimpleDNSserver

[Service]
User=root
WorkingDirectory=/usr/local/bin
LimitNOFILE=2097152
TasksMax=infinity
ExecStart=/usr/local/bin/simpledns /etc/simpledns/config.json
Restart=on-failure
StartLimitInterval=180
StartLimitBurst=30
RestartSec=5s

[Install]
WantedBy=multi-user.target

```
6. Port 53 (the DNS server port) is usually in use by default. To solve this, follow https://unix.stackexchange.com/a/676977/405697 then run `systemctl restart systemd-resolved`
7. Make sure that there are no other DNS servers (such as bind9) are running
8. Enable the Simple DNS server service on startup & start it now:
```
systemctl enable --now simpledns.service
```