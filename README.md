# portcullis-go
Portcullis Go Library

#### Usage
FromContext() accepts the context from your GRPC request
```go
import "github.com/kubex/portcullis-go"

project := portcullis.FromContext(ctx).ProjectID
```

Dependencies included in [Glide](https://glide.sh/).lock