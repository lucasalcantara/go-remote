package commands

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

type CommandRequest struct {
	Params []string `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
}

func (m *CommandRequest) Reset()                    { *m = CommandRequest{} }
func (m *CommandRequest) String() string            { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()               {}
func (*CommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CommandReply struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *CommandReply) Reset()                    { *m = CommandReply{} }
func (m *CommandReply) String() string            { return proto.CompactTextString(m) }
func (*CommandReply) ProtoMessage()               {}
func (*CommandReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*CommandRequest)(nil), "control.CommandRequest")
	proto.RegisterType((*CommandReply)(nil), "control.CommandReply")
}

type ControlServer interface {
	MouseCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	ClickCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	KeyboardCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	SpaceCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	EnterCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	LeftCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	RightCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	LowerAudioCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	UpCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	DownCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
	IncreaseAudioCommandEndPoint(context.Context, *CommandRequest) (*CommandReply, error)
}

type EndPoint func(context.Context, *CommandRequest) (*CommandReply, error)
