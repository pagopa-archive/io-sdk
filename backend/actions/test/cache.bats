#!/usr/bin/env bats
load util

@test "cache cli" {
    run python3 deploy/util/cache.py set:pippo=pluto
    ckline "set:pippo=pluto"
    run python3 deploy/util/cache.py "scan:*"
    ckline '"pippo"'
    run python3 deploy/util/cache.py get:pippo
    ckline '"pippo": "pluto"'
    run python3 deploy/util/cache.py del:pippo
    ckline '"del:pippo": true'
    run python3 deploy/util/cache.py get:pippo
    ckline '"error": "cannot find pippo"'
}

@test "cache action" {
    run  wsk action invoke util/cache -r -p set pippo=pluto 
    ckline "set:pippo=pluto" 
    run wsk action invoke util/cache -r -p get pippo
    ckline '"pippo": "pluto"'
    run wsk action invoke util/cache -r -p del pippo
    ckline '"del:pippo": true'
    run wsk action invoke util/cache -r -p get pippo
    ckline '"error": "cannot find pippo"'
}
