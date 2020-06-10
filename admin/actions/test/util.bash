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
  echo ">>> " python3 -m httpie --timeout=300 GET "$@" >>/tmp/debug.log
  run python3 -m httpie --ignore-stdin --timeout=300 GET "$@"
  echo "$output" >>/tmp/debug.log
}

function put { 
  echo ">>> " python3 -m httpie --timeout=300 PUT "$@" >>/tmp/debug.log
  run python3 -m httpie --ignore-stdin --timeout=300 PUT "$@"
  echo "$output" >>/tmp/debug.log
}

function post { 
  echo ">>> " python3 -m httpie --ignore-stdin --timeout=300 POST "$@" >>/tmp/debug.log
  run python3 -m httpie --timeout=300 POST "$@"
  echo "$output" >>/tmp/debug.log
}

function fpost { 
  echo ">>> " python3 -m httpie -f --ignore-stdin --timeout=300 POST "$@" >>/tmp/debug.log
  run python3 -m httpie -f --timeout=300 POST "$@"
  echo "$output" >>/tmp/debug.log
}
