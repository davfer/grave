Grave
----------------

Tool to track dead code empirically around your codebase. Despite the fact that
this tool can be seen useless in Go it can prove to be a good hand in:

- complex flow scenarios where asynchronicity is not in your favour
- having a bunch of undocumented endpoints you want to assess to remove
- tests you don't want to have covered

## WIP

- Under construction

## Installation

```bash
go get github.com/davfer/grave
```

## Usage

obscure_file.go:

```go
package somepackage

import (
	"context"
	"github.com/davfer/grave"
)

func OnObscureFlowEvent() {
	grave.Get().Bury(context.TODO(), grave.WithDate("2024-01-01"), grave.WithId("undesired-flow"))
}
```

