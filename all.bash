#!/bin/sh

set -eux

for d in iup demos
do
	gomake -j 4 -C $d "$@"
done

for d in iup
do
	(cd $d && gotest)
done
