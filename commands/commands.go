package commands

import (
	"github.com/go-vgo/robotgo"

	"golang.org/x/net/context"
	"strconv"
)

const speed = 8

type Command struct{}

func (Command) MouseCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	mouseX, mouseY := robotgo.GetMousePos()

	moveX, err := strconv.Atoi(command.Params[0])
	if err != nil {
		return &CommandReply{Status: "404"}, nil
	}

	moveY, err := strconv.Atoi(command.Params[1])
	if err != nil {
		return &CommandReply{Status: "404"}, nil
	}

	sumX := mouseX + (moveX * speed)
	sumY := mouseY + (moveY * speed)

	robotgo.MoveMouse(sumX, sumY)

	return &CommandReply{Status: "OK"}, nil
}

func (Command) ClickCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	robotgo.MouseClick()
	return &CommandReply{Status: "OK"}, nil
}

func (Command) KeyboardCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	for _, input := range command.Params {
		for _, letter := range input {
			robotgo.KeyTap(string(letter))
		}
	}

	return &CommandReply{Status: "OK"}, nil
}

func (Command) SpaceCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	robotgo.KeyTap("escape")
	return &CommandReply{Status: "OK"}, nil
}

func (Command) EnterCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	robotgo.KeyTap("enter")
	return &CommandReply{Status: "OK"}, nil
}

func (Command) LeftCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	robotgo.KeyTap("left")
	return &CommandReply{Status: "OK"}, nil
}

func (Command) RightCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	robotgo.KeyTap("right")
	return &CommandReply{Status: "OK"}, nil
}

func (Command) LowerAudioCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	robotgo.KeyTap("audio_vol_down")
	return &CommandReply{Status: "OK"}, nil
}

func (Command) UpCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	robotgo.KeyTap("up")
	return &CommandReply{Status: "OK"}, nil
}

func (Command) DownCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	robotgo.KeyTap("down")
	return &CommandReply{Status: "OK"}, nil
}

func (Command) IncreaseAudioCommandEndPoint(ctx context.Context, command *CommandRequest) (*CommandReply, error) {
	robotgo.KeyTap("audio_vol_up")
	return &CommandReply{Status: "OK"}, nil
}
