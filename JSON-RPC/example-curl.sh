#!/usr/bin/env bash
curl -X POST \
http://localhost:9000/rpc \
-H 'cache-control: no-cache' \
-H 'content-type: application/json' \
-d '{
    "method": "JSONServer.TwitterProfileDetail",
    "params": [
        {
            "name":"Gbedilodo"
        }
    ],
    "id": "2"
}'
