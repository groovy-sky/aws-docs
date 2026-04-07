This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::Alarm Dimension

Dimension is an embedded property of the `AWS::CloudWatch::Alarm` type. Dimensions
are name/value pairs that can be associated with a CloudWatch metric. You can
specify a maximum of 30 dimensions for a given metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name of the dimension, from 1–255 characters in length. This dimension
name must have been included when
the metric was published.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for the dimension, from 1–255 characters in length.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Two alarms with dimension values supplied by the Ref function

The [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html) and [GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html) intrinsic functions
are often used to supply values for CloudWatch metric dimensions.
Here is an example using the Ref function.

#### JSON

```json

{
    "CPUAlarmHigh": {
        "Type": "AWS::CloudWatch::Alarm",
        "Properties": {
            "AlarmDescription": "Scale-up if CPU is greater than 90% for 10 minutes",
            "MetricName": "CPUUtilization",
            "Namespace": "AWS/EC2",
            "Statistic": "Average",
            "Period": "300",
            "EvaluationPeriods": "2",
            "Threshold": "90",
            "AlarmActions": [
                {
                    "Ref": "WebServerScaleUpPolicy"
                }
            ],
            "Dimensions": [
                {
                    "Name": "AutoScalingGroupName",
                    "Value": {
                        "Ref": "WebServerGroup"
                    }
                }
            ],
            "ComparisonOperator": "GreaterThanThreshold"
        }
    },
    "CPUAlarmLow": {
        "Type": "AWS::CloudWatch::Alarm",
        "Properties": {
            "AlarmDescription": "Scale-down if CPU is less than 70% for 10 minutes",
            "MetricName": "CPUUtilization",
            "Namespace": "AWS/EC2",
            "Statistic": "Average",
            "Period": "300",
            "EvaluationPeriods": "2",
            "Threshold": "70",
            "AlarmActions": [
                {
                    "Ref": "WebServerScaleDownPolicy"
                }
            ],
            "Dimensions": [
                {
                    "Name": "AutoScalingGroupName",
                    "Value": {
                        "Ref": "WebServerGroup"
                    }
                }
            ],
            "ComparisonOperator": "LessThanThreshold"
        }
    }
}
```

#### YAML

```yaml

CPUAlarmHigh:
  Type: 'AWS::CloudWatch::Alarm'
  Properties:
    AlarmDescription: Scale-up if CPU is greater than 90% for 10 minutes
    MetricName: CPUUtilization
    Namespace: AWS/EC2
    Statistic: Average
    Period: '300'
    EvaluationPeriods: '2'
    Threshold: '90'
    AlarmActions:
      - !Ref WebServerScaleUpPolicy
    Dimensions:
      - Name: AutoScalingGroupName
        Value: !Ref WebServerGroup
    ComparisonOperator: GreaterThanThreshold
CPUAlarmLow:
  Type: 'AWS::CloudWatch::Alarm'
  Properties:
    AlarmDescription: Scale-down if CPU is less than 70% for 10 minutes
    MetricName: CPUUtilization
    Namespace: AWS/EC2
    Statistic: Average
    Period: '300'
    EvaluationPeriods: '2'
    Threshold: '70'
    AlarmActions:
      - !Ref WebServerScaleDownPolicy
    Dimensions:
      - Name: AutoScalingGroupName
        Value: !Ref WebServerGroup
    ComparisonOperator: LessThanThreshold
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudWatch::Alarm

Metric
