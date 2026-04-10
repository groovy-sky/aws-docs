This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::ComputationModel AssetModelPropertyBindingValue

Contains information about an `assetModelProperty` binding value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssetModelId" : String,
  "PropertyId" : String
}

```

### YAML

```yaml

  AssetModelId: String
  PropertyId: String

```

## Properties

`AssetModelId`

The ID of the asset model, in UUID format.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyId`

The ID of the asset model property used in data binding value.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnomalyDetectionComputationModelConfiguration

AssetPropertyBindingValue

All content copied from https://docs.aws.amazon.com/.
