# CLI 
A CLI library used to integrate command line interaction into your Go application. This is heavily influenced by [cobra cli](https://github.com/spf13/cobra). In fact, I want to use as much from that project as I can, making changes to suite my specific development needs which I will highlight along the way.

Using `cli` is easy. First, use `go get` to install the latest version
of the library.

```
go get -u github.com/rsb/cli-go@latest
```

Next, include Cobra in your application:

```go
import "github.com/rsb/cli-go"
```

# License
FuelCell is released under the MIT. See [LICENSE](https://github.com/rsb/fuelcell/blob/master/LICENSE)