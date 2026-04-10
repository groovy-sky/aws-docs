This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MemoryDB::SubnetGroup

Specifies a subnet group. A subnet group is a collection of subnets (typically private)
that you can designate for your clusters running in an Amazon Virtual Private Cloud (VPC) environment. When you create a cluster in an Amazon VPC,
you must specify a subnet group. MemoryDB uses that subnet group to choose a
subnet and IP addresses within that subnet to associate with your nodes. For more
information, see [Subnets and subnet\
groups](../../../memorydb/latest/devguide/subnetgroups.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MemoryDB::SubnetGroup",
  "Properties" : {
      "Description" : String,
      "SubnetGroupName" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MemoryDB::SubnetGroup
Properties:
  Description: String
  SubnetGroupName: String
  SubnetIds:
    - String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the subnet group.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetGroupName`

The name of the subnet group to be used for the cluster.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z][a-z0-9\-]*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

A list of Amazon VPC subnet IDs for the subnet group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-memorydb-subnetgroup-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`ARN`

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ARN of the subnet group,
such as `arn:aws:memorydb:us-east-1:123456789012:subnetgroup/my-subnet-group`

`SupportedNetworkTypes`

The network types supported by this subnet. Returns an array of strings that can include 'ipv4', 'ipv6', or both, indicating whether the subnet supports IPv4 only, IPv6 only, or dual-stack deployments.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
