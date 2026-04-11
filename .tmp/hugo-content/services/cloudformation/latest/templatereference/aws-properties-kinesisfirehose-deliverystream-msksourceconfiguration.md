This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream MSKSourceConfiguration

The configuration for the Amazon MSK cluster to be used as the source for a delivery
stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationConfiguration" : AuthenticationConfiguration,
  "MSKClusterARN" : String,
  "ReadFromTimestamp" : String,
  "TopicName" : String
}

```

### YAML

```yaml

  AuthenticationConfiguration:
    AuthenticationConfiguration
  MSKClusterARN: String
  ReadFromTimestamp: String
  TopicName: String

```

## Properties

`AuthenticationConfiguration`

The authentication configuration of the Amazon MSK cluster.

_Required_: Yes

_Type_: [AuthenticationConfiguration](aws-properties-kinesisfirehose-deliverystream-authenticationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MSKClusterARN`

The ARN of the Amazon MSK cluster.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReadFromTimestamp`

The start date and time in UTC for the offset position within your MSK topic from where
Firehose begins to read. By default, this is set to timestamp when Firehose becomes Active.

If you want to create a Firehose stream with Earliest start position from SDK or CLI,
you need to set the `ReadFromTimestamp` parameter to Epoch
(1970-01-01T00:00:00Z).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TopicName`

The topic name within the Amazon MSK cluster.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9\._\-]+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KMSEncryptionConfig

OpenXJsonSerDe

All content copied from https://docs.aws.amazon.com/.
