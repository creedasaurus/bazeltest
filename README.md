# bazeltest

I tried using the master branch of gazelle, which help to sort of some of the missing dependecy issues, but now I've got a strange issue when it comes to using the stackdriver exporter.

### Setup
Everything would build just fine until I did the following:

Adding `stackdriver`
```go
  _ "contrib.go.opencensus.io/exporter/stackdriver"
)
```
Then running to get the deps in the mod file
`go mod tidy`

Then:
```bash
❯ bazel run //:gazelle -- update-repos -from_file=go.mod
INFO: Analysed target //:gazelle (0 packages loaded, 2 targets configured).
INFO: Found 1 target...
Target //:gazelle up-to-date:
  bazel-bin/gazelle-runner.bash
  bazel-bin/gazelle
INFO: Elapsed time: 0.220s, Critical Path: 0.00s
INFO: 0 processes.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
```
And

```bash
❯ bazel run //:gazelle -- update
INFO: Analysed target //:gazelle (2 packages loaded, 110 targets configured).
INFO: Found 1 target...
Target //:gazelle up-to-date:
  bazel-bin/gazelle-runner.bash
  bazel-bin/gazelle
INFO: Elapsed time: 1.785s, Critical Path: 0.01s
INFO: 0 processes.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
```

Then the problem is:
```bash
❯ bazel build //...
INFO: Analysed 2 targets (0 packages loaded, 0 targets configured).
INFO: Found 2 targets...
ERROR: /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/BUILD.bazel:3:1: GoCompile external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/darwin_amd64_stripped/go_default_library%/github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1.a failed (Exit 1) builder failed: error executing command bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix darwin_amd64 -src external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/metrics.pb.go -arc ... (remaining 14 argument(s) skipped)

Use --sandbox_debug to see verbose messages from the sandbox
compile: missing strict dependencies:
	/private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/sandbox/darwin-sandbox/1/execroot/__main__/external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/metrics.pb.go: import of "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
Known dependencies are:
	github.com/golang/protobuf/proto
	github.com/golang/protobuf/ptypes/timestamp
	github.com/golang/protobuf/ptypes/wrappers
Check that imports in Go sources match importpath attributes in deps.
INFO: Elapsed time: 0.722s, Critical Path: 0.41s
INFO: 0 processes.
FAILED: Build did NOT complete successfully
```

Adding the `-s and --sandbox_debug` flags:
```bash
~/workspace/random/bazeltest master
❯ bazel build -s --sandbox_debug //...
INFO: Analysed 2 targets (0 packages loaded, 0 targets configured).
INFO: Found 2 targets...
SUBCOMMAND: # @go_googleapis//google/container/v1:container_go_proto [action 'GoCompile external/go_googleapis/google/container/v1/darwin_amd64_stripped/container_go_proto%/google.golang.org/genproto/googleapis/container/v1.a']
(cd /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/execroot/__main__ && \
  exec env - \
    APPLE_SDK_PLATFORM=MacOSX \
    APPLE_SDK_VERSION_OVERRIDE=10.14 \
    CGO_ENABLED=1 \
    GOARCH=amd64 \
    GOOS=darwin \
    GOROOT=external/go_sdk \
    GOROOT_FINAL=GOROOT \
    PATH=external/local_config_cc:/bin:/usr/bin \
    XCODE_VERSION_OVERRIDE=10.2.0 \
  bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix darwin_amd64 -src bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/container/v1/darwin_amd64_stripped/container_go_proto%/google.golang.org/genproto/googleapis/container/v1/cluster_service.pb.go -arc 'google.golang.org/genproto/googleapis/api/annotations=google.golang.org/genproto/googleapis/api/annotations=bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/annotations_go_proto%/google.golang.org/genproto/googleapis/api/annotations.a=' -arc 'github.com/golang/protobuf/proto=github.com/golang/protobuf/proto=bazel-out/darwin-fastbuild/bin/external/com_github_golang_protobuf/proto/darwin_amd64_stripped/go_default_library%/github.com/golang/protobuf/proto.a=' -arc 'google.golang.org/grpc=google.golang.org/grpc=bazel-out/darwin-fastbuild/bin/external/org_golang_google_grpc/darwin_amd64_stripped/go_default_library%/google.golang.org/grpc.a=' -arc 'golang.org/x/net/context=golang.org/x/net/context=bazel-out/darwin-fastbuild/bin/external/org_golang_x_net/context/darwin_amd64_stripped/go_default_library%/golang.org/x/net/context.a=' -arc 'github.com/golang/protobuf/ptypes/any=github.com/golang/protobuf/ptypes/any=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/any_go_proto%/github.com/golang/protobuf/ptypes/any.a=' -arc 'google.golang.org/genproto/protobuf/api=google.golang.org/genproto/protobuf/api=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/api_go_proto%/google.golang.org/genproto/protobuf/api.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/plugin=github.com/golang/protobuf/protoc-gen-go/plugin=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/compiler_plugin_go_proto%/github.com/golang/protobuf/protoc-gen-go/plugin.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/descriptor=github.com/golang/protobuf/protoc-gen-go/descriptor=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/descriptor_go_proto%/github.com/golang/protobuf/protoc-gen-go/descriptor.a=' -arc 'github.com/golang/protobuf/ptypes/duration=github.com/golang/protobuf/ptypes/duration=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/duration_go_proto%/github.com/golang/protobuf/ptypes/duration.a=' -arc 'github.com/golang/protobuf/ptypes/empty=github.com/golang/protobuf/ptypes/empty=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/empty_go_proto%/github.com/golang/protobuf/ptypes/empty.a=' -arc 'google.golang.org/genproto/protobuf/field_mask=google.golang.org/genproto/protobuf/field_mask=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/field_mask_go_proto%/google.golang.org/genproto/protobuf/field_mask.a=' -arc 'google.golang.org/genproto/protobuf/source_context=google.golang.org/genproto/protobuf/source_context=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/source_context_go_proto%/google.golang.org/genproto/protobuf/source_context.a=' -arc 'github.com/golang/protobuf/ptypes/struct=github.com/golang/protobuf/ptypes/struct=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/struct_go_proto%/github.com/golang/protobuf/ptypes/struct.a=' -arc 'github.com/golang/protobuf/ptypes/timestamp=github.com/golang/protobuf/ptypes/timestamp=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/timestamp_go_proto%/github.com/golang/protobuf/ptypes/timestamp.a=' -arc 'google.golang.org/genproto/protobuf/ptype=google.golang.org/genproto/protobuf/ptype=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/type_go_proto%/google.golang.org/genproto/protobuf/ptype.a=' -arc 'github.com/golang/protobuf/ptypes/wrappers=github.com/golang/protobuf/ptypes/wrappers=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/wrappers_go_proto%/github.com/golang/protobuf/ptypes/wrappers.a=' -o bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/container/v1/darwin_amd64_stripped/container_go_proto%/google.golang.org/genproto/googleapis/container/v1.a -package_list bazel-out/host/bin/external/go_sdk/packages.txt -p google.golang.org/genproto/googleapis/container/v1 -- -trimpath .)
SUBCOMMAND: # @go_googleapis//google/api:monitoredres_go_proto [action 'GoCompile external/go_googleapis/google/api/darwin_amd64_stripped/monitoredres_go_proto%/google.golang.org/genproto/googleapis/api/monitoredres.a']
(cd /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/execroot/__main__ && \
  exec env - \
    APPLE_SDK_PLATFORM=MacOSX \
    APPLE_SDK_VERSION_OVERRIDE=10.14 \
    CGO_ENABLED=1 \
    GOARCH=amd64 \
    GOOS=darwin \
    GOROOT=external/go_sdk \
    GOROOT_FINAL=GOROOT \
    PATH=external/local_config_cc:/bin:/usr/bin \
    XCODE_VERSION_OVERRIDE=10.2.0 \
  bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix darwin_amd64 -src bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/monitoredres_go_proto%/google.golang.org/genproto/googleapis/api/monitoredres/monitored_resource.pb.go -arc 'google.golang.org/genproto/googleapis/api/label=google.golang.org/genproto/googleapis/api/label=bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/label_go_proto%/google.golang.org/genproto/googleapis/api/label.a=' -arc 'github.com/golang/protobuf/proto=github.com/golang/protobuf/proto=bazel-out/darwin-fastbuild/bin/external/com_github_golang_protobuf/proto/darwin_amd64_stripped/go_default_library%/github.com/golang/protobuf/proto.a=' -arc 'github.com/golang/protobuf/ptypes/any=github.com/golang/protobuf/ptypes/any=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/any_go_proto%/github.com/golang/protobuf/ptypes/any.a=' -arc 'google.golang.org/genproto/protobuf/api=google.golang.org/genproto/protobuf/api=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/api_go_proto%/google.golang.org/genproto/protobuf/api.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/plugin=github.com/golang/protobuf/protoc-gen-go/plugin=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/compiler_plugin_go_proto%/github.com/golang/protobuf/protoc-gen-go/plugin.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/descriptor=github.com/golang/protobuf/protoc-gen-go/descriptor=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/descriptor_go_proto%/github.com/golang/protobuf/protoc-gen-go/descriptor.a=' -arc 'github.com/golang/protobuf/ptypes/duration=github.com/golang/protobuf/ptypes/duration=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/duration_go_proto%/github.com/golang/protobuf/ptypes/duration.a=' -arc 'github.com/golang/protobuf/ptypes/empty=github.com/golang/protobuf/ptypes/empty=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/empty_go_proto%/github.com/golang/protobuf/ptypes/empty.a=' -arc 'google.golang.org/genproto/protobuf/field_mask=google.golang.org/genproto/protobuf/field_mask=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/field_mask_go_proto%/google.golang.org/genproto/protobuf/field_mask.a=' -arc 'google.golang.org/genproto/protobuf/source_context=google.golang.org/genproto/protobuf/source_context=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/source_context_go_proto%/google.golang.org/genproto/protobuf/source_context.a=' -arc 'github.com/golang/protobuf/ptypes/struct=github.com/golang/protobuf/ptypes/struct=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/struct_go_proto%/github.com/golang/protobuf/ptypes/struct.a=' -arc 'github.com/golang/protobuf/ptypes/timestamp=github.com/golang/protobuf/ptypes/timestamp=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/timestamp_go_proto%/github.com/golang/protobuf/ptypes/timestamp.a=' -arc 'google.golang.org/genproto/protobuf/ptype=google.golang.org/genproto/protobuf/ptype=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/type_go_proto%/google.golang.org/genproto/protobuf/ptype.a=' -arc 'github.com/golang/protobuf/ptypes/wrappers=github.com/golang/protobuf/ptypes/wrappers=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/wrappers_go_proto%/github.com/golang/protobuf/ptypes/wrappers.a=' -o bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/monitoredres_go_proto%/google.golang.org/genproto/googleapis/api/monitoredres.a -package_list bazel-out/host/bin/external/go_sdk/packages.txt -p google.golang.org/genproto/googleapis/api/monitoredres -- -trimpath .)
SUBCOMMAND: # @io_opencensus_go//plugin/ocgrpc:go_default_library [action 'GoCompile external/io_opencensus_go/plugin/ocgrpc/darwin_amd64_stripped/go_default_library%/go.opencensus.io/plugin/ocgrpc.a']
(cd /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/execroot/__main__ && \
  exec env - \
    APPLE_SDK_PLATFORM=MacOSX \
    APPLE_SDK_VERSION_OVERRIDE=10.14 \
    CGO_ENABLED=1 \
    GOARCH=amd64 \
    GOOS=darwin \
    GOROOT=external/go_sdk \
    GOROOT_FINAL=GOROOT \
    PATH=external/local_config_cc:/bin:/usr/bin \
    XCODE_VERSION_OVERRIDE=10.2.0 \
  bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix darwin_amd64 -src external/io_opencensus_go/plugin/ocgrpc/client.go -src external/io_opencensus_go/plugin/ocgrpc/client_metrics.go -src external/io_opencensus_go/plugin/ocgrpc/client_stats_handler.go -src external/io_opencensus_go/plugin/ocgrpc/doc.go -src external/io_opencensus_go/plugin/ocgrpc/server.go -src external/io_opencensus_go/plugin/ocgrpc/server_metrics.go -src external/io_opencensus_go/plugin/ocgrpc/server_stats_handler.go -src external/io_opencensus_go/plugin/ocgrpc/stats_common.go -src external/io_opencensus_go/plugin/ocgrpc/trace_common.go -arc 'go.opencensus.io/stats=go.opencensus.io/stats=bazel-out/darwin-fastbuild/bin/external/io_opencensus_go/stats/darwin_amd64_stripped/go_default_library%/go.opencensus.io/stats.a=' -arc 'go.opencensus.io/stats/view=go.opencensus.io/stats/view=bazel-out/darwin-fastbuild/bin/external/io_opencensus_go/stats/view/darwin_amd64_stripped/go_default_library%/go.opencensus.io/stats/view.a=' -arc 'go.opencensus.io/tag=go.opencensus.io/tag=bazel-out/darwin-fastbuild/bin/external/io_opencensus_go/tag/darwin_amd64_stripped/go_default_library%/go.opencensus.io/tag.a=' -arc 'go.opencensus.io/trace=go.opencensus.io/trace=bazel-out/darwin-fastbuild/bin/external/io_opencensus_go/trace/darwin_amd64_stripped/go_default_library%/go.opencensus.io/trace.a=' -arc 'go.opencensus.io/trace/propagation=go.opencensus.io/trace/propagation=bazel-out/darwin-fastbuild/bin/external/io_opencensus_go/trace/propagation/darwin_amd64_stripped/go_default_library%/go.opencensus.io/trace/propagation.a=' -arc 'google.golang.org/grpc/codes=google.golang.org/grpc/codes=bazel-out/darwin-fastbuild/bin/external/org_golang_google_grpc/codes/darwin_amd64_stripped/go_default_library%/google.golang.org/grpc/codes.a=' -arc 'google.golang.org/grpc/grpclog=google.golang.org/grpc/grpclog=bazel-out/darwin-fastbuild/bin/external/org_golang_google_grpc/grpclog/darwin_amd64_stripped/go_default_library%/google.golang.org/grpc/grpclog.a=' -arc 'google.golang.org/grpc/metadata=google.golang.org/grpc/metadata=bazel-out/darwin-fastbuild/bin/external/org_golang_google_grpc/metadata/darwin_amd64_stripped/go_default_library%/google.golang.org/grpc/metadata.a=' -arc 'google.golang.org/grpc/stats=google.golang.org/grpc/stats=bazel-out/darwin-fastbuild/bin/external/org_golang_google_grpc/stats/darwin_amd64_stripped/go_default_library%/google.golang.org/grpc/stats.a=' -arc 'google.golang.org/grpc/status=google.golang.org/grpc/status=bazel-out/darwin-fastbuild/bin/external/org_golang_google_grpc/status/darwin_amd64_stripped/go_default_library%/google.golang.org/grpc/status.a=' -arc 'golang.org/x/net/context=golang.org/x/net/context=bazel-out/darwin-fastbuild/bin/external/org_golang_x_net/context/darwin_amd64_stripped/go_default_library%/golang.org/x/net/context.a=' -o bazel-out/darwin-fastbuild/bin/external/io_opencensus_go/plugin/ocgrpc/darwin_amd64_stripped/go_default_library%/go.opencensus.io/plugin/ocgrpc.a -package_list bazel-out/host/bin/external/go_sdk/packages.txt -p go.opencensus.io/plugin/ocgrpc -- -trimpath .)
SUBCOMMAND: # @go_googleapis//google/api:metric_go_proto [action 'GoCompile external/go_googleapis/google/api/darwin_amd64_stripped/metric_go_proto%/google.golang.org/genproto/googleapis/api/metric.a']
(cd /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/execroot/__main__ && \
  exec env - \
    APPLE_SDK_PLATFORM=MacOSX \
    APPLE_SDK_VERSION_OVERRIDE=10.14 \
    CGO_ENABLED=1 \
    GOARCH=amd64 \
    GOOS=darwin \
    GOROOT=external/go_sdk \
    GOROOT_FINAL=GOROOT \
    PATH=external/local_config_cc:/bin:/usr/bin \
    XCODE_VERSION_OVERRIDE=10.2.0 \
  bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix darwin_amd64 -src bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/metric_go_proto%/google.golang.org/genproto/googleapis/api/metric/metric.pb.go -arc 'google.golang.org/genproto/googleapis/api=google.golang.org/genproto/googleapis/api=bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/api_go_proto%/google.golang.org/genproto/googleapis/api.a=' -arc 'google.golang.org/genproto/googleapis/api/label=google.golang.org/genproto/googleapis/api/label=bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/label_go_proto%/google.golang.org/genproto/googleapis/api/label.a=' -arc 'github.com/golang/protobuf/proto=github.com/golang/protobuf/proto=bazel-out/darwin-fastbuild/bin/external/com_github_golang_protobuf/proto/darwin_amd64_stripped/go_default_library%/github.com/golang/protobuf/proto.a=' -arc 'github.com/golang/protobuf/ptypes/any=github.com/golang/protobuf/ptypes/any=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/any_go_proto%/github.com/golang/protobuf/ptypes/any.a=' -arc 'google.golang.org/genproto/protobuf/api=google.golang.org/genproto/protobuf/api=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/api_go_proto%/google.golang.org/genproto/protobuf/api.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/plugin=github.com/golang/protobuf/protoc-gen-go/plugin=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/compiler_plugin_go_proto%/github.com/golang/protobuf/protoc-gen-go/plugin.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/descriptor=github.com/golang/protobuf/protoc-gen-go/descriptor=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/descriptor_go_proto%/github.com/golang/protobuf/protoc-gen-go/descriptor.a=' -arc 'github.com/golang/protobuf/ptypes/duration=github.com/golang/protobuf/ptypes/duration=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/duration_go_proto%/github.com/golang/protobuf/ptypes/duration.a=' -arc 'github.com/golang/protobuf/ptypes/empty=github.com/golang/protobuf/ptypes/empty=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/empty_go_proto%/github.com/golang/protobuf/ptypes/empty.a=' -arc 'google.golang.org/genproto/protobuf/field_mask=google.golang.org/genproto/protobuf/field_mask=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/field_mask_go_proto%/google.golang.org/genproto/protobuf/field_mask.a=' -arc 'google.golang.org/genproto/protobuf/source_context=google.golang.org/genproto/protobuf/source_context=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/source_context_go_proto%/google.golang.org/genproto/protobuf/source_context.a=' -arc 'github.com/golang/protobuf/ptypes/struct=github.com/golang/protobuf/ptypes/struct=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/struct_go_proto%/github.com/golang/protobuf/ptypes/struct.a=' -arc 'github.com/golang/protobuf/ptypes/timestamp=github.com/golang/protobuf/ptypes/timestamp=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/timestamp_go_proto%/github.com/golang/protobuf/ptypes/timestamp.a=' -arc 'google.golang.org/genproto/protobuf/ptype=google.golang.org/genproto/protobuf/ptype=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/type_go_proto%/google.golang.org/genproto/protobuf/ptype.a=' -arc 'github.com/golang/protobuf/ptypes/wrappers=github.com/golang/protobuf/ptypes/wrappers=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/wrappers_go_proto%/github.com/golang/protobuf/ptypes/wrappers.a=' -o bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/metric_go_proto%/google.golang.org/genproto/googleapis/api/metric.a -package_list bazel-out/host/bin/external/go_sdk/packages.txt -p google.golang.org/genproto/googleapis/api/metric -- -trimpath .)
SUBCOMMAND: # @go_googleapis//google/devtools/cloudtrace/v2:cloudtrace_go_proto [action 'GoCompile external/go_googleapis/google/devtools/cloudtrace/v2/darwin_amd64_stripped/cloudtrace_go_proto%/google.golang.org/genproto/googleapis/devtools/cloudtrace/v2.a']
(cd /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/execroot/__main__ && \
  exec env - \
    APPLE_SDK_PLATFORM=MacOSX \
    APPLE_SDK_VERSION_OVERRIDE=10.14 \
    CGO_ENABLED=1 \
    GOARCH=amd64 \
    GOOS=darwin \
    GOROOT=external/go_sdk \
    GOROOT_FINAL=GOROOT \
    PATH=external/local_config_cc:/bin:/usr/bin \
    XCODE_VERSION_OVERRIDE=10.2.0 \
  bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix darwin_amd64 -src bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/devtools/cloudtrace/v2/darwin_amd64_stripped/cloudtrace_go_proto%/google.golang.org/genproto/googleapis/devtools/cloudtrace/v2/trace.pb.go -src bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/devtools/cloudtrace/v2/darwin_amd64_stripped/cloudtrace_go_proto%/google.golang.org/genproto/googleapis/devtools/cloudtrace/v2/tracing.pb.go -arc 'google.golang.org/genproto/googleapis/api/annotations=google.golang.org/genproto/googleapis/api/annotations=bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/annotations_go_proto%/google.golang.org/genproto/googleapis/api/annotations.a=' -arc 'google.golang.org/genproto/googleapis/rpc/status=google.golang.org/genproto/googleapis/rpc/status=bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/rpc/darwin_amd64_stripped/status_go_proto%/google.golang.org/genproto/googleapis/rpc/status.a=' -arc 'github.com/golang/protobuf/proto=github.com/golang/protobuf/proto=bazel-out/darwin-fastbuild/bin/external/com_github_golang_protobuf/proto/darwin_amd64_stripped/go_default_library%/github.com/golang/protobuf/proto.a=' -arc 'google.golang.org/grpc=google.golang.org/grpc=bazel-out/darwin-fastbuild/bin/external/org_golang_google_grpc/darwin_amd64_stripped/go_default_library%/google.golang.org/grpc.a=' -arc 'golang.org/x/net/context=golang.org/x/net/context=bazel-out/darwin-fastbuild/bin/external/org_golang_x_net/context/darwin_amd64_stripped/go_default_library%/golang.org/x/net/context.a=' -arc 'github.com/golang/protobuf/ptypes/any=github.com/golang/protobuf/ptypes/any=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/any_go_proto%/github.com/golang/protobuf/ptypes/any.a=' -arc 'google.golang.org/genproto/protobuf/api=google.golang.org/genproto/protobuf/api=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/api_go_proto%/google.golang.org/genproto/protobuf/api.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/plugin=github.com/golang/protobuf/protoc-gen-go/plugin=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/compiler_plugin_go_proto%/github.com/golang/protobuf/protoc-gen-go/plugin.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/descriptor=github.com/golang/protobuf/protoc-gen-go/descriptor=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/descriptor_go_proto%/github.com/golang/protobuf/protoc-gen-go/descriptor.a=' -arc 'github.com/golang/protobuf/ptypes/duration=github.com/golang/protobuf/ptypes/duration=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/duration_go_proto%/github.com/golang/protobuf/ptypes/duration.a=' -arc 'github.com/golang/protobuf/ptypes/empty=github.com/golang/protobuf/ptypes/empty=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/empty_go_proto%/github.com/golang/protobuf/ptypes/empty.a=' -arc 'google.golang.org/genproto/protobuf/field_mask=google.golang.org/genproto/protobuf/field_mask=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/field_mask_go_proto%/google.golang.org/genproto/protobuf/field_mask.a=' -arc 'google.golang.org/genproto/protobuf/source_context=google.golang.org/genproto/protobuf/source_context=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/source_context_go_proto%/google.golang.org/genproto/protobuf/source_context.a=' -arc 'github.com/golang/protobuf/ptypes/struct=github.com/golang/protobuf/ptypes/struct=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/struct_go_proto%/github.com/golang/protobuf/ptypes/struct.a=' -arc 'github.com/golang/protobuf/ptypes/timestamp=github.com/golang/protobuf/ptypes/timestamp=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/timestamp_go_proto%/github.com/golang/protobuf/ptypes/timestamp.a=' -arc 'google.golang.org/genproto/protobuf/ptype=google.golang.org/genproto/protobuf/ptype=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/type_go_proto%/google.golang.org/genproto/protobuf/ptype.a=' -arc 'github.com/golang/protobuf/ptypes/wrappers=github.com/golang/protobuf/ptypes/wrappers=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/wrappers_go_proto%/github.com/golang/protobuf/ptypes/wrappers.a=' -o bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/devtools/cloudtrace/v2/darwin_amd64_stripped/cloudtrace_go_proto%/google.golang.org/genproto/googleapis/devtools/cloudtrace/v2.a -package_list bazel-out/host/bin/external/go_sdk/packages.txt -p google.golang.org/genproto/googleapis/devtools/cloudtrace/v2 -- -trimpath .)
SUBCOMMAND: # @go_googleapis//google/api:distribution_go_proto [action 'GoCompile external/go_googleapis/google/api/darwin_amd64_stripped/distribution_go_proto%/google.golang.org/genproto/googleapis/api/distribution.a']
(cd /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/execroot/__main__ && \
  exec env - \
    APPLE_SDK_PLATFORM=MacOSX \
    APPLE_SDK_VERSION_OVERRIDE=10.14 \
    CGO_ENABLED=1 \
    GOARCH=amd64 \
    GOOS=darwin \
    GOROOT=external/go_sdk \
    GOROOT_FINAL=GOROOT \
    PATH=external/local_config_cc:/bin:/usr/bin \
    XCODE_VERSION_OVERRIDE=10.2.0 \
  bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix darwin_amd64 -src bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/distribution_go_proto%/google.golang.org/genproto/googleapis/api/distribution/distribution.pb.go -arc 'github.com/golang/protobuf/proto=github.com/golang/protobuf/proto=bazel-out/darwin-fastbuild/bin/external/com_github_golang_protobuf/proto/darwin_amd64_stripped/go_default_library%/github.com/golang/protobuf/proto.a=' -arc 'github.com/golang/protobuf/ptypes/any=github.com/golang/protobuf/ptypes/any=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/any_go_proto%/github.com/golang/protobuf/ptypes/any.a=' -arc 'google.golang.org/genproto/protobuf/api=google.golang.org/genproto/protobuf/api=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/api_go_proto%/google.golang.org/genproto/protobuf/api.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/plugin=github.com/golang/protobuf/protoc-gen-go/plugin=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/compiler_plugin_go_proto%/github.com/golang/protobuf/protoc-gen-go/plugin.a=' -arc 'github.com/golang/protobuf/protoc-gen-go/descriptor=github.com/golang/protobuf/protoc-gen-go/descriptor=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/descriptor_go_proto%/github.com/golang/protobuf/protoc-gen-go/descriptor.a=' -arc 'github.com/golang/protobuf/ptypes/duration=github.com/golang/protobuf/ptypes/duration=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/duration_go_proto%/github.com/golang/protobuf/ptypes/duration.a=' -arc 'github.com/golang/protobuf/ptypes/empty=github.com/golang/protobuf/ptypes/empty=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/empty_go_proto%/github.com/golang/protobuf/ptypes/empty.a=' -arc 'google.golang.org/genproto/protobuf/field_mask=google.golang.org/genproto/protobuf/field_mask=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/field_mask_go_proto%/google.golang.org/genproto/protobuf/field_mask.a=' -arc 'google.golang.org/genproto/protobuf/source_context=google.golang.org/genproto/protobuf/source_context=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/source_context_go_proto%/google.golang.org/genproto/protobuf/source_context.a=' -arc 'github.com/golang/protobuf/ptypes/struct=github.com/golang/protobuf/ptypes/struct=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/struct_go_proto%/github.com/golang/protobuf/ptypes/struct.a=' -arc 'github.com/golang/protobuf/ptypes/timestamp=github.com/golang/protobuf/ptypes/timestamp=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/timestamp_go_proto%/github.com/golang/protobuf/ptypes/timestamp.a=' -arc 'google.golang.org/genproto/protobuf/ptype=google.golang.org/genproto/protobuf/ptype=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/type_go_proto%/google.golang.org/genproto/protobuf/ptype.a=' -arc 'github.com/golang/protobuf/ptypes/wrappers=github.com/golang/protobuf/ptypes/wrappers=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/wrappers_go_proto%/github.com/golang/protobuf/ptypes/wrappers.a=' -o bazel-out/darwin-fastbuild/bin/external/go_googleapis/google/api/darwin_amd64_stripped/distribution_go_proto%/google.golang.org/genproto/googleapis/api/distribution.a -package_list bazel-out/host/bin/external/go_sdk/packages.txt -p google.golang.org/genproto/googleapis/api/distribution -- -trimpath .)
SUBCOMMAND: # @com_github_census_instrumentation_opencensus_proto//gen-go/metrics/v1:go_default_library [action 'GoCompile external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/darwin_amd64_stripped/go_default_library%/github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1.a']
(cd /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/execroot/__main__ && \
  exec env - \
    APPLE_SDK_PLATFORM=MacOSX \
    APPLE_SDK_VERSION_OVERRIDE=10.14 \
    CGO_ENABLED=1 \
    GOARCH=amd64 \
    GOOS=darwin \
    GOROOT=external/go_sdk \
    GOROOT_FINAL=GOROOT \
    PATH=external/local_config_cc:/bin:/usr/bin \
    XCODE_VERSION_OVERRIDE=10.2.0 \
  bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix darwin_amd64 -src external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/metrics.pb.go -arc 'github.com/golang/protobuf/proto=github.com/golang/protobuf/proto=bazel-out/darwin-fastbuild/bin/external/com_github_golang_protobuf/proto/darwin_amd64_stripped/go_default_library%/github.com/golang/protobuf/proto.a=' -arc 'github.com/golang/protobuf/ptypes/timestamp=github.com/golang/protobuf/ptypes/timestamp=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/timestamp_go_proto%/github.com/golang/protobuf/ptypes/timestamp.a=' -arc 'github.com/golang/protobuf/ptypes/wrappers=github.com/golang/protobuf/ptypes/wrappers=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/wrappers_go_proto%/github.com/golang/protobuf/ptypes/wrappers.a=' -o bazel-out/darwin-fastbuild/bin/external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/darwin_amd64_stripped/go_default_library%/github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1.a -package_list bazel-out/host/bin/external/go_sdk/packages.txt -p github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1 -- -trimpath .)
ERROR: /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/BUILD.bazel:3:1: GoCompile external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/darwin_amd64_stripped/go_default_library%/github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1.a failed (Exit 1) sandbox-exec failed: error executing command
  (cd /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/sandbox/darwin-sandbox/7/execroot/__main__ && \
  exec env - \
    APPLE_SDK_PLATFORM=MacOSX \
    APPLE_SDK_VERSION_OVERRIDE=10.14 \
    CGO_ENABLED=1 \
    DEVELOPER_DIR=/Applications/Xcode.app/Contents/Developer \
    GOARCH=amd64 \
    GOOS=darwin \
    GOROOT=external/go_sdk \
    GOROOT_FINAL=GOROOT \
    PATH=external/local_config_cc:/bin:/usr/bin \
    SDKROOT=/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.14.sdk \
    TMPDIR=/var/folders/dk/vl32dhy9627byxvq76xy37540000gn/T/ \
    XCODE_VERSION_OVERRIDE=10.2.0 \
  /usr/bin/sandbox-exec -f /private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/sandbox/darwin-sandbox/7/sandbox.sb /var/tmp/_bazel_creedh/install/9c6b7f144e998c67d1156beb135a9453/_embedded_binaries/process-wrapper '--timeout=0' '--kill_delay=15' bazel-out/host/bin/external/go_sdk/builder compile -sdk external/go_sdk -installsuffix darwin_amd64 -src external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/metrics.pb.go -arc 'github.com/golang/protobuf/proto=github.com/golang/protobuf/proto=bazel-out/darwin-fastbuild/bin/external/com_github_golang_protobuf/proto/darwin_amd64_stripped/go_default_library%/github.com/golang/protobuf/proto.a=' -arc 'github.com/golang/protobuf/ptypes/timestamp=github.com/golang/protobuf/ptypes/timestamp=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/timestamp_go_proto%/github.com/golang/protobuf/ptypes/timestamp.a=' -arc 'github.com/golang/protobuf/ptypes/wrappers=github.com/golang/protobuf/ptypes/wrappers=bazel-out/darwin-fastbuild/bin/external/io_bazel_rules_go/proto/wkt/darwin_amd64_stripped/wrappers_go_proto%/github.com/golang/protobuf/ptypes/wrappers.a=' -o bazel-out/darwin-fastbuild/bin/external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/darwin_amd64_stripped/go_default_library%/github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1.a -package_list bazel-out/host/bin/external/go_sdk/packages.txt -p github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1 -- -trimpath .)
compile: missing strict dependencies:
	/private/var/tmp/_bazel_creedh/41d88521a085dfa9f7efd48de5051492/sandbox/darwin-sandbox/7/execroot/__main__/external/com_github_census_instrumentation_opencensus_proto/gen-go/metrics/v1/metrics.pb.go: import of "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
Known dependencies are:
	github.com/golang/protobuf/proto
	github.com/golang/protobuf/ptypes/timestamp
	github.com/golang/protobuf/ptypes/wrappers
Check that imports in Go sources match importpath attributes in deps.
INFO: Elapsed time: 0.424s, Critical Path: 0.21s
INFO: 0 processes.
FAILED: Build did NOT complete successfully
```
