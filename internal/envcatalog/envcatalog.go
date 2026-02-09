package envcatalog

type VarInfo struct {
	Category    string
	Name        string
	Description string
	Dynamic     bool
	Internal    bool
}

func Catalog() []VarInfo {
	return []VarInfo{
		{
			Category:    "Config",
			Name:        "VERIFIER_CONFIG",
			Description: "Path to the verifier config file.",
		},
		{
			Category:    "Config",
			Name:        "VERIFIER_<FLAG>",
			Dynamic:     true,
			Description: "Set any verifier CLI flag via environment (hyphens become underscores). Example: VERIFIER_NAMESPACE=default.",
		},
		{
			Category:    "Secrets",
			Name:        "VERIFIER_SECRET_PROVIDER",
			Description: "Default secret provider name for resolving secret:// references in verifier apply/plan.",
		},
		{
			Category:    "Secrets",
			Name:        "VERIFIER_SECRET_CONFIG",
			Description: "Path to a secrets provider config file for resolving secret:// references.",
		},
		{
			Category:    "Output",
			Name:        "NO_COLOR",
			Description: "Disable ANSI color output (any non-empty value).",
		},
		{
			Category:    "CLI",
			Name:        "VERIFIER_YES",
			Description: "Auto-approve confirmations (equivalent to passing --yes).",
		},
		{
			Category:    "Logging",
			Name:        "VERIFIER_KUBE_LOG_LEVEL",
			Description: "Kubernetes client-go verbosity (klog -v). At >=6 enables HTTP request/response tracing.",
		},
		{
			Category:    "Kubernetes",
			Name:        "KUBERNETES_MASTER",
			Description: "Address of the Kubernetes API server (overrides kubeconfig).",
		},
		{
			Category:    "Kubernetes",
			Name:        "KUBECONFIG",
			Description: "Path to the kubeconfig file (defaults to ~/.kube/config).",
		},
		{
			Category:    "Profiling",
			Name:        "VERIFIER_PROFILE",
			Description: "Enable profiling modes for verifier itself (e.g. startup writes CPU/heap profiles to the working directory).",
		},
		{
			Category:    "Features",
			Name:        "VERIFIER_FEATURE_<FLAG>",
			Dynamic:     true,
			Description: "Enable an experimental feature flag (repeatable via env). Example: VERIFIER_FEATURE_DEPLOY_PLAN_HTML_V3=1.",
		},
		{
			Category:    "Build",
			Name:        "VERIFIER_BUILDKIT_HOST",
			Description: "Override the BuildKit address used by `verifier build`.",
		},
		{
			Category:    "Build",
			Name:        "VERIFIER_BUILDKIT_CACHE",
			Description: "Configure BuildKit cache import/export for `verifier build`.",
		},
		{
			Category:    "Build",
			Name:        "VERIFIER_DOCKER_CONTEXT",
			Description: "Docker context to use for Buildx fallback (when provisioning a Docker-backed BuildKit builder).",
		},
		{
			Category:    "Build",
			Name:        "VERIFIER_DOCKER_CONFIG",
			Description: "Override Docker config directory for Buildx fallback (equivalent to DOCKER_CONFIG).",
		},
		{
			Category:    "Registry",
			Name:        "VERIFIER_AUTHFILE",
			Description: "Path to a container registry auth file for `verifier build` (containers-auth.json).",
		},
		{
			Category:    "Registry",
			Name:        "VERIFIER_REGISTRY_AUTH_FILE",
			Description: "Alternate registry auth file path for `verifier build`.",
		},
		{
			Category:    "Sandbox",
			Name:        "VERIFIER_SANDBOX_DISABLE",
			Description: "Disable sandbox execution where supported (set to 1).",
		},
		{
			Category:    "Sandbox",
			Name:        "VERIFIER_SANDBOX_CONFIG",
			Description: "Path to the sandbox policy configuration file.",
		},
		{
			Category:    "Sandbox",
			Name:        "VERIFIER_SANDBOX_ACTIVE",
			Internal:    true,
			Description: "Internal marker set inside the sandbox runtime.",
		},
		{
			Category:    "Sandbox",
			Name:        "VERIFIER_SANDBOX_LOG_PATH",
			Internal:    true,
			Description: "Internal path used by the sandbox to mirror diagnostics/logs.",
		},
		{
			Category:    "Sandbox",
			Name:        "VERIFIER_SANDBOX_CONTEXT",
			Internal:    true,
			Description: "Internal sandbox context marker.",
		},
		{
			Category:    "Sandbox",
			Name:        "VERIFIER_SANDBOX_CACHE",
			Internal:    true,
			Description: "Internal sandbox cache marker.",
		},
		{
			Category:    "Sandbox",
			Name:        "VERIFIER_SANDBOX_BUILDER",
			Internal:    true,
			Description: "Internal sandbox builder marker.",
		},
		{
			Category:    "Sandbox (Legacy)",
			Name:        "VERIFIER_NSJAIL_DISABLE",
			Internal:    true,
			Description: "Legacy alias for VERIFIER_SANDBOX_DISABLE.",
		},
		{
			Category:    "Sandbox (Legacy)",
			Name:        "VERIFIER_NSJAIL_ACTIVE",
			Internal:    true,
			Description: "Legacy alias for VERIFIER_SANDBOX_ACTIVE.",
		},
		{
			Category:    "Sandbox (Legacy)",
			Name:        "VERIFIER_NSJAIL_LOG_PATH",
			Internal:    true,
			Description: "Legacy alias for VERIFIER_SANDBOX_LOG_PATH.",
		},
		{
			Category:    "Sandbox (Legacy)",
			Name:        "VERIFIER_NSJAIL_CONTEXT",
			Internal:    true,
			Description: "Legacy alias for VERIFIER_SANDBOX_CONTEXT.",
		},
		{
			Category:    "Sandbox (Legacy)",
			Name:        "VERIFIER_NSJAIL_CACHE",
			Internal:    true,
			Description: "Legacy alias for VERIFIER_SANDBOX_CACHE.",
		},
		{
			Category:    "Sandbox (Legacy)",
			Name:        "VERIFIER_NSJAIL_BUILDER",
			Internal:    true,
			Description: "Legacy alias for VERIFIER_SANDBOX_BUILDER.",
		},
		{
			Category:    "Capture",
			Name:        "VERIFIER_CAPTURE_QUEUE_SIZE",
			Description: "Capture recorder in-memory queue size.",
		},
		{
			Category:    "Capture",
			Name:        "VERIFIER_CAPTURE_BATCH_SIZE",
			Description: "Capture recorder flush batch size.",
		},
		{
			Category:    "Capture",
			Name:        "VERIFIER_CAPTURE_FLUSH_MS",
			Description: "Capture recorder flush interval in milliseconds.",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_ROOT",
			Description: "Default stack root for `verifier stack ...` when --root is not provided.",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_PROFILE",
			Description: "Default stack profile overlay for `verifier stack ...` when --profile is not provided.",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_OUTPUT",
			Description: "Default output format for `verifier stack` commands when --output is not provided (table|json).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_CLUSTER",
			Description: "Default cluster filter for `verifier stack` selection (comma-separated).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_TAG",
			Description: "Default tag selector for `verifier stack` selection (comma-separated).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_FROM_PATH",
			Description: "Default from-path selector for `verifier stack` selection (comma-separated).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_RELEASE",
			Description: "Default release selector for `verifier stack` selection (comma-separated).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_GIT_RANGE",
			Description: "Default git diff range selector for `verifier stack` selection (example: origin/main...HEAD).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_GIT_INCLUDE_DEPS",
			Description: "When using VERIFIER_STACK_GIT_RANGE, include dependencies (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_GIT_INCLUDE_DEPENDENTS",
			Description: "When using VERIFIER_STACK_GIT_RANGE, include dependents (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_INCLUDE_DEPS",
			Description: "Include dependencies in selection expansion (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_INCLUDE_DEPENDENTS",
			Description: "Include dependents in selection expansion (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_ALLOW_MISSING_DEPS",
			Description: "Allow missing dependencies when selecting a subset (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_INFER_DEPS",
			Description: "Enable inferred dependencies when not explicitly set via flags (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_INFER_CONFIG_REFS",
			Description: "Enable inferred ConfigMap/Secret reference edges when inferring deps (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_APPLY_DRY_RUN",
			Description: "Default `verifier stack apply --dry-run` value when the flag is not provided (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_APPLY_DIFF",
			Description: "Default `verifier stack apply --diff` value when the flag is not provided (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_DELETE_CONFIRM_THRESHOLD",
			Description: "Default delete confirmation threshold for `verifier stack delete` when the flag is not provided.",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_RESUME_ALLOW_DRIFT",
			Description: "Default `verifier stack --allow-drift` value when the flag is not provided (set to 1/true).",
		},
		{
			Category:    "Stack",
			Name:        "VERIFIER_STACK_RESUME_RERUN_FAILED",
			Description: "Default `verifier stack --rerun-failed` value when the flag is not provided (set to 1/true).",
		},
	}
}
