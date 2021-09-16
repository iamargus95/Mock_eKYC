# Mock_eKYC
One2N Consulting assignment 

# Build HTTP APIs for an eKYC use case

eKYC consists of the following functionality.

1. Signup a new api client to get access and secret keys.
2. Upload Images and Metadata.
3. Perform face matching between two image ids to get the face match score.
4. Perform Optical Character Recognition (OCR) on images.
5. Generate client-wise reports for billing purposes.

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