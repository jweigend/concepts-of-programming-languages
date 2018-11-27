# Exercise 8 - Distributed Programming in Go

If you do not finish during the lecture period, please finish it as homework.

### Excercise 7.0 - Warm Up with gRPC
The Go greeter example gets you started with gRPC in Go with a simple working example.
https://grpc.io/docs/quickstart/go.html
After a walk through the exercise you should know how to install and run the protoc compiler inclusive Go plugin. 
You will need this code for the next exercises.

## Exercise 7.1 - ID Generator (Local Implementation and Client)
Build an interface and a implementation for generating unique numbers. 
Build a test client which generates 10 IDs by accessing the interface.

## Exercise 7.2 - ID Generator (gRPC Protocol Definition)
Build a GRPC protocol definition and compile it with the Google protocol compiler protoc.

## Exercise 7.3 - Build a Proxy which encapsulates the gRPC client logic and a Stub which encapsulates the GRPC server logic.
Use the gRPC Greetings sample as starting point.

