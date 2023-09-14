package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/rs/zerolog/log"
	grpc "google.golang.org/grpc"
)

// TODO: refactor
// TODO: make it more generic
// TODO: .env
// TODO: changelog
// TODO: readme
// TODO: args
// TODO: better logging
// TODO: docker file
// TODO: task file
// TODO: docker compose
// TODO: tests

func main() {
	host, port := "localhost", "9004" // TODO: as arg
	filename := "dump.zip"            // TODO: as arg

	file, err := os.Open(filename)
	if err != nil {
		log.Error().Msgf("failed to open file: %v", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Error().Msgf("failed to close file: %v", err)
		}
	}()

	con, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error().Msgf("failed to dial to %s:%s :%v", host, port, err)
		return
	}
	defer func() {
		log.Info().Msg("closing gRPC connection..")
		if err := con.Close(); err != nil {
			log.Error().Err(err).Msg("failed to close connection")
		}
	}()

	client := proto.NewFileClient(con) // TODO: generate example proto files

	//_, name := path.Split(f.Name())
	stat, _ := file.Stat()

	log.Info().Msgf("sending file %s with size %d", filename, stat.Size())

	mimetype := "application/zip"

	req := &proto.FileRequest{ // TODO: generate example proto files
		Chunk:    []byte{},
		MimeType: &mimetype,
	}

	// TODO: chunk size as argument
	buf := make([]byte, 1048576*2) // chunk size 2mb

	ctx, canc := context.WithTimeout(context.Background(), time.Minute*60) //TODO: too short
	defer canc()

	stream, err := client.Seed(ctx)
	if err != nil {
		log.Error().Msgf("failed to call SendStream: %v", err)
		return
	}

	startTime := time.Now()
	var offset int64
	for {
		n, errRead := file.ReadAt(buf, offset)
		if errRead != nil && errRead != io.EOF {
			log.Error().Msgf("failed to read file part: %v", errRead)

			if errCloseSend := stream.CloseSend(); err != nil {
				log.Error().Msgf("failed closing send: %v", errCloseSend)
			}

			break
		}

		req.Chunk = buf[:n]

		if errSendMsg := stream.SendMsg(req); errSendMsg != nil {
			// receive status error message from server
			errSendMsg = stream.RecvMsg(nil)

			sErr, ok := status.FromError(err)
			if ok {
				switch sErr.Code() {
				case codes.DeadlineExceeded:
					log.Info().Msgf("deadline exceeded: %s", sErr.Message())
				case codes.Internal:
					log.Error().Msgf("server error: %s", sErr.Message())
				}
			} else {
				log.Error().Msgf("transport layer error: %v", errSendMsg)
			}

			// in any case close send and break loop
			if errCloseSend := stream.CloseSend(); err != nil {
				log.Error().Msgf("failed closing send: %v", errCloseSend)
			}

			break
		}

		if errRead != nil && errRead == io.EOF {
			// receive response and sending break loop
			if _, errCloseRecv := stream.CloseAndRecv(); errCloseRecv != nil {
				log.Error().Msgf("failed closing stream: %v", errCloseRecv)
			}

			// successfully sent file
			break
		}

		offset += int64(n)
	}

	log.Info().Msgf("task took %d seconds", int(time.Since(startTime).Seconds()))
}
