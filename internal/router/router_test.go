package router_test

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/PritOriginal/zero-agency-test/internal/router"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RouterSuite struct {
	suite.Suite
	log    *slog.Logger
	c      *router.MockClassifier
	skills []*router.MockSkill
	r      *router.Router
}

func (st *RouterSuite) SetupSuite() {
	st.log = slog.New(slog.DiscardHandler)
	st.c = router.NewMockClassifier(st.T())
	st.skills = []*router.MockSkill{
		router.NewMockSkill(st.T()),
	}
	st.r = router.New(st.log, st.c)
	st.r.RegisterSkill("test", st.skills[0])
}

func TestRouter(t *testing.T) {
	suite.Run(t, new(RouterSuite))
}

type method[T any] struct {
	data T
	err  error
}

func (st *RouterSuite) TestMessage() {
	tests := []struct {
		name            string
		classify        method[string]
		wantErrGetSkill bool
		skillExecute    method[string]
	}{
		{
			name: "Ok",
			classify: method[string]{
				data: "test",
			},
			skillExecute: method[string]{
				data: "Навых выполнен",
			},
		},
		{
			name: "ErrClassify",
			classify: method[string]{
				err: errors.New(""),
			},
		},
		{
			name: "ErrUnknownTag",
			classify: method[string]{
				data: "tag",
			},
			wantErrGetSkill: true,
		},
		{
			name: "ErrExecuteSkill",
			classify: method[string]{
				data: "test",
			},
			skillExecute: method[string]{
				err: errors.New(""),
			},
		},
	}

	for _, tt := range tests {
		st.Run(tt.name, func() {
			func() {
				st.c.On("Classify", mock.Anything, mock.AnythingOfType("string")).Once().
					Return(tt.classify.data, tt.classify.err)
				if tt.classify.err != nil {
					return
				}

				if tt.wantErrGetSkill {
					return
				}

				st.skills[0].On("Execute", mock.AnythingOfType("string")).Once().
					Return(tt.skillExecute.data, tt.skillExecute.err)
				if tt.skillExecute.err != nil {
					return
				}
			}()

			resp, err := st.r.Route("тест")

			if tt.classify.err == nil &&
				!tt.wantErrGetSkill &&
				tt.skillExecute.err == nil {
				st.Equal(resp, tt.skillExecute.data)
				st.NoError(err)
			} else {
				st.NotNil(err)
			}

			st.c.AssertExpectations(st.T())
			for _, skill := range st.skills {
				skill.AssertExpectations(st.T())
			}
		})
	}
}
