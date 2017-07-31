CPPFLAGS= -Wall -g -O0

SOURCES = $(shell echo *.cpp)
HEADERS = $(shell echo *.h)
OBJECTS=$(SOURCES:.cpp=.o)

TARGET=mmlib.a

all: $(TARGET)

clean:
	rm -f $(OBJECTS) $(TARGET)

$(TARGET) : $(OBJECTS)
	ar rcs $@ $^
