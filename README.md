# Question & Answer API

## Tech stack
1. API service: `go`
2. Database: `MongoDB`

## Testing the application
1. Run static analysis tool (`go vet`) 
```
make vet
```

2. Check for coding style mistakes (`golint`)
```
make lint
```

3. Run unittests
```
make test
```

## Run
1. Running API and Database
```
docker-compose up --build -d
```

2. Checking API logs
```
docker-compose logs -f api
```

2. Stopping API and Database
```
docker-compose down
```

## Usage

1. Create Question
```
curl -X POST \
    -H 'Content-Type: application/json' \
    --url http://localhost:8080/api/qa/questions \
    -d '
{
    "statement": "test?"
}'
```

2. Update Question
```
curl -X PUT \
    -H 'Content-Type: application/json' \
    --url http://localhost:8080/api/qa/questions/5f62bcd44a84a4700bb84e87 \
    -d '
{
    "id": "5f62bcd44a84a4700bb84e87",
    "statement": "test?"
    "answer": "abc"
}'
```

3. List all Questions
```
curl -X GET \
    -H 'Content-Type: application/json' \
    --url http://localhost:8080/api/qa/questions
```

4. Get one Question
```    
curl -X GET \
    -H 'Content-Type: application/json' \
    --url http://localhost:8080/api/qa/questions/5f62bac267ce62dc34dcbca8
```

5. Delete one Question
```
curl -X DELETE \
    -H 'Content-Type: application/json' \
    --url http://localhost:8080/api/qa/questions/5f62bac267ce62dc34dcbca8
```