This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel SetVariable

Information about the variable and its new value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Value" : String,
  "VariableName" : String
}

```

### YAML

```yaml

  Value: String
  VariableName: String

```

## Properties

`Value`

The new value of the variable.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VariableName`

The name of the variable.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SetTimer

Sns

All content copied from https://docs.aws.amazon.com/.
