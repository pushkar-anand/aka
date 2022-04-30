package shell

import "testing"

func Test_cleanShellName(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{name: "test bash", arg: "/bin/bash", want: "bash"},
		{name: "test bash", arg: "/bin/zsh", want: "zsh"},
		{name: "test bash with usr", arg: "/usr/bin/bash", want: "bash"},
		{name: "test zsh with usr", arg: "/usr/bin/zsh", want: "zsh"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanShellName(tt.arg); got != tt.want {
				t.Errorf("cleanShellName() = %v, want %v", got, tt.want)
			}
		})
	}
}
