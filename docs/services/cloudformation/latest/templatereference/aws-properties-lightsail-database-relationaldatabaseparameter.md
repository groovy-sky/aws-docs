This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Database RelationalDatabaseParameter

`RelationalDatabaseParameter` is a property of the [AWS::Lightsail::Database](../userguide/aws-resource-lightsail-database.md) resource. It describes parameters for the
database.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedValues" : String,
  "ApplyMethod" : String,
  "ApplyType" : String,
  "DataType" : String,
  "Description" : String,
  "IsModifiable" : Boolean,
  "ParameterName" : String,
  "ParameterValue" : String
}

```

### YAML

```yaml

  AllowedValues: String
  ApplyMethod: String
  ApplyType: String
  DataType: String
  Description: String
  IsModifiable: Boolean
  ParameterName: String
  ParameterValue: String

```

## Properties

`AllowedValues`

The valid range of values for the parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplyMethod`

Indicates when parameter updates are applied.

Can be `immediate` or `pending-reboot`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApplyType`

Specifies the engine-specific parameter type.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataType`

The valid data type of the parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsModifiable`

A Boolean value indicating whether the parameter can be modified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterName`

The name of the parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValue`

The value for the parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lightsail::Database

Tag

All content copied from https://docs.aws.amazon.com/.
