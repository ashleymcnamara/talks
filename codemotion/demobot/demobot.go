package main

import (
	"fmt"
	"os/exec"

	"github.com/micro/go-micro"
	"golang.org/x/net/context"

	proto "github.com/micro/go-bot/proto"
)

type Command struct{}

// Help returns the command usage
func (c *Command) Help(ctx context.Context, req *proto.HelpRequest, rsp *proto.HelpResponse) error {
	// Usage should include the name of the command
	rsp.Usage = "dockerps"
	rsp.Description = "This command returns the output of the `docker ps` command"
	return nil
}

// Exec executes the command
func (c *Command) Exec(ctx context.Context, req *proto.ExecRequest, rsp *proto.ExecResponse) error {
	cmd := exec.Command("docker", "ps")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		rsp.Result = []byte(err.Error())
		return err
	}
	fmt.Printf("%s\n", stdoutStderr)
	rsp.Result = []byte(stdoutStderr)
	// rsp.Error could be set to return an error instead
	// the function error would only be used for service level issues
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.bot.dockerps"),
	)

	service.Init()

	proto.RegisterCommandHandler(service.Server(), new(Command))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
