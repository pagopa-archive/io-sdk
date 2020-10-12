#!/usr/bin/env bats
load util

@test "util/sample post" {
    ipost http://localhost:3280/api/v1/web/guest/util/sample
    ckline '"fiscal_code": "AAAAAA00A00A000A:0"'
    ckline '"markdown": "# Benvenuto,'
    ckline '"subject": "Benvenuto'
    ipost http://localhost:3280/api/v1/web/guest/util/sample fiscal_code="PLMFNZ48R20I480G" due_date="2020-12-12" amount=1
    ckline '"amount": 1'
    ckline '"due_date": "2020-12-12T00:00:00"'
    ckline '"fiscal_code": "PLMFNZ48R20I480G:0"'
    ipost http://localhost:3280/api/v1/web/guest/util/sample count=2
    ckline '"fiscal_code": "AAAAAA00A00A000A:1"'
}

@test "util/sample get" {
    get http://localhost:3280/api/v1/web/guest/util/sample
    ckline '"fiscal_code": "AAAAAA00A00A000A:0"'
    ckline '"markdown": "# Benvenuto,'
    ckline '"subject": "Benvenuto'
}

@test "util/sample get count" {
    get "http://localhost:3280/api/v1/web/guest/util/sample?count=2"
    ckline '"fiscal_code": "AAAAAA00A00A000A:1"'
}

@test "util/sample get extra" {
    run http --timeout=300 GET 'http://localhost:3280/api/v1/web/guest/util/sample?fiscal_code=PLMFNZ48R20I480G&due_date=2020-12-12&amount=1'
    ckline '"amount": 1'
    ckline '"due_date": "2020-12-12T00:00:00"'
    ckline '"fiscal_code": "PLMFNZ48R20I480G:0"'
}