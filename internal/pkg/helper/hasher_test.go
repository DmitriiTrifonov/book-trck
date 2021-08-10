package helper

import "testing"

func TestHashHelper_Encode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{str: ""},
			want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name: "password",
			args: args{str: "password"},
			want: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hh := &HashHelper{}
			if got := hh.Encode(tt.args.str); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
