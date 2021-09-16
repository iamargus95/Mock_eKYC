# Mock_eKYC
One2N Consulting assignment 

# Build HTTP APIs for an eKYC use case

eKYC consists of the following functionality.

1. Signup a new api client to get access and secret keys.

2. Upload Images and Metadata.

3. Perform face matching between two image ids to get the face match score.

4. Perform Optical Character Recognition (OCR) on images.

5. Generate client-wise reports for billing purposes.


![overview](/images/assignment.png)

> This is the overall system design, for this exercise, you will primarily focus on the API layer. To start with, you may not need all components, the most basic ones are API layer and Database and Minio (to store images). RabbitMQ, Workers and Redis can be added in V2 implementation.


## Assumptions

1. We will not build the ML models for face match and OCR. Instead, for face match, return a random number between 0 and 100 as a score. For OCR, return a structured JSON via any [faker library](https://github.com/bxcodec/faker).

2. You will need to setup
    - Postgresql
    - Redis Cache
    - RabbitMQ Queue

3. To get familiarity with relational databases, try creating a sample db with some tables, Get familiarity with basics of SQL like
    - select
    - insert
    - update
    - delete
    - and where clauses.

> Complete the basics section [here](https://pgexercises.com/questions/basic/).

4. Get to know basic redis commands like
    - set
    - get
    - del
    - incr etc.

5. Understand basic terminology of RabbitMQ and how you can create queues, producers and consumers. Refer this [basic tutorial](https://www.rabbitmq.com/tutorials/tutorial-one-go.html) and then this [producer-consumer pattern tutorial](https://www.rabbitmq.com/tutorials/tutorial-two-go.html).


## Expections

1. Set up golang repo for http APIs. Learn and use any popular Golang WEB framework. (Example: [Gin](https://github.com/gin-gonic/gin) or [Beego](https://github.com/beego/beego)).

2. Use a migration tool (Eg. [migrate](https://github.com/golang-migrate/migrate)) for creating and managing DB migrations.

3. Write test cases for controllers (HTTP layer) using httptest.NewRecorder(). For basic examples, refer [testing section of gin docs](https://github.com/gin-gonic/gin#testing).

4. You are expected to come up with DB modelling and detailed API design. This includes designing of db tables for storing information of clients, api_requests, api_responses for ocr and facematch, image metadata, etc.

5. Use HMAC signature based authentication method. (Chinmay to add more details).

6. Always try to get something working first before making it better. Commit and push your code as you complete tasks. The git commit history should demonstrate the progress of the solution.
