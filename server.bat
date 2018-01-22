go install github.com/roscopecoltran/gcse/server
@if errorlevel 1 goto exit
%GOPATH%\bin\server

:exit
