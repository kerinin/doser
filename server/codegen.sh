#!/bin/bash -xe

go run "github.com/kerinin/doser/service/cmd/migrate" -data data.db
rm -rf models
sqlboiler sqlite3 -c sqlboiler.toml --wipe
go run github.com/99designs/gqlgen generate
