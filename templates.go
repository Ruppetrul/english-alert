package main

const serviceTemplate = `[Unit]
Description={{.Description}}
After=graphical-session.target

[Service]
Type=simple
WorkingDirectory={{.WorkingDirectory}}
ExecStart=/bin/sh -c 'eval $(dbus-launch --auto-syntax) && {{.ExecPath}}'
User=1000
Environment=DISPLAY=:0 DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus
`

const timerTemplate = `[Unit]
Description=Run {{.ServiceName}} every 5 minutes

[Timer]
OnCalendar=*:0/5
Unit={{.ServiceName}}.service

[Install]
WantedBy=timers.target`
