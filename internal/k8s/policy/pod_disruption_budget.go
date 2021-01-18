/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package policy

import (
	"context"

	"github.com/Uptycs/kubequery/internal/k8s"
	"github.com/kolide/osquery-go/plugin/table"
	v1beta1 "k8s.io/api/policy/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type podDisruptionBudget struct {
	k8s.CommonNamespacedFields
	v1beta1.PodDisruptionBudgetSpec
	v1beta1.PodDisruptionBudgetStatus
}

// PodDisruptionBudgetColumns TODO
func PodDisruptionBudgetColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&podDisruptionBudget{})
}

// PodDisruptionBudgetsGenerate TODO
func PodDisruptionBudgetsGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		pdbs, err := k8s.GetClient().PolicyV1beta1().PodDisruptionBudgets(metav1.NamespaceAll).List(context.TODO(), options)
		if err != nil {
			return nil, err
		}

		for _, pdb := range pdbs.Items {
			item := &podDisruptionBudget{
				CommonNamespacedFields:    k8s.GetCommonNamespacedFields(pdb.ObjectMeta),
				PodDisruptionBudgetSpec:   pdb.Spec,
				PodDisruptionBudgetStatus: pdb.Status,
			}
			results = append(results, k8s.ToMap(item))
		}

		if pdbs.Continue == "" {
			break
		}
		options.Continue = pdbs.Continue
	}

	return results, nil
}