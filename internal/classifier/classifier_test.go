package classifier

import (
	"context"
	"errors"
	"log/slog"
	"testing"

	"github.com/PritOriginal/zero-agency-test/internal/shared/tags"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ClassifierSuite struct {
	suite.Suite
	log        *slog.Logger
	client     *MockClient
	classifier *Classifier
}

func (st *ClassifierSuite) SetupSuite() {
	st.log = slog.New(slog.DiscardHandler)
	st.client = NewMockClient(st.T())
	st.classifier = New(st.log, st.client, []string{"support_urgency"})
}

func TestClassifier(t *testing.T) {
	suite.Run(t, new(ClassifierSuite))
}

type method[T any] struct {
	data T
	err  error
}

func (st *ClassifierSuite) TestСlassify() {
	tests := []struct {
		name                  string
		userInput             string
		wantEmptyInputErr     bool
		doRequest             method[string]
		wantErrUnmarshal      bool
		wantErrUnsupportedTag bool
	}{
		{
			name:      "Ok",
			userInput: "тест",
			doRequest: method[string]{
				data: `{"category": "support_urgency"}`,
			},
		},
		{
			name:              "Err-EmptyUserInput",
			userInput:         "",
			wantEmptyInputErr: true,
		},
		{
			name:      "Err-APIClassify",
			userInput: "тест",
			doRequest: method[string]{
				err: errors.New(""),
			},
		},
		{
			name:      "Err-Unmarshal",
			userInput: "тест",
			doRequest: method[string]{
				data: `{"category": "support_urgency"`,
			},
			wantErrUnmarshal: true,
		},
		{
			name:      "Err-NotSupportedTag",
			userInput: "тест",
			doRequest: method[string]{
				data: `{"category": "test"}`,
			},
			wantErrUnsupportedTag: true,
		},
	}

	for _, tt := range tests {
		st.Run(tt.name, func() {
			func() {
				if tt.wantEmptyInputErr {
					return
				}
				st.client.On("DoRequest", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Once().
					Return(tt.doRequest.data, tt.doRequest.err)
			}()

			tag, err := st.classifier.Classify(context.Background(), tt.userInput)

			if !tt.wantEmptyInputErr &&
				tt.doRequest.err == nil &&
				!tt.wantErrUnmarshal &&
				!tt.wantErrUnsupportedTag {
				st.Equal(tags.SupportUrgency, tag)
				st.NoError(err)
			} else {
				st.NotNil(err)
			}

			st.client.AssertExpectations(st.T())
		})
	}
}
