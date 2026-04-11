This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel EnforcedAssetModelInterfacePropertyMapping

Contains information about applied interface property and asset model property

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssetModelPropertyExternalId" : String,
  "AssetModelPropertyLogicalId" : String,
  "InterfaceAssetModelPropertyExternalId" : String
}

```

### YAML

```yaml

  AssetModelPropertyExternalId: String
  AssetModelPropertyLogicalId: String
  InterfaceAssetModelPropertyExternalId: String

```

## Properties

`AssetModelPropertyExternalId`

The external ID of the linked asset model property

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetModelPropertyLogicalId`

The logical ID of the linked asset model property

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InterfaceAssetModelPropertyExternalId`

The external ID of the applied interface property

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Attribute

EnforcedAssetModelInterfaceRelationship

All content copied from https://docs.aws.amazon.com/.
