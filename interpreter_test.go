package bf

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestInterpreter_Run(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name       string
		i          *Interpreter
		args       args
		wantOutput string
	}{
		{"Hello World Program",
			NewInterpreter("+[-[<<[+[--->]-[<<<]]]>>>-]>-.---.>..>.<<<<-.<+.>>>>>.>.<<.<-."),
			args{&bytes.Reader{}},
			"hello world"},
		{"Simon Says Program",
			NewInterpreter("+[>,.<]"),
			args{strings.NewReader("simon says\n")},
			"simon says\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			tt.i.Run(tt.args.input, output)
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("Interpreter.Run() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
