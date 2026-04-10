This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule AssetPropertyValue

An asset property value entry containing the following information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Quality" : String,
  "Timestamp" : AssetPropertyTimestamp,
  "Value" : AssetPropertyVariant
}

```

### YAML

```yaml

  Quality: String
  Timestamp:
    AssetPropertyTimestamp
  Value:
    AssetPropertyVariant

```

## Properties

`Quality`

Optional. A string that describes the quality of the value. Accepts substitution
templates. Must be `GOOD`, `BAD`, or `UNCERTAIN`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timestamp`

The asset property value timestamp.

_Required_: Yes

_Type_: [AssetPropertyTimestamp](aws-properties-iot-topicrule-assetpropertytimestamp.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the asset property.

_Required_: Yes

_Type_: [AssetPropertyVariant](aws-properties-iot-topicrule-assetpropertyvariant.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetPropertyTimestamp

AssetPropertyVariant

All content copied from https://docs.aws.amazon.com/.
