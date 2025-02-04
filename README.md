# Overview

TODO diagram

TODO Style notes
- NO package init() functions
- Dynamic behaviour must be explicit
 
## Cross-compile server on host, run it in container 
These are simple steps using a single container.

1.  build server on host

        GOOS=linux GOARCH=arm64 go build

2.  build docker image

        cd cmd/server
        docker build -t server-image .

3.  Start container with shared directory

    ```sh
    docker run -it \
           -v   /Users/hohn/work-gh/mrva/mrvacommander:/mrva/mrvacommander \
           server-image
    ```

4.  Run server in container

        cd /mrva/mrvacommander/cmd/server/ && ./server

## Using docker-compose
### Steps to build and run the server in a multi-container environment set up by docker-compose.

1.  Built the server-image, above

1.  Build server on host

        cd ~/work-gh/mrva/mrvacommander/cmd/server/
        GOOS=linux GOARCH=arm64 go build

1.  Start the containers

        cd ~/work-gh/mrva/mrvacommander/
        docker-compose down
        docker-compose up -d
    
4.  Run server in its container

        cd ~/work-gh/mrva/mrvacommander/
        docker exec -it server bash
        cd /mrva/mrvacommander/cmd/server/ 
        ./server -loglevel=debug -mode=container

1.  Test server via remote client by following the steps in [gh-mrva](https://github.com/hohn/gh-mrva/blob/connection-redirect/README.org#compacted-edit-run-debug-cycle)

### Some general docker-compose commands

2.  Get service status

        docker-compose ps
        
3.  Stop services

        docker-compose down
        
4.  View all logs

        docker-compose logs

5.  check containers from server container

        docker exec -it server bash
        curl -I http://rabbitmq:15672

### Use the minio ql database db

1.  Web access via

        open http://localhost:9001/login

    username / password are in `docker-compose.yml` for now.  The ql db listing 
    will be at

        http://localhost:9001/browser/qldb

1.  Populate the database by running

        ./populate-dbstore.sh
        
    from the host.

1.  The names in the bucket use the `owner_repo` format for now,
    e.g. `google_flatbuffers_db.zip`.
    TODO This will be enhanced to include other data later

1.  Test Go's access to the dbstore -- from the host -- via

        cd ./test
        go test -v

    This should produce

        === RUN   TestDBListing
        dbstore_test.go:44: Object Key: google_flatbuffers_db.zip
        dbstore_test.go:44: Object Key: psycopg_psycopg2_db.zip

### Use the minio query pack db

1.  Web access via

        open http://localhost:19001/login

    username / password are in `docker-compose.yml` for now.  The ql db listing 
    will be at

        http://localhost:19001/browser/qpstore

    
### To run Use the minio query pack db
