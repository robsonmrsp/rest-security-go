# Atempt to implement tree tier architecture.


## About The Project


- At this time we can could decouple the domain from the repository/data access, ie we can use 2 types of repositories one to postgres and another  for mongo, for example.
- Owr application still had a strong coupled to GIN. So, our http adapter need to be a GIN.
- we need to define a ControllerInterface to decouple the presentation/adapter/web/http
- all the above requirements are need just to proove the concept.

- but, but, but we will do this in the future! For now, we just make sure that decouple of domain and data. [TODO refactory the web adapter way]

Developed by -> robsonmrsp@gmail.com
