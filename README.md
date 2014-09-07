WorldWeather
============

This project has been created for gophercon india 2015 scholarship contest.


##Concurrency

Go supports lightweight threads/actors called goroutines which this project uses to limit the number of simultaneous
connections to the API server, we don't want to be bad consumers. :)


##Docs

using the godoc command to see the code documenation

    $ godoc -http=5050
