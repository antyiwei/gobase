# kubelet 结构细说
```go
const defaultRootDir = "/var/lib/kubelet"

//如果满足以下任何条件，配置字段应该放在KubeletFlags而不是KubeletConfiguration中：
//  - 在节点的生命周期中，它的值永远不会或不能安全地更改
//  - 无法在节点之间安全地共享其值（例如主机名）
// KubeletConfiguration旨在在节点之间共享
//一般情况下，请尽量避免添加标记或配置字段，
//我们已经有了大量令人困惑的东西。
type KubeletFlags struct {
	KubeConfig          string
	BootstrapKubeconfig string

	// 在调用主服务器期间插入随机错误的概率.
	ChaosChance float64
	// 立刻崩溃，而不是吃恐慌.
	ReallyCrashForTesting bool

	// TODO(mtaufen): It is increasingly looking like nobody actually uses the
	//                Kubelet's runonce mode anymore, so it may be a candidate
	//                for deprecation and removal.
	// If runOnce is true, the Kubelet will check the API server once for pods,
	// run those in addition to the pods specified by static pod files, and exit.
	RunOnce bool

	// enableServer enables the Kubelet's server
	EnableServer bool

	// HostnameOverride is the hostname used to identify the kubelet instead
	// of the actual hostname.
	HostnameOverride string
	// NodeIP is IP address of the node.
	// If set, kubelet will use this IP address for the node.
	NodeIP string

	// This flag, if set, sets the unique id of the instance that an external provider (i.e. cloudprovider)
	// can use to identify a specific node
	ProviderID string

	// Container-runtime-specific options.
	config.ContainerRuntimeOptions

	// certDirectory is the directory where the TLS certs are located (by
	// default /var/run/kubernetes). If tlsCertFile and tlsPrivateKeyFile
	// are provided, this flag will be ignored.
	CertDirectory string

	// cloudProvider is the provider for cloud services.
	// +optional
	CloudProvider string

	// cloudConfigFile is the path to the cloud provider configuration file.
	// +optional
	CloudConfigFile string

	// rootDirectory is the directory path to place kubelet files (volume
	// mounts,etc).
	RootDirectory string

	// The Kubelet will use this directory for checkpointing downloaded configurations and tracking configuration health.
	// The Kubelet will create this directory if it does not already exist.
	// The path may be absolute or relative; relative paths are under the Kubelet's current working directory.
	// Providing this flag enables dynamic kubelet configuration.
	// To use this flag, the DynamicKubeletConfig feature gate must be enabled.
	DynamicConfigDir flag.StringFlag

	// The Kubelet will load its initial configuration from this file.
	// The path may be absolute or relative; relative paths are under the Kubelet's current working directory.
	// Omit this flag to use the combination of built-in default configuration values and flags.
	KubeletConfigFile string

	// registerNode enables automatic registration with the apiserver.
	RegisterNode bool

	// registerWithTaints are an array of taints to add to a node object when
	// the kubelet registers itself. This only takes effect when registerNode
	// is true and upon the initial registration of the node.
	RegisterWithTaints []core.Taint

	// WindowsService should be set to true if kubelet is running as a service on Windows.
	// Its corresponding flag only gets registered in Windows builds.
	WindowsService bool

	// EXPERIMENTAL FLAGS
	// Whitelist of unsafe sysctls or sysctl patterns (ending in *).
	// +optional
	AllowedUnsafeSysctls []string
	// containerized should be set to true if kubelet is running in a container.
	Containerized bool
	// remoteRuntimeEndpoint is the endpoint of remote runtime service
	RemoteRuntimeEndpoint string
	// remoteImageEndpoint is the endpoint of remote image service
	RemoteImageEndpoint string
	// experimentalMounterPath is the path of mounter binary. Leave empty to use the default mount path
	ExperimentalMounterPath string
	// If enabled, the kubelet will integrate with the kernel memcg notification to determine if memory eviction thresholds are crossed rather than polling.
	// +optional
	ExperimentalKernelMemcgNotification bool
	// This flag, if set, enables a check prior to mount operations to verify that the required components
	// (binaries, etc.) to mount the volume are available on the underlying node. If the check is enabled
	// and fails the mount operation fails.
	ExperimentalCheckNodeCapabilitiesBeforeMount bool
	// This flag, if set, will avoid including `EvictionHard` limits while computing Node Allocatable.
	// Refer to [Node Allocatable](https://git.k8s.io/community/contributors/design-proposals/node-allocatable.md) doc for more information.
	ExperimentalNodeAllocatableIgnoreEvictionThreshold bool
	// Node Labels are the node labels to add when registering the node in the cluster
	NodeLabels map[string]string
	// volumePluginDir is the full path of the directory in which to search
	// for additional third party volume plugins
	VolumePluginDir string
	// lockFilePath is the path that kubelet will use to as a lock file.
	// It uses this file as a lock to synchronize with other kubelet processes
	// that may be running.
	LockFilePath string
	// ExitOnLockContention is a flag that signifies to the kubelet that it is running
	// in "bootstrap" mode. This requires that 'LockFilePath' has been set.
	// This will cause the kubelet to listen to inotify events on the lock file,
	// releasing it and exiting when another process tries to open that file.
	ExitOnLockContention bool
	// seccompProfileRoot is the directory path for seccomp profiles.
	SeccompProfileRoot string
	// bootstrapCheckpointPath is the path to the directory containing pod checkpoints to
	// run on restore
	BootstrapCheckpointPath string
	// NodeStatusMaxImages caps the number of images reported in Node.Status.Images.
	// This is an experimental, short-term flag to help with node scalability.
	NodeStatusMaxImages int32

	// DEPRECATED FLAGS
	// minimumGCAge is the minimum age for a finished container before it is
	// garbage collected.
	MinimumGCAge metav1.Duration
	// maxPerPodContainerCount is the maximum number of old instances to
	// retain per container. Each container takes up some disk space.
	MaxPerPodContainerCount int32
	// maxContainerCount is the maximum number of old instances of containers
	// to retain globally. Each container takes up some disk space.
	MaxContainerCount int32
	// masterServiceNamespace is The namespace from which the kubernetes
	// master services should be injected into pods.
	MasterServiceNamespace string
	// registerSchedulable tells the kubelet to register the node as
	// schedulable. Won't have any effect if register-node is false.
	// DEPRECATED: use registerWithTaints instead
	RegisterSchedulable bool
	// nonMasqueradeCIDR configures masquerading: traffic to IPs outside this range will use IP masquerade.
	NonMasqueradeCIDR string
	// This flag, if set, instructs the kubelet to keep volumes from terminated pods mounted to the node.
	// This can be useful for debugging volume related issues.
	KeepTerminatedPodVolumes bool
	// allowPrivileged enables containers to request privileged mode.
	// Defaults to true.
	AllowPrivileged bool
	// hostNetworkSources is a comma-separated list of sources from which the
	// Kubelet allows pods to use of host network. Defaults to "*". Valid
	// options are "file", "http", "api", and "*" (all sources).
	HostNetworkSources []string
	// hostPIDSources is a comma-separated list of sources from which the
	// Kubelet allows pods to use the host pid namespace. Defaults to "*".
	HostPIDSources []string
	// hostIPCSources is a comma-separated list of sources from which the
	// Kubelet allows pods to use the host ipc namespace. Defaults to "*".
	HostIPCSources []string
}

// NewKubeletFlags will create a new KubeletFlags with default values
func NewKubeletFlags() *KubeletFlags {
	remoteRuntimeEndpoint := ""
	if runtime.GOOS == "linux" {
		remoteRuntimeEndpoint = "unix:///var/run/dockershim.sock"
	} else if runtime.GOOS == "windows" {
		remoteRuntimeEndpoint = "tcp://localhost:3735"
	}

	return &KubeletFlags{
		EnableServer:                        true,
		ContainerRuntimeOptions:             *NewContainerRuntimeOptions(),
		CertDirectory:                       "/var/lib/kubelet/pki",
		RootDirectory:                       defaultRootDir,
		MasterServiceNamespace:              metav1.NamespaceDefault,
		MaxContainerCount:                   -1,
		MaxPerPodContainerCount:             1,
		MinimumGCAge:                        metav1.Duration{Duration: 0},
		NonMasqueradeCIDR:                   "10.0.0.0/8",
		RegisterSchedulable:                 true,
		ExperimentalKernelMemcgNotification: false,
		RemoteRuntimeEndpoint:               remoteRuntimeEndpoint,
		NodeLabels:                          make(map[string]string),
		VolumePluginDir:                     "/usr/libexec/kubernetes/kubelet-plugins/volume/exec/",
		RegisterNode:                        true,
		SeccompProfileRoot:                  filepath.Join(defaultRootDir, "seccomp"),
		HostNetworkSources:                  []string{kubetypes.AllSource},
		HostPIDSources:                      []string{kubetypes.AllSource},
		HostIPCSources:                      []string{kubetypes.AllSource},
		// TODO(#58010:v1.13.0): Remove --allow-privileged, it is deprecated
		AllowPrivileged: true,
		// prior to the introduction of this flag, there was a hardcoded cap of 50 images
		NodeStatusMaxImages: 50,
	}
}
```