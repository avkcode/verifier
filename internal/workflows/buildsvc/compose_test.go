package buildsvc

import (
	"bytes"
	"context"
	"testing"

	appcompose "verifier/pkg/compose"
)

type fakeComposeRunner struct {
	results []appcompose.ServiceBuildResult
	err     error
}

func (f fakeComposeRunner) BuildCompose(_ context.Context, _ appcompose.ComposeBuildOptions) ([]appcompose.ServiceBuildResult, error) {
	return append([]appcompose.ServiceBuildResult(nil), f.results...), f.err
}

func TestRunComposeBuild_SortsServiceOutput(t *testing.T) {
	var out bytes.Buffer
	svc := &service{
		composeRunner: fakeComposeRunner{
			results: []appcompose.ServiceBuildResult{
				{Service: "worker", Tags: []string{"verifier-test/worker:dev"}},
				{Service: "api", Tags: []string{"verifier-test/api:dev"}},
			},
		},
	}

	err := svc.runComposeBuild(context.Background(), []string{"docker-compose.yml"}, Options{}, nil, nil, false, nil, Streams{Out: &out, Err: &out})
	if err != nil {
		t.Fatalf("runComposeBuild() err = %v", err)
	}

	if got := out.String(); got != "api: verifier-test/api:dev\nworker: verifier-test/worker:dev\n" {
		t.Fatalf("unexpected output:\n%s", got)
	}
}
