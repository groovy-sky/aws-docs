This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::SecurityProfile PrimaryAttributeValue

A primary attribute value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessType" : String,
  "AttributeName" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  AccessType: String
  AttributeName: String
  Values:
    - String

```

## Properties

`AccessType`

The value's access type.

_Required_: Yes

_Type_: String

_Allowed values_: `ALLOW`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AttributeName`

The value's attribute name.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!aws:|connect:)[\p{L}\p{Z}\p{N}\-_.:=@'|]+$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The value's values.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `1000 | 2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PrimaryAttributeAccessControlConfigurationItem

Tag

All content copied from https://docs.aws.amazon.com/.
