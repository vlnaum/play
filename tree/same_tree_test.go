package tree

import "testing"

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *TreeNode
		q *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			want: true,
		},
		{
			name: "one empty",
			args: args{
				p: &TreeNode{Val: 1},
			},
			want: false,
		},
		{
			name: "only root eq",
			args: args{
				p: &TreeNode{Val: 1},
				q: &TreeNode{Val: 1},
			},
			want: true,
		},
		{
			name: "only root ne",
			args: args{
				p: &TreeNode{Val: 1},
				q: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			name: "complex eq",
			args: args{
				p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
				q: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			},
			want: true,
		},
		{
			name: "complex ne (diff leafs location)",
			args: args{
				p: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
				q: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
