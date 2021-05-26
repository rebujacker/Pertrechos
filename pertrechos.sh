#!/bin/sh

case "$1" in

	#Remove Golang and dependencies
    "clean" )
		rm -rf go*
		rm -rf butron/pkg
		rm -rf butron/src
		rm -rf butron/butron
		rm -rf falcata/pkg
		rm -rf falcata/src
		rm -rf falcata/falcata
		exit 1
        ;;

    #Get dependencies and compile for each tool

    "butron" )
		
		OS=$2
		ARCHITECTURE=$3

		#Required Software to compile go binaries
		sudo apt-get update
		sudo apt-get install gcc unzip

		#Download GO and set env vars. correctly to compile
		wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz
		tar xvf go1.13.3.linux-amd64.tar.gz -C .
		rm go1.13.3.linux-amd64.tar.gz
		export GOROOT="$(pwd)/go"
		export GOPATH="$(pwd)/butron/"
		export PATH=$(pwd)/go/bin:$PATH

		#Get Libraries for target tool
		cd butron
		go get "github.com/armon/go-socks5"
		go get "golang.org/x/crypto/ssh"
		
		#Compile
		GOOS=${OS} GOARCH=${ARCHITECTURE} go build -o butron butron.go


		exit 1
    	;;

    "falcata" )

		OS=$2
		ARCHITECTURE=$3

		#Required Software to compile go binaries
		sudo apt-get update
		sudo apt-get install gcc unzip

		#Download GO and set env vars. correctly to compile
		wget https://dl.google.com/go/go1.13.3.linux-amd64.tar.gz
		tar xvf go1.13.3.linux-amd64.tar.gz -C .
		rm go1.13.3.linux-amd64.tar.gz
		export GOROOT="$(pwd)/go"
		export GOPATH="$(pwd)/falcata/"
		export PATH=$(pwd)/go/bin:$PATH

		#Get Libraries for target tool
		cd falcata
		go get "golang.org/x/crypto/ssh"
		go get "golang.org/x/term"
		go get "github.com/kr/pty"
		
		#Compile
		GOOS=${OS} GOARCH=${ARCHITECTURE} go build -o falcata falcata.go
	
		exit 1
        ;;
esac

