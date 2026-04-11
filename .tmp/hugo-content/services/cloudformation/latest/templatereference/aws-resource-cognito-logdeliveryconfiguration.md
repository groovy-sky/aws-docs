This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::LogDeliveryConfiguration

Sets up or modifies the logging configuration of a user pool. User pools can export
user notification logs and, when threat protection is active, user-activity logs. For
more information, see [Exporting user\
pool logs](../../../cognito/latest/developerguide/exporting-quotas-and-usage.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::LogDeliveryConfiguration",
  "Properties" : {
      "LogConfigurations" : [ LogConfiguration, ... ],
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::LogDeliveryConfiguration
Properties:
  LogConfigurations:
    - LogConfiguration
  UserPoolId: String

```

## Properties

`LogConfigurations`

A logging destination of a user pool. User pools can have multiple logging
destinations for message-delivery and user-activity logs.

_Required_: No

_Type_: Array of [LogConfiguration](aws-properties-cognito-logdeliveryconfiguration-logconfiguration.md)

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The ID of the user pool where you configured logging.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a resource identifier. A log delivery configuration
attached to a user pool returns a user pool ID like
`us-east-1_EXAMPLE`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

A user pool ID, for example `us-east-1_EXAMPLE`.

## Examples

### Creating a new log delivery configuration for a user pool

The following example creates log delivery of user message-delivery errors to
a log group and threat-protection logs to a stream.

#### JSON

```json

{
    "LogDeliveryConfiguration": {
        "Properties": {
            "LogConfigurations": [
                {
                    "CloudWatchLogsConfiguration": {
                        "LogGroupArn": "arn:aws:logs:us-west-2:123456789012:log-group:cognito-exported"
                    },
                    "EventSource": "userNotification",
                    "LogLevel": "ERROR"
                },
                {
                    "EventSource": "userAuthEvents",
                    "FirehoseConfiguration": {
                        "StreamArn": "arn:aws:firehose:us-west-2:123456789012:deliverystream/test-deliverystream"
                    },
                    "LogLevel": "INFO"
                }
            ],
            "UserPoolId": "us-west-2_EXAMPLE"
        },
        "Type": "AWS::Cognito::LogDeliveryConfiguration"
    }
}
```

#### YAML

```yaml

LogDeliveryConfiguration:
    Type: AWS::Cognito::LogDeliveryConfiguration
    Properties:
      LogConfigurations:
        - CloudWatchLogsConfiguration:
            LogGroupArn: arn:aws:logs:us-west-2:123456789012:log-group:cognito-exported
          EventSource: userNotification
          LogLevel: ERROR
        - EventSource: userAuthEvents
          FirehoseConfiguration:
            StreamArn: arn:aws:firehose:us-west-2:123456789012:deliverystream/test-deliverystream
          LogLevel: INFO
      UserPoolId: us-west-2_EXAMPLE
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RulesConfigurationType

CloudWatchLogsConfiguration

All content copied from https://docs.aws.amazon.com/.
