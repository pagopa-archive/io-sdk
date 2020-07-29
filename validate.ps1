<#
.Synopsis
   Validate iosdk snapshot on windows
.DESCRIPTION
   Execute iosdk init, iosdk start and test installation
.EXAMPLE
   .\validate.ps1
.EXAMPLE
   .\validate.ps1 0.5.1
.OUTPUTS
   At the end you should see iosdk status, showing services running
.NOTES
   Powershell is awesome
.FUNCTIONALITY
   Validation of the iosdk installation
#>

if ($args[0] -eq $null) {
    $version = "0.5.1"
} else {
    $version = $args[0]
}

$file = "iosdk_$version.exe"

$url = "https://github.com/pagopa/io-sdk/releases/download/$version/$file"

Write-Output $url

$outpath = "$env:TEMP/$file"

Invoke-WebRequest -Uri $url -OutFile $outpath

$args = @("/S")
Start-Process -Filepath "$outpath" -ArgumentList $args -Verb RunAs

#Start-Sleep -s 5

$env:Path = [System.Environment]::GetEnvironmentVariable("Path","User")

if (Write-Output $env:Path | Select-String -Pattern iosdk) {
    Write-Output "PATH OK"
}

#Start-Process -NoNewWindow -Wait -Filepath "$env:ProgramFiles\IOSDK\iosdk.exe" -ArgumentList "init, $HOME/tmp-iosdk-validation, pagopa/io-sdk-javascript, --io-apikey=123456" 
& "$env:ProgramFiles\IOSDK\iosdk.exe" "init" "$HOME/tmp-iosdk-validation" "pagopa/io-sdk-javascript" "--io-apikey=123456"

#Start-Process -NoNewWindow -Wait -Filepath "$env:ProgramFiles\IOSDK\iosdk.exe" -ArgumentList "start"
& "$env:ProgramFiles\IOSDK\iosdk.exe" "start"

#Start-Process -NoNewWindow -Wait -Filepath "$env:ProgramFiles\IOSDK\iosdk.exe" -ArgumentList "status"
& "$env:ProgramFiles\IOSDK\iosdk.exe" "status"

#Start-Process -NoNewWindow -Wait -Filepath "$env:ProgramFiles\Docker\Docker\Resources\bin\docker.exe" -ArgumentList "exec, iosdk-theia, /home/project/build.sh"
& "docker" "exec" "iosdk-theia" "/home/project/build.sh"

#Start-Process -NoNewWindow -Wait -Filepath "$env:ProgramFiles\IOSDK\iosdk.exe" -ArgumentList "status"
& "$env:ProgramFiles\IOSDK\iosdk.exe" "status"

#Start-Process -NoNewWindow -Wait -Filepath "$env:ProgramFiles\IOSDK\iosdk.exe" -ArgumentList "stop"
& "$env:ProgramFiles\IOSDK\iosdk.exe" "stop"

Remove-Item -Recurse -Force "$HOME\tmp-iosdk-validation"

Start-Process -Filepath "$env:ProgramFiles/IOSDK/uninstall.exe" -ArgumentList $args -Verb RunAs -Wait

