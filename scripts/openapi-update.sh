#!/bin/bash

MYDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOTDIR="$MYDIR/.."

main() {
    source_dir="$ROOTDIR/embed/msgraph-metadata/openapi"
    echo "source_dir: $source_dir"
    target_dir="$ROOTDIR/embed/openapi"
    echo "target_dir: $target_dir"
    echo "removing all exist type files..."
    rm -r $target_dir/v1.0/openapi.yaml
    rm -r $target_dir/beta/openapi.yaml
    echo "done"
    echo "copying new type files..."
    cp -r $source_dir/v1.0/openapi.yaml $target_dir/v1.0/openapi.yaml
    cp -r $source_dir/beta/openapi.yaml $target_dir/beta/openapi.yaml
    echo "done"
}

main "$@"