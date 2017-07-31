
extern "C" {
    __declspec(dllexport) int GetValue(int v) {
        return v * 1000 + 1;
    }
}

