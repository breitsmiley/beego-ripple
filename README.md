# beego-ripple
Beego test

1. git clone https://github.com/breitsmiley/beego-ripple.git
2. generate ssl.crt and ssl.key, put them into conf/app
3. edit conf/app.conf

Run for prod env (access by https://domain.com:10443)

```
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

```

Run for prod env (access by http://domain.com:8080)

```
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d

```