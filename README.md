## Go API + Traefik (Load Balancer) + NATS-Streaming

![Docker Compose Actions Workflow](https://github.com/dxavx/go_traefik_nats-streaming/workflows/Docker%20Compose%20Actions%20Workflow/badge.svg?branch=master) 
[![Go Report Card](https://goreportcard.com/badge/github.com/dxavx/go_traefik_nats-streaming)](https://goreportcard.com/report/github.com/dxavx/go_traefik_nats-streaming)


There is an example of using the “Traefik"  load balancer,  message broker "NATS-Streaming" and Go code with the ability to increase the number of instances of the container with the Go code (scaling ).  

In the file __build.sh__ you can set the number of copies of the app: `docker-compose -f./docker-compose.yml up -d -- scale subscriber=3`

Start: ``` bash build.sh ```

Publishing a random string ``` http://localhost/v1/pub ```

You can view the message in the subscriber container console

Traefik GUI ``` http://localhost:8080 ```

Stop: ``` bash purge.sh``` 

*** 
## Go API + Traefik (Load Balancer) + NATS-Streaming


Пример использования балансировщика “Traefik“, брокера сообщений "NATS-Streaming" и кода на Go с возможностью увеличивать количество экземпляров контейнера с кодом Go ( скейлинг ).  

В файле __build.sh__ вы можете указать количество копий приложения: `docker-compose -f./docker-compose.yml up -d -- scale subscriber=3`

Start: ``` bash build.sh ```

Публикация сучайно строки ``` http://localhost/v1/pub ```

Сообщение можно просмотреть в консоле контейнеров subscriber

Traefik GUI ``` http://localhost:8080 ```

Stop: ``` bash purge.sh``` 