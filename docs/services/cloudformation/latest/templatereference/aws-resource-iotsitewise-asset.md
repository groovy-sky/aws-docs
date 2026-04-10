This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Asset

Creates an asset from an existing asset model. For more information, see [Creating assets](../../../iot-sitewise/latest/userguide/create-assets.md) in the
_AWS IoT SiteWise User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTSiteWise::Asset",
  "Properties" : {
      "AssetDescription" : String,
      "AssetExternalId" : String,
      "AssetHierarchies" : [ AssetHierarchy, ... ],
      "AssetModelId" : String,
      "AssetName" : String,
      "AssetProperties" : [ AssetProperty, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTSiteWise::Asset
Properties:
  AssetDescription: String
  AssetExternalId: String
  AssetHierarchies:
    - AssetHierarchy
  AssetModelId: String
  AssetName: String
  AssetProperties:
    - AssetProperty
  Tags:
    - Tag

```

## Properties

`AssetDescription`

The ID of the asset, in UUID format.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetExternalId`

The external ID of the asset model composite model. For more information, see [Using external IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-ids) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetHierarchies`

A list of asset hierarchies that each contain a `hierarchyId`. A hierarchy specifies allowed parent/child asset relationships.

_Required_: No

_Type_: Array of [AssetHierarchy](aws-properties-iotsitewise-asset-assethierarchy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetModelId`

The ID of the asset model from which to create the asset. This can be either the actual ID in UUID format, or else `externalId:` followed by the external ID, if it has one.
For more information, see [Referencing objects with external IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-id-references) in the _AWS IoT SiteWise User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetName`

A friendly name for the asset.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetProperties`

The list of asset properties for the asset.

This object doesn't include properties that you define in composite models. You can find
composite model properties in the `assetCompositeModels` object.

_Required_: No

_Type_: Array of [AssetProperty](aws-properties-iotsitewise-asset-assetproperty.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that contain metadata for the asset. For more information, see
[Tagging your AWS IoT SiteWise\
resources](../../../iot-sitewise/latest/userguide/tag-resources.md) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotsitewise-asset-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `AssetId` and `AssetArn`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssetArn`

The ARN of the asset.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`AssetId`

The ID of the asset.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

User

AssetHierarchy

All content copied from https://docs.aws.amazon.com/.
