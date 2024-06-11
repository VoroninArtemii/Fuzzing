all: fuzz

fuzz: fuzz.cc
	clang++ -fsanitize=address,fuzzer fuzz.cc

corpus1: a.out
	./a.out -seed=3918206239 corpus1/
	
corpus2: a.out
	./a.out -seed=3918206239 corpus2/

corpus3: a.out
	./a.out -seed=3918206239 corpus3/

corpus4: a.out
	./a.out -seed=3918206239 corpus4/

corpus5: a.out
	./a.out -seed=3918206239 corpus5/

clean:
	rm -rf *.out
	find corpus*/ -type f -not -name 'input*' -delete
