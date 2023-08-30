package users

import (
	"reflect"
	"testing"
)

func TestNewUserName(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserName
		wantErr bool
	}{
		{name: "success case", args: args{userName: "userName09+=,.@-"}, want: wantUserName("userName09+=,.@-"), wantErr: false},
		{name: "fail case", args: args{userName: "userName^"}, want: nil, wantErr: true},
		{name: "fail case", args: args{userName: "テスト"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserName(tt.args.userName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func wantUserName(userName string) *UserName {
	u, _ := NewUserName(userName)
	return u
}
