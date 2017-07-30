CPPFLAGS= -fPIC -g
LDFLAGS= -shared

SOURCES = $(shell echo *.cpp)
HEADERS = $(shell echo *.h)
OBJECTS=$(SOURCES:.cpp=.o)

TARGET=mmlib.so

all: $(TARGET)

clean:
	rm -f $(OBJECTS) $(TARGET)
	rm -f ../$(TARGET)

$(TARGET) : $(OBJECTS)
	$(CC) $(CFLAGS) $(OBJECTS) -o $@ $(LDFLAGS)
	cp $(TARGET) ..

