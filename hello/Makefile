.PHONY: all build run clean

CC = gcc
CFLAGS = -g -Wall -Wno-parentheses
TARGETS_c = hello_c args_c bell_c cat_c cat2_c head_c head2_c
TARGETS_go = hello_go args_go bell_go
TARGETS = $(TARGETS_c) $(TARGETS_go)

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

head_c: head.c
	gcc -Wall -O2 -o $@ $<

head2_c: head2.c
	gcc -Wall -O2 -o $@ $<

clean:
	go clean
	-rm -f $(TARGETS)
