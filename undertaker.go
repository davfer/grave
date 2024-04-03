package grave

import (
	"context"
	"errors"
	"github.com/davfer/archit/patterns/opts"
	"github.com/davfer/crudo"
	"github.com/davfer/crudo/criteria"
	"github.com/davfer/crudo/entity"
	"github.com/go-logr/logr"
)

type tomb struct {
	logger logr.Logger
}

type Undertaker struct {
	logger logr.Logger
	repo   crudo.Repository[*Grave]
}

func (u *Undertaker) Bury(ctx context.Context, o ...opts.Opt[Grave]) {

	grave := opts.New(o...)
	match, err := u.repo.MatchOne(ctx, criteria.Attr{
		Name:       "id",
		Value:      grave.id,
		Comparison: criteria.ComparisonEq,
	})
	if errors.Is(err, entity.ErrEntityNotFound) {

	} else if err != nil {
		u.logger.Error(err, "failed retrieving graves")
		return
	}

}
