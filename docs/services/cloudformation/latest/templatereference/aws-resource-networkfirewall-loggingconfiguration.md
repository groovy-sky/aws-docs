This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::LoggingConfiguration

Use the logging configuration to define the destinations and logging options for an firewall.

You must change the logging configuration by changing one `LogDestinationConfig` setting at a time in your `LogDestinationConfigs`.

You can make only one of the following changes to your logging configuration resource:

- Create a new log destination object by adding a single
`LogDestinationConfig` array element to
`LogDestinationConfigs`.

- Delete a log destination object by removing a single
`LogDestinationConfig` array element from
`LogDestinationConfigs`.

- Change the `LogDestination` setting in a single
`LogDestinationConfig` array element.

You can't change the `LogDestinationType` or `LogType` in a
`LogDestinationConfig`. To change these settings, delete the existing
`LogDestinationConfig` object and create a new one, in two separate modifications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkFirewall::LoggingConfiguration",
  "Properties" : {
      "EnableMonitoringDashboard" : Boolean,
      "FirewallArn" : String,
      "FirewallName" : String,
      "LoggingConfiguration" : LoggingConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::NetworkFirewall::LoggingConfiguration
Properties:
  EnableMonitoringDashboard: Boolean
  FirewallArn: String
  FirewallName: String
  LoggingConfiguration:
    LoggingConfiguration

```

## Properties

`EnableMonitoringDashboard`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirewallArn`

The Amazon Resource Name (ARN) of the firewallthat the logging configuration is associated with.
You can't change the firewall specification after you create the logging configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws.*$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FirewallName`

The name of the firewall that the logging configuration is associated with.
You can't change the firewall specification after you create the logging configuration.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoggingConfiguration`

Defines how AWS Network Firewall performs logging for a firewall.

_Required_: Yes

_Type_: [LoggingConfiguration](aws-properties-networkfirewall-loggingconfiguration-loggingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the firewall that the logging configuration is associated with. For example:

`{ "Ref": "arn:aws:network-firewall:us-east-1:012345678901:firewall/myFirewallName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

- [Create a logging configuration for CloudWatch Logs and Kinesis Data Firehose](#aws-resource-networkfirewall-loggingconfiguration--examples--Create_a_logging_configuration_for_and_Kinesis_Data_Firehose)

- [Create a logging configuration for Amazon S3](#aws-resource-networkfirewall-loggingconfiguration--examples--Create_a_logging_configuration_for_Amazon_S3)

### Create a logging configuration for CloudWatch Logs and Kinesis Data Firehose

The following shows example logging configuration specifications for alert logs that go to an Amazon CloudWatch Logs log group and flow logs that go to an Amazon Kinesis Data Firehose delivery stream.

#### JSON

```json

"SampleLoggingConfiguration": {
    "Type": "AWS::NetworkFirewall::LoggingConfiguration",
    "Properties": {
        "FirewallArn": {
            "Ref": "SampleFirewallArn"
        },
        "LoggingConfiguration": {
            "LogDestinationConfigs": [
                {
                    "LogType": "ALERT",
                    "LogDestinationType": "CloudWatchLogs",
                    "LogDestination": {
                        "logGroup": "SampleLogGroup"
                    }
                },
                {
                    "LogType": "FLOW",
                    "LogDestinationType": "KinesisDataFirehose",
                    "LogDestination": {
                        "deliveryStream": "SampleStream"
                    }
                }
            ]
        }
    }
}
```

#### YAML

```yaml

SampleLoggingConfiguration:
  Type: 'AWS::NetworkFirewall::LoggingConfiguration'
  Properties:
    FirewallArn: !Ref SampleFirewallArn
    LoggingConfiguration:
      LogDestinationConfigs:
        - LogType: ALERT
          LogDestinationType: CloudWatchLogs
          LogDestination:
            logGroup: SampleLogGroup
        - LogType: FLOW
          LogDestinationType: KinesisDataFirehose
          LogDestination:
            deliveryStream: SampleStream

```

### Create a logging configuration for Amazon S3

The following shows example logging configuration specifications for flow logs that go to an Amazon S3 bucket.

#### JSON

```json

"SampleLoggingConfiguration": {
    "Type": "AWS::NetworkFirewall::LoggingConfiguration",
    "Properties": {
        "FirewallArn": {
            "Ref": "SampleFirewallArn"
        },
        "LoggingConfiguration": {
            "LogDestinationConfigs": [
                {
                    "LogType": "FLOW",
                    "LogDestinationType": "S3",
                    "LogDestination": {
                        "bucketName": "sample-bucket-name",
                        "prefix": "sample/s3/prefix"
                    }
                }
            ]
        }
    }
}
```

#### YAML

```yaml

SampleLoggingConfiguration:
  Type: 'AWS::NetworkFirewall::LoggingConfiguration'
  Properties:
    FirewallArn: !Ref SampleFirewallArn
    LoggingConfiguration:
      LogDestinationConfigs:
        - LogType: FLOW
          LogDestinationType: S3
          LogDestination:
            bucketName: sample-bucket-name
            prefix: sample/s3/prefix

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

LogDestinationConfig

All content copied from https://docs.aws.amazon.com/.
