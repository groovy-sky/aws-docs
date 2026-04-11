This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Grafana::Workspace NetworkAccessControl

The configuration settings for in-bound network access to your workspace.

When this is configured, only listed IP addresses and VPC endpoints will be able to
access your workspace. Standard Grafana authentication and authorization are still
required.

Access is granted to a caller that is in either the IP address list or the VPC
endpoint list - they do not need to be in both.

If this is not configured, or is removed, then all IP addresses and VPC endpoints are
allowed. Standard Grafana authentication and authorization are still
required.

###### Note

While both `prefixListIds` and `vpceIds` are required, you
can pass in an empty array of strings for either parameter if you do not want to allow any
of that type.

If both are passed as empty arrays, no traffic is allowed to the workspace,
because only _explicitly_ allowed connections are accepted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PrefixListIds" : [ String, ... ],
  "VpceIds" : [ String, ... ]
}

```

### YAML

```yaml

  PrefixListIds:
    - String
  VpceIds:
    - String

```

## Properties

`PrefixListIds`

An array of prefix list IDs. A prefix list is a list of CIDR ranges of IP addresses.
The IP addresses specified are allowed to access your workspace. If the list is not
included in the configuration (passed an empty array) then no IP addresses are
allowed to access the workspace. You create a prefix list using the Amazon VPC
console.

Prefix list IDs have the format `pl-1a2b3c4d`.

For more information about prefix lists, see [Group CIDR blocks using managed\
prefix lists](../../../vpc/latest/userguide/managed-prefix-lists.md) in the _Amazon Virtual Private Cloud User_
_Guide_.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpceIds`

An array of Amazon VPC endpoint IDs for the workspace. You can create VPC
endpoints to your Amazon Managed Grafana workspace for access from within a VPC. If a
`NetworkAccessConfiguration` is specified then only VPC endpoints
specified here are allowed to access the workspace. If you pass in an empty array
of strings, then no VPCs are allowed to access the workspace.

VPC endpoint IDs have the format
`vpce-1a2b3c4d`.

For more information about creating an interface VPC endpoint, see [Interface VPC\
endpoints](../../../grafana/latest/userguide/vpc-endpoints.md) in the _Amazon Managed Grafana User_
_Guide_.

###### Note

The only VPC endpoints that can be specified here are interface VPC endpoints for
Grafana workspaces (using the `com.amazonaws.[region].grafana-workspace`
service endpoint). Other VPC endpoints are ignored.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdpMetadata

RoleValues

All content copied from https://docs.aws.amazon.com/.
