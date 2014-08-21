package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"code.google.com/p/go.crypto/ssh"
	"ria/cli/ssh_utils"
	"strings"
	"reflect"
	color "github.com/daviddengcn/go-colortext"
)

func Tail() cli.Command {
	command := cli.Command{
		Name: "tail",
		Usage:     "show last logs",
		Flags: []cli.Flag {
			cli.StringFlag{
				Name: "s",
				Value: "web01.mylan",
				Usage: "server",
			},
			cli.StringFlag{
				Name: "v",
				Value: "country",
				Usage: "my",
			},
			cli.IntFlag{
				Name: "n",
				Value: 100,
				Usage: "number of lines",
			},
			cli.BoolFlag{
				Name: "f",
				Usage: "watch results",
			},
		},
		Action: func(c *cli.Context) {

			config := &ssh.ClientConfig{
				User:
				"asharov",
				Auth: []ssh.AuthMethod{ssh_utils.MakeKeyring()},
			}


			type Host struct {
				name   string
				stdout chan string
			}
			var hosts []*Host;

			for _, hostName := range strings.Split(c.String("s"), ",") {
				hosts = append(hosts, &Host{
						name: hostName,
						stdout: make(chan string, 100),
					})
			}
			stderr := make(chan error)

			for _, host := range hosts {
				var flags string
				if c.Bool("f") {
					flags += " -f"
				} else {
					flags += " -"+c.String("n")
				}
				cmd := "tail " + flags + " /shop/logs/live/nginx/alice.access.log-20140821";
				go ssh_utils.ExecuteCmd(cmd, host.name, config, host.stdout, stderr)

			}

			go func() {
				cases := make([]reflect.SelectCase, len(hosts))
				for i, host := range hosts {
					cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(host.stdout)}
				}
				for {
					if chosen, value, ok := reflect.Select(cases); ok {
						host := hosts[chosen]
						res := value.String()

						color.ChangeColor(color.Blue, true, color.None, false)
						fmt.Print(host.name + ": ")
						color.ResetColor()
						fmt.Print(res)
					}
				}
			}()

			err := <-stderr;
			println(fmt.Sprintf("%q", err))
		},
	}
	return command
}

