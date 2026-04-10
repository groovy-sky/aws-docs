This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBShardGroup

Creates a new DB shard group for Aurora Limitless Database. You must enable Aurora Limitless Database to create a DB shard group.

Valid for: Aurora DB clusters only

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBShardGroup",
  "Properties" : {
      "ComputeRedundancy" : Integer,
      "DBClusterIdentifier" : String,
      "DBShardGroupIdentifier" : String,
      "MaxACU" : Number,
      "MinACU" : Number,
      "PubliclyAccessible" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBShardGroup
Properties:
  ComputeRedundancy: Integer
  DBClusterIdentifier: String
  DBShardGroupIdentifier: String
  MaxACU: Number
  MinACU: Number
  PubliclyAccessible: Boolean
  Tags:
    - Tag

```

## Properties

`ComputeRedundancy`

Specifies whether to create standby standby DB data access shard for the DB shard group.
Valid values are the following:

- 0 - Creates a DB shard group without a standby DB data access shard. This is the default value.

- 1 - Creates a DB shard group with a standby DB data access shard in a different Availability Zone (AZ).

- 2 - Creates a DB shard group with two standby DB data access shard in two different AZs.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBClusterIdentifier`

The name of the primary DB cluster for the DB shard group.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBShardGroupIdentifier`

The name of the DB shard group.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxACU`

The maximum capacity of the DB shard group in Aurora capacity units (ACUs).

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinACU`

The minimum capacity of the DB shard group in Aurora capacity units (ACUs).

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PubliclyAccessible`

Specifies whether the DB shard group is publicly accessible.

When the DB shard group is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from
within the DB shard group's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB shard group's VPC.
Access to the DB shard group is ultimately controlled by the security group it uses.
That public access is not permitted if the security group assigned to the DB shard group doesn't permit it.

When the DB shard group isn't publicly accessible, it is an internal DB shard group with a DNS name that resolves to a private IP address.

Default: The default behavior varies depending on whether `DBSubnetGroupName` is specified.

If `DBSubnetGroupName` isn't specified, and `PubliclyAccessible` isn't specified, the following applies:

- If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB shard group is private.

- If the default VPC in the target Region has an internet gateway attached to it, the DB shard group is public.

If `DBSubnetGroupName` is specified, and `PubliclyAccessible` isn't specified, the following applies:

- If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB shard group is private.

- If the subnets are part of a VPC that has an internet gateway attached to it, the DB shard group is public.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An optional set of key-value pairs to associate arbitrary data of your choosing with the DB shard group.

_Required_: No

_Type_: Array of [Tag](aws-properties-rds-dbshardgroup-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name ( `DBClusterIdentifier`) of the DB cluster.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`DBShardGroupResourceId`

The AWS Region-unique, immutable identifier for the DB shard group.

`Endpoint`

This data type represents the information you need to connect to an Amazon RDS DB instance.
This data type is used as a response element in the following actions:

- `CreateDBInstance`

- `DescribeDBInstances`

- `DeleteDBInstance`

For the data structure that represents Amazon Aurora DB cluster endpoints,
see `DBClusterEndpoint`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::RDS::DBSecurityGroupIngress

Tag

All content copied from https://docs.aws.amazon.com/.
