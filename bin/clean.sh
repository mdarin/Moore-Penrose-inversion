#!/bin/bash

progname=""
if [ -n "$1" ]
then
  progname=$1
  shift
  echo "progname: $progname"
else
  echo "You should define progname!"     
	echo "Usage: cmd <progname>"
  exit 1
fi

export GOPATH=$(pwd) 
go clean src/main.go 
rm -f $progname 
exit 0

