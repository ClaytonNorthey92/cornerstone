# cornerstone
Simple AWS MFA Authentication

# How to build and install

To build cornerstone, you will need [Go installed](https://go.dev/doc/install).

Cornerstone has been tested on Go 1.20, use that version or later to be safe.

Then you can simply run `CGO_ENABLED=0 go build -o cornerstone ./...`

That will create a cornerstone executable to your current directory.  You can then add that to your `$PATH`.


# How to use

Ensure you set the env variable `CORNERSTONE_AWS_ARN` to your IAM ARN.

Then you can simply run cornerstone like this...

```bash
cornerstone NNNNNN
```

Where `NNNNNN` is your 6-digit MFA code.

Your output will look like this:
```bash
export AWS_ACCESS_KEY_ID=...
export AWS_SECRET_ACCESS_KEY=...
export AWS_SESSION_TOKEN=...
```

For convenience, you can do this to set your env...

```
$(cornerstone NNNNNN)
```

Then you're good to go!