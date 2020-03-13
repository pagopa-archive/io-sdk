!include LogicLib.nsh
; The name of the installer
Name "io-sdk"

; The file to write
OutFile "io-sdk.exe"

; Request application privileges for Windows Vista
RequestExecutionLevel admin

; Build Unicode installer
Unicode True

; The default installation directory
InstallDir $PROGRAMFILES64\io-sdk

; Registry key to check for directory (so if you install again, it will 
; overwrite the old one automatically)
InstallDirRegKey HKLM "Software\io-sdk" "Install_Dir"

;--------------------------------

; Pages

;Page components
Page directory
Page instfiles

UninstPage uninstConfirm
UninstPage instfiles

;--------------------------------

; The stuff to install
Section "io-sdk (required)"

  SectionIn RO
  
  ; Set output path to the installation directory.
  SetOutPath $INSTDIR
  
  ; Put file there
  File "iosdk_win.exe"
  
  ; Write the installation path into the registry
  WriteRegStr HKLM SOFTWARE\io-sdk "Install_Dir" "$INSTDIR"
  
  ; Write the uninstall keys for Windows
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\io-sdk" "DisplayName" "io-sdk"
  WriteRegStr HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\io-sdk" "UninstallString" '"$INSTDIR\uninstall.exe"'
  WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\io-sdk" "NoModify" 1
  WriteRegDWORD HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\io-sdk" "NoRepair" 1
  WriteUninstaller "$INSTDIR\uninstall.exe"

; Check if the path entry already exists and write result to $0
nsExec::Exec 'echo %PATH% | find "$INSTDIR"'
Pop $0   ; gets result code

${If} $0 = 0
    nsExec::Exec 'SETX PATH %PATH%;$INSTDIR'
${EndIf}

SectionEnd

; Optional section (can be disabled by the user)
Section "Start Menu Shortcuts"

  CreateDirectory "$SMPROGRAMS\io-sdk"
  CreateShortcut "$SMPROGRAMS\io-sdk\Uninstall.lnk" "$INSTDIR\uninstall.exe" "" "$INSTDIR\uninstall.exe" 0
  CreateShortcut "$SMPROGRAMS\io-sdk\io-sdk.lnk" "$INSTDIR\iosdk_win.exe" "" "$INSTDIR\io-sdk.nsi" 0
  
SectionEnd

;--------------------------------

; Uninstaller

Section "Uninstall"
  
  ; Remove registry keys
  DeleteRegKey HKLM "Software\Microsoft\Windows\CurrentVersion\Uninstall\io-sdk"
  DeleteRegKey HKLM SOFTWARE\io-sdk

  ; Remove files and uninstaller
  Delete $INSTDIR\io-sdk.nsi
  Delete $INSTDIR\uninstall.exe

  ; Remove shortcuts, if any
  Delete "$SMPROGRAMS\io-sdk\*.*"

  ; Remove directories used
  RMDir "$SMPROGRAMS\io-sdk"
  RMDir "$INSTDIR"

SectionEnd
