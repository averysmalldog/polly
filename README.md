> # Note: `polly` is being deprecated
> Turns out, sometimes one repo is better than two. I created this repo as an exercise for what I imagined might be some kind of grand, sprawling heap of complex functionality down the line. That's not exactly what happened. Check out the release note for averysmalldog/tesla-gen3wc-monitor.
> 
> Long story short, I'm going to put this repo in read-only mode (archive it) since it is essentially a historical artifact.

# polly
A simple Go module for periodically GET-ing from an API (For now, a Tesla Gen 3 Wall Connector) and dumping the output into a datastore (for now, InfluxDB).

## Prerequisites

_Note: these prerequisites will be reduced if and when I genericize this tool._

- You will need a local instance of InfluxDB running on `localhost:8086` with a database called `tesla` that has no auth.
- You will need a [Tesla HPWC Gen3](https://shop.tesla.com/product/wall-connector) installed, provisioned, and joined to your home (or business) wifi.
- You will need to be able to reach the HPWC from whatever computer you run this code on.

## Using polly

`polly` is a Go module. You need to import it into your Go code to use it. `polly`'s only input is the environment variable `HPWC_IP`, which you will need to have set in the environment where it is called.

Here's a quick example of how you can run a very simple program using the `polly` package:

```go
package main

import "github.com/averysmalldog/polly"

func main() {
	polly.Execute()
}
```

If you were to save that as `main.go`, you would then follow these steps to get things working:

1. Set the environment `HPWC_IP` to be the IP address of your Tesla Gen 3 Wall Connector, e.g. `export HPWC_IP="192.168.1.10"`.
2. Run it, e.g. `go run main.go`.

You will see one dot for each successful GET against the Wall Connector, as you can see in the example below.

```shell
% go run main.go
...........
```

If `polly` can't reach your Gen 3 Wall Connector, you'll get this error for each time the GET fails:

```shell
error - during GET of hpwc. Do you have the right IP?
```
