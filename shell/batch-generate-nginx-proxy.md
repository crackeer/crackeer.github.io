# 快速生成nginx代理配置

```bash
#!/bin/bash
httpConf=$(cat << 'EOF'
server {
    listen 80;
    server_name HOSTNAME;

    location / {
        proxy_pass http://HOST_IP;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
EOF
)

streamConf=$(cat << 'EOF'
server {
    listen LISTEN_PORT;
    server_name HOSTNAME;
    proxy_pass HOST_IP:TARGET_PORT;
}
EOF
)

function get_ip() {
    name=$1
    host_ip=$(dig $name | grep $name | grep  -v ';' | awk '{print $5}')
}

hostNames=(
    i.meta.test.com
)
for name in ${hostNames[@]}; do
    get_ip $name
    echo "Host: $name, IP: $host_ip"
    echo "$httpConf" > /etc/nginx/conf.d/$name.conf
    sed -i "s/HOSTNAME/$name/g" /etc/nginx/conf.d/$name.conf 
    sed -i "s/HOST_IP/$host_ip/g" /etc/nginx/conf.d/$name.conf
done

streamHost=(
    "443,github.com,443"
)

if [ ! -d "/etc/nginx/stream.d" ]; then
    mkdir /etc/nginx/stream.d
fi
for item in ${streamHost[@]}; do
    IFS=',' read -r -a arr <<< "$item"
    listen_port=${arr[0]}
    host_name=${arr[1]}
    target_port=${arr[2]}
    get_ip $host_name
    echo "Stream Host: $host_name, IP: $host_ip, Listen Port: $listen_port, Target Port: $target_port"
    echo "$streamConf" > /etc/nginx/stream.d/$host_name-$listen_port.conf
    sed -i "s/HOSTNAME/$host_name/g" /etc/nginx/stream.d/$host_name-$listen_port.conf 
    sed -i "s/HOST_IP/$host_ip/g" /etc/nginx/stream.d/$host_name-$listen_port.conf
    sed -i "s/LISTEN_PORT/$listen_port/g" /etc/nginx/stream.d/$host_name-$listen_port.conf
    sed -i "s/TARGET_PORT/$target_port/g" /etc/nginx/stream.d/$host_name-$listen_port.conf
done
```