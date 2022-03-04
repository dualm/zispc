package zispc

import "testing"

func Test_getRandStr(t *testing.T) {
	t.Parallel()

	type args struct {
		l int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				l: 8,
			},

			want: "",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("getRandStr() = %v", getRandStr(tt.args.l))
		})
	}
}
