#!/bin/bash

for file in example0{0..8}.txt; do
    echo "$file"
    go run . "./examples/$file"
    echo ""
    echo ""
    echo ""
    echo ""
done
