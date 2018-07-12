#!/bin/sh

cd `dirname $0`

# baseDir="/Users/NaokiKoreeda/dir1/"
# targetDir="/Users/NaokiKoreeda/dir2/"
baseDir="/Users/gurubimakku/work/dir1/"
targetDir="/Users/gurubimakku/work/dir2/"

./goCheckDiff ${baseDir} ${targetDir}
