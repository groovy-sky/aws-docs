This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RolesAnywhere::Profile MappingRule

A single mapping entry for each supported specifier or sub-field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Specifier" : String
}

```

### YAML

```yaml

  Specifier: String

```

## Properties

`Specifier`

Specifier within a certificate field, such as CN, OU, or UID from the Subject field.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttributeMapping

Tag

All content copied from https://docs.aws.amazon.com/.
