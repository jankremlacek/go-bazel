# go-bazel

A Bazel scaffold template for `golang` / `protobuf+grpc` mono repository.

## Features

- Multi `golang` microservices in the mono repo
- Single build pipeline
- Shared `golang` library
- Shared `grpc/protobuf` library
- Automatic release GitHub workflow action
    - macOS and Linux builds

## A hypothetical setup

- A cli `ServiceA` calling `ServiceB` grpc method `Sum(int32, inte32)`
- A `ServiceB` serving the `Sum` method over grpc.

### Step 1: Install Bazel

- **macOS**: [https://bazel.build/install/os-x](https://bazel.build/install/os-x)
- **Ubuntu**: [https://bazel.build/install/ubuntu](https://bazel.build/install/ubuntu)

### Step 2: Build it

Run the complete build pipeline using:

```sh
bazel build //...
```

> **Note:** The first build will take a moment. No worries, you will see the Bazel mono repo benefits later.

### Step 3: Run to see the result

In one terminal, run the grpc serving `ServiceB`:

```sh
bazel run services/serviceb :42042
```

In another terminal, run the client `ServiceA`:

```sh
bazel run services/servicea :42042
```

You should see the result. Protobuf + grpc built, services binaries built as well.
The built binaries are in the `/bazel-bin` directory in their respective sub
directories.

## Protobuf/grpc caveats

There is a little caveat with proto when using Bazel. The `golang` generated proto
files are in the `bazel-bin` build folder (as they are result of the build, not a
source). So they are not accessible to your IDE. [There is an official issue](https://github.com/bazelbuild/rules_go/issues/512).
Until solved, the workaround is:

- For every proto service:
- Run Bazel to generate the proto `golang` implementation:

  ```sh
  bazel run //proto/[service]
  ```

- Manually copy generated implementation back to the `proto/[service]` directory:

  ```sh
  cp ./bazel-bin/proto/[service]/[service]_go_proto_/github.com/kikotxyz/babykktd/proto/[service]/*.pb.go ./proto/[service]/
  ```

- Also, you have to exclude the copied file from the Bazel build. Create file `/proto/[service]/.bazelignore` and put there all generated `[filename].pb.go` files.

> **Note:** Of course, this process can be automated.

## Additonal commands

- Build `BUILD.bazel` build files using [Gazelle](https://github.com/bazelbuild/bazel-gazelle):
    ```sh
    bazel run //:gazelle
    ```
- Update Gazelle `golang` dependencies using:
    ```sh
    bazel run //:gazelle-update-repos
    ```
- Test whole pipeline using:
    ```sh
    bazel test //...
    ```

---

## Additional resources

- [Bazel official website](https://bazel.build/)
- [Bazel at GitHub](https://github.com/bazelbuild/bazel)
- [Gazelle at GitHub](https://github.com/bazelbuild/bazel-gazelle)
