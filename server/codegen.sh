#!/bin/bash

sqlboiler sqlite3 -c sqlboiler.toml --wipe
go run github.com/99designs/gqlgen generate
