This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::ComputationModel AssetPropertyBindingValue

Represents a data binding value referencing a specific asset property. It's used to bind
computation model variables to actual asset property values for processing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssetId" : String,
  "PropertyId" : String
}

```

### YAML

```yaml

  AssetId: String
  PropertyId: String

```

## Properties

`AssetId`

The ID of the asset containing the property. This identifies the specific asset instance's
property value used in the computation model.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyId`

The ID of the property within the asset. This identifies the specific property's value
used in the computation model.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetModelPropertyBindingValue

ComputationModelConfiguration

All content copied from https://docs.aws.amazon.com/.
