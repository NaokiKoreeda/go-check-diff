# go-check-diff
Simply, Check the difference between two directories.

# Usage
## example

Compile according to your env.  
For the env that can be specified, see below.
[https://golang.org/doc/install/source#environment](https://golang.org/doc/install/source#environment)

```console
$ cd go-check-diff
$ GOOS=darwin GOARCH=amd64 go build goCheckDiff.go
```

Create a shell to run.

```console
$ touch goCheckDiff.sh
```

Write as follows.  
Change the description of `baseDir` and `targetDir` to the directory you want to check.

```sh
#!/bin/sh

cd `dirname $0`

baseDir="/Users/NaokiKoreeda/dir1/"
targetDir="/Users/NaokiKoreeda/dir2/"

./goCheckDiff ${baseDir} ${targetDir}
```

Give permission to shell.

```console
$ chmod 755 goCheckDiff.sh
```

Execute.

```console
$ ./goCheckDiff.sh
```

The following result is output.

```
2018-07-12 12:24:44 Start.
File is different. -> /Users/NaokiKoreeda/dir2/test/test.txt
File not found. -> /Users/NaokiKoreeda/dir2/test2/test.txt
Directory not found. -> /Users/NaokiKoreeda/dir2/test3/
2018-07-12 12:24:44 End.
```
