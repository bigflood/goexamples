CPPFLAGS= -Wall -fPIC
LDFLAGS= -shared -lstdc++

SOURCES = $(shell echo *.cpp)
HEADERS = $(shell echo *.h)
OBJECTS=$(SOURCES:.cpp=.o)

TARGET=mmlib.so

all: $(TARGET)

clean:
	rm -f $(OBJECTS) $(TARGET)

$(TARGET) : $(OBJECTS)
	$(CC) $(CFLAGS) $(OBJECTS) -o $@ $(LDFLAGS)

