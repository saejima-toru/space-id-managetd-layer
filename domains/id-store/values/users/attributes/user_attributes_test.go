package attributes

import (
	"reflect"
	"testing"
)

func TestNewUserAttributeName(t *testing.T) {
	type args struct {
		userAttrName string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserAttributeName
		wantErr bool
	}{
		{name: "boundary test: success",
			args: args{userAttrName: "A"}, wantErr: false, want: silentNewUserName("A")},
		{name: "boundary minimum: failed",
			args: args{userAttrName: ""}, wantErr: true, want: nil},
		{name: "boundary over: failed",
			args: args{userAttrName: "AAAAAAAAAAAAAAAAAAAAA"}, wantErr: true, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserAttributeName(tt.args.userAttrName)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserAttributeName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserAttributeName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func silentNewUserName(userName string) *UserAttributeName {
	u, _ := NewUserAttributeName(userName)
	return u
}

func TestNewUserAttributeType(t *testing.T) {
	type args struct {
		userAttrType string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserAttributeType
		wantErr bool
	}{
		{
			name: "allowed boundary type :success",
			args: args{userAttrType: "string"}, want: silentNewUserAttributeType("string"),
			wantErr: false},
		{
			name: "allowed boundary type :success",
			args: args{userAttrType: "number"}, want: silentNewUserAttributeType("number"),
			wantErr: false},
		{
			name: "denied boundary type :fail",
			args: args{userAttrType: "array"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserAttributeType(tt.args.userAttrType)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserAttributeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserAttributeType() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func silentNewUserAttributeType(userAttrType string) *UserAttributeType {
	u, _ := NewUserAttributeType(userAttrType)
	return u
}

func TestNewUserAttributeValidate(t *testing.T) {
	type args struct {
		minBytes int
		maxBytes int
	}
	tests := []struct {
		name    string
		args    args
		want    *UserAttributeValidate
		wantErr bool
	}{
		{
			name: "minBytes in minimum boundary: success", args: args{minBytes: 0, maxBytes: 2048},
			want:    silentNewUserAttributeValidate(0, 2048),
			wantErr: false,
		},
		{
			name: "minBytes in max boundary: success", args: args{minBytes: 2048, maxBytes: 2048},
			want:    silentNewUserAttributeValidate(2048, 2048),
			wantErr: false,
		},
		{
			name: "minBytes in minimum boundary: fail", args: args{minBytes: -1, maxBytes: 2048},
			want:    silentNewUserAttributeValidate(-1, 2048),
			wantErr: true,
		},
		{
			name: "minBytes in over boundary: fail", args: args{minBytes: 2049, maxBytes: 2048},
			want:    silentNewUserAttributeValidate(2049, 2048),
			wantErr: true,
		},
		// Max
		{
			name: "maxBytes in minimum boundary: success", args: args{minBytes: 0, maxBytes: 0},
			want:    silentNewUserAttributeValidate(0, 0),
			wantErr: false,
		},
		{
			name: "maxBytes in max boundary: success", args: args{minBytes: 0, maxBytes: 2048},
			want:    silentNewUserAttributeValidate(0, 2048),
			wantErr: false,
		},
		{
			name: "maxBytes in minimum boundary: fail", args: args{minBytes: 0, maxBytes: -1},
			want:    silentNewUserAttributeValidate(0, -1),
			wantErr: true,
		},
		{
			name: "maxBytes in over boundary: fail", args: args{minBytes: 0, maxBytes: 2049},
			want:    silentNewUserAttributeValidate(0, 2049),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUserAttributeValidate(&tt.args.minBytes, &tt.args.maxBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUserAttributeValidate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserAttributeValidate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func silentNewUserAttributeValidate(minBytes int, maxBytes int) *UserAttributeValidate {
	u, _ := NewUserAttributeValidate(&minBytes, &maxBytes)
	return u
}

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

func TestUserAttributes_AddUserAttributeRecord(t *testing.T) {
	type fields struct {
		userAttributeRecords []UserAttributeRecord
	}
	type args struct {
		addAttrRecord UserAttributeRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UserAttributes
		wantErr bool
	}{
		{
			name: "success added",
			fields: fields{
				userAttributeRecords: []UserAttributeRecord{
					{
						userAttributeName:     UserAttributeName{userAttrName: "FiledA"},
						userAttributeType:     UserAttributeType{userAttrType: "string"},
						userAttributeValidate: UserAttributeValidate{minBytes: 0, maxBytes: 2048},
						userAttributeValue:    nil,
						canUpdate:             false,
					},
				},
			},
			args: args{addAttrRecord: UserAttributeRecord{
				userAttributeName:     UserAttributeName{userAttrName: "FiledB"},
				userAttributeType:     UserAttributeType{userAttrType: "string"},
				userAttributeValidate: UserAttributeValidate{minBytes: 0, maxBytes: 2048},
				userAttributeValue:    nil,
				canUpdate:             false,
			}},
			want: &UserAttributes{
				userAttributeRecords: []UserAttributeRecord{
					{
						userAttributeName:     UserAttributeName{userAttrName: "FiledA"},
						userAttributeType:     UserAttributeType{userAttrType: "string"},
						userAttributeValidate: UserAttributeValidate{minBytes: 0, maxBytes: 2048},
						userAttributeValue:    nil,
						canUpdate:             false,
					},
					{
						userAttributeName:     UserAttributeName{userAttrName: "FiledB"},
						userAttributeType:     UserAttributeType{userAttrType: "string"},
						userAttributeValidate: UserAttributeValidate{minBytes: 0, maxBytes: 2048},
						userAttributeValue:    nil,
						canUpdate:             false,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "fail added",
			fields: fields{
				userAttributeRecords: []UserAttributeRecord{
					{
						userAttributeName:     UserAttributeName{userAttrName: "FiledA"},
						userAttributeType:     UserAttributeType{userAttrType: "string"},
						userAttributeValidate: UserAttributeValidate{minBytes: 0, maxBytes: 2048},
						userAttributeValue:    nil,
						canUpdate:             false,
					},
				},
			},
			args: args{addAttrRecord: UserAttributeRecord{
				userAttributeName:     UserAttributeName{userAttrName: "FiledA"},
				userAttributeType:     UserAttributeType{userAttrType: "string"},
				userAttributeValidate: UserAttributeValidate{minBytes: 0, maxBytes: 2048},
				userAttributeValue:    nil,
				canUpdate:             false,
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAttributes{
				userAttributeRecords: tt.fields.userAttributeRecords,
			}
			got, err := u.AddUserAttributeRecord(tt.args.addAttrRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUserAttributeRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUserAttributeRecord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAttributes_contains(t *testing.T) {
	type fields struct {
		userAttributeRecords []UserAttributeRecord
	}
	type args struct {
		mixed UserAttributeRecord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAttributes{
				userAttributeRecords: tt.fields.userAttributeRecords,
			}
			if got := u.contains(tt.args.mixed); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAttributes_findByAttrName(t *testing.T) {
	type fields struct {
		userAttributeRecords []UserAttributeRecord
	}
	type args struct {
		findAttrName UserAttributeName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UserAttributeRecord
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAttributes{
				userAttributeRecords: tt.fields.userAttributeRecords,
			}
			if got := u.findByAttrName(tt.args.findAttrName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findByAttrName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAttributes_indexOf(t *testing.T) {
	type fields struct {
		userAttributeRecords []UserAttributeRecord
	}
	type args struct {
		mixed UserAttributeRecord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAttributes{
				userAttributeRecords: tt.fields.userAttributeRecords,
			}
			if got := u.indexOf(tt.args.mixed); got != tt.want {
				t.Errorf("indexOf() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.v, tt.args.slices); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_duplicate(t *testing.T) {
	type args struct {
		records []UserAttributeRecord
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := duplicate(tt.args.records); got != tt.want {
				t.Errorf("duplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
