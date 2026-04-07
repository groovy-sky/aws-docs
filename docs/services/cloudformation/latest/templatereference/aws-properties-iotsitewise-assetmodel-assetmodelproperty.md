This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel AssetModelProperty

Contains information about an asset model property.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataType" : String,
  "DataTypeSpec" : String,
  "ExternalId" : String,
  "Id" : String,
  "LogicalId" : String,
  "Name" : String,
  "Type" : PropertyType,
  "Unit" : String
}

```

### YAML

```yaml

  DataType: String
  DataTypeSpec: String
  ExternalId: String
  Id: String
  LogicalId: String
  Name: String
  Type:
    PropertyType
  Unit: String

```

## Properties

`DataType`

The data type of the asset model property.

If you specify `STRUCT`, you must also specify `dataTypeSpec` to
identify the type of the structure for this property.

_Required_: Yes

_Type_: String

_Allowed values_: `STRING | INTEGER | DOUBLE | BOOLEAN | STRUCT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataTypeSpec`

The data type of the structure for this property. This parameter exists on properties that
have the `STRUCT` data type.

_Required_: No

_Type_: String

_Allowed values_: `AWS/ALARM_STATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID of the asset property. For more information, see [Using external\
IDs](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-ids) in the _AWS IoT SiteWise User Guide_.

###### Note

One of `ExternalId` or `LogicalId` must be specified.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID of the property.

###### Note

This is a return value and can't be set.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogicalId`

The `LogicalID` of the asset model property.

###### Note

One of `ExternalId` or `LogicalId` must be specified.

_Required_: No

_Type_: String

_Pattern_: `[^\u0000-\u001F\u007F]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the asset model property.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Contains a property type, which can be one of `attribute`,
`measurement`, `metric`, or `transform`.

_Required_: Yes

_Type_: [PropertyType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-assetmodel-propertytype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit of the asset model property, such as `Newtons` or
`RPM`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssetModelHierarchy

Attribute
