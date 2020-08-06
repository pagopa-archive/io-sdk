<#
.Synopsis
   Validate iosdk snapshot on windows
.DESCRIPTION
   Execute iosdk init, iosdk start and test installation
.EXAMPLE
   .\validate.ps1 0.5.1
.OUTPUTS
   At the end you should see iosdk status, showing services running
.NOTES
   Powershell is awesome
.FUNCTIONALITY
   Validation of the iosdk installation
#>

$version = "2020.0805.1627-snapshot"
if ($null -eq $args[0]) {
    Write-Output "usage: validate.ps1 <version>"
    exit
}
$version = $args[0]

$file = "iosdk_$version.exe"
$url = "https://github.com/pagopa/io-sdk/releases/download/$version/$file"
$outpath = "$env:TEMP/$file"

Invoke-WebRequest -Uri $url -OutFile $outpath
Start-Process -Filepath "$outpath" -ArgumentList @("/S") -Verb RunAs

$env:Path = [System.Environment]::GetEnvironmentVariable("Path","User")

if (! (Write-Output $env:Path | Select-String -Pattern iosdk)) {
    Write-Output "PATH not set properly"
    Write-Output "FAIL"
    exit
}

& "$env:ProgramFiles\IOSDK\iosdk.exe" "init" "$HOME/tmp-iosdk-validation" "pagopa/io-sdk-javascript" "--io-apikey=123456"
& "$env:ProgramFiles\IOSDK\iosdk.exe" "start" "--skip-open-browser"

& "$env:ProgramFiles\IOSDK\iosdk.exe" "status"

& "docker" "exec" "iosdk-theia" "/home/project/build.sh"

$dataFile = "$HOME/tmp-iosdk-validation/data/data.xlsx" 
$bytes = [System.IO.File]::ReadAllBytes($dataFile)
$encoded = [System.Convert]::ToBase64String($bytes)
$testJson = "{`"file`": `"'$encoded`"}"
$testUrl = "http://localhost:3280/api/v1/web/guest/iosdk/import"

$out = Invoke-RestMethod -Uri $testUrl -Method Post -ContentType "application/json" -Body $testJson 

if($out.data[0].fiscal_code -eq "ISPXNB32R82Y766F") {
   Write-Output "SUCCESS"
} else {
   Write-Output "FAIL"
}

Remove-Item -Recurse -Force "$HOME\tmp-iosdk-validation"

& Start-Process -Filepath "$env:ProgramFiles/IOSDK/uninstall.exe" -ArgumentList @("/S") -Verb RunAs -Wait
