This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::ComputationModel

Create a computation model with a configuration and data binding.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTSiteWise::ComputationModel",
  "Properties" : {
      "ComputationModelConfiguration" : ComputationModelConfiguration,
      "ComputationModelDataBinding" : {Key: Value, ...},
      "ComputationModelDescription" : String,
      "ComputationModelName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTSiteWise::ComputationModel
Properties:
  ComputationModelConfiguration:
    ComputationModelConfiguration
  ComputationModelDataBinding:
    Key: Value
  ComputationModelDescription: String
  ComputationModelName: String
  Tags:
    - Tag

```

## Properties

`ComputationModelConfiguration`

The configuration for the computation model.

_Required_: Yes

_Type_: [ComputationModelConfiguration](aws-properties-iotsitewise-computationmodel-computationmodelconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputationModelDataBinding`

The data binding for the computation model. Key is a variable name defined in configuration.
Value is a `ComputationModelDataBindingValue` referenced by the variable.

_Required_: Yes

_Type_: Object of [ComputationModelDataBindingValue](aws-properties-iotsitewise-computationmodel-computationmodeldatabindingvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputationModelDescription`

The description of the computation model.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9 _\-#$*!@]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputationModelName`

The name of the computation model.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9 _\-#$*!@]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that contain metadata for the asset. For more information, see
[Tagging your AWS IoT SiteWise\
resources](../../../iot-sitewise/latest/userguide/tag-resources.md) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotsitewise-computationmodel-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ComputationModelId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ComputationModelArn`

The ARN of the computation model, which has the following format.

`arn:${Partition}:iotsitewise:${Region}:${Account}:computation-model/${ComputationModelId}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`ComputationModelId`

The ID of the computation model.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VariableValue

AnomalyDetectionComputationModelConfiguration

All content copied from https://docs.aws.amazon.com/.
