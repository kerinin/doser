#!/bin/bash

sqlboiler sqlite3 -c sqlboiler.toml
go run github.com/99designs/gqlgen generate
