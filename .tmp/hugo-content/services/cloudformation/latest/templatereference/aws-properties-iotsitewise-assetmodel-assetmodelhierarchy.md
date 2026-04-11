This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel AssetModelHierarchy

Describes an asset hierarchy that contains a hierarchy's name, ID, and child asset model
ID that specifies the type of asset that can be in this hierarchy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChildAssetModelId" : String,
  "ExternalId" : String,
  "Id" : String,
  "LogicalId" : String,
  "Name" : String
}

```

### YAML

```yaml

  ChildAssetModelId: String
  ExternalId: String
  Id: String
  LogicalId: String
  Name: String

```

## Properties

`ChildAssetModelId`

The ID of the asset model, in UUID format. All assets in this hierarchy must be instances of the
`childAssetModelId` asset model. AWS IoT SiteWise will always return the actual
asset model ID for this value. However, when you are specifying this value as part of a call to
[UpdateAssetModel](../../../../reference/iot-sitewise/latest/apireference/api-updateassetmodel.md), you may provide either the asset model ID or else `externalId:`
followed by the asset model's external ID. For more information, see [Using external IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-ids) in the _AWS IoT SiteWise User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID (if any) provided in the [CreateAssetModel](../../../../reference/iot-sitewise/latest/apireference/api-createassetmodel.md) or [UpdateAssetModel](../../../../reference/iot-sitewise/latest/apireference/api-updateassetmodel.md) operation. You can assign an external ID by specifying
this value as part of a call to [UpdateAssetModel](../../../../reference/iot-sitewise/latest/apireference/api-updateassetmodel.md). However, you can't change the external ID if one is
already assigned. For more information, see [Using external\
IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-ids) in the _AWS IoT SiteWise User Guide_.

###### Note

One of `ExternalId` or `LogicalId` must be specified.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID of the asset model hierarchy. This ID is a `hierarchyId`.

###### Note

This is a return value and can't be set.

- If you are callling [UpdateAssetModel](../../../../reference/iot-sitewise/latest/apireference/api-updateassetmodel.md) to create a _new_ hierarchy:
You can specify its ID here, if desired. AWS IoT SiteWise automatically
generates a unique ID for you, so this parameter is never required. However, if
you prefer to supply your own ID instead, you can specify it here in UUID
format. If you specify your own ID, it must be globally unique.

- If you are calling UpdateAssetModel to modify an _existing_
hierarchy: This can be either the actual ID in UUID format, or else
`externalId:` followed by the external ID, if it has one. For
more information, see [Referencing objects with external IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-id-references) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogicalId`

The `LogicalID` of the asset model hierarchy. This ID is a
`hierarchyLogicalId`.

###### Note

One of `ExternalId` or `LogicalId` must be specified.

_Required_: No

_Type_: String

_Pattern_: `[^\u0000-\u001F\u007F]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the asset model hierarchy that you specify by using the [CreateAssetModel](../../../../reference/iot-sitewise/latest/apireference/api-createassetmodel.md) or
[UpdateAssetModel](../../../../reference/iot-sitewise/latest/apireference/api-updateassetmodel.md) API operation.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetModelCompositeModel

AssetModelProperty

All content copied from https://docs.aws.amazon.com/.
