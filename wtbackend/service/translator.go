package service

import (
	"context"
	"log"

	"github.com/Rhaqim/wtbackend/domain"
	pb "github.com/Rhaqim/wtbackend/proto/translate"
	"google.golang.org/grpc"

	// "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Translator struct {
	client pb.TranslatorClient
}

func NewTranslatorClient(address string) (domain.TranslationService, error) {
	// creds := credentials.NewClientTLSFromCert(nil, "")
	creds := insecure.NewCredentials()

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	return &Translator{pb.NewTranslatorClient(conn)}, nil
}

func (t *Translator) Translate(content, sourceLanguage, targetLanguage string) (string, error) {

	req := &pb.TranslateRequest{
		Message:        content,
		LanguageSource: sourceLanguage,
		LanguageTarget: targetLanguage,
	}

	resp, err := t.client.Translate(context.Background(), req)
	if err != nil {
		return "", err
	}

	return resp.TranslatedMessage, nil
}
