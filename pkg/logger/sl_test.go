package logger

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErr(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want slog.Attr
	}{
		{
			name: "Test1",
			args: args{
				err: errors.New("test"),
			},
			want: slog.Attr{
				Key:   "error",
				Value: slog.StringValue(errors.New("test").Error()),
			},
		},
		{
			name: "Test2",
			args: args{
				err: errors.New("1234"),
			},
			want: slog.Attr{
				Key:   "error",
				Value: slog.StringValue(errors.New("1234").Error()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Err(tt.args.err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestSetupLogger(t *testing.T) {
	tests := []struct {
		name    string
		env     Environment
		wantErr bool
	}{
		{
			name: "Ok-Local",
			env:  Local,
		},
		{
			name: "Ok-Dev",
			env:  Dev,
		},
		{
			name: "Ok-Prod",
			env:  Prod,
		},
		{
			name:    "Err-InvalidEnv",
			env:     "test",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := SetupLogger(tt.env)
			if !tt.wantErr {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
