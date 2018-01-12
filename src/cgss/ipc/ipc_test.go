package ipc

import "testing"

type EchoServer struct {
}

func (server *EchoServer) Handle(request, params string) *Response {
	return &Response{"200", "ECHO: " + request}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NesIpcClient(server)
	client2 := NesIpcClient(server)

	resp1, _ := client1.Call("From client1", "test")
	resp2, _ := client1.Call("From client2", "test")

	if resp1.Body != "ECHO: From client1" || resp2.Body != "ECHO: From client2" {
		t.Error("IpcClient call failed. resp1: ", resp1, " resp2: ", resp2)
	}

	client1.Close()
	client2.Close()
}
