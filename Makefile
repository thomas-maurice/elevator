all:
	go get
	go test ./scheduler
	/bin/sh -c "cd scheduler && go fmt"
	/bin/sh -c "cd algorithms && go fmt"
	/bin/sh -c " go fmt"
	go build -o sched
