#!/usr/bin/env bash

VERSION=1.0.0


mkdir -p dist



function buildAndDist {
    export GOOS=$1
    export GOARCH=$2
    mkdir -p target/$1-$2
    cp README.md LICENSE.md target/$1-$2
    go build -o target/$1-$2/pdfmerge
    (cd target/$1-$2; zip pdfmerge-$1-$2-${VERSION}.zip *)
    mv target/$1-$2/pdfmerge-$1-$2-${VERSION}.zip dist
}

buildAndDist linux amd64
buildAndDist windows amd64
buildAndDist darwin amd64

