package grpcApi

import (
	"context"
	"cryptoService/cryptoService/grpcApi"
	"cryptoService/cryptography/utils"
)

type Server struct {
	grpcApi.UnimplementedCryptoServiceServer
}

func (s *Server) EncryptData(ctx context.Context, req *grpcApi.EncryptDecryptRequest) (*grpcApi.EncryptDecryptResponse, error) {
	request := req.GetMessage()
	pad := utils.GetPaddingByName(request.Padding)
	algo := utils.GetAlgorithmByName(request.CryptoAlgorithm, utils.StringToByteArray(request.Key))
	mod, _ := utils.GetModeByName(request.CipherMode, pad, algo)
	encrypted, err := mod.Encrypt(utils.StringToByteArray(request.Content))

	if err != nil {
		return nil, err
	}

	request.Content = utils.BytesToString(encrypted)

	return &grpcApi.EncryptDecryptResponse{
		Message: "File data received successfully",
		Data:    request,
	}, nil
}

func (s *Server) DecryptData(ctx context.Context, req *grpcApi.EncryptDecryptRequest) (*grpcApi.EncryptDecryptResponse, error) {
	request := req.GetMessage()
	pad := utils.GetPaddingByName(request.Padding)
	algo := utils.GetAlgorithmByName(request.CryptoAlgorithm, utils.StringToByteArray(request.Key))
	mod, _ := utils.GetModeByName(request.CipherMode, pad, algo)
	decrypted, err := mod.Decrypt(utils.StringToByteArray(request.Content))

	if err != nil {
		return nil, err
	}

	request.Content = utils.BytesToString(decrypted)

	return &grpcApi.EncryptDecryptResponse{
		Message: "File data received successfully",
		Data:    request,
	}, nil
}

func (s *Server) EncryptFileFromComputer(ctx context.Context, req *grpcApi.EncryptDecryptFileRequest) (*grpcApi.EncryptDecryptFileResponse, error) {
	request := req.GetMessage().Message
	format := req.GetMessage().Format
	pad := utils.GetPaddingByName(request.Padding)
	algo := utils.GetAlgorithmByName(request.CryptoAlgorithm, utils.StringToByteArray(request.Key))
	mod, _ := utils.GetModeByName(request.CipherMode, pad, algo)

	plaintext, err := utils.ReadFileToBytes(request.Content)
	if err != nil {
		return nil, err
	}
	encrypted, err := mod.Encrypt(plaintext)

	if err != nil {
		return nil, err
	}

	return &grpcApi.EncryptDecryptFileResponse{
		Message: "File data received successfully",
		Data: &grpcApi.MessageLikeFile{
			Message: &grpcApi.Message{
				Key:             request.Key,
				CryptoAlgorithm: request.CryptoAlgorithm,
				Padding:         request.Padding,
				CipherMode:      request.CipherMode,
				Content:         utils.BytesToString(encrypted),
			},
			Format: format,
		},
	}, nil
}

func (s *Server) DecryptFileFromComputer(ctx context.Context, req *grpcApi.EncryptDecryptFileRequest) (*grpcApi.EncryptDecryptFileResponse, error) {
	request := req.GetMessage().Message
	format := req.GetMessage().Format
	pad := utils.GetPaddingByName(request.Padding)
	algo := utils.GetAlgorithmByName(request.CryptoAlgorithm, utils.StringToByteArray(request.Key))
	mod, _ := utils.GetModeByName(request.CipherMode, pad, algo)

	encrypted, err := utils.ReadFileToBytes(request.Content)
	if err != nil {
		return nil, err
	}

	decrypted, err := mod.Decrypt(encrypted)
	if err != nil {
		return nil, err
	}

	return &grpcApi.EncryptDecryptFileResponse{
		Message: "File data received successfully",
		Data: &grpcApi.MessageLikeFile{
			Message: &grpcApi.Message{
				Key:             request.Key,
				CryptoAlgorithm: request.CryptoAlgorithm,
				Padding:         request.Padding,
				CipherMode:      request.CipherMode,
				Content:         utils.BytesToString(decrypted),
			},
			Format: format,
		},
	}, nil
}
