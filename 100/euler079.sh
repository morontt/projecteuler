#!/bin/bash

digit=0

while [ $digit -lt 10 ]; do
    count_digit=`cat ./p079_keylog.txt | grep $digit | wc -l`
    echo $digit" - "$count_digit
    digit=$(( $digit + 1 ))
done
