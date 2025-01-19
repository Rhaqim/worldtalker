package service

import (
	"context"
	"testing"

	pb "github.com/Rhaqim/wtbackend/proto/translate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

// MockTranslatorClient is a mock implementation of the TranslatorClient interface
type MockTranslatorClient struct {
	mock.Mock
}

func (m *MockTranslatorClient) Translate(ctx context.Context, in *pb.TranslateRequest, opts ...grpc.CallOption) (*pb.TranslateResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.TranslateResponse), args.Error(1)
}

func TestTranslate(t *testing.T) {
	mockClient := new(MockTranslatorClient)
	translator := &Translator{client: mockClient}

	testCases := []struct {
		content        string
		sourceLanguage string
		targetLanguage string
		expectedResult string
		expectedError  error
	}{
		{"Hello", "en", "es", "Hola", nil},
		{"Bonjour", "fr", "en", "Hello", nil},
	}

	for _, tc := range testCases {
		req := &pb.TranslateRequest{
			Message:        tc.content,
			LanguageSource: tc.sourceLanguage,
			LanguageTarget: tc.targetLanguage,
		}
		resp := &pb.TranslateResponse{TranslatedMessage: tc.expectedResult}

		mockClient.On("Translate", context.Background(), req).Return(resp, tc.expectedError)

		result, err := translator.Translate(tc.content, tc.sourceLanguage, tc.targetLanguage)

		assert.Equal(t, tc.expectedResult, result)
		assert.Equal(t, tc.expectedError, err)

		mockClient.AssertExpectations(t)
	}
}
