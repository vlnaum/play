package tree

import (
	"reflect"
	"testing"
)

func Test_bft(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty",
			want: nil,
		},
		{
			name: "only root",
			args: args{
				&TreeNode{Val: 1},
			},
			want: []int{1},
		},
		{
			name: "1 2 3 4",
			args: args{
				&TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bft(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dft() = %v, want %v", got, tt.want)
			}
		})
	}
}
