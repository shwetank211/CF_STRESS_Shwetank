Ensure that you have `docker` and `docker-compose` installed. To analyze MongoDB database, you will also need to install `MongoDB Compass`.


Start all the services via the command:
```shell
docker-compose up
```

To stop all services, press `Ctrl+C`.

If you want to keep running the services even after closing the terminal, use the `-d` or `--detach` flag:
```shell
docker-compose up -d
```

The services created are:
* **web**
* **asynq-server**
* **asynqmon**
* **prometheus**
* **redis**
* **mongo**
* **nginx**

To simulate a distributed system, where there are multiple copies of judge/server running on different systems, use the `scale` flag from `docker-compose`. For example, if you want the web server to be running on 2 machines  and the judge to be running on 3 different machines, use the command

```shell
docker-compose up --scale web=2 --scale judge=3
```

Once the application has started, a sample request to the server might looks like:

```shell
localhost:80/internal/simulate-concurrent-users
```

Note: `nginx` runs on port 80. You don't need to worry about which port the server runs on.

To access the application stats, visit `Prometheus` homepage at

```shell
localhost:9090
```

To access the judge statistics, visit the Judge Monitor at 

```shell
localhost:8080
```

To access the databse, open up MongoDB Compass and use this connection string

```shell
mongodb://localhost:27017
```
