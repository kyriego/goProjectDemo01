package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l))
}

func (u unary) String() string {
	switch u.op {
	case '+':
		return fmt.Sprintf("+%s", u.x.String())
	case '-':
		return fmt.Sprintf("-%s", u.x.String())
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) String() string {
	switch b.op {
	case '+':
		return fmt.Sprintf("%s + %s", b.x.String(), b.y.String())
	case '-':
		return fmt.Sprintf("%s - %s", b.x.String(), b.y.String())
	case '*':
		return fmt.Sprintf("%s * %s", b.x.String(), b.y.String())
	case '/':
		return fmt.Sprintf("%s / %s", b.x.String(), b.y.String())
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) String() string {
	var buf bytes.Buffer
	buf.WriteString(c.fn)
	buf.WriteByte('(')
	for _, arg := range c.args {
		buf.WriteString(arg.String())
		buf.WriteByte(',')
		buf.WriteByte(' ')
	}
	buf.Truncate(buf.Len() - 2)
	buf.WriteByte(')')
	return buf.String()
}
