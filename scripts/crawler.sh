#!/bin/sh

go install github.com/roscopecoltran/gcse/crawler
@if errorlevel 1 goto exit
%GOPATH%\bin\crawler

:exit
