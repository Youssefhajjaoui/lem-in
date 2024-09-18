#!/bin/bash

for file in example0{0..9}.txt; do
    echo "$file"
    go run . "./examples/$file"
done
