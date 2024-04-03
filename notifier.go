package grave

import (
	"fmt"
	"github.com/go-logr/logr"
)

type Sink interface {
	Notify(Zombie) error
}

type LoggerNotifier struct {
	logger logr.Logger
	msg    string
	err    error
	level  int
	fields map[string]any
}

func (l *LoggerNotifier) Notify(z Zombie) error {
	if l.err != nil {
		l.logger.Error(l.err, l.getMessage(z), l.getFields(z)...)
	} else {
		l.logger.V(l.level).Info(l.getMessage(z), l.getFields(z)...)
	}

	return nil
}

func (l *LoggerNotifier) getFields(z Zombie) []any {
	var res []any
	for k, v := range l.fields {
		res = append(res, k, v)
	}

	return res
}

func (l *LoggerNotifier) getMessage(z Zombie) string {
	return fmt.Sprintf("zombie %s is live", z.id)
}
