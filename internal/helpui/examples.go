package helpui

// curatedExamples supplements Cobra's .Example fields with task-based golden paths.
// Keys are Cobra command paths (CommandPath()).
var curatedExamples = map[string][]string{
	"verifier logs": {
		"# Tail pods matching a regex in a namespace\nverifier logs 'checkout-.*' -n prod-payments",
		"# Highlight errors\nverifier logs 'checkout-.*' -n prod-payments --highlight ERROR",
	},
	"verifier init": {
		"# Create a repo-local .verifier.yaml\nverifier init",
		"# Preview the config without writing\nverifier init --dry-run",
		"# Run the interactive wizard\nverifier init --interactive",
		"# Use an opinionated preset\nverifier init --preset prod",
		"# Apply a built-in template\nverifier init --template platform",
		"# Apply a template from a URL\nverifier init --template https://example.com/verifier-init.yaml",
		"# Merge defaults into an existing config\nverifier init --merge",
		"# Scaffold chart/ and values/ plus gitignore\nverifier init --layout --gitignore",
		"# Scaffold Vault-backed secrets\nverifier init --secrets-provider vault",
		"# Emit JSON for automation\nverifier init --output json --dry-run",
		"# Write a replayable init plan\nverifier init --plan --plan-output .verifier/init-plan.json",
		"# Apply a saved init plan\nverifier init --apply-plan .verifier/init-plan.json",
		"# Initialize another path\nverifier init ./services/api",
		"# Overwrite existing config\nverifier init --force",
	},
	"verifier build": {
		"# Build an image from a directory\nverifier build --context . --tag ghcr.io/acme/app:dev",
		"# Share the build stream over WebSocket\nverifier build --context . --ws-listen :9085",
	},
	"verifier help": {
		"# Launch the interactive help UI\nverifier help --ui",
		"# Show help for a specific command\nverifier help apply",
	},
	"verifier apply plan": {
		"# Preview a Helm upgrade\nverifier apply plan --chart ./chart --release foo -n default",
		"# Render a shareable HTML visualization\nverifier apply plan --visualize --chart ./chart --release foo -n default",
		"# Preview with secret references\nverifier apply plan --chart ./chart --release foo -n default --secret-provider local",
		"# Preview with Vault-backed secrets\nverifier apply plan --chart ./chart --release foo -n default --secret-provider vault",
		"# Compare against a saved baseline\nverifier apply plan --chart ./chart --release foo -n default --compare-to ./plan.json",
		"# Write a baseline snapshot\nverifier apply plan --chart ./chart --release foo -n default --baseline ./plan.json",
	},
	"verifier apply": {
		"# Deploy a chart\nverifier apply --chart ./chart --release foo -n default",
		"# Run the deploy viewer\nverifier apply --chart ./chart --release foo -n default --ui",
		"# Deploy with secret references\nverifier apply --chart ./chart --release foo -n default --secret-provider local",
		"# Deploy with Vault-backed secrets\nverifier apply --chart ./chart --release foo -n default --secret-provider vault",
	},
	"verifier delete": {
		"# Delete a release\nverifier delete --release foo -n default",
		"# Run the destroy viewer\nverifier delete --release foo -n default --ui",
	},
	"verifier revert": {
		"# Revert a release to the last known-good revision\nverifier revert --release foo -n default",
	},
	"verifier env": {
		"# Show env var reference (machine-readable)\nverifier env --format json",
	},
	"verifier secrets": {
		"# Validate a secret reference\nverifier secrets test --secret-provider vault --ref secret://vault/app/db#password",
		"# List secrets under a provider prefix\nverifier secrets list --secret-provider local --path app --format json",
		"# Discover secret refs across the repo\nverifier secrets discover --scope repo",
		"# Discover secret refs for a chart\nverifier secrets discover --scope chart --chart ./chart --values values/dev.yaml",
		"# Discover secret refs for a stack\nverifier secrets discover --scope stack --config ./stacks/prod",
	},
	"verifier version": {
		"# Print version information\nverifier version",
	},
	"verifier stack": {
		"# Plan the stack (default: read-only, like `verifier stack plan`)\nverifier stack --config ./stacks/prod",
		"# Restrict selection via environment defaults\nVERIFIER_STACK_TAG=critical VERIFIER_STACK_CLUSTER=prod-us verifier stack --config ./stacks/prod",
		"# Emit a machine-readable plan for tooling\nverifier stack --config ./stacks/prod --output json",
	},
	"verifier stack plan": {
		"# Write a reproducible plan bundle for review/CI\nverifier stack plan --config ./stacks/prod --bundle ./stack-plan.tgz",
		"# Embed a live diff summary in the bundle (requires cluster access)\nverifier stack plan --config ./stacks/prod --bundle ./stack-plan.tgz --bundle-diff-summary",
	},
	"verifier stack graph": {
		"# Render a Graphviz DOT graph\nverifier stack graph --config ./stacks/prod > stack.dot",
		"# Render a Mermaid graph\nverifier stack graph --config ./stacks/prod --format mermaid > stack.mmd",
	},
	"verifier stack explain": {
		"# Explain why a release is selected (by name)\nverifier stack explain --config ./stacks/prod api",
		"# Print only selection reasons\nverifier stack explain --config ./stacks/prod api --why",
	},
	"verifier stack apply": {
		"# Apply the selected releases (DAG order)\nverifier stack apply --config ./stacks/prod --yes",
		"# Resume the most recent run (uses stored frozen plan unless --replan is set)\nverifier stack apply --config ./stacks/prod --resume --yes",
		"# Enable manifest diffs (defaulted via env)\nVERIFIER_STACK_APPLY_DIFF=1 verifier stack apply --config ./stacks/prod --yes",
		"# Apply with secret references\nverifier stack apply --config ./stacks/prod --secret-provider vault --yes",
	},
	"verifier stack delete": {
		"# Delete the selected releases (reverse DAG order)\nverifier stack delete --config ./stacks/prod --yes",
		"# Prompt only when deleting 50+ releases\nverifier stack delete --config ./stacks/prod --delete-confirm-threshold 50",
	},
	"verifier stack status": {
		"# Tail the most recent run\nverifier stack status --config ./stacks/prod --follow",
		"# Show a specific run ID (see `verifier stack runs`)\nverifier stack status --config ./stacks/prod --run-id 2025-12-30T12-34-56.000000000Z --follow",
	},
	"verifier stack runs": {
		"# List recent runs\nverifier stack runs --config ./stacks/prod --limit 50",
	},
	"verifier stack audit": {
		"# Show audit table for the most recent run\nverifier stack audit --config ./stacks/prod",
		"# Export a shareable HTML report\nverifier stack audit --config ./stacks/prod --output html > stack-audit.html",
	},
	"verifier stack export": {
		"# Export the most recent run as a portable bundle\nverifier stack export --config ./stacks/prod",
		"# Export a specific run ID\nverifier stack export --config ./stacks/prod --run-id 2025-12-30T12-34-56.000000000Z --out ./exports/run.tgz",
	},
	"verifier stack seal": {
		"# Seal a plan directory for CI (includes inputs bundle by default)\nverifier stack seal --config ./stacks/prod --out ./.verifier/stack/sealed --command apply",
		"# Seal without bundling inputs\nverifier stack seal --config ./stacks/prod --out ./.verifier/stack/sealed --bundle=false --command apply",
	},
	"verifier stack rerun-failed": {
		"# Resume the most recent run and schedule only failed nodes\nverifier stack rerun-failed --config ./stacks/prod --yes",
	},
	"verify": {
		"# Verify a chart render (inline)\nverify --chart ./chart --release foo -n default",
		"# Verify a chart render with cluster lookups\nverify --chart ./chart --release foo -n default --use-cluster --context my-context",
		"# Verify a live namespace\nverify --namespace default --context my-context",
		"# Discover builtin rules\nverify rules list\nverify rules show k8s/container_is_privileged",
		"# Generate a starter config for scripting\nverify init --chart ./chart --release foo -n default --write verify.yaml\nverify verify.yaml",
		"# Run the bundled verify showcase (includes a CRITICAL rule)\nverify testdata/verify/showcase/verify.yaml",
		"# Compare against a baseline report\nverify verify.yaml --compare-to ./baseline.json",
		"# Write a baseline report\nverify verify.yaml --baseline ./baseline.json",
		"# HTML report (lux minimal)\nverify --manifest ./rendered.yaml --format html --report ./verify-report.html --open",
		"# Print a suggested fix plan (table output only)\nverify verify.yaml --fix",
	},
	"package": {
		"# Package a chart directory\npackage ./chart --output dist/chart.sqlite",
		"# Verify an existing archive\npackage --verify dist/chart.sqlite",
		"# Package then verify (quiet with SHA)\npackage ./chart --output dist/chart.sqlite --print-sha --quiet && package --verify dist/chart.sqlite",
		"# Stream an archive over ssh\npackage ./chart --output - | ssh host \"cat > chart.sqlite\"",
		"# Unpack an archive into a directory\npackage --unpack dist/chart.sqlite --destination ./chart-unpacked",
	},
}
