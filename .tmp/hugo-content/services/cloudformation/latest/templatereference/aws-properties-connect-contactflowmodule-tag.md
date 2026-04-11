This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::ContactFlowModule Tag

A key-value pair to associate with a resource.

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

The key name of the tag. You can specify a value that is 1 to 128 Unicode characters
in length and cannot be prefixed with aws:. You can use any of the following characters:
the set of Unicode letters, digits, whitespace, \_, ., /, =, +, and -

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for the tag. You can specify a value that is 0 to 256 Unicode characters in
length and cannot be prefixed with aws:. You can use any of the following characters:
the set of Unicode letters, digits, whitespace, \_, ., /, =, +, and -

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExternalInvocationConfiguration

AWS::Connect::ContactFlowModuleAlias

All content copied from https://docs.aws.amazon.com/.
