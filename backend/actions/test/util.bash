function ckdiff {
    cat | diff <(echo "$output") -
}

function ckline {
    for line in "${lines[@]}"
    do if echo $line | grep "$@"
       then return 0
       fi
    done
    return 1
}

function filter {
    for line in "${lines[@]}"
    do echo $line | $*
    done
}

function setup {
    export H=$BATS_TEST_DIRNAME
    export AUTH=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP
    wsk property set --apihost http://localhost:3280 --auth $AUTH >/dev/null
    export URL=$(wsk action get util/echo --url | tail +2 | sed -e 's!/util/echo!!')
    [ -n "$URL" ] || exit 1
}