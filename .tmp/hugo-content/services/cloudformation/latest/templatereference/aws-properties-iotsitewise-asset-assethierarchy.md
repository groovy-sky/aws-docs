This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Asset AssetHierarchy

Describes an asset hierarchy that contains a hierarchy's name and ID.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChildAssetId" : String,
  "ExternalId" : String,
  "Id" : String,
  "LogicalId" : String
}

```

### YAML

```yaml

  ChildAssetId: String
  ExternalId: String
  Id: String
  LogicalId: String

```

## Properties

`ChildAssetId`

The Id of the child asset.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID of the hierarchy, if it has one. When you update an asset hierarchy, you
may assign an external ID if it doesn't already have one. You can't change the external ID of
an asset hierarchy that already has one. For more information, see [Using external IDs](../../../iot-sitewise/latest/userguide/object-ids.md#external-ids) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_][a-zA-Z_\-0-9.:]*[a-zA-Z0-9_]+`

_Minimum_: `2`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The ID of the hierarchy. This ID is a `hierarchyId`.

###### Note

This is a return value and can't be set.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogicalId`

The ID of the hierarchy. This ID is a `hierarchyId`.

_Required_: No

_Type_: String

_Pattern_: `[^\u0000-\u001F\u007F]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTSiteWise::Asset

AssetProperty

All content copied from https://docs.aws.amazon.com/.
