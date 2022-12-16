# e-shop-microservices

This project is a microservice project designed with event driven architecture.


## Features

The project consists of identity, catalog, order, basket and email services. 
The services communicate with each other via an event bus implemented with rabbitmq. 
The client makes rest calls to these services via an api-gateway. 
Services record their ip addresses to consul service discovery as soon as they run for the first time, and the api gateway gets the ip of these services from service discovery.
This process is not very costly because Consul has a cache mechanism in itself.


### Architecture

![event bus (2)](https://user-images.githubusercontent.com/64227421/208158347-a457a3f0-1bbd-4dd0-b064-ec76270f29fb.png)


|         Event       |     Sender      |        Sender Description     |          Receivers            | Receiver Description
| ------------------- | --------------- | ----------------------------- | ----------------------------- | --------------------------------
| UserCreatedEvent    | identityService | Sent when new user is created | emailService, basketService   | A verification code is sent to the email of the created user and a basket is created for the user
| BasketVerifiedEvent | basketService   | Sent when basket is verified  | orderService                  | An order is created with the products in the verified basket.
| OrderCompletedEvent | orderService    | Sent when order is complete   | catalogService, basketService | After the order is completed, the number of products is reduced from stock and the basket is emptied   


### Technologies
* Go
* Rabbit MQ
* MongoDB
* Mysql
* Redis
* Consul
* Docker

### Run

```
docker-compose up -d
```


