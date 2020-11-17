.PHONY: all build run clean

all:
	make build

build:
	make hello_c
	make hello_go
	make args_c
	make args_go

hello_c: hello.c
	gcc -Wall -O2 -o hello_c hello.c

hello_go: hello.go
	go build -o hello_go hello.go

args_c: args.c
	gcc -Wall -O2 -o args_c args.c

args_go: args.go
	go build -o args_go args.go

run:
	./hello_c
	./hello_go
	./args_c
	./args_go

clean:
	-rm hello_c
	-rm hello_go
	-rm args_c
	-rm args_go
