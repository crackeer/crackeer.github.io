# filestash

```yml
version: '2'
services:
  app:
    container_name: filestash
    image: machines/filestash:latest
    restart: always
    environment:
    - APPLICATION_URL=
    - CANARY=true
    - OFFICE_URL=http://wopi_server:9980
    - OFFICE_FILESTASH_URL=http://app:8334
    - OFFICE_REWRITE_URL=http://127.0.0.1:9980
    ports:
    - "8334:8334"
    volumes:
    - filestash:/app/data/state/

  wopi_server:
    container_name: filestash_wopi
    image: collabora/code:24.04.10.2.1
    restart: always
    environment:
    - "extra_params=--o:ssl.enable=false"
    - aliasgroup1="https://.*:443"
    command:
    - /bin/bash
    - -c
    - |
         curl -o /usr/share/coolwsd/browser/dist/branding-desktop.css https://gist.githubusercontent.com/mickael-kerjean/bc1f57cd312cf04731d30185cc4e7ba2/raw/d706dcdf23c21441e5af289d871b33defc2770ea/destop.css
         /bin/su -s /bin/bash -c '/start-collabora-online.sh' cool
    user: root
    ports:
    - "9980:9980"

volumes:
    filestash: {}
```

# FTP服务

```yml
services:
  ftp-server:
    image: fauria/vsftpd:latest
    container_name: ftp-server
    ports:
      - "8090:21"
      - "8091-8098:8091-8098"
    volumes:
      - ./data:/home/vsftpd
      - ./log:/var/log/vsftpd
      - ./run-vsftpd.sh:/usr/sbin/run-vsftpd.sh
      - ./preset_virtual_users.txt:/etc/vsftpd/preset_virtual_users.txt
    environment:
      FTP_USER: "ftpuser"
      FTP_PASS: "ftpuser"
      PASV_ADDRESS: "10.33.207.152"
      PASV_MIN_PORT: "8091"
      PASV_MAX_PORT: "8098"
      REVERSE_LOOKUP_ENABLE: "NO"
    restart: always
```

## run-vsftpd.sh

```shell
#!/bin/bash

# If no env var for FTP_USER has been specified, use 'admin':
if [ "$FTP_USER" = "**String**" ]; then
    export FTP_USER='admin'
fi

# If no env var has been specified, generate a random password for FTP_USER:
if [ "$FTP_PASS" = "**Random**" ]; then
    export FTP_PASS=`cat /dev/urandom | tr -dc A-Z-a-z-0-9 | head -c${1:-16}`
fi

export LOG_STDOUT='Yes.'

# Create home dir and update vsftpd user db:

echo -e "${FTP_USER}\n${FTP_PASS}" > /etc/vsftpd/virtual_users.txt
if [ -f /etc/vsftpd/preset_virtual_users.txt ]; then
    cat /etc/vsftpd/preset_virtual_users.txt >> /etc/vsftpd/virtual_users.txt
fi
if [ -f /etc/vsftpd/virtual_users.db ]; then
    rm -f /etc/vsftpd/virtual_users.db
fi
/usr/bin/db_load -T -t hash -f /etc/vsftpd/virtual_users.txt /etc/vsftpd/virtual_users.db

# Create home directories for each virtual user.
# `virtual_users.txt` contains alternating lines: username, password.
if [ -f /etc/vsftpd/virtual_users.txt ]; then
    while IFS= read -r user; do
        if [ -n "$user" ]; then
            mkdir -p "/home/vsftpd/$user"
            chown ftp:ftp "/home/vsftpd/$user"
        fi
    done < <(awk 'NR%2==1 && NF' /etc/vsftpd/virtual_users.txt)
fi

# Set passive mode parameters:
if [ "$PASV_ADDRESS" = "**IPv4**" ]; then
    export PASV_ADDRESS=$(/sbin/ip route|awk '/default/ { print $3 }')
fi

echo "pasv_address=${PASV_ADDRESS}" >> /etc/vsftpd/vsftpd.conf
echo "pasv_max_port=${PASV_MAX_PORT}" >> /etc/vsftpd/vsftpd.conf
echo "pasv_min_port=${PASV_MIN_PORT}" >> /etc/vsftpd/vsftpd.conf
echo "pasv_addr_resolve=${PASV_ADDR_RESOLVE}" >> /etc/vsftpd/vsftpd.conf
echo "pasv_enable=${PASV_ENABLE}" >> /etc/vsftpd/vsftpd.conf
echo "file_open_mode=${FILE_OPEN_MODE}" >> /etc/vsftpd/vsftpd.conf
echo "local_umask=${LOCAL_UMASK}" >> /etc/vsftpd/vsftpd.conf
echo "xferlog_std_format=${XFERLOG_STD_FORMAT}" >> /etc/vsftpd/vsftpd.conf
echo "reverse_lookup_enable=${REVERSE_LOOKUP_ENABLE}" >> /etc/vsftpd/vsftpd.conf
echo "pasv_promiscuous=${PASV_PROMISCUOUS}" >> /etc/vsftpd/vsftpd.conf
echo "port_promiscuous=${PORT_PROMISCUOUS}" >> /etc/vsftpd/vsftpd.conf

# Get log file path
export LOG_FILE=`grep xferlog_file /etc/vsftpd/vsftpd.conf|cut -d= -f2`

# stdout server info:
if [ ! $LOG_STDOUT ]; then
cat << EOB
        *************************************************
        *                                               *
        *    Docker image: fauria/vsftpd                *
        *    https://github.com/fauria/docker-vsftpd    *
        *                                               *
        *************************************************

        SERVER SETTINGS
        ---------------
        · FTP User: $FTP_USER
        · FTP Password: $FTP_PASS
        · Log file: $LOG_FILE
        · Redirect vsftpd log to STDOUT: No.
EOB
else
    /usr/bin/ln -sf /dev/stdout $LOG_FILE
fi

# Run vsftpd:
&>/dev/null /usr/sbin/vsftpd /etc/vsftpd/vsftpd.conf
```

## preset_virtual_users.txt

```txt
ftpuser1
password1
ftpuser2
password2
```



# Minio

```yml
version: '3.8'
services:
  minio:
    image: minio:priv
    container_name: minio
    ports:
      - "4500:9000"
      - "4501:9001"
    volumes:
      - ./data:/data
    environment:
      MINIO_ROOT_USER: admin
      MINIO_ROOT_PASSWORD: 12345678
    command: server /data --console-address ":9001"
    restart: always
```

# MySQL8

## docker-composer.yml

```yam
version: "3.9"
services:
  mysql:
    image: mysql:8.0
    container_name: mysql8
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: MySQL-ROOT-PASSWORD
      MYSQL_USER: rushi 
      MYSQL_PASSWORD: RUSHI-PASSWORD
    ports:
      - "3306:3306"
    volumes:
      - /data1/mysql/data:/var/lib/mysql
      - /data1/mysql/conf:/etc/mysql/conf.d
    command: >
      --default-authentication-plugin=mysql_native_password
```

## my.conf

```conf
# /data1/mysql/conf/my.conf
[mysqld]
character-set-server=utf8mb4
collation-server=utf8mb4_unicode_ci
default-time-zone='+08:00'
secure-file-priv=/var/lib/mysql
sql_mode=''
max_connections=512
default_authentication_plugin=mysql_native_password
```

