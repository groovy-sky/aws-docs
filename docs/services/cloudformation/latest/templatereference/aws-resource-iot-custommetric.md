This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::CustomMetric

Use the `AWS::IoT::CustomMetric` resource to define a custom metric
published by your devices to Device Defender. For API reference, see [CreateCustomMetric](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateCustomMetric.html) and for general information, see [Custom\
metrics](https://docs.aws.amazon.com/iot/latest/developerguide/dd-detect-custom-metrics.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::CustomMetric",
  "Properties" : {
      "DisplayName" : String,
      "MetricName" : String,
      "MetricType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoT::CustomMetric
Properties:
  DisplayName: String
  MetricName: String
  MetricType: String
  Tags:
    - Tag

```

## Properties

`DisplayName`

The friendly name in the console for the custom metric. This name doesn't have to be
unique. Don't use this name as the metric identifier in the device metric report. You can
update the friendly name after you define it.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the custom metric. This will be used in the metric report submitted from
the device/thing. The name can't begin with `aws:`. You can’t change the name
after you define it.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricType`

The type of the custom metric. Types include `string-list`,
`ip-address-list`, `number-list`, and `number`.

###### Important

The type `number` only takes a single metric value as an input, but when
you submit the metrics value in the DeviceMetrics report, you must pass it as an array
with a single value.

_Required_: Yes

_Type_: String

_Allowed values_: `string-list | ip-address-list | number-list | number`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that can be used to manage the custom metric.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-custommetric-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the custom metric name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`MetricArn`

The Amazon Resource Number (ARN) of the custom metric; for example,
`arn:aws-partition:iot:region:accountId:custommetric/metricName`.

## Examples

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Amazon Web Services IoT CustomMetric Sample Template",
  "Resources": {
    "BatteryPercentageMetric": {
      "Type": "AWS::IoT::CustomMetric",
      "Properties": {
        "MetricName": "batteryPercentage",
        "DisplayName": "Remaining battery percentage",
        "MetricType": "number"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Amazon Web Services IoT CustomMetric Sample Template
Resources:
  BatteryPercentageMetric:
    Type: AWS::IoT::CustomMetric
    Properties:
      MetricName: batteryPercentage
      DisplayName: Remaining battery percentage
      MetricType: number
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
