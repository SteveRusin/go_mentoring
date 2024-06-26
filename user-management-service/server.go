package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	users_rpc "github.com/SteveRusin/go_mentoring/generated"
	"github.com/SteveRusin/go_mentoring/user-management-service/randomId"
	"github.com/SteveRusin/go_mentoring/user-management-service/repository"
)

type Server struct {
	users_rpc.UnimplementedUserMangmentServer
	repo repository.UserRepository
}

func newServer() *Server {
	return &Server{
		repo: repository.NewUserPgRepository(),
	}
}

func (s *Server) StoreUser(ctx context.Context, req *users_rpc.StoreUserRequest) (*users_rpc.StoreUserReply, error) {
	user := &repository.User{
		Id:       randomId.New(),
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}

	result, err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}

	reply := users_rpc.StoreUserReply{
		Id:   result.Id,
		Name: result.Name,
	}

	return &reply, nil
}

func (s *Server) GetUser(ctx context.Context, req *users_rpc.GetUserRequest) (*users_rpc.GetUserReply, error) {
	creds := &repository.UserCreds{
		Name:     req.GetName(),
		Password: req.GetPassword(),
	}

	res, err := s.repo.FindUserByCreds(creds)
	if err != nil {
		return nil, err
	}

	reply := &users_rpc.GetUserReply{
		Id:   res.Id,
		Name: res.Name,
	}

	return reply, nil
}

func (s *Server) UploadImage(stream users_rpc.UserMangment_UploadImageServer) error {
	log.Println("Starting image uploading")
	totalReceivedSize := uint32(0)
	fileId := randomId.New()

	req, err := stream.Recv()
	if err != nil {
		log.Println("error", err)
		return stream.SendAndClose(&users_rpc.UploadImageResponse{
			Response: &users_rpc.UploadImageResponse_Error{
				Error: "Something went wrong",
			},
		})
	}

	info := req.GetInfo()
	imgType := info.GetImageType()

	if imgType != "jpeg" && imgType != "png" {
		log.Println("Invalid image type")
		return stream.SendAndClose(&users_rpc.UploadImageResponse{
			Response: &users_rpc.UploadImageResponse_Error{
				Error: "Invalid image type",
			},
		})
	}

	err = os.MkdirAll("images", os.ModePerm)
	if err != nil {
		println("Error creating directory", err.Error())
		return stream.SendAndClose(&users_rpc.UploadImageResponse{
			Response: &users_rpc.UploadImageResponse_Error{
				Error: "Error creating directory",
			},
		})
	}

	filePath := "images/" + fileId + "." + imgType
	file, err := os.Create(filePath)
	if err != nil {
		println("Error creating file", err.Error())
		return stream.SendAndClose(&users_rpc.UploadImageResponse{
			Response: &users_rpc.UploadImageResponse_Error{
				Error: "Something went wrong",
			},
		})
	}
	defer file.Close()

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			log.Println("EOF: Done receiving chunks")
			break
		}

		if err != nil {
			log.Println("Error reading chunk", err)
			cleanUpFile(filePath)
			return stream.SendAndClose(&users_rpc.UploadImageResponse{
				Response: &users_rpc.UploadImageResponse_Error{
					Error: "Error reading chunk",
				},
			})
		}

		chunk := req.GetChunkData()
		totalReceivedSize += uint32(len(chunk))
		log.Println("Received", totalReceivedSize, "bytes", "of", info.GetFileSize())

		_, err = file.Write(chunk)
		if err != nil {
			log.Println("Error writing to file", err)
			cleanUpFile(filePath)
			return stream.SendAndClose(&users_rpc.UploadImageResponse{
				Response: &users_rpc.UploadImageResponse_Error{
					Error: "Something went wrong while writing to file",
				},
			})
		}
	}

	return stream.SendAndClose(&users_rpc.UploadImageResponse{
		Response: &users_rpc.UploadImageResponse_Id{
			Id: fileId + "." + imgType,
		},
	})
}

func (s *Server) FetchImage(req *users_rpc.FetchImageRequest, stream users_rpc.UserMangment_FetchImageServer) error {
	image := strings.Trim(req.GetId(), " ")
	log.Println("Fetching image", image)

	if image == "" {
		return stream.Send(&users_rpc.FetchImageResponse{
			Data: &users_rpc.FetchImageResponse_Error{
				Error: "Image id is epmty",
			},
		})
	}

	file, err := os.Open("images/" + image)
	if err != nil {
		log.Println("Error opening file", err)
		return stream.Send(&users_rpc.FetchImageResponse{
			Data: &users_rpc.FetchImageResponse_Error{
				Error: "Image does not exist",
			},
		})
	}
	defer file.Close()

	chunkSize := 1024

	buffer := make([]byte, chunkSize)

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error:", err)
			return stream.Send(&users_rpc.FetchImageResponse{
				Data: &users_rpc.FetchImageResponse_Error{
					Error: "Error reading chunk",
				},
			})
		}

		if bytesRead == 0 || err == io.EOF {
			println("No more data to read")
			break
		}

		err = stream.Send(&users_rpc.FetchImageResponse{
			Data: &users_rpc.FetchImageResponse_ChunkData{ChunkData: buffer[:bytesRead]},
		})
		if err != nil {
			fmt.Println("Error sending chunk:", err)
			break
		}
	}

	return nil
}

func cleanUpFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Println("Error removing file", err)
	}
}
