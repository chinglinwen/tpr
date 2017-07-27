# tpr
tcp, http(s) reverse proxy

## Usage

```
Usage of ./tpr:
  -c string
        toml config file (default "config.toml")

```

## Example complete config file

```
[services.a]
port=":80"
from="foo.com"
to="localhost:9000"

[services.b]
port=":80"
from="a.foo.com"
to="localhost:9000"

[services.example1]
port=":80"
from="e1.foo.com"
to="localhost:80"

[services.example2]
type="https"
port=":443"
from="e2.foo.com"
to="localhost:443"
```

> ** Note **
> the from value's domain needs to resolvable
> say domain foo.com need resolve to machine which run tpr.