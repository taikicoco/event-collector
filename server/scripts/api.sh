#!/bin/sh

chmod +x ../.git/hooks/pre-commit
figlet Go API
reflex -r '(\.go|go\.mod)' -s go run main.go