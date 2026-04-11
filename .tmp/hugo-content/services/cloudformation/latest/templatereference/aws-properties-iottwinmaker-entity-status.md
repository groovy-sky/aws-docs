This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity Status

The current status of the entity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Error" : Error,
  "State" : String
}

```

### YAML

```yaml

  Error:
    Error
  State: String

```

## Properties

`Error`

The error message.

_Required_: No

_Type_: [Error](aws-properties-iottwinmaker-entity-error.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The current state of the entity, component, component type, or workspace.

Valid Values: `CREATING | UPDATING | DELETING | ACTIVE | ERROR`

_Required_: No

_Type_: String

_Allowed values_: `CREATING | UPDATING | DELETING | ACTIVE | ERROR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RelationshipValue

AWS::IoTTwinMaker::Scene

All content copied from https://docs.aws.amazon.com/.
