
# GolangAWSPipeline playground
---

This project is a learning example. Its aim is to get an AWS pipeline for
Continual Integration testing using:

* Golang
* Docker Compose

The application itself is just a graphql server serving up static content.
The test integration spins up a docker compose, docerising both the integration test as well as the application. The exit code for the integration container will signal for the pipeline to fail or not.

---
### Lessons learnt

* There are steps in the lifecycle of a maintaining source code. I think Maven does this well, in Golang it feels like it is project by project
    1. Clone repo 
    2. Panic compliation errors
    3. Generate stubs / models
    4. Feel relife
    5. Carry on



