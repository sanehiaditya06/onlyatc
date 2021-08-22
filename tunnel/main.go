import "github.com/jweslley/localtunnel"

...

var port := 8000
var subdomain := "onlyatc"
var tunnel := localtunnel.NewLocalTunnel(port)
var err := tunnel.OpenAs(onlyatc)
if (err != nil) {
	fmt.Printf("your url is: %s\n", tunnel.URL())
}

...

tunnel.Close()
