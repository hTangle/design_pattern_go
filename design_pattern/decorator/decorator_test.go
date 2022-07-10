package decorator

import "testing"

func TestEncryptionDecorator_WriteDate(t *testing.T) {
	type fields struct {
		Wrapper Source
	}
	type args struct {
		p    string
		data string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "test",
			fields: fields{Wrapper: &FileSource{}},
			args: args{
				p:    "./test.log",
				data: "hello world",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EncryptionDecorator{
				Wrapper: tt.fields.Wrapper,
			}
			if err := s.WriteDate(tt.args.p, []byte(tt.args.data)); (err != nil) != tt.wantErr {
				t.Errorf("WriteDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
