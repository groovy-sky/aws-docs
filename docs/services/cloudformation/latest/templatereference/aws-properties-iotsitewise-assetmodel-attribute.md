This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel Attribute

Contains an asset attribute property. For more information, see
[Attributes](../../../iot-sitewise/latest/userguide/asset-properties.md#attributes) in the _AWS IoT SiteWise User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValue" : String
}

```

### YAML

```yaml

  DefaultValue: String

```

## Properties

`DefaultValue`

The default value of the asset model property attribute. All assets that you create from
the asset model contain this attribute value. You can update an attribute's value after you
create an asset. For more information, see [Updating attribute values](../../../iot-sitewise/latest/userguide/update-attribute-values.md) in the
_AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetModelProperty

EnforcedAssetModelInterfacePropertyMapping

All content copied from https://docs.aws.amazon.com/.
