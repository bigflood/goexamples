
copy .\mmlib\Release\mmlib.dll .
 
dlltool -dllname mmlib.dll --def .\mmlib\mmlib.def --output-lib mmlib.a

go build -o example_run.exe github.com/bigflood/goexamples/cgo_win64_dll_file
