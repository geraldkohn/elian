build:
	go build -o bin/httpServer cmd/run/httpAndClient/main.go 
	go build -o bin/agencyServer cmd/run/agencyServer/main.go
	go build -o bin/patientServer  cmd/run/patientServer/main.go
	go build -o bin/recordServer cmd/run/recordServer/main.go
	go build -o bin/staffServer  cmd/run/staffServer/main.go

clean:	
	rm -rf bin
 