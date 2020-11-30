#!/bin/bash

function usage {
    echo "usage ${1}"
    exit ${1}
}

[[ $# -eq 1 ]] || usage 1
[[ ${1} =~ ^[0-9]+$ ]] || usage 2
([[ ${1} -ge 0 ]] && [[ ${1} -le 25 ]]) || usage 3

padded=$(printf "%02d" ${1})

base="day_${padded}"
mkdir -p ${base}
mkdir -p ${base}/inputs
mkdir -p ${base}/cmd
mkdir -p ${base}/cmd/run

touch ${base}/inputs/real_a.txt
touch ${base}/inputs/real_b.txt

cp main.template "${base}/${base}.go"
cp test.template "${base}/${base}_test.go"
cp run.template "${base}/cmd/run/main.go"
cp run_sh.template "${base}/run.sh"

for f in $(find ${base} -type f)
do
    sed -i "s/##DAY##/${padded}/g" $f
done
