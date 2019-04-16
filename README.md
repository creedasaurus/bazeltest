# bazeltest

I'm having issues getting things going with Bazel in a repo at work. This repo is simply to test out the depenency with the particular issue to see if I can isolate it and find the proper way of using Bazel and bazel_go rules.

### Setup

`go mod init github.com/creedasaurus/bazeltest`
...
write a main file and add the dependencies to go mod
`go mod tidy`

Now use the basic `WORKSPACE/BUILD` file examples from [gazelle](https://github.com/bazelbuild/bazel-gazelle) get started with bazel.
 
Then:
```bash
❯ bazel version
Build label: 0.23.2-homebrew
```

```bash
❯ bazel run //:gazelle
Starting local Bazel server and connecting to it...
INFO: Analysed target //:gazelle (53 packages loaded, 6733 targets configured).
INFO: Found 1 target...
Target //:gazelle up-to-date:
  bazel-bin/gazelle-runner.bash
  bazel-bin/gazelle
INFO: Elapsed time: 22.928s, Critical Path: 5.90s
INFO: 33 processes: 32 darwin-sandbox, 1 local.
INFO: Build completed successfully, 46 total actions
INFO: Build completed successfully, 46 total actions
```
```bash
❯ bazel run //:gazelle -- update-repos -from_file=go.mod
INFO: Analysed target //:gazelle (1 packages loaded, 2 targets configured).
INFO: Found 1 target...
Target //:gazelle up-to-date:
  bazel-bin/gazelle-runner.bash
  bazel-bin/gazelle
INFO: Elapsed time: 0.232s, Critical Path: 0.00s
INFO: 0 processes.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
```

Then the problem is always:
```
❯ bazel build //...
ERROR: /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/external/com_google_cloud_go/storage/BUILD.bazel:3:1: no such package '@com_github_googleapis_gax_go//v2': BUILD file not found on package path and referenced by '@com_google_cloud_go//storage:go_default_library'
ERROR: Analysis of target '//:go_default_library' failed; build aborted: no such package '@com_github_googleapis_gax_go//v2': BUILD file not found on package path
INFO: Elapsed time: 57.803s
INFO: 0 processes.
FAILED: Build did NOT complete successfully (16 packages loaded, 1 target configured)
    Fetching @org_golang_x_net; Cloning 16b79f2e4e95ea23b2bf9903c9809ff7b013ce85 of https://go.googlesource.com/net
    Fetching @org_golang_x_oauth2; fetching
    Fetching @org_golang_google_grpc; Cloning 2fdaae294f38ed9a121193c51ec99fecd3b13eb7 of https://github.com/grpc/grpc-go
    Fetching @go_googleapis; Cloning 41d72d444fbe445f4da89e13be02078734fb7875 of https://github.com/googleapis/googleapis
```

I've also tried building from the bazel container with the same results. Some missing `BUILD` file from `@com_github_googleapis_gax_go//v2` or something.
