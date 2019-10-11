# 2.1

Remove Bazel for testing purposes.

# 2.0

- rest.Error.StatusCode has been renamed to rest.Error.Status to match the
  change in the accepted RFC.

- rest.Client no longer has a default timeout. Use context.Context to specify
  a timeout for HTTP requests.

- Add rest.Gone for 410 responses.
