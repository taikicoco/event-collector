#!/bin/sh

curl -XPOST "http://localhost:1222/log/detail" \
	-d '{"name": "log_test_name", "page_url": "www.exsampl.com", "version":1 }' \
	-H "Content-Type: application/json" \
