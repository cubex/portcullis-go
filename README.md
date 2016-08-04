# portcullis-go
Portcullis Go Library

#### Usage
FromContext() accepts the context from your GRPC request
```go
import "github.com/fortifi/portcullis-go"

org := portcullis.FromContext(ctx).OrganisationID
```

Dependencies included in [Glide](https://glide.sh/).lock