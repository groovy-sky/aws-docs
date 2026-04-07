This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::DeviceFleet

The `AWS::SageMaker::DeviceFleet` resource is an Amazon SageMaker resource type that allows you to
create a DeviceFleet that manages your SageMaker Edge Manager Devices. You must register your devices against the
`DeviceFleet` separately.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::DeviceFleet",
  "Properties" : {
      "Description" : String,
      "DeviceFleetName" : String,
      "OutputConfig" : EdgeOutputConfig,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::DeviceFleet
Properties:
  Description: String
  DeviceFleetName: String
  OutputConfig:
    EdgeOutputConfig
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the fleet.

_Required_: No

_Type_: String

_Pattern_: `[\S\s]+`

_Minimum_: `0`

_Maximum_: `800`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceFleetName`

Name of the device fleet.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*_*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputConfig`

The output configuration for storing sample data collected by the fleet.

_Required_: Yes

_Type_: [EdgeOutputConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-devicefleet-edgeoutputconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) that has access to AWS Internet of Things (IoT).

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z\-]*:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs that contain metadata to help you categorize and organize your device fleets. Each
tag consists of a key and a value, both of which you define.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sagemaker-devicefleet-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DeviceFleetName.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

EdgeOutputConfig
