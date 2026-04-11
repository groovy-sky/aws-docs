This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel EnforcedAssetModelInterfaceRelationship

Contains information about applied interface hierarchy and asset model
hierarchy

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InterfaceAssetModelId" : String,
  "PropertyMappings" : [ EnforcedAssetModelInterfacePropertyMapping, ... ]
}

```

### YAML

```yaml

  InterfaceAssetModelId: String
  PropertyMappings:
    - EnforcedAssetModelInterfacePropertyMapping

```

## Properties

`InterfaceAssetModelId`

The ID of the asset model that has the interface applied to it.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyMappings`

A list of property mappings between the interface asset model and the asset model
where the interface is applied.

_Required_: No

_Type_: Array of [EnforcedAssetModelInterfacePropertyMapping](aws-properties-iotsitewise-assetmodel-enforcedassetmodelinterfacepropertymapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnforcedAssetModelInterfacePropertyMapping

ExpressionVariable

All content copied from https://docs.aws.amazon.com/.
