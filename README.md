# forum

Purpose of this project is creating a web forum that allows:

- creating posts and comments to posts;
- associating categories to posts;
- liking and disliking posts and comments;
- filtering user created or liked posts.

The objectives of project is [here](https://github.com/01-edu/public/tree/master/subjects/forum).

## Chapters

- [How to run?](#how-to-run)
- [Authors](#authors)

## How to run?

### Prerequisites

- Your environment must need at least [**Docker**](#docker). If you running project by [default](#default) you need **Go** installed in your system.
- You can [configurate](#database-configuration) your database connection. But this is optional.

### Database configuration

In the `configs` directory exists example of `.env` file. There settings that you can change by yourself.

### Default

1. Download necessary Go modules:
```bash
go mod download 
```

2. Run:

Easy server start:
```bash
make run
```

OR build the application and run:
```bash
make build
./main
```

3. Open forum in browser by address [`locahost:8080`](http://localhost:8080)

### Docker

1. Run docker:
```bash
make docker
```
2. Open forum in browser by address [`locahost:8080`](http://localhost:8080)
3. Closing server when done:
```bash
make prune
```

## Authors
- [nrblzn](https://github.com/RaevNur)
- [damirkap89](https://github.com/KarbozovDamir)