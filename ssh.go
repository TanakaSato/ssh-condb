package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/manifoldco/promptui"

	db "./db"
)

func main() {

	host := db.GetAnyHost()

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "> {{ .Hostname | cyan }}",
		Inactive: "  {{ .Hostname | white }}",
		Selected: "> {{ .Hostname | cyan }}",
		Details: `
--------- Host Info ----------
{{ "Hostname:" | faint }}	{{ .Hostname }}
{{ "Username:" | faint }}	{{ .Username }}
{{ "Proxy:" | faint }}	{{ .Proxy }}
{{ "Port:" | faint }}	{{ .Port }}`,
	}

	searcher := func(input string, index int) bool {
		s := host[index]
		hostname := strings.Replace(strings.ToLower(s.Hostname), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(hostname, input)
	}

	prompt := promptui.Select{
		Label:     "Select Host",
		Items:     host,
		Templates: templates,
		Size:      5,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()
	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		return
	}

	sc := db.GetSingleHost(host[i].Hostname)
	if sc.Hostname == "" {
		log.Fatal("invalid hostname.\n")
	}

	client := anyProxy(sc)

	session, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		log.Println(err)
	}
	defer terminal.Restore(fd, state)

	w, h, err := terminal.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm", h, w, modes)
	if err != nil {
		log.Println(err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	err = session.Shell()
	if err != nil {
		log.Fatal(err)
	}

	signalchan := make(chan os.Signal, 1)
	signal.Notify(signalchan, syscall.SIGWINCH)
	go func() {
		for {
			s := <-signalchan
			switch s {
			case syscall.SIGWINCH:
				fd := int(os.Stdout.Fd())
				w, h, _ = terminal.GetSize(fd)
				session.WindowChange(h, w)
			}
		}
	}()

	err = session.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func anyProxy(sshconf db.Sshconfig) *ssh.Client {
	sc := makeSSHConfig(sshconf)

	if sshconf.Proxy != 0 {
		proxy := db.GetID(sshconf.Proxy)

		c := anyProxy(proxy)

		proxyConn, err := c.Dial("tcp", net.JoinHostPort(sshconf.Hostname, strconv.Itoa(sshconf.Port)))
		if err != nil {
			log.Println(err)
		}
		pConnect, pChans, pReqs, err := ssh.NewClientConn(proxyConn, net.JoinHostPort(sshconf.Hostname, strconv.Itoa(sshconf.Port)), sc)
		if err != nil {
			log.Println(err)
		}
		log.Println("Connect: " + sshconf.Hostname)
		return ssh.NewClient(pConnect, pChans, pReqs)
	} else {
		log.Println("Connect: " + sshconf.Hostname)

		client, err := ssh.Dial("tcp", net.JoinHostPort(sshconf.Hostname, strconv.Itoa(sshconf.Port)), sc)
		if err != nil {
			log.Println(err)
		}
		return client
	}
}

func makeSSHConfig(sc db.Sshconfig) *ssh.ClientConfig {

	auth := []ssh.AuthMethod{}

	if sc.Authkey != "" {
		k := []byte(sc.Authkey)
		signer, err := ssh.ParsePrivateKey(k)
		if err != nil {
			log.Fatal(err)
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	if sc.Password != "" {
		auth = append(auth, ssh.Password(sc.Password))
	}

	sshConfig := &ssh.ClientConfig{
		User:            sc.Username,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	return sshConfig
}

func fp() {
	// TODO peco, fzf like
}
