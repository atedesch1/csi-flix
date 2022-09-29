# CSI-FLIX

## Run

### Requirements

- Api Key from http://www.omdbapi.com/
- docker
- docker-compose

### Configuration

Replace the api key environment variable inside the `docker-compose.yml` file (API_KEY=REPLACEME) with the api key from omdbAPI. 

### Spin up the containers

To run the application run the following command in the root folder:

```bash
docker-compose up
```

This will spin up postgres, go and vue containers and setup a network between them and outside in order to communicate.

Access the app on http://localhost:9000