This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::ComponentType CompositeComponentType

Specifies the ID of the composite component type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentTypeId" : String
}

```

### YAML

```yaml

  ComponentTypeId: String

```

## Properties

`ComponentTypeId`

The ID of the component type.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z_\.\-0-9:]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTTwinMaker::ComponentType

DataConnector

All content copied from https://docs.aws.amazon.com/.
