function ckdiff {
    cat | diff <(echo "$output") -
}

function filter {
    for line in "${lines[@]}"
    do echo $line | $*
    done
}