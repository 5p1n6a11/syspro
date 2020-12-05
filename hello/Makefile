.PHONY: all build run clean

CC = gcc
CFLAGS = -g -Wall -Wno-parentheses
TARGETS_c = hello_c args_c bell_c cat_c cat2_c
TARGETS_go = hello_go args_go bell_go
TARGETS = $(TARGETS_c) $(TARGETS_go)

all:
	make build

build:
	make hello_c
	make hello_go
	make args_c
	make args_go
	make bell_c
	make bell_go
	make cat_c
	make cat2_c

hello_c: hello.c
	gcc -Wall -O2 -o $@ $<

hello_go: hello.go
	go build -o $@ $<

args_c: args.c
	gcc -Wall -O2 -o $@ $<

args_go: args.go
	go build -o $@ $<

bell_c: bell.c
	gcc -Wall -O2 -o $@ $<

bell_go: bell.go
	go build -o $@ $<

cat_c: cat.c
	gcc -Wall -O2 -o $@ $<

cat2_c: cat2.c
	gcc -Wall -O2 -o $@ $<

run:
	./hello_c
	./hello_go
	./args_c
	./args_go
	./bell_c
	./bell_go
	./cat_c
	./cat2_c

clean:
	go clean
	-rm -f $(TARGETS)
