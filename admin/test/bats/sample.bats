#!/usr/bin/env bats
load util

@test "util/sample post" {
    ipost "$SAMPLE"
    ckline '"fiscal_code": "AAAAAA00A00A000A:1"'
    ckline '"markdown": "# Benvenuto,'
    ckline '"subject": "Benvenuto'
    ipost "$SAMPLE" fiscal_code="PLMFNZ48R20I480G" due_date="2020-12-12" amount=1
    ckline '"amount": 1'
    ckline '"due_date": "2020-12-12T00:00:00"'
    ckline '"fiscal_code": "PLMFNZ48R20I480G:1"'
    ipost "$SAMPLE" count=2
    ckline '"fiscal_code": "AAAAAA00A00A000A:2"'
}

@test "util/sample get" {
    get "$SAMPLE"
    ckline '"fiscal_code": "AAAAAA00A00A000A:1"'
    ckline '"markdown": "# Benvenuto,'
    ckline '"subject": "Benvenuto'
}

@test "util/sample get count" {
    get "$SAMPLE?count=2"
    ckline '"fiscal_code": "AAAAAA00A00A000A:2"'
}

@test "util/sample get extra" {
    get "$SAMPLE"'?fiscal_code=PLMFNZ48R20I480G&due_date=2020-12-12&amount=1'
    ckline '"amount": 1'
    ckline '"due_date": "2020-12-12T00:00:00"'
    ckline '"fiscal_code": "PLMFNZ48R20I480G:1"'
}