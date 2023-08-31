package attribute

import (
	"testing"
)

func TestUserAttributeName_EqualTo(t *testing.T) {
	type fields struct {
		userAttrName string
	}
	type args struct {
		v UserAttributeName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "equal",
			fields: fields{"userName"},
			args:   args{UserAttributeName{userAttrName: "userName"}},
			want:   true,
		},
		{
			name:   "not-equal",
			fields: fields{"userName"},
			args:   args{UserAttributeName{userAttrName: "userName-B"}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAttributeName{
				userAttrName: tt.fields.userAttrName,
			}
			if got := u.EqualTo(tt.args.v); got != tt.want {
				t.Errorf("EqualTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAttributeRecord_UpdateUserAttributeValue(t *testing.T) {
	type fields struct {
		userAttributeName     UserAttributeName
		userAttributeType     UserAttributeType
		userAttributeValidate UserAttributeValidate
		userAttributeValue    *UserAttributeValue
		canUpdate             bool
	}
	type args struct {
		attrValue UserAttributeValue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "update success: immutable",
			fields: fields{
				userAttributeName:     UserAttributeName{userAttrName: "FiledA"},
				userAttributeType:     UserAttributeType{userAttrType: "string"},
				userAttributeValidate: UserAttributeValidate{minBytes: 0, maxBytes: 2048},
				userAttributeValue:    nil,
				canUpdate:             false,
			},
			args:    args{attrValue: UserAttributeValue{userAttributeValue: "Updated Value"}},
			wantErr: false,
		},
		{
			name: "update fail: immutable",
			fields: fields{
				userAttributeName:     UserAttributeName{userAttrName: "FiledA"},
				userAttributeType:     UserAttributeType{userAttrType: "string"},
				userAttributeValidate: UserAttributeValidate{minBytes: 0, maxBytes: 2048},
				userAttributeValue:    &UserAttributeValue{userAttributeValue: "Immutable"},
				canUpdate:             false,
			},
			args:    args{attrValue: UserAttributeValue{userAttributeValue: "Fail Update Value"}},
			wantErr: true,
		},
		{
			name: "update success: un immutable",
			fields: fields{
				userAttributeName:     UserAttributeName{userAttrName: "FiledA"},
				userAttributeType:     UserAttributeType{userAttrType: "string"},
				userAttributeValidate: UserAttributeValidate{minBytes: 0, maxBytes: 2048},
				userAttributeValue:    &UserAttributeValue{userAttributeValue: "Un immutable"},
				canUpdate:             true,
			},
			args:    args{attrValue: UserAttributeValue{userAttributeValue: "Updated Value"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &UserAttributeRecord{
				userAttributeName:     tt.fields.userAttributeName,
				userAttributeType:     tt.fields.userAttributeType,
				userAttributeValidate: tt.fields.userAttributeValidate,
				userAttributeValue:    tt.fields.userAttributeValue,
				canUpdate:             tt.fields.canUpdate,
			}
			if err := a.UpdateUserAttributeValue(tt.args.attrValue); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUserAttributeValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		v      string
		slices []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not contains", args: args{v: "c", slices: []string{"a", "b"}},
			want: false,
		},
		{
			name: "contains", args: args{v: "a", slices: []string{"a", "b"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.v, tt.args.slices); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
