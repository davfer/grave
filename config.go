package grave

import (
	"github.com/davfer/archit/patterns/opts"
	"github.com/davfer/crudo/entity"
	"time"
)

type Grave struct {
	id              string
	buried          time.Time
	firstTimeAwaken time.Time
	lastTimeAwaken  time.Time
}

func (g *Grave) GetId() entity.Id {
	return entity.NewIdFromString(g.id)
}

func (g *Grave) SetId(entity.Id) error {
	return nil
}

func (g *Grave) GetResourceId() (string, error) {
	return g.id, nil
}

func (g *Grave) SetResourceId(s string) error {
	return nil
}

func (g *Grave) PreCreate() error {
	g.firstTimeAwaken = time.Now()
	return nil
}

func (g *Grave) PreUpdate() error {
	g.lastTimeAwaken = time.Now()
	return nil
}

func WithId(id string) opts.Opt[Grave] {
	return func(c Grave) Grave {
		c.id = id
		return c
	}
}

func WithDate(buried string) opts.Opt[Grave] {
	return func(c Grave) Grave {
		c.buried, _ = time.Parse(time.RFC3339, buried)

		return c
	}
}

type cemeteryConfig struct {
	sink Sink
}
