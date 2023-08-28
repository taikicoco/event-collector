#!/bin/sh

cd ../db/migrations

goose mysql "user:password@tcp(mysql)/db?charset=utf8mb4&parseTime=true" up
