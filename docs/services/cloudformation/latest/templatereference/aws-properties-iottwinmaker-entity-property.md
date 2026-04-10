This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity Property

An object that sets information about a property.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Definition" : Definition,
  "Value" : DataValue
}

```

### YAML

```yaml

  Definition:
    Definition
  Value:
    DataValue

```

## Properties

`Definition`

An object that specifies information about a property.

_Required_: No

_Type_: [Definition](aws-properties-iottwinmaker-entity-definition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

An object that contains information about a value for a time series property.

_Required_: No

_Type_: [DataValue](aws-properties-iottwinmaker-entity-datavalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Error

PropertyGroup

All content copied from https://docs.aws.amazon.com/.
