package main

import (
	"cryptoService/cryptoService/grpcApi"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	//router := gin.Default()
	//
	//router.POST("/encryptMsg", contollers.EncryptDataHandler)
	//router.POST("/decryptMsg", contollers.DecryptDataHandler)
	//router.POST("/encryptFile", contollers.EncryptFileFromComputerHandler)
	//router.POST("/decryptFile", contollers.DecryptFileFromComputerHandler)
	//
	//err := router.Run(":8090")
	//if err != nil {
	//	_ = fmt.Errorf("error with start GO seervice %s", err)
	//	return
	//}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}

	s := grpc.NewServer()
	grpcApi.RegisterCryptoServiceServer(s, grpcApi.UnimplementedCryptoServiceServer{})

	fmt.Println("gRPC server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
