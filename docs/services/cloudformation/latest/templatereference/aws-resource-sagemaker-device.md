This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Device

The `AWS::SageMaker::Device` resource is an Amazon SageMaker resource type that allows you to
register your Devices against an existing SageMaker Edge Manager DeviceFleet. Each device must be listed individually
in the CFN specification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::Device",
  "Properties" : {
      "Device" : Device,
      "DeviceFleetName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::Device
Properties:
  Device:
    Device
  DeviceFleetName: String
  Tags:
    - Tag

```

## Properties

`Device`

Edge device you want to create.

_Required_: No

_Type_: [Device](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-device-device.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceFleetName`

The name of the fleet the device belongs to.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*_*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs that contain metadata to help you categorize and organize your devices. Each tag
consists of a key and a value, both of which you define.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-device-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DeviceFleetName.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

Device
