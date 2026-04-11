This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::ComputationModel ComputationModelDataBindingValue

Contains computation model data binding value information, which can be one of
`assetModelProperty`, `list`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssetModelProperty" : AssetModelPropertyBindingValue,
  "AssetProperty" : AssetPropertyBindingValue,
  "List" : [ ComputationModelDataBindingValue, ... ]
}

```

### YAML

```yaml

  AssetModelProperty:
    AssetModelPropertyBindingValue
  AssetProperty:
    AssetPropertyBindingValue
  List:
    - ComputationModelDataBindingValue

```

## Properties

`AssetModelProperty`

Specifies an asset model property data binding value.

_Required_: No

_Type_: [AssetModelPropertyBindingValue](aws-properties-iotsitewise-computationmodel-assetmodelpropertybindingvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetProperty`

The asset property value used for computation model data binding.

_Required_: No

_Type_: [AssetPropertyBindingValue](aws-properties-iotsitewise-computationmodel-assetpropertybindingvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`List`

Specifies a list of data binding value.

_Required_: No

_Type_: Array of [ComputationModelDataBindingValue](aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComputationModelConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
