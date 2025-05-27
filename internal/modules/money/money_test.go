package money

import (
	"testing"
)

func TestStringToInt(t *testing.T) {
	type args struct {
		val   string
		scale int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "error - alpha",
			args:    args{val: "aaa", scale: 5},
			want:    0,
			wantErr: true,
		},
		{
			name:    "error - alpha numeric",
			args:    args{val: "111aaa", scale: 5},
			want:    0,
			wantErr: true,
		},
		{
			name:    "success",
			args:    args{val: "100.23344", scale: 5},
			want:    10_023_344,
			wantErr: false,
		},
		{
			name:    "success - int number",
			args:    args{val: "100", scale: 5},
			want:    10_000_000,
			wantErr: false,
		},
		{
			name:    "success - negative number",
			args:    args{val: "-100", scale: 5},
			want:    -10_000_000,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt(tt.args.val, tt.args.scale)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
