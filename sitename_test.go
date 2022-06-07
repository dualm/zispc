package zispc

import "testing"

func TestGetSitename(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: "S01",
		},
		{
			name: "2",
			args: args{
				n: 1000,
			},
			want: "S1000",
		},
	}
	SetWithS()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSiteName(tt.args.n); got != tt.want {
				t.Errorf("GetSitename() = %v, want %v", got, tt.want)
			}
		})
	}
	tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "3",
			args: args{
				n: 1,
			},
			want: "001",
		},
		{
			name: "4",
			args: args{
				n: 1000,
			},
			want: "1000",
		},
	}
	UnsetWithS()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSiteName(tt.args.n); got != tt.want {
				t.Errorf("GetSitename() = %v, want %v", got, tt.want)
			}
		})
	}
}
