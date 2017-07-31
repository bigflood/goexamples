
#include <iostream>

using namespace std;

extern "C" {
    void PrintValue(int v) {
        cout << "PrintValue: " << v << " ~ " << endl;
    }
}

