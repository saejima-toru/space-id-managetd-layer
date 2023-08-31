package attributes

import (
	"identity-management/domains/id-store/values/users/attributes/attribute"
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
		want    *attribute.UserAttributeName
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
			got, err := attribute.NewUserAttributeName(tt.args.userAttrName)
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
func silentNewUserName(userName string) *attribute.UserAttributeName {
	u, _ := attribute.NewUserAttributeName(userName)
	return u
}

func TestNewUserAttributeType(t *testing.T) {
	type args struct {
		userAttrType string
	}
	tests := []struct {
		name    string
		args    args
		want    *attribute.UserAttributeType
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
			got, err := attribute.NewUserAttributeType(tt.args.userAttrType)
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
func silentNewUserAttributeType(userAttrType string) *attribute.UserAttributeType {
	u, _ := attribute.NewUserAttributeType(userAttrType)
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
		want    *attribute.UserAttributeValidate
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
			got, err := attribute.NewUserAttributeValidate(&tt.args.minBytes, &tt.args.maxBytes)
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
func silentNewUserAttributeValidate(minBytes int, maxBytes int) *attribute.UserAttributeValidate {
	u, _ := attribute.NewUserAttributeValidate(&minBytes, &maxBytes)
	return u
}

func TestUserAttributes_contains(t *testing.T) {
	type fields struct {
		userAttributeRecords []attribute.UserAttributeRecord
	}
	type args struct {
		mixed attribute.UserAttributeRecord
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
		userAttributeRecords []attribute.UserAttributeRecord
	}
	type args struct {
		findAttrName attribute.UserAttributeName
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *attribute.UserAttributeRecord
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
		userAttributeRecords []attribute.UserAttributeRecord
	}
	type args struct {
		mixed attribute.UserAttributeRecord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "not found",
			fields: fields{
				userAttributeRecords: []attribute.UserAttributeRecord{
					*silentCreateAttributeRecord(
						"attrNameA", "string", nil, false,
					),
					*silentCreateAttributeRecord(
						"attrNameB", "string", nil, false,
					),
				},
			},
			args: args{mixed: *silentCreateAttributeRecord(
				"attrNameC", "string", nil, false,
			)},
			want: -1,
		},
		{
			name: "founded",
			fields: fields{
				userAttributeRecords: []attribute.UserAttributeRecord{
					*silentCreateAttributeRecord(
						"attrNameA", "string", nil, false,
					),
					*silentCreateAttributeRecord(
						"attrNameB", "string", nil, false,
					),
				},
			},
			args: args{mixed: *silentCreateAttributeRecord(
				"attrNameB", "string", nil, false,
			)},
			want: 1,
		},
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

func TestUserAttributes_AddUserAttributeRecord(t *testing.T) {
	type fields struct {
		userAttributeRecords []attribute.UserAttributeRecord
	}
	type args struct {
		addAttrRecord attribute.UserAttributeRecord
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
				userAttributeRecords: []attribute.UserAttributeRecord{
					*silentCreateAttributeRecord(
						"FiledA", "string", nil, false,
					),
				},
			},
			args: args{
				addAttrRecord: *silentCreateAttributeRecord(
					"FiledB", "string", nil, false,
				),
			},
			want: &UserAttributes{
				userAttributeRecords: []attribute.UserAttributeRecord{
					*silentCreateAttributeRecord(
						"FiledA", "string", nil, false,
					),
					*silentCreateAttributeRecord(
						"FiledB", "string", nil, false,
					),
				},
			},
			wantErr: false,
		},
		{
			name: "fail added",
			fields: fields{
				userAttributeRecords: []attribute.UserAttributeRecord{
					*silentCreateAttributeRecord(
						"FiledA", "string", nil, false,
					),
				},
			},
			args: args{
				addAttrRecord: *silentCreateAttributeRecord(
					"FiledA", "string", nil, false,
				),
			},
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

func silentCreateAttributeRecord(
	attrName string,
	attrType string,
	attrValue *string,
	canUpdate bool,
) *attribute.UserAttributeRecord {
	var setAttrValue *attribute.UserAttributeValue
	setAttrName := silentNewUserName(attrName)
	setAttrType := silentNewUserAttributeType(attrType)
	setValidate := silentNewUserAttributeValidate(0, 2048)

	setAttrValue = nil
	if attrValue != nil {
		setAttrValue = attribute.NewUserAttributeValue(*attrValue)
	}
	return attribute.NewUserAttributeRecord(
		*setAttrName, *setAttrType, setAttrValue, *setValidate, canUpdate)
}

func Test_duplicate(t *testing.T) {
	type args struct {
		records []attribute.UserAttributeRecord
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := duplicate(tt.args.records); got != tt.want {
				t.Errorf("duplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAttributes_MargeUserAttribute(t *testing.T) {
	type fields struct {
		userAttributeRecords []attribute.UserAttributeRecord
	}
	type args struct {
		userAttrs []attribute.UserAttributeRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *UserAttributes
		wantErr bool
	}{
		{
			name: "all marge complete",
			fields: fields{
				[]attribute.UserAttributeRecord{
					*silentCreateAttributeRecord("attrNameA", "string", nil, true),
					*silentCreateAttributeRecord("attrNameB", "string", nil, true),
				},
			},
			args: args{
				userAttrs: []attribute.UserAttributeRecord{
					*silentCreateAttributeRecord("attrNameA", "string", tString("margeA"), true),
				},
			},
			want: &UserAttributes{
				userAttributeRecords: []attribute.UserAttributeRecord{
					*silentCreateAttributeRecord("attrNameA", "string", tString("margeA"), true),
					*silentCreateAttributeRecord("attrNameB", "string", nil, true),
				},
			},
			wantErr: false,
		},
		{
			name: "not defined error",
			fields: fields{
				[]attribute.UserAttributeRecord{
					*silentCreateAttributeRecord("attrNameA", "string", nil, true),
					*silentCreateAttributeRecord("attrNameB", "string", nil, true),
				},
			},
			args: args{
				userAttrs: []attribute.UserAttributeRecord{
					*silentCreateAttributeRecord("attrNameC", "string", tString("margeC"), true),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserAttributes{
				userAttributeRecords: tt.fields.userAttributeRecords,
			}
			got, err := u.MargeUserAttribute(tt.args.userAttrs)
			if (err != nil) != tt.wantErr {
				t.Errorf("MargeUserAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MargeUserAttribute() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func tString(v string) *string {
	return &v
}
