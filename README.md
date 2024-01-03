# git hook
git config core.hooksPath .githooks

# clean
```
❯ make clean
  >  Cleaning build cache
```

# compile
```
❯ make compile
  >  Cleaning build cache
  >  Building binary...
  >  /Library/Developer/CommandLineTools/usr/bin/make inner_tool_windows_amd64.exe
  >  /Library/Developer/CommandLineTools/usr/bin/make inner_tool_linux_amd64
  >  /Library/Developer/CommandLineTools/usr/bin/make inner_tool_darwin_amd64
  >  Building binary end
```

# execute
```
  ./inner_tool_windows_amd64.exe .c .h
  ./inner_tool_linux_amd64 .c .h
  ./inner_tool_darwin_amd64 .c .h
```
