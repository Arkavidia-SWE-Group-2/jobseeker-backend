# Jobseeker Backend

An backend application using Golang for Arkavidia - Academya Final Project, that focus on Jobseeker application

## Tech Stack

- Golang

  - Fiber: Web Framework
  - GORM: Database ORM

- PostgreSQL

## Getting Started

Quick guide how to get started with this project.

## Prerequisites

- Go (>= go1.22)
- PostgreSQL
- Git

## Installation

1. Clone repository

```sh
git clone https://github.com/Arkavidia-SWE-Group-2/jobseeker-backend
```

2. Move into directory

```sh
cd jobseeker-backend
```

3. Copy the config yaml file

```sh
cp config.example.yaml config.yaml
```

4. Generate private and public key for jwt authentication

```sh
make genkey
```

5. Install all dependencies

```sh
go mod tidy
```

6. Run application

> Using golang

```sh
go run cmd/api/main.go
```

> Using makefile

```sh
make run
```
