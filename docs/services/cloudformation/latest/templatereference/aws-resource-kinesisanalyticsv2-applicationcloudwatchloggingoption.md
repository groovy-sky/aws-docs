This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption

Adds an Amazon CloudWatch log stream to monitor application configuration
errors.

###### Note

Only one _ApplicationCloudWatchLoggingOption_ resource can be
attached per application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption",
  "Properties" : {
      "ApplicationName" : String,
      "CloudWatchLoggingOption" : CloudWatchLoggingOption
    }
}

```

### YAML

```yaml

Type: AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption
Properties:
  ApplicationName: String
  CloudWatchLoggingOption:
    CloudWatchLoggingOption

```

## Properties

`ApplicationName`

The name of the application.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CloudWatchLoggingOption`

Provides a description of Amazon CloudWatch logging options, including the log stream
Amazon Resource Name (ARN).

_Required_: Yes

_Type_: [CloudWatchLoggingOption](aws-properties-kinesisanalyticsv2-applicationcloudwatchloggingoption-cloudwatchloggingoption.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

## Examples

### Create an ApplicationCloudWatchLoggingOption resource

#### JSON

```json

{
    "BasicApplicationV2CloudWatchLoggingOption": {
        "Type": "AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption",
        "Properties": {
            "ApplicationName": {
                "Ref": "BasicApplication"
            },
            "CloudWatchLoggingOption": {
                "LogStreamARN": {
                    "Fn::Join": [
                        ":",
                        [
                            "arn:aws:logs",
                            {
                                "Ref": "AWS::Region"
                            },
                            {
                                "Ref": "AWS::AccountId"
                            },
                            "log-group",
                            {
                                "Ref": "TestCWLogGroup"
                            },
                            "log-stream",
                            {
                                "Ref": "TestCWLogStream"
                            }
                        ]
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

BasicApplicationV2CloudWatchLoggingOption:
  Type: AWS::KinesisAnalyticsV2::ApplicationCloudWatchLoggingOption
  Properties:
    ApplicationName:
      Ref: BasicApplication
    CloudWatchLoggingOption:
      LogStreamARN:
        Fn::Join:
        - ":"
        - - arn:aws:logs
          - Ref: AWS::Region
          - Ref: AWS::AccountId
          - log-group
          - Ref: TestCWLogGroup
          - log-stream
          - Ref: TestCWLogStream

```

## See also

- [AddApplicationCloudWatchLoggingOption](../../../managed-flink/latest/apiv2/api-addapplicationcloudwatchloggingoption.md) in the _Amazon_
_Kinesis Data Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ZeppelinMonitoringConfiguration

CloudWatchLoggingOption

All content copied from https://docs.aws.amazon.com/.
