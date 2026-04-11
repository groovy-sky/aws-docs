This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity Error

The entity error.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Code" : String,
  "Message" : String
}

```

### YAML

```yaml

  Code: String
  Message: String

```

## Properties

`Code`

The entity error code.

_Required_: No

_Type_: String

_Allowed values_: `VALIDATION_ERROR | INTERNAL_FAILURE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Message`

The entity error message.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Definition

Property

All content copied from https://docs.aws.amazon.com/.
