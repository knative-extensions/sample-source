# CLAUDE.md - knative-source-template

## What this repo is

`knative.dev/sample-source` (at `knative-extensions/sample-source`) is the canonical
Knative Eventing source template. It implements a heartbeat source that fires
CloudEvents at a configurable interval. Use it as a starting point for new sources.

## Repo layout

```
cmd/{controller,receive_adapter,webhook}/  binary entry points
pkg/adapter/           receive adapter: sends events on a tick
pkg/apis/samples/v1alpha1/  SampleSource CRD type, lifecycle, validation
pkg/client/            generated clientset/informers/listers/injection
pkg/reconciler/        controller reconciler + deployment helper
config/                Kubernetes manifests (CRD, RBAC, Deployments, Webhooks)
hack/                  codegen, dep-update, release scripts
```

## How to instantiate

1. Fork/copy the repo.
2. Update the module path: edit `go.mod`, then
   `sed -i 's|knative.dev/sample-source|knative.dev/YOUR-SOURCE|g'` on all Go files.
3. Rename `SampleSource` / `samplesource` / `samples` in `pkg/apis/` and `config/`.
4. Replace the heartbeat logic in `pkg/adapter/adapter.go`.
5. Run `./hack/update-codegen.sh` to regenerate `pkg/client/`.
6. Update `config/300-samplesource.yaml` with your CRD name and event types.

## Build and test

```sh
gofmt -l .           # must be silent
go vet ./...
go build ./...
go test ./...        # adapter test uses a local HTTP sink, no cluster needed
golangci-lint run ./...
```

## Deploy

```sh
ko apply -f config/  # builds images and applies all manifests
```

The controller reads `SAMPLE_SOURCE_RA_IMAGE` from its own env to know which image
to create for each adapter pod.

## Codegen

Do not hand-edit `pkg/client/` or `zz_generated.deepcopy.go`. After changing types:

```sh
./hack/update-codegen.sh
```

## Gotchas

- `serviceAccountName` is **required** by the validator. `example.yaml` omits it; the
  webhook will reject that resource. Add it or relax `samplesource_validation.go`.
- `admissionWebhookName` in `cmd/webhook/main.go` is the TLS `ServiceName`, not the
  admission webhook registration name (which is derived from `system.Namespace()`).
- `500-webhook-configration.yaml` has a typo in the filename ("configration").
- Static webhook names in manifests assume deployment to the `knative-samples` namespace.

## Writing standard

No em-dashes. No slop: delve, leverage, robust, seamless, elevate, unleash, harness,
realm, tapestry, testament, figurative landscape/navigate, boasts, empower, foster,
streamline, cutting-edge. Terse engineer voice.
