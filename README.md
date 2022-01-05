# clean 
❯ make clean
  >  Cleaning build cache

# compile
❯ make compile
  >  Cleaning build cache
  >  Building binary...
  >  /Library/Developer/CommandLineTools/usr/bin/make conv_windows_amd64.exe
  >  /Library/Developer/CommandLineTools/usr/bin/make conv_linux_amd64
  >  /Library/Developer/CommandLineTools/usr/bin/make conv_darwin_amd64
  >  Building binary end

# windows
V:\gpm\gpm_demo\pp\conv>conv_windows_amd64.exe -h
platform:windows+amd64
versionString:
commitString:
Usage of conv_windows_amd64.exe:
  -e, --exclude string   exclude file (default "CMSIS")
  -t, --type string      toUTF8, dos2unix, all (default "dos2unix")
  -v, --verbose string   verbose message (default "false")
pflag: help requested