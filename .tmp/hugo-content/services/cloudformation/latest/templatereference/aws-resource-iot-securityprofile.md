This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile

Use the `AWS::IoT::SecurityProfile` resource to create a Device Defender
security profile. For API reference, see [CreateSecurityProfile](../../../../reference/iot/latest/apireference/api-createsecurityprofile.md) and for general information, see [Detect](../../../iot/latest/developerguide/device-defender-detect.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::SecurityProfile",
  "Properties" : {
      "AdditionalMetricsToRetainV2" : [ MetricToRetain, ... ],
      "AlertTargets" : {Key: Value, ...},
      "Behaviors" : [ Behavior, ... ],
      "MetricsExportConfig" : MetricsExportConfig,
      "SecurityProfileDescription" : String,
      "SecurityProfileName" : String,
      "Tags" : [ Tag, ... ],
      "TargetArns" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoT::SecurityProfile
Properties:
  AdditionalMetricsToRetainV2:
    - MetricToRetain
  AlertTargets:
    Key: Value
  Behaviors:
    - Behavior
  MetricsExportConfig:
    MetricsExportConfig
  SecurityProfileDescription: String
  SecurityProfileName: String
  Tags:
    - Tag
  TargetArns:
    - String

```

## Properties

`AdditionalMetricsToRetainV2`

A list of metrics whose data is retained (stored). By default, data is retained for any
metric used in the profile's `behaviors`, but it's also retained for any metric
specified here. Can be used with custom metrics; can't be used with dimensions.

_Required_: No

_Type_: Array of [MetricToRetain](aws-properties-iot-securityprofile-metrictoretain.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlertTargets`

Specifies the destinations to which alerts are sent. (Alerts are always sent to the
console.) Alerts are generated when a device (thing) violates a behavior.

_Required_: No

_Type_: Object of [AlertTarget](aws-properties-iot-securityprofile-alerttarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Behaviors`

Specifies the behaviors that, when violated by a device (thing), cause an alert.

_Required_: No

_Type_: Array of [Behavior](aws-properties-iot-securityprofile-behavior.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricsExportConfig`

Specifies the MQTT topic and role ARN required for metric export.

_Required_: No

_Type_: [MetricsExportConfig](aws-properties-iot-securityprofile-metricsexportconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityProfileDescription`

A description of the security profile.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityProfileName`

The name you gave to the security profile.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that can be used to manage the security profile.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-securityprofile-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArns`

The ARN of the target (thing group) to which the security profile is attached.

_Required_: No

_Type_: Array of String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the security profile name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`SecurityProfileArn`

The Amazon Resource Name (ARN) of the security profile.

## Examples

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Amazon Web Services IoT SecurityProfile Sample Template",
  "Resources": {
    "MySecurityProfile": {
      "Type": "AWS::IoT::SecurityProfile",
      "Properties": {
        "AdditionalMetricsToRetainV2": [
          {
            "Metric": "aws:num-messages-received"
          },
          {
            "Metric": "aws:num-disconnects"
          }
        ],
        "AlertTargets": {
          "SNS": {
            "AlertTargetArn": "arn:aws:sns:us-east-1:123456789012:DeviceDefenderDetectAlerts",
            "RoleArn": "arn:aws:iam::123456789012:role/RoleForDefenderAlerts"
          }
        },
        "Behaviors": [
          {
            "Name": "MaxMessageSize",
            "Metric": "aws:message-byte-size",
            "Criteria": {
              "ConsecutiveDatapointsToAlarm": 1,
              "ConsecutiveDatapointsToClear": 1,
              "ComparisonOperator": "less-than-equals",
              "Value": {
                "Count": 5
              }
            }
          },
          {
            "Name": "OutboundMessageCount",
            "Metric": "aws:num-messages-sent",
            "Criteria": {
              "DurationSeconds": 300,
              "ComparisonOperator": "less-than-equals",
              "Value": {
                "Count": 50
              }
            }
          },
          {
            "Name": "AuthFailuresStatThreshold",
            "Metric": "aws:num-authorization-failures",
            "Criteria": {
              "ComparisonOperator": "less-than-equals",
              "DurationSeconds": 300,
              "StatisticalThreshold": {
                "Statistic": "p90"
              }
            }
          }
        ],
        "SecurityProfileDescription": "Contains expected behaviors for connected devices",
        "SecurityProfileName": "ProfileForConnectedDevices",
        "Tags": [
          {
            "Key": "Application",
            "Value": "SmartHome"
          }
        ],
        "TargetArns": [
          "arn:aws:iot:us-east-1:123456789012:all/things"
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Amazon Web Services IoT SecurityProfile Sample Template
Resources:
  MySecurityProfile:
    Type: 'AWS::IoT::SecurityProfile'
    Properties:
      AdditionalMetricsToRetainV2:
        - Metric: 'aws:num-messages-received'
        - Metric: 'aws:num-disconnects'
      AlertTargets:
        SNS:
          AlertTargetArn: 'arn:aws:sns:us-east-1:123456789012:DeviceDefenderDetectAlerts'
          RoleArn: 'arn:aws:iam::123456789012:role/RoleForDefenderAlerts'
      Behaviors:
        - Name: MaxMessageSize
          Metric: 'aws:message-byte-size'
          Criteria:
            ConsecutiveDatapointsToAlarm: 1
            ConsecutiveDatapointsToClear: 1
            ComparisonOperator: less-than-equals
            Value:
              Count: 5
        - Name: OutboundMessageCount
          Metric: 'aws:num-messages-sent'
          Criteria:
            DurationSeconds: 300
            ComparisonOperator: less-than-equals
            Value:
              Count: 50
        - Name: AuthFailuresStatThreshold
          Metric: 'aws:num-authorization-failures'
          Criteria:
            ComparisonOperator: less-than-equals
            DurationSeconds: 300
            StatisticalThreshold:
              Statistic: p90
      SecurityProfileDescription: Contains expected behaviors for connected devices
      SecurityProfileName: ProfileForConnectedDevices
      Tags:
        - Key: Application
          Value: SmartHome
      TargetArns:
        - 'arn:aws:iot:us-east-1:123456789012:all/things'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AlertTarget

All content copied from https://docs.aws.amazon.com/.
