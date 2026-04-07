This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::ComponentType PropertyDefinition

PropertyDefinition is an object that maps strings to the property definitions in the component type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configurations" : {Key: Value, ...},
  "DataType" : DataType,
  "DefaultValue" : DataValue,
  "IsExternalId" : Boolean,
  "IsRequiredInEntity" : Boolean,
  "IsStoredExternally" : Boolean,
  "IsTimeSeries" : Boolean
}

```

### YAML

```yaml

  Configurations:
    Key: Value
  DataType:
    DataType
  DefaultValue:
    DataValue
  IsExternalId: Boolean
  IsRequiredInEntity: Boolean
  IsStoredExternally: Boolean
  IsTimeSeries: Boolean

```

## Properties

`Configurations`

A mapping that specifies configuration information about the property.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_\-0-9]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataType`

_Required_: No

_Type_: [DataType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iottwinmaker-componenttype-datatype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

A boolean value that specifies whether the property ID comes from an external data store.

_Required_: No

_Type_: [DataValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iottwinmaker-componenttype-datavalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsExternalId`

A Boolean value that specifies whether the property ID comes from an external
data source.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsRequiredInEntity`

A boolean value that specifies whether the property is required in an entity.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsStoredExternally`

A boolean value that specifies whether the property is stored externally.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsTimeSeries`

A boolean value that specifies whether the property consists of time series data.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LambdaFunction

PropertyGroup
