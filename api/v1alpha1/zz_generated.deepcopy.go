//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDAppCheck) DeepCopyInto(out *ArgoCDAppCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDAppCheck.
func (in *ArgoCDAppCheck) DeepCopy() *ArgoCDAppCheck {
	if in == nil {
		return nil
	}
	out := new(ArgoCDAppCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDAppUpdate) DeepCopyInto(out *ArgoCDAppUpdate) {
	*out = *in
	if in.SourceUpdates != nil {
		in, out := &in.SourceUpdates, &out.SourceUpdates
		*out = make([]ArgoCDSourceUpdate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDAppUpdate.
func (in *ArgoCDAppUpdate) DeepCopy() *ArgoCDAppUpdate {
	if in == nil {
		return nil
	}
	out := new(ArgoCDAppUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDHelm) DeepCopyInto(out *ArgoCDHelm) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]ArgoCDHelmImageUpdate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDHelm.
func (in *ArgoCDHelm) DeepCopy() *ArgoCDHelm {
	if in == nil {
		return nil
	}
	out := new(ArgoCDHelm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDHelmImageUpdate) DeepCopyInto(out *ArgoCDHelmImageUpdate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDHelmImageUpdate.
func (in *ArgoCDHelmImageUpdate) DeepCopy() *ArgoCDHelmImageUpdate {
	if in == nil {
		return nil
	}
	out := new(ArgoCDHelmImageUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDKustomize) DeepCopyInto(out *ArgoCDKustomize) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDKustomize.
func (in *ArgoCDKustomize) DeepCopy() *ArgoCDKustomize {
	if in == nil {
		return nil
	}
	out := new(ArgoCDKustomize)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoCDSourceUpdate) DeepCopyInto(out *ArgoCDSourceUpdate) {
	*out = *in
	if in.Kustomize != nil {
		in, out := &in.Kustomize, &out.Kustomize
		*out = new(ArgoCDKustomize)
		(*in).DeepCopyInto(*out)
	}
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		*out = new(ArgoCDHelm)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoCDSourceUpdate.
func (in *ArgoCDSourceUpdate) DeepCopy() *ArgoCDSourceUpdate {
	if in == nil {
		return nil
	}
	out := new(ArgoCDSourceUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BookkeeperPromotionMechanism) DeepCopyInto(out *BookkeeperPromotionMechanism) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BookkeeperPromotionMechanism.
func (in *BookkeeperPromotionMechanism) DeepCopy() *BookkeeperPromotionMechanism {
	if in == nil {
		return nil
	}
	out := new(BookkeeperPromotionMechanism)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chart) DeepCopyInto(out *Chart) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chart.
func (in *Chart) DeepCopy() *Chart {
	if in == nil {
		return nil
	}
	out := new(Chart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSubscription) DeepCopyInto(out *ChartSubscription) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSubscription.
func (in *ChartSubscription) DeepCopy() *ChartSubscription {
	if in == nil {
		return nil
	}
	out := new(ChartSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(EnvironmentSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	if in.Subscriptions != nil {
		in, out := &in.Subscriptions, &out.Subscriptions
		*out = new(Subscriptions)
		(*in).DeepCopyInto(*out)
	}
	if in.PromotionMechanisms != nil {
		in, out := &in.PromotionMechanisms, &out.PromotionMechanisms
		*out = new(PromotionMechanisms)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthChecks != nil {
		in, out := &in.HealthChecks, &out.HealthChecks
		*out = new(HealthChecks)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentState) DeepCopyInto(out *EnvironmentState) {
	*out = *in
	if in.FirstSeen != nil {
		in, out := &in.FirstSeen, &out.FirstSeen
		*out = (*in).DeepCopy()
	}
	if in.Commits != nil {
		in, out := &in.Commits, &out.Commits
		*out = make([]GitCommit, len(*in))
		copy(*out, *in)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]Image, len(*in))
		copy(*out, *in)
	}
	if in.Charts != nil {
		in, out := &in.Charts, &out.Charts
		*out = make([]Chart, len(*in))
		copy(*out, *in)
	}
	if in.Health != nil {
		in, out := &in.Health, &out.Health
		*out = new(Health)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentState.
func (in *EnvironmentState) DeepCopy() *EnvironmentState {
	if in == nil {
		return nil
	}
	out := new(EnvironmentState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EnvironmentStateStack) DeepCopyInto(out *EnvironmentStateStack) {
	{
		in := &in
		*out = make(EnvironmentStateStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStateStack.
func (in EnvironmentStateStack) DeepCopy() EnvironmentStateStack {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStateStack)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
	*out = *in
	if in.AvailableStates != nil {
		in, out := &in.AvailableStates, &out.AvailableStates
		*out = make(EnvironmentStateStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.States != nil {
		in, out := &in.States, &out.States
		*out = make(EnvironmentStateStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSubscription) DeepCopyInto(out *EnvironmentSubscription) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSubscription.
func (in *EnvironmentSubscription) DeepCopy() *EnvironmentSubscription {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitCommit) DeepCopyInto(out *GitCommit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitCommit.
func (in *GitCommit) DeepCopy() *GitCommit {
	if in == nil {
		return nil
	}
	out := new(GitCommit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitRepoUpdate) DeepCopyInto(out *GitRepoUpdate) {
	*out = *in
	if in.Bookkeeper != nil {
		in, out := &in.Bookkeeper, &out.Bookkeeper
		*out = new(BookkeeperPromotionMechanism)
		**out = **in
	}
	if in.Kustomize != nil {
		in, out := &in.Kustomize, &out.Kustomize
		*out = new(KustomizePromotionMechanism)
		(*in).DeepCopyInto(*out)
	}
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		*out = new(HelmPromotionMechanism)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitRepoUpdate.
func (in *GitRepoUpdate) DeepCopy() *GitRepoUpdate {
	if in == nil {
		return nil
	}
	out := new(GitRepoUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitSubscription) DeepCopyInto(out *GitSubscription) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitSubscription.
func (in *GitSubscription) DeepCopy() *GitSubscription {
	if in == nil {
		return nil
	}
	out := new(GitSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Health) DeepCopyInto(out *Health) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Health.
func (in *Health) DeepCopy() *Health {
	if in == nil {
		return nil
	}
	out := new(Health)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthChecks) DeepCopyInto(out *HealthChecks) {
	*out = *in
	if in.ArgoCDAppChecks != nil {
		in, out := &in.ArgoCDAppChecks, &out.ArgoCDAppChecks
		*out = make([]ArgoCDAppCheck, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthChecks.
func (in *HealthChecks) DeepCopy() *HealthChecks {
	if in == nil {
		return nil
	}
	out := new(HealthChecks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartDependencyUpdate) DeepCopyInto(out *HelmChartDependencyUpdate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartDependencyUpdate.
func (in *HelmChartDependencyUpdate) DeepCopy() *HelmChartDependencyUpdate {
	if in == nil {
		return nil
	}
	out := new(HelmChartDependencyUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmImageUpdate) DeepCopyInto(out *HelmImageUpdate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmImageUpdate.
func (in *HelmImageUpdate) DeepCopy() *HelmImageUpdate {
	if in == nil {
		return nil
	}
	out := new(HelmImageUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmPromotionMechanism) DeepCopyInto(out *HelmPromotionMechanism) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]HelmImageUpdate, len(*in))
		copy(*out, *in)
	}
	if in.Charts != nil {
		in, out := &in.Charts, &out.Charts
		*out = make([]HelmChartDependencyUpdate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmPromotionMechanism.
func (in *HelmPromotionMechanism) DeepCopy() *HelmPromotionMechanism {
	if in == nil {
		return nil
	}
	out := new(HelmPromotionMechanism)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSubscription) DeepCopyInto(out *ImageSubscription) {
	*out = *in
	if in.IgnoreTags != nil {
		in, out := &in.IgnoreTags, &out.IgnoreTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSubscription.
func (in *ImageSubscription) DeepCopy() *ImageSubscription {
	if in == nil {
		return nil
	}
	out := new(ImageSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizeImageUpdate) DeepCopyInto(out *KustomizeImageUpdate) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizeImageUpdate.
func (in *KustomizeImageUpdate) DeepCopy() *KustomizeImageUpdate {
	if in == nil {
		return nil
	}
	out := new(KustomizeImageUpdate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KustomizePromotionMechanism) DeepCopyInto(out *KustomizePromotionMechanism) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]KustomizeImageUpdate, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KustomizePromotionMechanism.
func (in *KustomizePromotionMechanism) DeepCopy() *KustomizePromotionMechanism {
	if in == nil {
		return nil
	}
	out := new(KustomizePromotionMechanism)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Promotion) DeepCopyInto(out *Promotion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Promotion.
func (in *Promotion) DeepCopy() *Promotion {
	if in == nil {
		return nil
	}
	out := new(Promotion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Promotion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionList) DeepCopyInto(out *PromotionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Promotion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionList.
func (in *PromotionList) DeepCopy() *PromotionList {
	if in == nil {
		return nil
	}
	out := new(PromotionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PromotionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionMechanisms) DeepCopyInto(out *PromotionMechanisms) {
	*out = *in
	if in.GitRepoUpdates != nil {
		in, out := &in.GitRepoUpdates, &out.GitRepoUpdates
		*out = make([]GitRepoUpdate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ArgoCDAppUpdates != nil {
		in, out := &in.ArgoCDAppUpdates, &out.ArgoCDAppUpdates
		*out = make([]ArgoCDAppUpdate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionMechanisms.
func (in *PromotionMechanisms) DeepCopy() *PromotionMechanisms {
	if in == nil {
		return nil
	}
	out := new(PromotionMechanisms)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionSpec) DeepCopyInto(out *PromotionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionSpec.
func (in *PromotionSpec) DeepCopy() *PromotionSpec {
	if in == nil {
		return nil
	}
	out := new(PromotionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotionStatus) DeepCopyInto(out *PromotionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotionStatus.
func (in *PromotionStatus) DeepCopy() *PromotionStatus {
	if in == nil {
		return nil
	}
	out := new(PromotionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepoSubscriptions) DeepCopyInto(out *RepoSubscriptions) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = make([]GitSubscription, len(*in))
		copy(*out, *in)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]ImageSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Charts != nil {
		in, out := &in.Charts, &out.Charts
		*out = make([]ChartSubscription, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepoSubscriptions.
func (in *RepoSubscriptions) DeepCopy() *RepoSubscriptions {
	if in == nil {
		return nil
	}
	out := new(RepoSubscriptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Subscriptions) DeepCopyInto(out *Subscriptions) {
	*out = *in
	if in.Repos != nil {
		in, out := &in.Repos, &out.Repos
		*out = new(RepoSubscriptions)
		(*in).DeepCopyInto(*out)
	}
	if in.UpstreamEnvs != nil {
		in, out := &in.UpstreamEnvs, &out.UpstreamEnvs
		*out = make([]EnvironmentSubscription, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Subscriptions.
func (in *Subscriptions) DeepCopy() *Subscriptions {
	if in == nil {
		return nil
	}
	out := new(Subscriptions)
	in.DeepCopyInto(out)
	return out
}
