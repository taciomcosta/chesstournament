# Chess Tournament - Web API Example in Golang

A Go sample project to summarize some of my thoughts and practices when developing software. :)

This project was inspired by the data model below. (Find more [here](http://www.databaseanswers.org/data_models/chess_tournaments/index.htm))

*add image here*


## Running
#### docker-compose
We are going to need a Postgres instance. You can set up the API and Postgres with a single docker-compose command
`docker-compose up --build`

#### make
Alternatively, you can set up your own Postgres instance and change the database configuration in `.env`, in the root folder.

You can also benefit from other useful `make` commands already configured, like:
`make test`, `make coverage`, `make build`, `make build:windows`

## Structure
```
├── cmd // contains apps entrypoints (delivery mechanisms)
│   └── webapi // REST app 
│       ├── handlers
│       │   ├── chessclub.go
│       │   └── chessclub_test.go
│       └── main.go
├── internal // contains reusable packages
│   └── chessclub 
│       ├── chessclub.go
│       ├── chessclub_mock.go
│       └── chessclub_test.go
└── README.md
```

## Features
- [ ] Basic CRUD operations for *Chess Club, Player, Tournament, Match, Organizer and  Sponsor*
- [ ] REST API & Swagger docs
- [ ] CLI API
- [ ] GraphQL API
- [ ] Postgres integration
- [ ] 

Test-driven developed :hearts:
