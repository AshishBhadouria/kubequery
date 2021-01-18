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

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type daemonSet struct {
	k8s.CommonNamespacedFields
	k8s.CommonPodFields
	v1.DaemonSetStatus
	Selector             *metav1.LabelSelector
	UpdateStrategy       v1.DaemonSetUpdateStrategy
	MinReadySeconds      int32
	RevisionHistoryLimit *int32
}

// DaemonSetColumns TODO
func DaemonSetColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&daemonSet{})
}

// DaemonSetsGenerate TODO
func DaemonSetsGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		dss, err := k8s.GetClient().AppsV1().DaemonSets(metav1.NamespaceAll).List(context.TODO(), options)
		if err != nil {
			return nil, err
		}

		for _, ds := range dss.Items {
			item := &daemonSet{
				CommonNamespacedFields: k8s.GetCommonNamespacedFields(ds.ObjectMeta),
				CommonPodFields:        k8s.GetCommonPodFields(ds.Spec.Template.Spec),
				DaemonSetStatus:        ds.Status,
				Selector:               ds.Spec.Selector,
				UpdateStrategy:         ds.Spec.UpdateStrategy,
				MinReadySeconds:        ds.Spec.MinReadySeconds,
				RevisionHistoryLimit:   ds.Spec.RevisionHistoryLimit,
			}
			results = append(results, k8s.ToMap(item))
		}

		if dss.Continue == "" {
			break
		}
		options.Continue = dss.Continue
	}

	return results, nil
}

type daemonSetContainer struct {
	k8s.CommonNamespacedFields
	k8s.CommonContainerFields
	DaemonSetName string
	ContainerType string
}

// DaemonSetContainerColumns TODO
func DaemonSetContainerColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&daemonSetContainer{})
}

// DaemonSetContainersGenerate TODO
func DaemonSetContainersGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		dss, err := k8s.GetClient().AppsV1().DaemonSets(metav1.NamespaceAll).List(context.TODO(), options)
		if err != nil {
			return nil, err
		}

		for _, ds := range dss.Items {
			for _, c := range ds.Spec.Template.Spec.InitContainers {
				item := &daemonSetContainer{
					CommonNamespacedFields: k8s.GetCommonNamespacedFields(ds.ObjectMeta),
					CommonContainerFields:  k8s.GetCommonContainerFields(c),
					DaemonSetName:          ds.Name,
					ContainerType:          "init",
				}
				item.Name = c.Name
				results = append(results, k8s.ToMap(item))
			}
			for _, c := range ds.Spec.Template.Spec.Containers {
				item := &daemonSetContainer{
					CommonNamespacedFields: k8s.GetCommonNamespacedFields(ds.ObjectMeta),
					CommonContainerFields:  k8s.GetCommonContainerFields(c),
					DaemonSetName:          ds.Name,
					ContainerType:          "container",
				}
				item.Name = c.Name
				results = append(results, k8s.ToMap(item))
			}
			for _, c := range ds.Spec.Template.Spec.EphemeralContainers {
				item := &daemonSetContainer{
					CommonNamespacedFields: k8s.GetCommonNamespacedFields(ds.ObjectMeta),
					CommonContainerFields:  k8s.GetCommonEphemeralContainerFields(c),
					DaemonSetName:          ds.Name,
					ContainerType:          "ephemeral",
				}
				item.Name = c.Name
				results = append(results, k8s.ToMap(item))
			}
		}

		if dss.Continue == "" {
			break
		}
		options.Continue = dss.Continue
	}

	return results, nil
}

type daemonSetVolume struct {
	k8s.CommonNamespacedFields
	k8s.CommonVolumeFields
	DaemonSetName string
}

// DaemonSetVolumeColumns TODO
func DaemonSetVolumeColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&daemonSetVolume{})
}

// DaemonSetVolumesGenerate TODO
func DaemonSetVolumesGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		dss, err := k8s.GetClient().AppsV1().DaemonSets(metav1.NamespaceAll).List(context.TODO(), options)
		if err != nil {
			return nil, err
		}

		for _, ds := range dss.Items {
			for _, v := range ds.Spec.Template.Spec.Volumes {
				item := &daemonSetVolume{
					CommonNamespacedFields: k8s.GetCommonNamespacedFields(ds.ObjectMeta),
					CommonVolumeFields:     k8s.GetCommonVolumeFields(v),
					DaemonSetName:          ds.Name,
				}
				item.Name = v.Name
				results = append(results, k8s.ToMap(item))
			}
		}

		if dss.Continue == "" {
			break
		}
		options.Continue = dss.Continue
	}

	return results, nil
}