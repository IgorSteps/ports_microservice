# Golang microservices assignment

Your code structure should follow microservices best practices and our evaluation will focus primarily on your ability to follow good design principles and less on correctness and completeness of algorithms. During the face to face interview you will have the opportunity to explain your design choices and provide justifications for the parts that you omitted.

Avoid using frameworks such as go-kit and go-micro since one of the purposes of the assignment is to evaluate the candidate ability of structuring the solution in their own way.
If you have questions about the test, please draw your own conclusions.

## Technical test

- Given a file with ports data (ports.json), write 2 services
- The first service (Client API) should parse the JSON file and have REST interface
- While reading the file, it should call a second service (PortDomainService), that either creates a new record in a database, or updates the existing one
- The end result should be a database containing the ports, representing the latest version found in the JSON. Database can be Map in memory
- The first service (Client API) should provide an endpoint to retrieve the data from the second service (PortDomainService)
- Provide all tests that you think are needed for your assignment. This will allow the reviewer to evaluate your critical thinking as well as your knowledge about testing
- Use gRPC as a transport between services
- The readme should explain how to run your program and test it

Choose the approach that you think is best (i.e. most flexible).

## Note

We are looking for the ClientAPI (the service reading the JSON) to be written in a way that is easy to reuse, give or take a few customisations.
