# Test assignement 
Determine download/upload speed using fast.com/speedtest.net

## Installation
You can clone the repo and use the `main.go` to check the full example: `git clone https://github.com/vtoma/test-speed.git`
OR
use the library in your Go code:

```golang
import (
    speed "github.com/vtoma/test-speed/lib"
)
```

## Testing & Benchmarking
```bash
go test ./...
go test -cover ./...
go test -bench=. ./...
```

## Author
Valentin Toma - valentin.toma.md@gmail.com