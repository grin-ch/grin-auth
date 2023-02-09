#!/bin/bash

servers=(
    "captcha"
    "account"
)
wkdir=$(cd $(dirname $0); pwd)
cd $wkdir/..

pid=$$
target=$1

if [ ! -n "$target" ] ;then
    echo "missing target"
    exit 1
fi

if [[ "${servers[@]}"  =~ "${target}" ]]; then
    echo "run.sh pid:${pid}"
    echo "server:${target} in serve"
    go run ./cmd/$target
elif [[ ! "${servers[@]}"  =~ "${target}" ]]; then
    echo "target $target not exists"
fi