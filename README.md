# Go Web

This is a web app generator written in Go. It uses the following packages:

- [github.com/labstack/echo](https://godoc.org/github.com/labstack/echo) for the web app
- [github.com/jmoiron/sqlx](https://godoc.org/github.com/jmoiron/sqlx) for the database access
  - [github.com/lib/pq](https://godoc.org/github.com/lib/pq) if you prefer PostgreSQL
  - [github.com/go-sql-driver/mysql](https://godoc.org/github.com/go-sql-driver/mysql) if you prefer MySQL
- [html/template](https://godoc.org/html/template) for the templating engine

The aim is to scaffold an entire web app from scratch and include basic
generator functions for new handlers, etc. This is driven by a desire to
remove some of the boilerplate typing that goes into building a Go web app,
without going too far towards a Rails-like framework.
