package realisations

import (
	"github.com/NicoNex/echotron/v3"
	"testing"
)

func TestEchoRealisations_SendMsg(t *testing.T) {
	token := "TOKEN"
	type fields struct {
		API echotron.API
	}
	type args struct {
		chatId int64
		msg    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				API: echotron.NewAPI(token),
			},
			args: args{
				chatId: 7138652177,
				msg:    "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echoRealisations{
				API: tt.fields.API,
			}
			if err := e.SendMsg(tt.args.chatId, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("EchoRealisations.SendMsg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEchoRealisations_SendPhoto(t *testing.T) {
	token := "TOKEN"
	type fields struct {
		API echotron.API
	}
	type args struct {
		chatId int64
		file   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				API: echotron.NewAPI(token),
			},
			args: args{
				chatId: 7138652177,
				file:   []byte("FILE"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echoRealisations{
				API: tt.fields.API,
			}
			if err := e.SendPhoto(tt.args.chatId, tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("EchoRealisations.SendPhoto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
