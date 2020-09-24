# Question & Answer API

Coding Challenge - Golang:

You are to design a mini system for the following business idea.

We want to build a site called QuestionsAndAnswers.com that will compete with Quora/Stackoverflow and others with 1 major difference.  We will only allow 1 answer per question.  If someone thinks they have a better answer, they will have to update the existing answer for that question instead of adding another answer.  In essence, each question can only have 0 or 1 answer.

The MVP site will be very simple consisting of two pages:

LIST Page: Contains the list of all the ”question and answer” pairs displayed in 2 columns.  At the top of the page there is also a link/button called “Create” which takes you to the edit/create page.  Also each question is clickable which takes you to the edit/create page as well.

EDIT/CREATE PAGE: This is the page where a new question can be created. If coming from an existing question link, the answers to the question can be edited.

No user tracking or security needed for this version.  We would like to receive code that runs.  The other parts that we’d like to see are (including the reason(s) why you made those decisions):

Documentation on how to set up the solution and run it

Database design

API Interface design 

Backend Code.  Frontend is optional - please specify how to make backend calls if FE is missing

What’s missing that you wish you had more time for?  Please think about the different problems you might encounter if the business idea is successful.  This would include considerations such as increased load, increased data and an upvoting feature.

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
