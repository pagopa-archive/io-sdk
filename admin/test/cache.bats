#!/usr/bin/env bats
load util

@test "cache action" {
    run wsk action invoke util/cache -r -p set pippo=pluto 
    ckline "set:pippo=pluto" 
    run wsk action invoke util/cache -r -p get pippo
    ckline '"pippo": "pluto"'
    run wsk action invoke util/cache -r -p del pippo
    ckline '"del:pippo": true'
    run wsk action invoke util/cache -r -p get pippo
    ckline '"error": "cannot find pippo"'
}
