This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterNetworkInterface VpcRouterNetworkInterfaceConfiguration

The configuration settings for a router network interface within a VPC, including the security group IDs and subnet ID.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "SubnetId" : String
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  SubnetId: String

```

## Properties

`SecurityGroupIds`

The IDs of the security groups to associate with the router network interface within the VPC.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The ID of the subnet within the VPC to associate the router network interface with.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::MediaConnect::RouterOutput

All content copied from https://docs.aws.amazon.com/.
