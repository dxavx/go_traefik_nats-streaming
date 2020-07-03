## Go API + Traefik (Load Balancer)

![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/pdacity/docker_gc) 

There is an example of using the “Traefik " load balancer and Go code with the ability to increase the number of instances of the container with the Go code (scaling ).  

In the file __build.sh__ you can set the number of copies of the app: `docker-compose -f./docker-compose.yml up -d -- scale api=3`

Start: ``` bash build.sh ```

API  ``` http://localhost/v1/ping ```

Traefik GUI ``` http://localhost:8080 ```

Stop: ``` bash purge.sh``` 

*** 
## Go API + Traefik (Load Balancer)

![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/pdacity/docker_gc)

Пример использования балансировщика “Traefik“ и кода на Go с возможностью увеличивать количество экземпляров контейнера с кодом Go ( скейлинг ).  

В файле __build.sh__ вы можете указать количество копий приложения : `docker-compose -f./docker-compose.yml up -d -- scale api=3`

Start: ``` bash build.sh ```

API  ``` http://localhost/v1/ping ```

Traefik GUI ``` http://localhost:8080 ```

Stop: ``` bash purge.sh``` 
