#|/bin/bash

sudo su
cd /root
apt update -y

wget -O proxysocks https://github.com
chmod +x proxysocks

cat > /etc/systemd/system/proxysocks.service <<-END
[Unit]
Description=ProxySOCKS
After=network.target

[Service]
ExecStart=/root/proxysocks
WorkingDirectory=/root/
StandardOutput=inherit
StandardError=inherit
Restart=always

[Install]
WantedBy=multi-user.target
END

systemctl daemon-reload

systemctl enable proxysocks
systemctl restart proxysocks
