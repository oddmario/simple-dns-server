# Simple DNS Server

## Installation
1. Create a file with the path `/etc/systemd/system/argondns.service` and with the following content:
```
[Unit]
Description=ArgonDNS

[Service]
User=root
WorkingDirectory=/root/argondns
LimitNOFILE=2097152
TasksMax=infinity
ExecStart=/root/argondns/simpledns_linux_amd64
Restart=on-failure
StartLimitInterval=180
StartLimitBurst=30
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

2. Put the `simpledns_linux_amd64` executable file at a directory with path `/root/argondns/` - so the final path of the executable will be `/root/argondns/simpledns_linux_amd64`

3. Run `chmod -R 777 /root/argondns/simpledns_linux_amd64`

4. Place the `config.json` file found in this repository at `/root/argondns` along with the `simpledns_linux_amd64` executable file

5. Run `systemctl enable argondns`

6. Port 53 (the DNS server port) is usually in use by default. To solve this, follow https://unix.stackexchange.com/a/676977/405697 then run `systemctl restart systemd-resolved`

7. Make sure that there are no other DNS servers (such as bind9) are running, then run `systemctl start argondns` to start our DNS server :) 