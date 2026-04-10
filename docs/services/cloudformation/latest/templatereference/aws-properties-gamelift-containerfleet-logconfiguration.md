This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::ContainerFleet LogConfiguration

A method for collecting container logs for the fleet. Amazon GameLift Servers saves all standard
output for each container in logs, including game session logs. You can select from the
following methods:

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogDestination" : String,
  "LogGroupArn" : String,
  "S3BucketName" : String
}

```

### YAML

```yaml

  LogDestination: String
  LogGroupArn: String
  S3BucketName: String

```

## Properties

`LogDestination`

The type of log collection to use for a fleet.

- `CLOUDWATCH` \-\- (default value) Send logs to an Amazon CloudWatch log group that you define. Each container emits a log stream, which is organized in the log group.

- `S3` \-\- Store logs in an Amazon S3 bucket that you define. This bucket must reside in the fleet's home AWS Region.

- `NONE` \-\- Don't collect container logs.

_Required_: No

_Type_: String

_Allowed values_: `NONE | CLOUDWATCH | S3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupArn`

If log destination is `CLOUDWATCH`, logs are sent to the specified log group in Amazon CloudWatch.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:/\-\*]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3BucketName`

If log destination is `S3`, logs are sent to the specified Amazon S3 bucket name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocationConfiguration

ManagedCapacityConfiguration

All content copied from https://docs.aws.amazon.com/.
