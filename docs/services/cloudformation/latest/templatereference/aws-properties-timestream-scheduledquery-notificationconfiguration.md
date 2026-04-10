This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery NotificationConfiguration

Notification configuration for a scheduled query. A notification is sent by Timestream
when a scheduled query is created, its state is updated or when it is deleted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SnsConfiguration" : SnsConfiguration
}

```

### YAML

```yaml

  SnsConfiguration:
    SnsConfiguration

```

## Properties

`SnsConfiguration`

Details on SNS configuration.

_Required_: Yes

_Type_: [SnsConfiguration](aws-properties-timestream-scheduledquery-snsconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiMeasureMappings

S3Configuration

All content copied from https://docs.aws.amazon.com/.
