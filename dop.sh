#!/bin/sh
if [ $1 == "start" ]
then
  echo "do start"
  # shellcheck disable=SC2092
  # shellcheck disable=SC2006
  `go run cmd/main.go`
fi