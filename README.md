Live Streaming Recorder
======================

# Development

**Installing steps**

Copy database

```bash
$ cp -pf src/data/sqlite.db.brank src/data/sqlite.db
```

Start Up docker container

```bash
$ ./docker.sh -u
```

- Front end: http://localhost:3000/
- Back end: http://localhost:5050/
- sqlite-web: http://localhost:7000/


Shutdown docker container

```bash
$ ./docker.sh -d
```

# License

[MIT](https://choosealicense.com/licenses/mit/)
