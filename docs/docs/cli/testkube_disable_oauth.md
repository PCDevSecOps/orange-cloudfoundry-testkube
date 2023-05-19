## testkube disable oauth

disable oauth authentication for direct api

```
testkube disable oauth [flags]
```

### Options

```
  -h, --help   help for oauth
```

### Options inherited from parent commands

```
  -a, --api-uri string     api uri, default value read from config if set (default "https://demo.testkube.io/results/v1")
  -c, --client string      client used for connecting to Testkube API one of proxy|direct (default "proxy")
      --namespace string   Kubernetes namespace, default value read from config if set (default "testkube")
      --oauth-enabled      enable oauth
      --verbose            show additional debug messages
```

### SEE ALSO

* [testkube disable](testkube_disable.md)	 - Disable feature
