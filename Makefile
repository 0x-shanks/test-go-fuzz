.PHONY: atoi
atoi:
	go test -fuzz=FuzzMyAtoi -fuzztime=5m

.PHONY: fibo
fibo:
	go test -fuzz=FuzzFibo -fuzztime=5m