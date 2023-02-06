#!/bin/bash

wkdir=$(cd $(dirname $0); pwd)
cd $wkdir/..

go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/model/schema