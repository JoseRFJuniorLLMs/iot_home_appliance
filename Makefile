CFLAGS += `pkg-config --cflags upm-i2clcd`
LDFLAGS += `pkg-config --libs upm-i2clcd`

arauto:
	gcc -c -Wall -O2 $(CFLAGS) $(LDFLAGS) edison.cpp
	ar rvs edison.a edison.o
	go build arauto.go

clean:
	rm -f edison.a edison.o
