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
  minio:
    image: fauria/vsftpd:latest
    container_name: ftp-server
    ports:
      - "8089:21"
      - "8091-8092:8091-8092"
    volumes:
      - ./data:/home/vsftpd
      - ./log:/var/log/vsftpd
    environment:
      FTP_USER: "ftpuser"
      FTP_PASS: "ftpuser"
      PASV_ADDRESS: "10.33.207.152"
      PASV_MIN_PORT: "8091"
      PASV_MAX_PORT: "8092"
      LOG_STDOUT: "1"
    restart: always
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

