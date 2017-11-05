package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	os.Setenv("Foo", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR", os.Getenv("BAR"))
	os.Unsetenv("FOO")

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

/*

// windows 下测试
$ go run Basic/environment-variables.go
FOO: 1
BAR

ACLOCAL_PATH
ALLUSERSPROFILE
ANDROID_HOME
APPDATA
CLASSPATH
COMMONPROGRAMFILES
COMPUTERNAME
COMSPEC
CONFIG_SITE
CommonProgramFiles(x86)
CommonProgramW6432
DISPLAY
EXEPATH
GIT_HOME
GIT_LFS_PATH
GOARCH
GOBIN
GOOS
GOPATH
GOROOT
Gradle
HOME
HOMEDRIVE
HOMEPATH
HOSTNAME
INFOPATH
JAVA_HOME
LANG
LOCALAPPDATA
LOGONSERVER
MANPATH
MEmuHyperv_Path
MEmu_Path
MINGW_CHOST
MINGW_PACKAGE_PREFIX
MINGW_PREFIX
MOZ_PLUGIN_PATH
MSYSTEM
MSYSTEM_CARCH
MSYSTEM_CHOST
MSYSTEM_PREFIX
MYSQL_HOME
NPM_HOME
NUMBER_OF_PROCESSORS
NVM_HOME
NVM_SYMLINK
OLDPWD
ORIGINAL_PATH
ORIGINAL_TEMP
ORIGINAL_TMP
OS
OneDrive
PATH
PATHEXT
PKG_CONFIG_PATH
PLINK_PROTOCOL
PRINTER
PROCESSOR_ARCHITECTURE
PROCESSOR_IDENTIFIER
PROCESSOR_LEVEL
PROCESSOR_REVISION
PROGRAMFILES
PS1
PSModulePath
PUBLIC
PWD
ProgramData
ProgramFiles(x86)
ProgramW6432
SHELL
SHLVL
SSH_ASKPASS
SYSTEMDRIVE
SYSTEMROOT
TEMP
TERM
TMP
TMPDIR
USERDOMAIN
USERDOMAIN_ROAMINGPROFILE
USERNAME
USERPROFILE
VS110COMNTOOLS
VS120COMNTOOLS
VS140COMNTOOLS
WINDIR
_
asl.log

*/
