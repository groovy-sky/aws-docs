This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Neptune::DBSubnetGroup

The `AWS::Neptune::DBSubnetGroup` type creates an Amazon Neptune
DB subnet group. Subnet groups must contain at least two subnets in two different
Availability Zones in the same AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Neptune::DBSubnetGroup",
  "Properties" : {
      "DBSubnetGroupDescription" : String,
      "DBSubnetGroupName" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Neptune::DBSubnetGroup
Properties:
  DBSubnetGroupDescription: String
  DBSubnetGroupName: String
  SubnetIds:
    - String
  Tags:
    - Tag

```

## Properties

`DBSubnetGroupDescription`

Provides the description of the DB subnet group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBSubnetGroupName`

The name of the DB subnet group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The Amazon EC2 subnet IDs for the DB subnet group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags that you want to attach to the DB subnet group.

_Required_: No

_Type_: Array of [Tag](aws-properties-neptune-dbsubnetgroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
