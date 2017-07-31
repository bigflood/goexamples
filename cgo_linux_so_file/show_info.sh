#!/bin/bash



echo "======= ======="
file mmlib.so
ldd mmlib.so
nm mmlib.so | grep PrintValue

echo "======= ======="
file run_example
LD_LIBRARY_PATH=. ldd run_example
nm run_example | grep PrintValue
