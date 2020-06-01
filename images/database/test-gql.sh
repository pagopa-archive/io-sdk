#!/bin/bash
cd "$(dirname $0)"
curl -s -X POST -H "Content-Type: application/json" \
  --data '{ "query": "{ messages { destinatario scadenza testo titolo } }"}' \
  http://localhost:8080/v1/graphql | jq .

