Gusset
======

There are two solutions:

* Simple - just logs out
* Dynamo - write to a dynamo db

Compile the app using a command like:

`
env GOOS=linux GOARCH=amd64 go build -o ../bin/simple main.go
`

Create a zip using:

`
zip -j  ../bin/simple.zip ../bin/simple
`