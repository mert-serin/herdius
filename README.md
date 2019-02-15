# Herdius Assignment

In this project I tried to create a bidirectional stream with Golang and GRPC.

## Proto template

Our server.proto service definition has only one function named, CheckMax. Since we only want to stream integer between server and client, our MaxRequest and MaxResponse has only variable for value. We could have used same response but to make our code readable, I defined two different message

## Server

Our server has got certificated that I created from my terminal, it'll be registered with these credentials. CheckMax function will be have the all logic. It'll read the value that we received from client, compare with the previous value, if it's bigger than the previous value, it'll response this new value to clients.

## Client

Our client has got public key, it'll signed its request with this public key. Client has 3 different goroutines:

First one will listen for users input. It'll convert input to integer and stream this value to server.

Second goroutine will listen server's message. It'll log the new max value when it's changed on server.

Third goroutine will listen done message. If we get done message from server, client will receive done message

TO-DO

## Testing

I've never written testing scenarios for grpc, so it's going to be the hardest part of this assingment.