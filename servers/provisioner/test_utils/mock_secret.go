// © Copyright 2024 Hewlett Packard Enterprise Development LP

// Utils for test classes
package test_utils

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	k8sCorev1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// Implements corev1.SecretInterface to mock its methods
type MockSecret struct {
	Secret *v1.Secret
}

func (m MockSecret) Create(ctx context.Context, secret *v1.Secret, opts metav1.CreateOptions) (*v1.Secret, error) {
	return m.Secret, nil
}

func (m MockSecret) Update(ctx context.Context, secret *v1.Secret, opts metav1.UpdateOptions) (*v1.Secret, error) {
	return m.Secret, nil
}

func (m MockSecret) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return nil
}

func (m MockSecret) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return nil
}

func (m MockSecret) Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Secret, error) {
	return m.Secret, nil
}

func (m MockSecret) List(ctx context.Context, opts metav1.ListOptions) (*v1.SecretList, error) {
	return nil, nil
}

func (m MockSecret) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return nil, nil
}

func (m MockSecret) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Secret, err error) {
	return nil, nil
}

func (m MockSecret) Apply(ctx context.Context, secret *k8sCorev1.SecretApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Secret, err error) {
	return nil, nil
}
