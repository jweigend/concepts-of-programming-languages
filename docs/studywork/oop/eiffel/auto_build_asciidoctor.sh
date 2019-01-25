#!/usr/bin/env bash

case "$1" in
    html) CMD="asciidoctor";;
    epub) CMD="asciidoctor-epub";;
    *) CMD="asciidoctor-pdf";;
esac

echo "Using cmd: $CMD"

while true; do
    inotifywait -e modify *.adoc
    $CMD 00_*.adoc
done
