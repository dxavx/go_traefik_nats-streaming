## Go API + Traefik (Load Balancer)
There is an example of using the “Traefik " load balancer and Go code with the ability to increase the number of instances of the container with the Go code (scaling ).  

In the file __build.sh__ you can set the number of copies of the app: `docker-compose -f./docker-compose.yml up -d -- scale api=3`

Start: ``` bash build.sh ```

API  ``` http://localhost/v1/ping ```

Traefik GUI ``` http://localhost:8080 ```

Stop: ``` bash purge.sh``` 
