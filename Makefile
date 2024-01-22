# Detect operating system in Makefile.
# Author: He Tao
# Date: 2015-05-30

ifeq ($(OS),Windows_NT)
    # Note: .exe is *required* on Windows; otherwise, 
    #       make.exe quietly falls back to cmd.exe
    SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
endif

test:
ifeq ($(OS),Windows_NT)
	go test @(go list -f '{{.Dir}}' -m)
else
	go list -f '{{.Dir}}' -m | xargs go test
endif