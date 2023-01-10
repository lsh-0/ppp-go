#!/bin/bash
set -e

cmd="$1"

if test ! "$cmd"; then
    echo "command required."
    echo
    echo "available commands:"
    echo "  build-all"
    echo "  build"
    echo "  test"
    exit 1
fi

shift
rest=$*

if test "$cmd" = "build-all"; then
    # -o write to file
    # -v verbose
    # -a build all
    go build \
        -a \
        -v \
        -o ppp
    exit 0

elif test "$cmd" = "build"; then
    # -o write to file
    # -v verbose
    # -a build all
    go build \
        -v \
        -o ppp
    exit 0

elif test "$cmd" = "test"; then
    source_path="$1" # "./manage.sh test ./internal/utils"
    if [ -n "$source_path" ]; then
        go test -v "$source_path"
    else
        # -v verbose
        go test \
            -v \
            ./...
    fi
    exit 0

fi

echo "unknown command: $cmd"
exit 1
