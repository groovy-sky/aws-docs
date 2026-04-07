This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::Detector

The `AWS::GuardDuty::Detector` resource specifies a new GuardDuty
detector. A detector is an object that represents the GuardDuty service. A
detector is required for GuardDuty to become operational.

Make sure you use either `DataSources` or `Features` in a one
request, and not both.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GuardDuty::Detector",
  "Properties" : {
      "DataSources" : CFNDataSourceConfigurations,
      "Enable" : Boolean,
      "Features" : [ CFNFeatureConfiguration, ... ],
      "FindingPublishingFrequency" : String,
      "Tags" : [ TagItem, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::GuardDuty::Detector
Properties:
  DataSources:
    CFNDataSourceConfigurations
  Enable: Boolean
  Features:
    - CFNFeatureConfiguration
  FindingPublishingFrequency: String
  Tags:
    - TagItem

```

## Properties

`DataSources`

Describes which data sources will be enabled for the detector.

_Required_: No

_Type_: [CFNDataSourceConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-guardduty-detector-cfndatasourceconfigurations.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enable`

Specifies whether the detector is to be enabled on creation.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Features`

A list of features that will be configured for the detector.

_Required_: No

_Type_: Array of [CFNFeatureConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-guardduty-detector-cfnfeatureconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingPublishingFrequency`

Specifies how frequently updated findings are exported.

_Required_: No

_Type_: String

_Allowed values_: `FIFTEEN_MINUTES | ONE_HOUR | SIX_HOURS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies tags added to a new detector resource. Each tag consists of a key and an
optional value, both of which you define.

Currently, support is available only for creating and deleting a tag. No support exists
for updating the tags.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [TagItem](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-guardduty-detector-tagitem.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique ID of the detector.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The unique ID of the detector.

## Examples

### Declare a Detector Resource

The following example shows how to declare a GuardDuty `Detector` resource:

#### JSON

```json

"mydetector": {
    "Type" : "AWS::GuardDuty::Detector",
    "Properties" : {
        "Enable" : True,
        "FindingPublishingFrequency" : "FIFTEEN_MINUTES"
    }
}
```

#### YAML

```yaml

mydetector:
    Type: AWS::GuardDuty::Detector
    Properties:
        Enable: True
        FindingPublishingFrequency: FIFTEEN_MINUTES
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon GuardDuty

CFNDataSourceConfigurations
