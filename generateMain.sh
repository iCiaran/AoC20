#!/bin/bash

import_file=$(mktemp)
challenges_file=$(mktemp)
output_file=main.go

for day in $(ls -d day_* 2> /dev/null)
do
    i=${day##*_}
    j=$((10#$i))
    echo "challenges[${j}] = day_${i}.New()" >> ${challenges_file}
    echo "\"github.com/iCiaran/AoC20/day_${i}"\" >> ${import_file}
done

cat templates/main_01.template ${import_file} templates/main_02.template ${challenges_file} templates/main_03.template > $output_file
gofmt -w $output_file

rm -f ${import_file}
rm -f ${challenges_file}
