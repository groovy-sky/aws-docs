This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster RemotePodNetwork

A network CIDR that can contain pods that run Kubernetes webhooks on hybrid nodes.

These CIDR blocks are determined by configuring your Container Network Interface (CNI)
plugin. We recommend the Calico CNI or Cilium CNI. Note that the Amazon VPC CNI plugin for Kubernetes isn't
available for on-premises and edge locations.

Enter one or more IPv4 CIDR blocks in decimal dotted-quad notation (for example,
` 10.2.0.0/16`).

It must satisfy the following requirements:

- Each block must be within an `IPv4` RFC-1918 network range. Minimum
allowed size is /32, maximum allowed size is /8. Publicly-routable addresses
aren't supported.

- Each block cannot overlap with the range of the VPC CIDR blocks for your EKS
resources, or the block of the Kubernetes service IP range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cidrs" : [ String, ... ]
}

```

### YAML

```yaml

  Cidrs:
    - String

```

## Properties

`Cidrs`

A network CIDR that can contain pods that run Kubernetes webhooks on hybrid nodes.

These CIDR blocks are determined by configuring your Container Network Interface (CNI)
plugin. We recommend the Calico CNI or Cilium CNI. Note that the Amazon VPC CNI plugin for Kubernetes isn't
available for on-premises and edge locations.

Enter one or more IPv4 CIDR blocks in decimal dotted-quad notation (for example,
` 10.2.0.0/16`).

It must satisfy the following requirements:

- Each block must be within an `IPv4` RFC-1918 network range. Minimum
allowed size is /32, maximum allowed size is /8. Publicly-routable addresses
aren't supported.

- Each block cannot overlap with the range of the VPC CIDR blocks for your EKS
resources, or the block of the Kubernetes service IP range.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RemoteNodeNetwork

ResourcesVpcConfig

All content copied from https://docs.aws.amazon.com/.
