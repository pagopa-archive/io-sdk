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
  echo ">>> " http --timeout=300 GET "$@" >>/tmp/debug.log
  run http --ignore-stdin --timeout=300 GET "$@"
  echo "$output" >>/tmp/debug.log
}

function put { 
  echo ">>> " http --timeout=300 PUT "$@" >>/tmp/debug.log
  run http --timeout=300 PUT "$@"
  echo "$output" >>/tmp/debug.log
}

function iput { 
  echo ">>> " http --ignore-stdin --timeout=300 PUT "$@" >>/tmp/debug.log
  run http --ignore-stdin --timeout=300 PUT "$@"
  echo "$output" >>/tmp/debug.log
}


function post { 
  echo ">>> " http --timeout=300 POST "$@" >>/tmp/debug.log
  run http --timeout=300 POST "$@"
  echo "$output" >>/tmp/debug.log
}

function ipost { 
  echo ">>> " http --ignore-stdin --timeout=300 POST "$@" >>/tmp/debug.log
  run http --ignore-stdin --timeout=300 POST "$@"
  echo "$output" >>/tmp/debug.log
}


function fpost { 
  echo ">>> " http -f --ignore-stdin --timeout=300 POST "$@" >>/tmp/debug.log
  run http -f --ignore-stdin --timeout=300 POST "$@"
  echo "$output" >>/tmp/debug.log
}
