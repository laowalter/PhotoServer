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

# ToDo

* Test jquery
* picupdate 
    - check if there no command line argument.
    - check if input argument is a file or directory, can import single file.
    - add a swith -f to force update all the infomation of a photo or a directory.
* Add thumb/frame shot for mp4 mov 
* Parse raw format for Sony, Cannon and Nikon.
    - need to decide if use the raw file directly (time cost) to base64 
    - or save a large thumbnail to DB in advance(space cost) when show single pic.
    - or generate a large size from raw format and save it in jpg.
* Mobile client UI idea.
* Add User management
    - user management

* Other feature
    - Photos select from browser
    - Slide
    - GPS
        - google map api
        - select multple photos, use one gps position
        - update the original photo file, which doesnt have GPS info
    - PhotoRotation
    - tags belong per user.
