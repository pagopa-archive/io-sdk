docker-compose up -d
sleep 5
cat messages.psql | docker exec -i database_postgres_1 su - postgres  -c 'psql'
curl -X POST -d @track.json -H "Content-Type: application/json" -H "X-Hasura-Role: admin" http://localhost:8080/v1/query
