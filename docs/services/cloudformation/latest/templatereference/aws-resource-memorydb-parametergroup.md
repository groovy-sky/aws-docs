This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MemoryDB::ParameterGroup

Specifies a new MemoryDB parameter group. A parameter group is a collection of
parameters and their values that are applied to all of the nodes in any cluster. For more information, see [Configuring engine parameters using\
parameter groups](../../../memorydb/latest/devguide/parametergroups.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MemoryDB::ParameterGroup",
  "Properties" : {
      "Description" : String,
      "Family" : String,
      "ParameterGroupName" : String,
      "Parameters" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MemoryDB::ParameterGroup
Properties:
  Description: String
  Family: String
  ParameterGroupName: String
  Parameters: Json
  Tags:
    - Tag

```

## Properties

`Description`

A description of the parameter group.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Family`

The name of the parameter group family that this parameter group is compatible with.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParameterGroupName`

The name of the parameter group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

Returns the detailed parameter list for the parameter group.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-memorydb-parametergroup-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`ARN`

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ARN of the parameter group,
such as `arn:aws:memorydb:us-east-1:123456789012:parametergroup/my-parameter-group`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
