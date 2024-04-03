package grave

import (
	"context"
	"github.com/davfer/archit/patterns/opts"
)

type Burier[K any] interface {
	Bury(context.Context, ...opts.Opt[K])
}

var burier Burier[Grave]

func init() {
	burier = New()
}

func Get() Burier[Grave] {
	return burier
}

func New(o ...opts.Opt[Undertaker]) Undertaker {
	u := opts.New(o...)

	return &u
}

func Init(ctx context.Context, o ...opts.Opt[Undertaker]) Burier[Grave] {
	burier = New(o...)

	return burier
}
