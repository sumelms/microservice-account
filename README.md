<p align="center">
  <img src=".github/sumelms.png" />
</p>

<p align="center">
  <a href="https://travis-ci.org/sumelms/microservice-user">
    <img alt="Travis (.org)" src="https://travis-ci.org/sumelms/microservice-user.svg?branch=main">
  </a>  
  <a href="https://codecov.io/gh/sumelms/microservice-user">
    <img src="https://codecov.io/gh/sumelms/backend/microservice-user/main/graph/badge.svg?token=8E92BS3SR9" />
  </a>
  <img alt="GitHub" src="https://img.shields.io/github/license/sumelms/microservice-user">
  <a href="https://discord.gg/Yh9q9cd">
    <img alt="Discord" src="https://img.shields.io/discord/726500188021063682">
  </a>
</p>

## About Sumé LMS

> Note: This repository contains the **user microservice** of the Sumé LMS. If you are looking for more information 
> about the application, we strongly recommend you to [check the documentation](https://www.sumelms.com/docs).

Sumé LMS is a modern and open-source learning management system that uses modern technologies to deliver performance 
and scalability to your learning environment.

  * Compatible with SCORM and xAPI (TinCan)
  * Flexible and modular
  * Open-source and Free
  * Fast and modern
  * Easy to install and run
  * Designed for microservices
  * REST API based application
  * and more.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Prepare](#prepare)
- [Building](#building)
- [Running](#running)
- [Configuring](#configuring)
- [Testing](#testing)
- [Contributing](#contributing)
- [Team](#team)
- [Support](#support)
- [License](#license)

## Prerequisites

- Go >= 1.14.6
- PostgreSQL >= 9.5

## Prepare

Clone the repository

```bash
$ git clone [git@github.com](mailto:git@github.com):sumelms/microservice-user.git
```

Access the project folder, and download the Go dependencies

```bash
$ go get
```

It may take a while to download all the dependencies, then you are [ready to build](#building).

## Building

There are two ways that you can use to build this microservice. The first one will build it using your own machine, 
while the second one will build it using a docker instance. Also, you can build the docker image to use it with 
[Docker](https://www.docker.com/) and [Kubernetes](https://kubernetes.io/), but it is up to you.

Here are the following instructions for each available option:

### Local build

It should be pretty simple, once all the dependencies are download just run the following command:

```bash
$ make build
```

It will generate an executable file at the `/bin` directory inside the project folder, and If everything works, you can 
now [run the microservice](#local-run).

### Docker build

At this point, I'll assume that you have installed and configure the Docker in your system, but if it is not the case, 
visit the [https://docs.docker.com/get-started/](https://docs.docker.com/get-started/).

```bash
$ make docker-build
```

If everything works, you can now [run the microservice using the docker image](#docker-run).

## Running

OK! Now you build it you need to run the microservice. That should also be pretty easy.

### Local run

If you want to run the microservice locally, you may need to first [configure it](#configuring). 
Then you can run it, you just need to execute the following command:

```bash
$ make run
```

### Docker run

If you want to run the microservice using Docker, the easiest way to do it is setting it up with `docker-compose`, 
here is a basic example for a `docker-compose.yml` file:

```bash
# sumelmes/microservice-user/docker-compose.yml
version: '3'
services:
  microservice:
    container_name: microservice-user
    build: sumelms/sumelms-user
    ports: 
      - 8080:8080 
    restart: on-failure
    volumes:
      - sumelms_user:/usr/src/sumelms-user/
    depends_on:
      - postgres
    networks:
      - sumelms

  postgres:
    image: postgres:latest
    container_name: microservice-user-postgres
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=password
      - POSTGRES_DB=microservice_user
    ports:
      - '5432:5432'
    volumes:
      - database_postgres:/var/lib/postgresql/data
    networks:
      - sumelms

volumes:
  microservice:
  database_postgres:                  

networks:
  sumelms:
    driver: bridge
```

You don't need to create this file inside the project folder, but in order to simplify it, let assume that you did it, 
so, you just need to run the following command:

```bash
$ docker-composer up
```

Keep in mind that it will load the `config/config.yml` file from the project. If you want to change some 
configurations you can set the environment variables in your `docker-compose.yml` file, or edit the configuration file.

## Configuring

You can easily configure the application editing the `config/config.yml` file or using environment variables. We do 
strongly recommend that you use the configuration file instead of the environment variables. Again, it is up to you 
to decide. If you want to use the variables, be sure to prefix it all with `SUMELMS_`. 

The list of the environment variables and it's default values:

```bash
SUMELMS_SERVER_HTTP_PORT = 8080
SUMELMS_DATABASE_DRIVER = "postgres"
SUMELMS_DATABASE_HOST = "localhost"
SUMELMS_DATABASE_PORT = 5432
SUMELMS_DATABASE_USER = nil
SUMELMS_DATABASE_PASSWORD = nil
SUMELMS_DATABASE_DATABASE = "sumelms_user"
```

> We are using [configuro](https://github.com/sherifabdlnaby/configuro) to manage the configuration, so the precedence 
> order to configuration is: *Environment variables > .env > Config File > Value set in Struct before loading.*

## Testing

You can run all the tests with one single command:

```bash
$ make test
```

## Contributing

Thank you for considering contributing to the project. In order to ensure that the Sumé LMS community is welcome to 
all make sure to read our [Contributor Guideline](https://www.sumelms.com/docs/contributing).

## Team

### Core

- Ricardo Lüders (@rluders)
- Ariane Rocha (@arianerocha)

### Contributors

...

## Support

Do you need any help? 

### Discussion

You can reach us or get community support in our [Discord server](https://discord.gg/nRVVeWR). This is the best way to 
find help and get in touch with the community.

### Bugs or feature requests

If you found a bug or have a feature request, the best way to do 
it is [opening an issue](https://github.com/sumelms/sumelms/issues).

## License

This project licensed by the Apache License 2.0. For more information check the LICENSE file.
