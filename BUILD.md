# Build steps

1. Prepare Binaries for a Release
```
go install -ldflags "-X main.version=XXX" github.com/mmcquillan/protocli
cd $GOBIN
tar -cvzf protocli-XXX.tar.gz ./protocli
shasum -a 256 protocli-XXX.tar.gz
```


2. Create a github release.

https://github.com/mmcquillan/protocli/releases


3. Update Homebrew Forumla if needed.

https://github.com/mmcquillan/homebrew-tools/blob/master/protocli.rb
