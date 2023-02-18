#!/bin/sh

chmod +x ../.git/hooks/pre-commit
figlet WEB API
reflex -r '(\.go|go\.mod)' -s go run main.go