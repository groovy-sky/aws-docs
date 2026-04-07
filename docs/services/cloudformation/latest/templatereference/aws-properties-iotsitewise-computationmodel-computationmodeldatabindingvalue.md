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

_Type_: [AssetModelPropertyBindingValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-computationmodel-assetmodelpropertybindingvalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssetProperty`

The asset property value used for computation model data binding.

_Required_: No

_Type_: [AssetPropertyBindingValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-computationmodel-assetpropertybindingvalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`List`

Specifies a list of data binding value.

_Required_: No

_Type_: Array of [ComputationModelDataBindingValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ComputationModelConfiguration

Tag
