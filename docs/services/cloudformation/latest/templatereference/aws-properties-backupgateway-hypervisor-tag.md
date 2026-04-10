This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BackupGateway::Hypervisor Tag

A key-value pair you can use to manage, filter, and search for your resources. Allowed
characters include UTF-8 letters, numbers, and the following characters: + - = . \_ :
/. Spaces are not allowed in tag values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key part of a tag's key-value pair. The key can't start with `aws:`.

_Required_: Yes

_Type_: String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The value part of a tag's key-value pair.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\x00]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::BackupGateway::Hypervisor

Next

All content copied from https://docs.aws.amazon.com/.
