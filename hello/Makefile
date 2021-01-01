.PHONY: all build run clean

CC = gcc
CFLAGS = -g -Wall -Wno-parentheses
TARGETS_c = hello_c args_c bell_c \
			cat_c cat2_c \
			head_c head2_c head4_c \
			grep_c \
			ls_c \
			mkdir_c rmdir_c \
			ln_c symlink_c rm_c \
			mv_c stat_c chmod_c \

TARGETS_go = hello_go args_go bell_go
TARGETS = $(TARGETS_c) $(TARGETS_go)

hello_c: hello.c
	$(CC) -Wall -o $@ $<

hello_go: hello.go
	go build -o $@ $<

args_c: args.c
	$(CC) -Wall -o $@ $<

args_go: args.go
	go build -o $@ $<

bell_c: bell.c
	$(CC) -Wall -o $@ $<

bell_go: bell.go
	go build -o $@ $<

cat_c: cat.c
	$(CC) -Wall -o $@ $<

cat2_c: cat2.c
	$(CC) -Wall -o $@ $<

head_c: head.c
	$(CC) -Wall -o $@ $<

head2_c: head2.c
	$(CC) -Wall -o $@ $<

head4_c: head4.c
	$(CC) -Wall -o $@ $<

grep_c: grep.c
	$(CC) -Wall -o $@ $<

ls_c: ls.c
	$(CC) -Wall -o $@ $<

mkdir_c: mkdir.c
	$(CC) -Wall -o $@ $<

rmdir_c: rmdir.c
	$(CC) -Wall -o $@ $<

ln_c: ln.c
	$(CC) -Wall -o $@ $<

symlink_c: symlink.c
	$(CC) -Wall -o $@ $<

rm_c: rm.c
	$(CC) -Wall -o $@ $<

mv_c: mv.c
	$(CC) -Wall -o $@ $<

stat_c: stat.c
	$(CC) -Wall -o $@ $<

chmod_c: chmod.c
	$(CC) -Wall -o $@ $<

clean:
	go clean
	-rm -f $(TARGETS_c)
