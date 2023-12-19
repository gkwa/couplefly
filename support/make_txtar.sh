#!/usr/bin/env bash

set -e

tmp=$(mktemp -d couplefly.XXXXX)

if [ -z "${tmp+x}" ] || [ -z "$tmp" ]; then
    echo "Error: \$tmp is not set or is an empty string."
    exit 1
fi

{
    rg --files . \
        | grep -v $tmp/filelist.txt \
        | grep -vE 'couplefly$' \
        | grep -v README.org \
        | grep -v make_txtar.sh \
        | grep -v go.sum \
        | grep -v go.mod \
        | grep -v Makefile \
        | grep -v cmd/main.go \
        | grep -v logger.go \
        # | grep -v couplefly.go \

} | tee $tmp/filelist.txt
tar -cf $tmp/couplefly.tar -T $tmp/filelist.txt
mkdir -p $tmp/couplefly
tar xf $tmp/couplefly.tar -C $tmp/couplefly
rg --files $tmp/couplefly
txtar-c $tmp/couplefly | pbcopy

rm -rf $tmp
