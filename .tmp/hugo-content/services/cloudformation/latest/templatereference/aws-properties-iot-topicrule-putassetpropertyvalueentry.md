This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule PutAssetPropertyValueEntry

An asset property value entry containing the following information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssetId" : String,
  "EntryId" : String,
  "PropertyAlias" : String,
  "PropertyId" : String,
  "PropertyValues" : [ AssetPropertyValue, ... ]
}

```

### YAML

```yaml

  AssetId: String
  EntryId: String
  PropertyAlias: String
  PropertyId: String
  PropertyValues:
    - AssetPropertyValue

```

## Properties

`AssetId`

The ID of the AWS IoT SiteWise asset. You must specify either a `propertyAlias`
or both an `aliasId` and a `propertyId`. Accepts substitution
templates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntryId`

Optional. A unique identifier for this entry that you can define to better track which
message caused an error in case of failure. Accepts substitution templates. Defaults to a new
UUID.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyAlias`

The name of the property alias associated with your asset property. You must specify
either a `propertyAlias` or both an `aliasId` and a
`propertyId`. Accepts substitution templates.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyId`

The ID of the asset's property. You must specify either a `propertyAlias` or
both an `aliasId` and a `propertyId`. Accepts substitution
templates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyValues`

A list of property values to insert that each contain timestamp, quality, and value (TQV)
information.

_Required_: Yes

_Type_: Array of [AssetPropertyValue](aws-properties-iot-topicrule-assetpropertyvalue.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenSearchAction

PutItemInput

All content copied from https://docs.aws.amazon.com/.
