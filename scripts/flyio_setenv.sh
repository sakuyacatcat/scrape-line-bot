#!/bin/bash

while read -r line || [ -n "$line" ]; do
    varname=$(echo $line | cut -d = -f 1)
    varvalue=$(echo $line | cut -d = -f 2-)
    echo "Setting $varname with value $varvalue"
    flyctl secrets set $varname=$varvalue -a scrape-line-bot
done <.env
