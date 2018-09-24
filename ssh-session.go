package icinga

import (
	"fmt"
	"strconv"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"io/ioutil"
	"os/user"
)

// NewSshSession creates a new instance of ssh.Session
func NewSshSession(host string, port int, ssh_user string) (*ssh.Session, error) {

	usr, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	if len(host) == 0 {
		return nil, fmt.Errorf("host is missing: [%s]", host)
	}
	if port <= 0 {
		return nil, fmt.Errorf("ssh port is invalid: [%v]", port)
	}

	signer, err := createSigner(usr.HomeDir + "/.ssh/id_rsa")
	if err != nil {
		return nil, fmt.Errorf("failed to create a signer for user [%s]: %v", usr.Name, err)
	}

	hostKeyCallback, err := knownhosts.New(usr.HomeDir + "/.ssh/known_hosts")
	if err != nil {
		return nil, fmt.Errorf("failed to get the host key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: ssh_user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}

	address := host + ":" + strconv.Itoa(port)
	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial host at [%s]: %v", address, err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("failed to create an ssh session %v", err)
	}

	return session, err
}

func createSigner(privKeyFile string) (ssh.Signer, error) {
	key, err := ioutil.ReadFile(fmt.Sprintf(privKeyFile))
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}

	return signer, nil
}