all: fuzz

fuzz: fuzz.cc
	clang++ -fsanitize=address,fuzzer fuzz.cc

run: a.out
	./a.out -seed=3918206239

clean:
	rm -rf *.out
	find corpus/ -type f -not -name 'input*' -delete
