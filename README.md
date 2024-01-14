# white-page
 
## 프로젝트 설정

### git ignore

https://www.toptal.com/developers/gitignore 에서 intellij-iml & go 설정 으로 생성 

### Debug configuration

* Run -> Edit Configurations -> + -> Go Build -> Configuration -> Environment 에 다음과 같이 입력
  * GOOS=windows
  * GOARCH=amd64
  * CGO_ENABLED=1

* gcc 컴파일러(C컴파일러) 설치
  * https://jmeubank.github.io/tdm-gcc/
  * 설치 후 재시작 필수

* 윈도우 WSL (PowerShell 관리자 권한으로 실행)
  * dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart
  * dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart
