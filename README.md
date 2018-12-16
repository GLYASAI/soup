# soup

install pkgconfig

install instant client

go get -u golang.org/x/sync

```
dyld: Library not loaded: @rpath/libclntsh.dylib.12.1
  Referenced from: /private/var/folders/9l/4ggdpsmx7cx17hvmfj0j6dkh0000gn/T/___go_build_main_go
  Reason: image not found
```

@cjbj :
```
Since recent versions of macOS ignore DYLD_LIBRARY_PATH in subshells, and therefore it isn't useful or reliable, you may just want to install Instant Client following the instructions and put it in ~/lib or /usr/local/lib
```