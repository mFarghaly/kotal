// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Account) DeepCopyInto(out *Account) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make(map[HexString]HexString, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account.
func (in *Account) DeepCopy() *Account {
	if in == nil {
		return nil
	}
	out := new(Account)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AvailabilityConfig) DeepCopyInto(out *AvailabilityConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AvailabilityConfig.
func (in *AvailabilityConfig) DeepCopy() *AvailabilityConfig {
	if in == nil {
		return nil
	}
	out := new(AvailabilityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Clique) DeepCopyInto(out *Clique) {
	*out = *in
	out.PoA = in.PoA
	if in.Signers != nil {
		in, out := &in.Signers, &out.Signers
		*out = make([]EthereumAddress, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Clique.
func (in *Clique) DeepCopy() *Clique {
	if in == nil {
		return nil
	}
	out := new(Clique)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ethash) DeepCopyInto(out *Ethash) {
	*out = *in
	if in.FixedDifficulty != nil {
		in, out := &in.FixedDifficulty, &out.FixedDifficulty
		*out = new(uint)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ethash.
func (in *Ethash) DeepCopy() *Ethash {
	if in == nil {
		return nil
	}
	out := new(Ethash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Forks) DeepCopyInto(out *Forks) {
	*out = *in
	if in.DAO != nil {
		in, out := &in.DAO, &out.DAO
		*out = new(uint)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Forks.
func (in *Forks) DeepCopy() *Forks {
	if in == nil {
		return nil
	}
	out := new(Forks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Genesis) DeepCopyInto(out *Genesis) {
	*out = *in
	if in.Accounts != nil {
		in, out := &in.Accounts, &out.Accounts
		*out = make([]Account, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Ethash != nil {
		in, out := &in.Ethash, &out.Ethash
		*out = new(Ethash)
		(*in).DeepCopyInto(*out)
	}
	if in.Clique != nil {
		in, out := &in.Clique, &out.Clique
		*out = new(Clique)
		(*in).DeepCopyInto(*out)
	}
	if in.IBFT2 != nil {
		in, out := &in.IBFT2, &out.IBFT2
		*out = new(IBFT2)
		(*in).DeepCopyInto(*out)
	}
	if in.Forks != nil {
		in, out := &in.Forks, &out.Forks
		*out = new(Forks)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Genesis.
func (in *Genesis) DeepCopy() *Genesis {
	if in == nil {
		return nil
	}
	out := new(Genesis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IBFT2) DeepCopyInto(out *IBFT2) {
	*out = *in
	out.PoA = in.PoA
	if in.Validators != nil {
		in, out := &in.Validators, &out.Validators
		*out = make([]EthereumAddress, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IBFT2.
func (in *IBFT2) DeepCopy() *IBFT2 {
	if in == nil {
		return nil
	}
	out := new(IBFT2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImportedAccount) DeepCopyInto(out *ImportedAccount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImportedAccount.
func (in *ImportedAccount) DeepCopy() *ImportedAccount {
	if in == nil {
		return nil
	}
	out := new(ImportedAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfig) DeepCopyInto(out *NetworkConfig) {
	*out = *in
	if in.Genesis != nil {
		in, out := &in.Genesis, &out.Genesis
		*out = new(Genesis)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfig.
func (in *NetworkConfig) DeepCopy() *NetworkConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Node) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeList) DeepCopyInto(out *NodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Node, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeList.
func (in *NodeList) DeepCopy() *NodeList {
	if in == nil {
		return nil
	}
	out := new(NodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	in.NetworkConfig.DeepCopyInto(&out.NetworkConfig)
	out.AvailabilityConfig = in.AvailabilityConfig
	if in.Import != nil {
		in, out := &in.Import, &out.Import
		*out = new(ImportedAccount)
		**out = **in
	}
	if in.Bootnodes != nil {
		in, out := &in.Bootnodes, &out.Bootnodes
		*out = make([]Enode, len(*in))
		copy(*out, *in)
	}
	if in.StaticNodes != nil {
		in, out := &in.StaticNodes, &out.StaticNodes
		*out = make([]Enode, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CORSDomains != nil {
		in, out := &in.CORSDomains, &out.CORSDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RPCAPI != nil {
		in, out := &in.RPCAPI, &out.RPCAPI
		*out = make([]API, len(*in))
		copy(*out, *in)
	}
	if in.WSAPI != nil {
		in, out := &in.WSAPI, &out.WSAPI
		*out = make([]API, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoA) DeepCopyInto(out *PoA) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoA.
func (in *PoA) DeepCopy() *PoA {
	if in == nil {
		return nil
	}
	out := new(PoA)
	in.DeepCopyInto(out)
	return out
}
