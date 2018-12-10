IWP
===

IWP is a search engine software for immoweb that is extensible to other websites. It is designed for professional usage. 

## Build

```
# Linux
$ go build -o webview-example && ./webview-example

# MacOS uses app bundles for GUI apps
$ mkdir -p example.app/Contents/MacOS
$ go build -o example.app/Contents/MacOS/example
$ open example.app # Or click on the app in Finder

# Windows requires special linker flags for GUI apps.
# It's also recommended to use TDM-GCC-64 compiler for CGo.
# http://tdm-gcc.tdragon.net/download
$ go build -ldflags="-H windowsgui" -o webview-example.exe
```

## Project state
I do not actively develop this project anymore, it is currently a working application designed for one specific users group.  
It does not properly implement error handling and certain edge cases due time constraints.  
While I do not develop this project anymore I am happy to review and accapt pull requests.
