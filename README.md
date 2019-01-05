# recorder
## for recording and relaying inspection results

### Usage
```
$ make build

$ docker-compose build

$ docker-compose up
```

The service is now running. Get some inspections:
```
curl -d '{"limit": 2}' 0.0.0.0:<localport>/inspections
```

