This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Type Field

The name and data type of an individual field in a user-defined type (UDT).
In addition to a Cassandra data type, you can also use another UDT. When you nest another UDT or collection data type, you have to
declare them with the `FROZEN` keyword.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldName" : String,
  "FieldType" : String
}

```

### YAML

```yaml

  FieldName: String
  FieldType: String

```

## Properties

`FieldName`

The name of the field.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FieldType`

The data type of the field. This can be any Cassandra data type or another user-defined type.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cassandra::Type

Next

All content copied from https://docs.aws.amazon.com/.
