arauto:
	gcc -c -Wall -O2 edison.cpp
	ar rvs edison.a edison.o

clean:
	rm edison.a
	rm edison.o
