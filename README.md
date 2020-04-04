# PhotoServer

## Install Mongodb first

```bash
$ docker pull mongodb
$ docker run --rm -p 127.0.0.1:27017-27019:27017-27019 --name mongo -v /yourDBdirectory:/data/db -d mongo --wiredTigerCacheSizeGB 1.5
```

## Install photoserver

```
$ go get -u github.com/laowalter/photoserver
```

## Complile tools

1. Insert photo data to DB.

```bash
$ cd photoserver
$ cd util/picupdate
$ go build picupdate.go
```

2. DB clean tools for outdate data.

```bash
$ cd photoserver
$ cd util/dbclean
$ go build dbclean.go
```

## Useage

1. Insert photo album located in "/data/album" in to DB

```bash
$ picupdate -path /data/ablum
```

2. Run dbclean

```bash
$ dbclean
```

## Start photo server

```bash
$ cd photoserver
$ go build main.go
$ ./main
```

Visit http://localhost:8080/

