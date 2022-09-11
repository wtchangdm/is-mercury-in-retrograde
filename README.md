# Is Mercury in retrograde?
Who knows?

## Run with Docker

### Build
```bash
$ docker build -t is-mercury-in-retrograde .
```

Or use pre-build OCI image: `wtchangdm/is-mercury-in-retrograde`

### Run as CLI

```bash
$ docker run --rm -it is-mercury-in-retrograde
```

### Run as API

```bash
$ docker run --rm -it -p 8080:8080 is-mercury-in-retrograde --server
```
