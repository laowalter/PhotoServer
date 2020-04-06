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

1. Compile program for uploading photo to DB.

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

## Usage

1. Upload photo album located in "/data/album" in to DB

```bash
$ picupdate -path /data/ablum
```

2. Run dbclean

Remove the photos moved or deleted from the previous upload directory.

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

