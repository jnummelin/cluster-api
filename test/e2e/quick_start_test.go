//go:build e2e
// +build e2e

/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/utils/ptr"

	"sigs.k8s.io/cluster-api/test/framework"
	"sigs.k8s.io/cluster-api/test/framework/kubetest"
)

var _ = Describe("When following the Cluster API quick-start", func() {
	QuickStartSpec(ctx, func() QuickStartSpecInput {
		return QuickStartSpecInput{
			E2EConfig:              e2eConfig,
			ClusterctlConfigPath:   clusterctlConfigPath,
			BootstrapClusterProxy:  bootstrapClusterProxy,
			ArtifactFolder:         artifactFolder,
			SkipCleanup:            skipCleanup,
			InfrastructureProvider: ptr.To("docker"),
			PostMachinesProvisioned: func(proxy framework.ClusterProxy, namespace, clusterName string) {
				// This check ensures that owner references are resilient - i.e. correctly re-reconciled - when removed.
				framework.ValidateOwnerReferencesResilience(ctx, proxy, namespace, clusterName,
					framework.CoreOwnerReferenceAssertion,
					framework.ExpOwnerReferenceAssertions,
					framework.DockerInfraOwnerReferenceAssertions,
					framework.KubeadmBootstrapOwnerReferenceAssertions,
					framework.KubeadmControlPlaneOwnerReferenceAssertions,
					framework.KubernetesReferenceAssertions,
				)
				// This check ensures that owner references are correctly updated to the correct apiVersion.
				framework.ValidateOwnerReferencesOnUpdate(ctx, proxy, namespace, clusterName,
					framework.CoreOwnerReferenceAssertion,
					framework.ExpOwnerReferenceAssertions,
					framework.DockerInfraOwnerReferenceAssertions,
					framework.KubeadmBootstrapOwnerReferenceAssertions,
					framework.KubeadmControlPlaneOwnerReferenceAssertions,
					framework.KubernetesReferenceAssertions,
				)
			},
		}
	})
})

var _ = Describe("When following the Cluster API quick-start with ClusterClass [PR-Blocking] [ClusterClass]", func() {
	QuickStartSpec(ctx, func() QuickStartSpecInput {
		return QuickStartSpecInput{
			E2EConfig:              e2eConfig,
			ClusterctlConfigPath:   clusterctlConfigPath,
			BootstrapClusterProxy:  bootstrapClusterProxy,
			ArtifactFolder:         artifactFolder,
			SkipCleanup:            skipCleanup,
			Flavor:                 ptr.To("topology"),
			InfrastructureProvider: ptr.To("docker"),
			// This check ensures that owner references are resilient - i.e. correctly re-reconciled - when removed.
			PostMachinesProvisioned: func(proxy framework.ClusterProxy, namespace, clusterName string) {
				framework.ValidateOwnerReferencesResilience(ctx, proxy, namespace, clusterName,
					framework.CoreOwnerReferenceAssertion,
					framework.ExpOwnerReferenceAssertions,
					framework.DockerInfraOwnerReferenceAssertions,
					framework.KubeadmBootstrapOwnerReferenceAssertions,
					framework.KubeadmControlPlaneOwnerReferenceAssertions,
					framework.KubernetesReferenceAssertions,
				)
				// This check ensures that owner references are correctly updated to the correct apiVersion.
				framework.ValidateOwnerReferencesOnUpdate(ctx, proxy, namespace, clusterName,
					framework.CoreOwnerReferenceAssertion,
					framework.ExpOwnerReferenceAssertions,
					framework.DockerInfraOwnerReferenceAssertions,
					framework.KubeadmBootstrapOwnerReferenceAssertions,
					framework.KubeadmControlPlaneOwnerReferenceAssertions,
					framework.KubernetesReferenceAssertions,
				)
			},
		}
	})
})

// NOTE: This test requires an IPv6 management cluster (can be configured via IP_FAMILY=IPv6).
var _ = Describe("When following the Cluster API quick-start with IPv6 [IPv6]", func() {
	QuickStartSpec(ctx, func() QuickStartSpecInput {
		return QuickStartSpecInput{
			E2EConfig:              e2eConfig,
			ClusterctlConfigPath:   clusterctlConfigPath,
			BootstrapClusterProxy:  bootstrapClusterProxy,
			ArtifactFolder:         artifactFolder,
			SkipCleanup:            skipCleanup,
			Flavor:                 ptr.To("ipv6"),
			InfrastructureProvider: ptr.To("docker"),
		}
	})
})

var _ = Describe("When following the Cluster API quick-start with Ignition", func() {
	QuickStartSpec(ctx, func() QuickStartSpecInput {
		return QuickStartSpecInput{
			E2EConfig:              e2eConfig,
			ClusterctlConfigPath:   clusterctlConfigPath,
			BootstrapClusterProxy:  bootstrapClusterProxy,
			ArtifactFolder:         artifactFolder,
			SkipCleanup:            skipCleanup,
			Flavor:                 ptr.To("ignition"),
			InfrastructureProvider: ptr.To("docker"),
		}
	})
})

var _ = Describe("When following the Cluster API quick-start with dualstack and ipv4 primary [IPv6]", func() {
	QuickStartSpec(ctx, func() QuickStartSpecInput {
		return QuickStartSpecInput{
			E2EConfig:              e2eConfig,
			ClusterctlConfigPath:   clusterctlConfigPath,
			BootstrapClusterProxy:  bootstrapClusterProxy,
			ArtifactFolder:         artifactFolder,
			SkipCleanup:            skipCleanup,
			Flavor:                 ptr.To("topology-dualstack-ipv4-primary"),
			InfrastructureProvider: ptr.To("docker"),
			PostMachinesProvisioned: func(proxy framework.ClusterProxy, namespace, clusterName string) {
				By("Running kubetest dualstack tests")
				// Start running the dualstack test suite from kubetest.
				Expect(kubetest.Run(
					ctx,
					kubetest.RunInput{
						ClusterProxy:       proxy.GetWorkloadCluster(ctx, namespace, clusterName),
						ArtifactsDirectory: artifactFolder,
						ConfigFilePath:     "./data/kubetest/dualstack.yaml",
					},
				)).To(Succeed())
			},
		}
	})
})

var _ = Describe("When following the Cluster API quick-start with dualstack and ipv6 primary [IPv6]", func() {
	QuickStartSpec(ctx, func() QuickStartSpecInput {
		return QuickStartSpecInput{
			E2EConfig:              e2eConfig,
			ClusterctlConfigPath:   clusterctlConfigPath,
			BootstrapClusterProxy:  bootstrapClusterProxy,
			ArtifactFolder:         artifactFolder,
			SkipCleanup:            skipCleanup,
			Flavor:                 ptr.To("topology-dualstack-ipv6-primary"),
			InfrastructureProvider: ptr.To("docker"),
			PostMachinesProvisioned: func(proxy framework.ClusterProxy, namespace, clusterName string) {
				By("Running kubetest dualstack tests")
				// Start running the dualstack test suite from kubetest.
				Expect(kubetest.Run(
					ctx,
					kubetest.RunInput{
						ClusterProxy:       proxy.GetWorkloadCluster(ctx, namespace, clusterName),
						ArtifactsDirectory: artifactFolder,
						ConfigFilePath:     "./data/kubetest/dualstack.yaml",
					},
				)).To(Succeed())
			},
		}
	})
})

var _ = Describe("When following the Cluster API quick-start check finalizers resilience after deletion", func() {
	QuickStartSpec(ctx, func() QuickStartSpecInput {
		return QuickStartSpecInput{
			E2EConfig:              e2eConfig,
			ClusterctlConfigPath:   clusterctlConfigPath,
			BootstrapClusterProxy:  bootstrapClusterProxy,
			ArtifactFolder:         artifactFolder,
			SkipCleanup:            skipCleanup,
			InfrastructureProvider: ptr.To("docker"),
			PostMachinesProvisioned: func(proxy framework.ClusterProxy, namespace, clusterName string) {
				// This check ensures that finalizers are resilient - i.e. correctly re-reconciled - when removed.
				framework.ValidateFinalizersResilience(ctx, proxy, namespace, clusterName)
			},
		}
	})
})

var _ = Describe("When following the Cluster API quick-start with ClusterClass check finalizers resilience after deletion [ClusterClass]", func() {
	QuickStartSpec(ctx, func() QuickStartSpecInput {
		return QuickStartSpecInput{
			E2EConfig:              e2eConfig,
			ClusterctlConfigPath:   clusterctlConfigPath,
			BootstrapClusterProxy:  bootstrapClusterProxy,
			ArtifactFolder:         artifactFolder,
			SkipCleanup:            skipCleanup,
			Flavor:                 ptr.To("topology"),
			InfrastructureProvider: ptr.To("docker"),
			PostMachinesProvisioned: func(proxy framework.ClusterProxy, namespace, clusterName string) {
				// This check ensures that finalizers are resilient - i.e. correctly re-reconciled - when removed.
				framework.ValidateFinalizersResilience(ctx, proxy, namespace, clusterName)
			},
		}
	})
})
