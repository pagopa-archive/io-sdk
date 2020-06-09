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
    do echo $line | "$@"
    done
}

function setup {
    export H=${BATS_TEST_DIRNAME:=$PWD}
    source $HOME/.wskprops
    export URL=$(wsk action get util/echo --url | tail +2 | sed -e 's!/util/echo!!')
    [ -n "$URL" ] || exit 1
}

function get { 
  run http --timeout=300 GET "$@"
}

function put { 
  run http --timeout=300 PUT "$@"
}

function post { 
  run http --timeout=300 POST "$@"
}

function fpost { 
  run http -f --timeout=300 POST "$@"
}

