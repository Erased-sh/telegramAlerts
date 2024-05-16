package tick

import (
	"context"
	"testing"
	"time"
)

func TestGenerateTicker(t *testing.T) {
	type args struct {
		amount  int64
		timeTag string
		action  func() error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				amount:  1,
				timeTag: "s",
				action: func() error {
					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateTicker(context.Background(), tt.args.amount, tt.args.timeTag, tt.args.action)
		})
	}
}

func Test_convertDuration(t *testing.T) {
	type args struct {
		amount  int64
		timeTag string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				amount:  1,
				timeTag: "s",
			},
			wantErr: false,
		},
		{
			name: "test2",
			args: args{
				amount:  1,
				timeTag: "m",
			},
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				amount:  1,
				timeTag: "h",
			},
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				amount:  1,
				timeTag: "ms",
			},
			wantErr: false,
		},
		{
			name: "test5",
			args: args{
				amount:  1,
				timeTag: "wrong",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := convertDuration(time.Duration(tt.args.amount), tt.args.timeTag)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertDuration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
