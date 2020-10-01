# Atempt to implement tree tier architecture or DDD or hexagonal  or ports and adapters.


## About The Project
- This is just a way to learn Golang. concepts like interface, pointers and Clear Archtecture will be used here.

- At this time we can decouple the domain from the repository/data access, ie we can use 2 types of repositories one for postgres and another for mongo, for example.

- Our application still has a strong coupled to GIN. So, our http adapter needs to be a GIN.

- We need to define a Controller Interface to decouple the presentation/adapter/web/http

- all the above requirements are need just to prove the concept.

- But, but, but we will do this in the future! For now, we just make sure that the decoupling of domain and data is working. [TODO refactory the web adapter way]

Developed by -> robsonmrsp@gmail.com
