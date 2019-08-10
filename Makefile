.PHONY: clean

TARGET=mythes-c

$(TARGET): libmythes.a
	go build .

libmythes.a: mythes.o
	ar r $@ $^

%.o: %.cpp
	g++ -E -O2 -o $@ -c $^

clean:
	rm -f *.o *.so *.a $(TARGET)