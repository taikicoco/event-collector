#!/bin/sh

chmod +x ../.git/hooks/pre-commit
figlet GoLogCollector
reflex -r '(\.go|go\.mod)' -s go run main.go