#!/bin/sh

cd `dirname $0`

baseDir="/Users/NaokiKoreeda/dir1/"
targetDir="/Users/NaokiKoreeda/dir2/"

./goCheckDiff ${baseDir} ${targetDir}
