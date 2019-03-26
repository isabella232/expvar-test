#!/bin/sh
go run app1/app1.go &
go run app2/app2.go &
go run app3/app3.go &
go run client/client.go
