package gen

import "io"

type marshalJsonMsgGen struct {
	passes
	p    printer
	fuse []byte
	ctx  *Context
}

func marshalJsonOrMsg(w io.Writer) *marshalJsonMsgGen {
	return &marshalJsonMsgGen{
		p: printer{w: w},
	}
}

func (m *marshalJsonMsgGen) Method() Method { return UnmarshalJsonOrMsg }

func (m *marshalJsonMsgGen) Apply(dirs []string) error {
	return nil
}

func (u *marshalJsonMsgGen) Execute(p Elem) error {
	if !u.p.ok() {
		return u.p.err
	}
	p = u.applyall(p)
	if p == nil {
		return nil
	}
	if !IsPrintable(p) {
		return nil
	}

	u.ctx = &Context{}

	u.p.comment("自動解析Json或MSG")
	u.p.printf("\nfunc (%s %s) UnmarshalJsonOrMsg(bts []byte) (o []byte, err error) {", p.Varname(), methodReceiver(p))
	u.p.printf("\n	_,msgErr := %s.UnmarshalMsg(b)", p.Varname())
	u.p.printf("\n 	if msgErr != nil {")
	u.p.printf("\n 		var json = jsoniter.ConfigCompatibleWithStandardLibrary")
	u.p.printf("\n 		jErr := json.Unmarshal(bts, z)")
	u.p.printf("\n 			if jErr != nil {")
	u.p.printf("\n 				err = jErr")
	u.p.printf("\n 			}")
	u.p.printf("\n 	}")

	u.p.print("\no = bts")
	u.p.nakedReturn()
	unsetReceiver(p)
	return u.p.err
}
