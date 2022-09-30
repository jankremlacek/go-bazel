workspace(name = "go-bazel")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "099a9fb96a376ccbbb7d291ed4ecbdfd42f6bc822ab77ae6f1b5cb9e914e94fa",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.35.0/rules_go-v0.35.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.35.0/rules_go-v0.35.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "efbbba6ac1a4fd342d5122cbdfdb82aeb2cf2862e35022c752eaddffada7c3f3",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.27.0/bazel-gazelle-v0.27.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.27.0/bazel-gazelle-v0.27.0.tar.gz",
    ],
)

http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

http_archive(
    name = "aspect_rules_js",
    sha256 = "0707a425093704fab05fb91c3a4b62cf22dca18ea334d8a72f156d4c18e8db90",
    strip_prefix = "rules_js-1.3.1",
    url = "https://github.com/aspect-build/rules_js/archive/refs/tags/v1.3.1.tar.gz",
)

http_archive(
    name = "aspect_rules_rollup",
    sha256 = "0554230e8323ec0e81d62c4fad0aea88c1b80c91483dfa76bdf71f0ad8a8fb53",
    strip_prefix = "rules_rollup-0.11.0",
    url = "https://github.com/aspect-build/rules_rollup/archive/refs/tags/v0.11.0.tar.gz",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

load("@aspect_rules_js//js:repositories.bzl", "rules_js_dependencies")
rules_js_dependencies()

load("@rules_nodejs//nodejs:repositories.bzl", "DEFAULT_NODE_VERSION", "nodejs_register_toolchains")
nodejs_register_toolchains(
    name = "nodejs",
    node_version = DEFAULT_NODE_VERSION,
)

load("@aspect_rules_js//npm:npm_import.bzl", "npm_translate_lock")
npm_translate_lock(
    name = "npm",
    npm_package_lock = "//:package-lock.json",
    package_json = "//:package.json",
    verify_node_modules_ignored = "//:.bazelignore",
)

load("@npm//:repositories.bzl", "npm_repositories")
npm_repositories()

load("@aspect_rules_rollup//rollup:repositories.bzl", "LATEST_VERSION", "rollup_register_toolchains")
rollup_register_toolchains(
    name = "rollup",
    # TODO: use the same rollup version from package.json
    rollup_version = LATEST_VERSION,
)

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

# gazelle:repository go_repository name=org_golang_x_xerrors importpath=golang.org/x/xerrors

load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.19.1")

gazelle_dependencies()

protobuf_deps()
