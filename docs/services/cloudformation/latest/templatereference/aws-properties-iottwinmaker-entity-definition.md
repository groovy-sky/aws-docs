This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity Definition

The entity definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configuration" : {Key: Value, ...},
  "DataType" : DataType,
  "DefaultValue" : DataValue,
  "IsExternalId" : Boolean,
  "IsFinal" : Boolean,
  "IsImported" : Boolean,
  "IsInherited" : Boolean,
  "IsRequiredInEntity" : Boolean,
  "IsStoredExternally" : Boolean,
  "IsTimeSeries" : Boolean
}

```

### YAML

```yaml

  Configuration:
    Key: Value
  DataType:
    DataType
  DefaultValue:
    DataValue
  IsExternalId: Boolean
  IsFinal: Boolean
  IsImported: Boolean
  IsInherited: Boolean
  IsRequiredInEntity: Boolean
  IsStoredExternally: Boolean
  IsTimeSeries: Boolean

```

## Properties

`Configuration`

The configuration.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_\-0-9]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataType`

The data type

_Required_: No

_Type_: [DataType](aws-properties-iottwinmaker-entity-datatype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

The default value.

_Required_: No

_Type_: [DataValue](aws-properties-iottwinmaker-entity-datavalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsExternalId`

Displays if the entity has a external Id.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsFinal`

Displays if the entity is final.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsImported`

Displays if the entity is imported.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsInherited`

Displays if the entity is inherited.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsRequiredInEntity`

Displays if the entity is a required entity.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsStoredExternally`

Displays if the entity is tored externally.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsTimeSeries`

Displays if the entity

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataValue

Error

All content copied from https://docs.aws.amazon.com/.
