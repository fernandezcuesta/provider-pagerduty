/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/crossplane-contrib/provider-pagerduty/apis/service/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Extension.
func (mg *Extension) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.ExtensionObjects),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.ExtensionObjectsRefs,
		Selector:      mg.Spec.ForProvider.ExtensionObjectsSelector,
		To: reference.To{
			List:    &v1alpha1.ServiceList{},
			Managed: &v1alpha1.Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ExtensionObjects")
	}
	mg.Spec.ForProvider.ExtensionObjects = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.ExtensionObjectsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.ExtensionObjects),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.InitProvider.ExtensionObjectsRefs,
		Selector:      mg.Spec.InitProvider.ExtensionObjectsSelector,
		To: reference.To{
			List:    &v1alpha1.ServiceList{},
			Managed: &v1alpha1.Service{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ExtensionObjects")
	}
	mg.Spec.InitProvider.ExtensionObjects = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.ExtensionObjectsRefs = mrsp.ResolvedReferences

	return nil
}
