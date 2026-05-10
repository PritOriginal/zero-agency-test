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
	st.classifier = New(st.log, st.client, []string{"test"})
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
		name             string
		doRequest        method[string]
		wantErrUnmarshal bool
	}{
		{
			name: "Ok",
			doRequest: method[string]{
				data: `{"category": "support_urgency"}`,
			},
		},
		{
			name: "Err-APIClassify",
			doRequest: method[string]{
				err: errors.New(""),
			},
		},
		{
			name: "Err-Unmarshal",
			doRequest: method[string]{
				data: `{"category": "support_urgency"`,
			},
			wantErrUnmarshal: true,
		},
	}

	for _, tt := range tests {
		st.Run(tt.name, func() {
			func() {
				st.client.On("DoRequest", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).Once().
					Return(tt.doRequest.data, tt.doRequest.err)
			}()

			tag, err := st.classifier.Classify(context.Background(), "тест")

			if tt.doRequest.err == nil && !tt.wantErrUnmarshal {
				st.Equal(tags.SupportUrgency, tag)
				st.NoError(err)
			} else {
				st.NotNil(err)
			}

			st.client.AssertExpectations(st.T())
		})
	}
}
