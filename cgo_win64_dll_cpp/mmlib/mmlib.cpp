
#include <iostream>

using namespace std;


extern "C" {
    __declspec(dllexport) void PrintValue(int v) {
        cout << "Value: " << (v * 10000) << " !!" << endl;        
    }
}

