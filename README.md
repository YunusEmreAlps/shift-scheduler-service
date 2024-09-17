# Shift Scheduluer Service

This service that allows users to manage their shift periods and create their shift lists. This service allows users to select which days and times they will be on duty in a given date range and create their shifts. In addition, users can view the shift roster for their shift period, and makeshift change requests, and have the option to delete or update existing shifts. Through notifications, users can be notified of updates related to their shift periods. The Shift Service provides a user-friendly platform that helps users to schedule shifts in an organized manner and manage their roster.

## Table of Contents

- [Shift Scheduluer Service](#shift-scheduluer-service)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Quick Start](#quick-start)
    - [Run with Docker](#run-with-docker)
      - [Docker development usage](#docker-development-usage)
    - [SWAGGER UI](#swagger-ui)
    - [Jaeger UI](#jaeger-ui)
    - [Prometheus UI](#prometheus-ui)
    - [Grafana UI](#grafana-ui)
    - [Project structure](#project-structure)
      - [`assets`](#assets)
      - [`config`](#config)
      - [`docs`](#docs)
      - [`/internal`](#internal)
      - [`/pkg`](#pkg)
      - [`handlers`](#handlers)
    - [Major Packages used in this project](#major-packages-used-in-this-project)
    - [Database Diagram](#database-diagram)
    - [Flow](#flow)
    - [Postman Collection](#postman-collection)
    - [Swagger Documentation](#swagger-documentation)
    - [Contact](#contact)

## Features

- RESTful API endpoints for CRUD operations.
- JWT Authentication. (**Not implemented yet**)
- Rate Limiting.
- Swagger Documentation.
- PostgreSQL database integration using GORM.
- Redis cache.
- Dockerized application for easy setup and deployment.
- Grafana, Prometheus and Jaeger

## Quick start

We can run this Shift Scheduler Service project with Docker. Here, I am providing both ways to run this project.

- Clone this project

```bash
# Move to your workspace
cd your-workspace

# Clone this project into your workspace
git clone https://github.com/YunusEmreAlps/shift-scheduler-service.git

# Move to the project root directory
cd shift-scheduler-service
```

### Run with Docker

- Create a file `.env` similar to `.env.example` at the root directory with your configuration.
- Install Docker and Docker Compose.
- Run `make` or `docker` commands.

```bash
make local // run all containers
make run // it's easier way to attach debugger or rebuild/rerun project
```

```bash
docker-compose -f docker-compose.local.yml up -d
```

- Access API using `http://localhost:9097`

#### Docker development usage

```bash
make docker
```

### SWAGGER UI

```bash
https://localhost:5000/swagger/index.html
```

### Jaeger UI

```bash
http://localhost:16686
```

### Prometheus UI

```bash
http://localhost:9090
```

### Grafana UI

```bash
http://localhost:3000
```

Unless you have configured Grafana differently, it is set to use [localhost](http://localhost:3000) by default. On the signin page, enter **admin** for username and password.

## Project structure

```bash
shift-scheduler-service/
|-- assets/
|   |-- v1_db.png
|-- config/
|   |-- config.go
|   |-- sample.env.yaml
|-- docker/
|-- internal/
|   |-- handlers/
|       |-- handlers.go
|   |-- middleware/
|       |-- jwt_middleware.go
|   |-- models/
|       |-- shift_schedule.go
|       |-- jsonb.go
|       |-- pagination.go
|       |-- jwt_response.go
|       |-- s3.go
|-- pkg/
|   |-- db/
|       |-- aws/
|           |-- aws.go
|       |-- postgres/
|           |-- db_conn.go
|       |-- redis/
|           |-- conn.go
|   |-- httpErrors/
|       |-- http_errors.go
|   |-- logger/
|       |-- logger.go
|   |-- metric/
|       |-- metric.go
|   |-- postman/
|       |-- Shift_Scheduler_Service.postman_collection.json
|   |-- sanitize/
|       |-- sanitize.go
|   |-- utils/
|       |-- helpers.go
|       |-- http.go
|       |-- pagination.go
|       |-- validator.go
|-- Dockerfile
|-- go.mod
|-- go.sum
|-- main.go
|-- Makefile
|-- README.md
```

### `assets`

Other assets to go along with your repository (images, logos, etc).

### `config`

Configuration. First, `config.yml` is read, then environment variables overwrite the yaml config if they match.
The config structure is in the `config.go`.
The `env-required: true` tag obliges you to specify a value (either in yaml, or in environment variables).

Reading the config from yaml contradicts the ideology of 12 factors, but in practice, it is more convenient than
reading the entire config from ENV.
It is assumed that default values are in yaml, and security-sensitive variables are defined in ENV.

### `docs`

Swagger documentation. Auto-generated by [swag](https://github.com/swaggo/swag) library.
You don't need to correct anything by yourself.

### `/internal`

Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 release notes for more details. Note that you are not limited to the top level internal directory. You can have more than one internal directory at any level of your project tree.

You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. Your actual application code can go in the /internal/app directory (e.g., /internal/app/myapp) and the code shared by those apps in the /internal/pkg directory (e.g., /internal/pkg/myprivlib).

### `/pkg`

Library code that's ok to use by external applications (e.g., /pkg/mypubliclib). Other projects will import these libraries expecting them to work, so think twice before you put something here :-) Note that the internal directory is a better way to ensure your private packages are not importable because it's enforced by Go. The /pkg directory is still a good way to explicitly communicate that the code in that directory is safe for use by others. The I'll take pkg over internal blog post by Travis Jeffery provides a good overview of the pkg and internal directories and when it might make sense to use them.

It's also a way to group Go code in one place when your root directory contains lots of non-Go components and directories making it easier to run various Go tools (as mentioned in these talks: Best Practices for Industrial Programming from GopherCon EU 2018, GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps and GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go).

See the /pkg directory if you want to see which popular Go repos use this project layout pattern. This is a common layout pattern, but it's not universally accepted and some in the Go community don't recommend it.

It's ok not to use it if your app project is really small and where an extra level of nesting doesn't add much value (unless you really want to :-)). Think about it when it's getting big enough and your root directory gets pretty busy (especially if you have a lot of non-Go app components).

The pkg directory origins: The old Go source code used to use pkg for its packages and then various Go projects in the community started copying the pattern (see this Brad Fitzpatrick's tweet for more context).

### `handlers`

The handlers directory contains the main handlers or controllers for the project. These handlers handle incoming requests, perform necessary actions, and return appropriate responses. They encapsulate the business logic and interact with other components of the project, such as services and data repositories.

It is important to note that the project structure described here may not include all the directories and files present in the actual project. The provided overview focuses on the key directories relevant to understanding the structure and organization of the project.

## Major Packages used in this project

- **gin**: Gin is an HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need a smashing performance, get yourself some Gin.
- **gorm**: Object Relational Mapping (ORM) library for Golang. ORM converts data between incompatible type systems using object-oriented programming languages.
- **postgreSQL go driver**: The Official Golang driver for PostgreSQL.
- **jwt**: JSON Web Tokens are an open, industry-standard RFC 7519 method for representing claims securely between two parties. Used for Access Token and Refresh Token.
- **viper**: For loading configuration from the `.env` file. Go configuration with fangs. Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, INI, envfile, or Java properties formats.
- **bcrypt**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- **testify**: A toolkit with common assertions and mocks that plays nicely with the standard library.
- Check more packages in `go.mod`.

## Database Diagram

![DB](/assets/v1_db.png)

## Flow

- ![Public API Request Flow](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-public-api-request-flow.png?raw=true)
- ![Private API Request Flow](https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/assets/go-arch-private-api-request-flow.png?raw=true)

## Postman Collection

The Postman collection for the shift scheduler project can be found in the postman directory. It includes a set of API requests that can be used to interact with the application. You can import the collection into Postman by following these steps:

1. Open Postman
2. Click on the "Import" button in the top left corner
3. Drag and drop the [Shift_Scheduler_Service.postman_collection.json](/pkg/postman/Shift_Scheduler_Service.postman_collection.json) file into the import window
4. The collection will be imported into Postman, and you can start using the API requests to interact with the application

## Swagger Documentation

The Swagger documentation for the shift scheduler project can be accessed by following these steps:

1. After running the project, open a web browser and navigate to `http://localhost:9097/shift-scheduler-service/swagger/index.html` to access the Swagger documentation
2. When you add a new API endpoint to the project, you need to run the `swag init` command to generate the Swagger documentation for the new endpoint. This command will update the `docs` directory with the new Swagger documentation.

## Contact

For any questions or support, please contact us at:

- Linkedin at [Yunus Emre Alpu](https://www.linkedin.com/in/yunus-emre-alpu-5b1496151/)
