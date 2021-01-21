/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package apps

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/fake"
)

func getDaemonSet(t *testing.T) *v1.DaemonSet {
	data, err := ioutil.ReadFile("daemon_set_test.json")
	assert.Nil(t, err)

	ds := &v1.DaemonSet{}
	err = json.Unmarshal(data, ds)
	assert.Nil(t, err)

	return ds
}

func TestDaemonSetsGenerate(t *testing.T) {
	k8s.SetClient(fake.NewSimpleClientset(getDaemonSet(t)), types.UID("blah"))
	dss, err := DaemonSetsGenerate(context.TODO(), table.QueryContext{})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]string{
		{
			"annotations":                      "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"cluster_uid":                      "blah",
			"creation_timestamp":               "1610476216",
			"current_number_scheduled":         "1",
			"desired_number_scheduled":         "1",
			"dns_policy":                       "ClusterFirst",
			"host_ipc":                         "0",
			"host_network":                     "1",
			"host_pid":                         "0",
			"labels":                           "{\"k8s-app\":\"calico-node\"}",
			"min_ready_seconds":                "0",
			"name":                             "calico-node",
			"namespace":                        "kube-system",
			"node_selector":                    "{\"kubernetes.io/os\":\"linux\"}",
			"number_available":                 "1",
			"number_misscheduled":              "0",
			"number_ready":                     "1",
			"number_unavailable":               "0",
			"observed_generation":              "1",
			"priority_class_name":              "system-node-critical",
			"restart_policy":                   "Always",
			"revision_history_limit":           "10",
			"scheduler_name":                   "default-scheduler",
			"selector":                         "{\"matchLabels\":{\"k8s-app\":\"calico-node\"}}",
			"service_account_name":             "calico-node",
			"termination_grace_period_seconds": "0",
			"tolerations":                      "[{\"operator\":\"Exists\",\"effect\":\"NoSchedule\"},{\"key\":\"CriticalAddonsOnly\",\"operator\":\"Exists\"},{\"operator\":\"Exists\",\"effect\":\"NoExecute\"}]",
			"uid":                              "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"update_strategy":                  "{\"type\":\"RollingUpdate\",\"rollingUpdate\":{\"maxUnavailable\":1}}",
			"updated_number_scheduled":         "1",
		},
	}, dss)
}

func TestDaemonSetContainersGenerate(t *testing.T) {
	k8s.SetClient(fake.NewSimpleClientset(getDaemonSet(t)), types.UID("blah"))
	dss, err := DaemonSetContainersGenerate(context.TODO(), table.QueryContext{})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]string{
		{
			"annotations":                "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"cluster_uid":                "blah",
			"command":                    "[\"/opt/cni/bin/calico-ipam\",\"-upgrade\"]",
			"container_type":             "init",
			"creation_timestamp":         "1610476216",
			"daemon_set_name":            "calico-node",
			"env":                        "[{\"name\":\"KUBERNETES_NODE_NAME\",\"valueFrom\":{\"fieldRef\":{\"apiVersion\":\"v1\",\"fieldPath\":\"spec.nodeName\"}}},{\"name\":\"CALICO_NETWORKING_BACKEND\",\"valueFrom\":{\"configMapKeyRef\":{\"name\":\"calico-config\",\"key\":\"calico_backend\"}}}]",
			"image":                      "calico/cni:v3.13.2",
			"image_pull_policy":          "IfNotPresent",
			"labels":                     "{\"k8s-app\":\"calico-node\"}",
			"name":                       "upgrade-ipam",
			"namespace":                  "kube-system",
			"privileged":                 "1",
			"resources":                  "{}",
			"stdin":                      "0",
			"stdin_once":                 "0",
			"termination_message_path":   "/dev/termination-log",
			"termination_message_policy": "File",
			"tty":                        "0",
			"uid":                        "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_mounts":              "[{\"name\":\"host-local-net-dir\",\"mountPath\":\"/var/lib/cni/networks\"},{\"name\":\"cni-bin-dir\",\"mountPath\":\"/host/opt/cni/bin\"}]",
		},
		{
			"annotations":                "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"cluster_uid":                "blah",
			"command":                    "[\"/install-cni.sh\"]",
			"container_type":             "init",
			"creation_timestamp":         "1610476216",
			"daemon_set_name":            "calico-node",
			"env":                        "[{\"name\":\"CNI_CONF_NAME\",\"value\":\"10-calico.conflist\"},{\"name\":\"CNI_NETWORK_CONFIG\",\"valueFrom\":{\"configMapKeyRef\":{\"name\":\"calico-config\",\"key\":\"cni_network_config\"}}},{\"name\":\"KUBERNETES_NODE_NAME\",\"valueFrom\":{\"fieldRef\":{\"apiVersion\":\"v1\",\"fieldPath\":\"spec.nodeName\"}}},{\"name\":\"CNI_MTU\",\"valueFrom\":{\"configMapKeyRef\":{\"name\":\"calico-config\",\"key\":\"veth_mtu\"}}},{\"name\":\"SLEEP\",\"value\":\"false\"},{\"name\":\"CNI_NET_DIR\",\"value\":\"/var/snap/microk8s/current/args/cni-network\"}]",
			"image":                      "calico/cni:v3.13.2",
			"image_pull_policy":          "IfNotPresent",
			"labels":                     "{\"k8s-app\":\"calico-node\"}",
			"name":                       "install-cni",
			"namespace":                  "kube-system",
			"privileged":                 "1",
			"resources":                  "{}",
			"stdin":                      "0",
			"stdin_once":                 "0",
			"termination_message_path":   "/dev/termination-log",
			"termination_message_policy": "File",
			"tty":                        "0",
			"uid":                        "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_mounts":              "[{\"name\":\"cni-bin-dir\",\"mountPath\":\"/host/opt/cni/bin\"},{\"name\":\"cni-net-dir\",\"mountPath\":\"/host/etc/cni/net.d\"}]",
		},
		{
			"annotations":                "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"cluster_uid":                "blah",
			"container_type":             "init",
			"creation_timestamp":         "1610476216",
			"daemon_set_name":            "calico-node",
			"image":                      "calico/pod2daemon-flexvol:v3.13.2",
			"image_pull_policy":          "IfNotPresent",
			"labels":                     "{\"k8s-app\":\"calico-node\"}",
			"name":                       "flexvol-driver",
			"namespace":                  "kube-system",
			"privileged":                 "1",
			"resources":                  "{}",
			"stdin":                      "0",
			"stdin_once":                 "0",
			"termination_message_path":   "/dev/termination-log",
			"termination_message_policy": "File",
			"tty":                        "0",
			"uid":                        "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_mounts":              "[{\"name\":\"flexvol-driver-host\",\"mountPath\":\"/host/driver\"}]",
		},
		{
			"annotations":                "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"cluster_uid":                "blah",
			"container_type":             "container",
			"creation_timestamp":         "1610476216",
			"daemon_set_name":            "calico-node",
			"env":                        "[{\"name\":\"DATASTORE_TYPE\",\"value\":\"kubernetes\"},{\"name\":\"WAIT_FOR_DATASTORE\",\"value\":\"true\"},{\"name\":\"NODENAME\",\"valueFrom\":{\"fieldRef\":{\"apiVersion\":\"v1\",\"fieldPath\":\"spec.nodeName\"}}},{\"name\":\"CALICO_NETWORKING_BACKEND\",\"valueFrom\":{\"configMapKeyRef\":{\"name\":\"calico-config\",\"key\":\"calico_backend\"}}},{\"name\":\"CLUSTER_TYPE\",\"value\":\"k8s,bgp\"},{\"name\":\"IP\",\"value\":\"autodetect\"},{\"name\":\"IP_AUTODETECTION_METHOD\",\"value\":\"first-found\"},{\"name\":\"CALICO_IPV4POOL_VXLAN\",\"value\":\"Always\"},{\"name\":\"FELIX_IPINIPMTU\",\"valueFrom\":{\"configMapKeyRef\":{\"name\":\"calico-config\",\"key\":\"veth_mtu\"}}},{\"name\":\"CALICO_IPV4POOL_CIDR\",\"value\":\"10.1.0.0/16\"},{\"name\":\"CALICO_DISABLE_FILE_LOGGING\",\"value\":\"true\"},{\"name\":\"FELIX_DEFAULTENDPOINTTOHOSTACTION\",\"value\":\"ACCEPT\"},{\"name\":\"FELIX_IPV6SUPPORT\",\"value\":\"false\"},{\"name\":\"FELIX_LOGSEVERITYSCREEN\",\"value\":\"error\"},{\"name\":\"FELIX_HEALTHENABLED\",\"value\":\"true\"}]",
			"image":                      "calico/node:v3.13.2",
			"image_pull_policy":          "IfNotPresent",
			"labels":                     "{\"k8s-app\":\"calico-node\"}",
			"liveness_probe":             "{\"exec\":{\"command\":[\"/bin/calico-node\",\"-felix-live\"]},\"initialDelaySeconds\":10,\"timeoutSeconds\":1,\"periodSeconds\":10,\"successThreshold\":1,\"failureThreshold\":6}",
			"name":                       "calico-node",
			"namespace":                  "kube-system",
			"privileged":                 "1",
			"readiness_probe":            "{\"exec\":{\"command\":[\"/bin/calico-node\",\"-felix-ready\"]},\"timeoutSeconds\":1,\"periodSeconds\":10,\"successThreshold\":1,\"failureThreshold\":3}",
			"resources":                  "{\"requests\":{\"cpu\":\"250m\"}}",
			"stdin":                      "0",
			"stdin_once":                 "0",
			"termination_message_path":   "/dev/termination-log",
			"termination_message_policy": "File",
			"tty":                        "0",
			"uid":                        "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_mounts":              "[{\"name\":\"lib-modules\",\"readOnly\":true,\"mountPath\":\"/lib/modules\"},{\"name\":\"xtables-lock\",\"mountPath\":\"/run/xtables.lock\"},{\"name\":\"var-run-calico\",\"mountPath\":\"/var/run/calico\"},{\"name\":\"var-lib-calico\",\"mountPath\":\"/var/lib/calico\"},{\"name\":\"policysync\",\"mountPath\":\"/var/run/nodeagent\"}]",
		},
	}, dss)
}

func TestDaemonSetVolumesGenerate(t *testing.T) {
	k8s.SetClient(fake.NewSimpleClientset(getDaemonSet(t)), types.UID("blah"))
	dss, err := DaemonSetVolumesGenerate(context.TODO(), table.QueryContext{})
	assert.Nil(t, err)
	assert.Equal(t, []map[string]string{
		{
			"annotations":                       "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"aws_elastic_block_store_partition": "0",
			"cluster_uid":                       "blah",
			"creation_timestamp":                "1610476216",
			"daemon_set_name":                   "calico-node",
			"gce_persistent_disk_partition":     "0",
			"host_path_path":                    "/lib/modules",
			"iscsi_discovery_chap_auth":         "0",
			"iscsi_lun":                         "0",
			"iscsi_session_chap_auth":           "0",
			"labels":                            "{\"k8s-app\":\"calico-node\"}",
			"name":                              "lib-modules",
			"namespace":                         "kube-system",
			"scale_iossl_enabled":               "0",
			"uid":                               "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_type":                       "host_path",
		},
		{
			"annotations":                       "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"aws_elastic_block_store_partition": "0",
			"cluster_uid":                       "blah",
			"creation_timestamp":                "1610476216",
			"daemon_set_name":                   "calico-node",
			"gce_persistent_disk_partition":     "0",
			"host_path_path":                    "/var/snap/microk8s/current/var/run/calico",
			"iscsi_discovery_chap_auth":         "0",
			"iscsi_lun":                         "0",
			"iscsi_session_chap_auth":           "0",
			"labels":                            "{\"k8s-app\":\"calico-node\"}",
			"name":                              "var-run-calico",
			"namespace":                         "kube-system",
			"scale_iossl_enabled":               "0",
			"uid":                               "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_type":                       "host_path",
		},
		{
			"annotations":                       "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"aws_elastic_block_store_partition": "0",
			"cluster_uid":                       "blah",
			"creation_timestamp":                "1610476216",
			"daemon_set_name":                   "calico-node",
			"gce_persistent_disk_partition":     "0",
			"host_path_path":                    "/var/snap/microk8s/current/var/lib/calico",
			"iscsi_discovery_chap_auth":         "0",
			"iscsi_lun":                         "0",
			"iscsi_session_chap_auth":           "0",
			"labels":                            "{\"k8s-app\":\"calico-node\"}",
			"name":                              "var-lib-calico",
			"namespace":                         "kube-system",
			"scale_iossl_enabled":               "0",
			"uid":                               "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_type":                       "host_path",
		},
		{
			"annotations":                       "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"aws_elastic_block_store_partition": "0",
			"cluster_uid":                       "blah",
			"creation_timestamp":                "1610476216",
			"daemon_set_name":                   "calico-node",
			"gce_persistent_disk_partition":     "0",
			"host_path_path":                    "/run/xtables.lock",
			"host_path_type":                    "FileOrCreate",
			"iscsi_discovery_chap_auth":         "0",
			"iscsi_lun":                         "0",
			"iscsi_session_chap_auth":           "0",
			"labels":                            "{\"k8s-app\":\"calico-node\"}",
			"name":                              "xtables-lock",
			"namespace":                         "kube-system",
			"scale_iossl_enabled":               "0",
			"uid":                               "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_type":                       "host_path",
		},
		{
			"annotations":                       "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"aws_elastic_block_store_partition": "0",
			"cluster_uid":                       "blah",
			"creation_timestamp":                "1610476216",
			"daemon_set_name":                   "calico-node",
			"gce_persistent_disk_partition":     "0",
			"host_path_path":                    "/var/snap/microk8s/current/opt/cni/bin",
			"iscsi_discovery_chap_auth":         "0",
			"iscsi_lun":                         "0",
			"iscsi_session_chap_auth":           "0",
			"labels":                            "{\"k8s-app\":\"calico-node\"}",
			"name":                              "cni-bin-dir",
			"namespace":                         "kube-system",
			"scale_iossl_enabled":               "0",
			"uid":                               "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_type":                       "host_path",
		},
		{
			"annotations":                       "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"aws_elastic_block_store_partition": "0",
			"cluster_uid":                       "blah",
			"creation_timestamp":                "1610476216",
			"daemon_set_name":                   "calico-node",
			"gce_persistent_disk_partition":     "0",
			"host_path_path":                    "/var/snap/microk8s/current/args/cni-network",
			"iscsi_discovery_chap_auth":         "0",
			"iscsi_lun":                         "0",
			"iscsi_session_chap_auth":           "0",
			"labels":                            "{\"k8s-app\":\"calico-node\"}",
			"name":                              "cni-net-dir",
			"namespace":                         "kube-system",
			"scale_iossl_enabled":               "0",
			"uid":                               "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_type":                       "host_path",
		},
		{
			"annotations":                       "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"aws_elastic_block_store_partition": "0",
			"cluster_uid":                       "blah",
			"creation_timestamp":                "1610476216",
			"daemon_set_name":                   "calico-node",
			"gce_persistent_disk_partition":     "0",
			"host_path_path":                    "/var/snap/microk8s/current/var/lib/cni/networks",
			"iscsi_discovery_chap_auth":         "0",
			"iscsi_lun":                         "0",
			"iscsi_session_chap_auth":           "0",
			"labels":                            "{\"k8s-app\":\"calico-node\"}",
			"name":                              "host-local-net-dir",
			"namespace":                         "kube-system",
			"scale_iossl_enabled":               "0",
			"uid":                               "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_type":                       "host_path",
		},
		{
			"annotations":                       "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"aws_elastic_block_store_partition": "0",
			"cluster_uid":                       "blah",
			"creation_timestamp":                "1610476216",
			"daemon_set_name":                   "calico-node",
			"gce_persistent_disk_partition":     "0",
			"host_path_path":                    "/var/snap/microk8s/current/var/run/nodeagent",
			"host_path_type":                    "DirectoryOrCreate",
			"iscsi_discovery_chap_auth":         "0",
			"iscsi_lun":                         "0",
			"iscsi_session_chap_auth":           "0",
			"labels":                            "{\"k8s-app\":\"calico-node\"}",
			"name":                              "policysync",
			"namespace":                         "kube-system",
			"scale_iossl_enabled":               "0",
			"uid":                               "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_type":                       "host_path",
		},
		{
			"annotations":                       "{\"deprecated.daemonset.template.generation\":\"1\"}",
			"aws_elastic_block_store_partition": "0",
			"cluster_uid":                       "blah",
			"creation_timestamp":                "1610476216",
			"daemon_set_name":                   "calico-node",
			"gce_persistent_disk_partition":     "0",
			"host_path_path":                    "/usr/libexec/kubernetes/kubelet-plugins/volume/exec/nodeagent~uds",
			"host_path_type":                    "DirectoryOrCreate",
			"iscsi_discovery_chap_auth":         "0",
			"iscsi_lun":                         "0",
			"iscsi_session_chap_auth":           "0",
			"labels":                            "{\"k8s-app\":\"calico-node\"}",
			"name":                              "flexvol-driver-host",
			"namespace":                         "kube-system",
			"scale_iossl_enabled":               "0",
			"uid":                               "e6fed7f0-f79a-464f-a3d2-b63247a1f590",
			"volume_type":                       "host_path",
		},
	}, dss)
}
