This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::SignalCatalog Branch

A group of signals that are defined in a hierarchical structure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "FullyQualifiedName" : String
}

```

### YAML

```yaml

  Description: String
  FullyQualifiedName: String

```

## Properties

`Description`

A brief description of the branch.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000-\u001F\u007F]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FullyQualifiedName`

The fully qualified name of the branch. For example, the fully qualified name of a
branch might be `Vehicle.Body.Engine`.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_.]+`

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Attribute

Node

All content copied from https://docs.aws.amazon.com/.
