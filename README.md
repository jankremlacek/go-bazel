# go-bazel

A sample scaffold golang + Bazel mono repository.

## Features

- Multi microservices in the mono repo
- Single build pipeline

## CLI

```sh
bazel run //:gazelle
```

```sh
bazel run //:gazelle-update-repos
```

```sh
bazel test //... --test_output=errors --nocache_test_results
```

```sh
bazel build //...
```

```sh
bazel run services/a
```

```sh
bazel run services/b
```
