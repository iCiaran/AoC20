#!/bin/bash

function usage {
    echo "usage: addDay.sh day_number"
    case ${1} in
        1)
            echo -e "\e[31mday_number is mandatory\e[0m"
            ;;
        2)
            echo -e "\e[31mday_number must be an integer\e[0m"
            ;;
        3)
            echo -e "\e[31mday_number must be in the range [1,25]\e[0m"
            ;;
    esac
    exit ${1}
}

function check_input {
    [[ $# -eq 1 ]] || usage 1
    [[ ${1} =~ ^[0-9]+$ ]] || usage 2
    ([[ ${1} -gt 0 ]] && [[ ${1} -le 25 ]]) || usage 3
}

function create_directories {
    base=${1}
    mkdir -p ${base}
    mkdir -p ${base}/inputs
    mkdir -p ${base}/cmd
    mkdir -p ${base}/cmd/run
}

function create_files {
    base=${1}
    padded=${2}

    touch ${base}/inputs/real_a.txt
    touch ${base}/inputs/real_b.txt
    echo "# Day ${padded}" > ${base}/README.md
}

function copy_templates {
    base=${1}
    cp templates/day.template "${base}/${base}.go"
    cp templates/test.template "${base}/${base}_test.go"
    cp templates/run.template "${base}/cmd/run/main.go"
    cp templates/run_sh.template "${base}/run.sh"
}

function replace_templates {
    base=${1}
    padded=${2}
    for f in $(find ${base} -type f)
    do
        sed -i "s/##DAY##/${padded}/g" $f
    done
}


# Main
check_input "$@"

padded=$(printf "%02d" ${1})
base="day_${padded}"

create_directories ${base}
create_files ${base} ${padded}
copy_templates ${base}
replace_templates ${base} ${padded}
