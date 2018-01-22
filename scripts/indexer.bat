go install github.com/roscopecoltran/gcse/indexer
@if errorlevel 1 goto exit
%GOPATH%\bin\indexer

:exit
