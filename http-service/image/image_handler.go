package image

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	users_rpc "github.com/SteveRusin/go_mentoring/generated"
	"github.com/SteveRusin/go_mentoring/http-service/middlewares"
	"github.com/SteveRusin/go_mentoring/http-service/users"
)

type imageHandlers struct {
	userClient users.UserClient
}

func NewImageHandlers() *imageHandlers {
	return &imageHandlers{
		userClient: users.NewUserRpcClient(),
	}
}

func (handler *imageHandlers) PostImage(w http.ResponseWriter, r *http.Request) *middlewares.HttpError {
	if r.Method != http.MethodPost {
		return middlewares.NewNotImplementedError()
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10 MB limit
	defer r.Body.Close()

	imageSize := r.ContentLength

	chunkSize := 1024 * 1024 // 1 MB chunk size
	buffer := make([]byte, chunkSize)
	client, err := handler.userClient.NewUploadImageClient(context.Background())
	if err != nil {
		log.Println("error creating client", err)
		return middlewares.NewBadRequestError()
	}

	totalSent := uint32(0)
	fileTypeSlice := strings.Split(r.Header.Get("Content-Type"), "/")
	if len(fileTypeSlice) != 2 || fileTypeSlice[0] != "image" {
		log.Println("Expected image type but got:", r.Header.Get("Content-Type"))
		return middlewares.NewBadRequestError()
	}

	imgInfo := &users_rpc.UploadImageRequest{
		Data: &users_rpc.UploadImageRequest_Info{
			Info: &users_rpc.ImageInfo{
				FileSize:  uint32(imageSize),
				ImageType: fileTypeSlice[1],
			},
		},
	}

	err = client.Send(imgInfo)
	if err != nil {
		log.Println("error sending image info", err)
		return middlewares.NewBadRequestError()
	}

	for {
		bytesRead, err := r.Body.Read(buffer)
		totalSent += uint32(bytesRead)

		if err != nil && err != io.EOF {
			log.Println("error reading chunk", err)
			return middlewares.NewBadRequestError()
		}

		if bytesRead == 0 {
			log.Println("No more data to read")
			break
		}

		req := &users_rpc.UploadImageRequest{
			Data: &users_rpc.UploadImageRequest_ChunkData{ChunkData: buffer[:bytesRead]},
		}

		log.Println("Sending chunk", bytesRead)
		err = client.Send(req)

		log.Println("Sent", totalSent, "bytes", "of", imageSize)
		if err == io.EOF {
			log.Println("EOF")
			break
		}

		if err != nil {
			log.Println("error sending chunk", err)
			return middlewares.NewBadRequestError()
		}

	}

	log.Println("Done uploading; Retrieving response")
	// without it postman receives EPIPE error
	// what's the reason?
	time.Sleep(time.Second)

	res, err := client.CloseAndRecv()
	if err != nil {
		println("Error closing send", err)
		return middlewares.NewBadRequestError()
	}

	if res.GetError() != "" {
		return middlewares.NewBadRequestError()
	}

	w.Write([]byte(res.GetId()))

	return nil
}

func (hander *imageHandlers) GetImage(w http.ResponseWriter, r *http.Request) *middlewares.HttpError {
	if r.Method != http.MethodGet {
		return middlewares.NewNotImplementedError()
	}

	pathSlice := strings.Split(r.URL.Path, "/")
	// e.g. /image/123
	if len(pathSlice) != 3 {
		return middlewares.NewBadRequestError()
	}

	imgId := pathSlice[2]

	client, err := hander.userClient.NewFetchImageClient(context.TODO(), &users_rpc.FetchImageRequest{Id: imgId})
	if err != nil {
		return middlewares.NewBadRequestError()
	}

	for {
		resp, err := client.Recv()

		if err == io.EOF {
			log.Println("EOF: Done receiving chunks")
			break
		}

		if err != nil {
			log.Println("Error reading chunk", err)

			return middlewares.NewBadRequestError()
		}

		fetchError := resp.GetError()
		if fetchError != "" {
			println("Error receiving image:", fetchError)
			return middlewares.NewBadRequestError()
		}
		chunk := resp.GetChunkData()

		if chunk == nil {
			println("No more data to fetch")
			break
		}

		_, err = w.Write(chunk)
		if err != nil {
			log.Println("Error writing to response", err)
			return middlewares.NewBadRequestError()
		}
	}

	return nil
}
