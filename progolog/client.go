package progolog

import (
	"bufio"
	"context"
	"io"

	"github.com/pkg/errors"
	pb "github.com/progoci/progo-log/progolog"
	"google.golang.org/grpc"
)

// Client is the client for progo log.
//
// This was kept as a exported variable to allow testing. Dependency injection,
// although feasible, requires to pass the client down the chain, which makes
// it very confusing why the methods in between needs a client to the logs.
var Client pb.LoggerClient

// SendOpts are the options for sending a log.
type SendOpts struct {
	BuildID     string
	ServiceName string
	StepName    string
	StepNumber  int32
	Command     string
}

// Start starts the progolog client
func Start(conn *grpc.ClientConn) {
	Client = pb.NewLoggerClient(conn)
}

// Send sends a new message.
func Send(buffer *bufio.Reader, opts *SendOpts) error {
	stream, err := Client.Store(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to store log")
	}

	var buf []byte

	for {
		buf = make([]byte, buffer.Size())
		_, err := buffer.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "failed to log piece")
		}

		// Prepares and send message.
		log := &pb.Log{
			BuildID:     opts.BuildID,
			ServiceName: opts.ServiceName,
			StepName:    opts.StepName,
			StepNumber:  opts.StepNumber,
			Command:     opts.Command,
			Body:        buf,
		}
		if err := stream.Send(log); err != nil {
			return errors.Wrap(err, "failed to final log piece")
		}
	}

	// Send the last piece.
	log := &pb.Log{
		Body: buf,
	}
	if err := stream.Send(log); err != nil {
		return errors.Wrap(err, "failed to final log piece")
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		return errors.Wrap(err, "failed to close streaming")
	}

	return nil
}
