#!/bin/bash

vars=$(env | grep ^KAFKA_)

for var in $vars; do
    key=${var%%=*}  # the left part of first symbol `=`
    value=${var#*=} # the right part of first symbol `=`
    key=${key#*_}   # the right part of first symbol `_`
    key=${key,,}    # to lowercase
    key=${key//_/.} # replace `_` to `.`

    sed -i '/^'$key'=/d' config/server.properties
    echo "$key=$value" >> config/server.properties
    echo "kafka reconfigured $key=$value"
done
