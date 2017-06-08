package client

import (
	"log"

	"golang.org/x/net/context"

	"commands"
	"google.golang.org/grpc"
)

var emptyParams = []string{}

const (
	port        = ":8086"
	serviceName = "/control.commands/"
)

func MouseCommandEndPoint(ip, x, y string) string {
	params := []string{x, y}

	return sendCommand("MouseCommandEndPoint", ip, params)
}

func ClickCommandEndPoint(ip string) string {
	return sendCommand("ClickCommandEndPoint", ip, emptyParams)
}

func KeyboardCommandEndPoint(ip, text string) string {
	return sendCommand("KeyboardCommandEndPoint", ip, []string{text})
}

func SpaceCommandEndPoint(ip string) string {
	return sendCommand("SpaceCommandEndPoint", ip, emptyParams)
}

func EnterCommandEndPoint(ip string) string {
	return sendCommand("EnterCommandEndPoint", ip, emptyParams)
}

func LeftCommandEndPoint(ip string) string {
	return sendCommand("LeftCommandEndPoint", ip, emptyParams)
}
func RightCommandEndPoint(ip string) string {
	return sendCommand("RightCommandEndPoint", ip, emptyParams)
}

func LowerAudioCommandEndPoint(ip string) string {
	return sendCommand("LowerAudioCommandEndPoint", ip, emptyParams)
}

func UpCommandEndPoint(ip string) string {
	return sendCommand("UpCommandEndPoint", ip, emptyParams)
}

func DownCommandEndPoint(ip string) string {
	return sendCommand("DownCommandEndPoint", ip, emptyParams)
}

func IncreaseAudioCommandEndPoint(ip string) string {
	return sendCommand("IncreaseAudioCommandEndPoint", ip, emptyParams)
}

func sendCommand(method, ip string, params []string) string {
	// Set up a connection to the server.
	address := ip + port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//c := pb.NewGreeterClient(conn)

	//r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
	method = serviceName + method
	out := new(commands.CommandReply)
	err = grpc.Invoke(context.Background(), method, &commands.CommandRequest{Params: params},
		out, conn, grpc.EmptyCallOption{})
	if err != nil {
		out.Status = "500"
	}

	return out.Status
}
