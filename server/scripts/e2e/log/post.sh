#!/bin/sh

curl -XPOST "http://localhost:1222/log/1" \
	-d '{"access": 1, "conversion": 1}' \
	-H "Content-Type: application/json" \
