SHELL = /bin/bash

# expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,panics,mem:memstats.Alloc"

tidy:
	go mod tidy
	go mod vendor

run:
	go run app/services/sales-api/main.go