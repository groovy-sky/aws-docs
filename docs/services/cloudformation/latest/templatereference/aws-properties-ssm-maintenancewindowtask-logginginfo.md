This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask LoggingInfo

###### Note

`LoggingInfo` has been deprecated. To specify an Amazon S3
bucket to contain logs, instead use the `OutputS3BucketName` and
`OutputS3KeyPrefix` options in the
`TaskInvocationParameters` structure. For information about how
Systems Manager handles these options for the supported maintenance window
task types, see [AWS::SSM::MaintenanceWindowTask\
MaintenanceWindowRunCommandParameters](../userguide/aws-properties-ssm-maintenancewindowtask-maintenancewindowruncommandparameters.md).

The `LoggingInfo` property type specifies information about the Amazon S3 bucket to write instance-level logs to.

`LoggingInfo` is a property of the [AWS::SSM::MaintenanceWindowTask](../userguide/aws-resource-ssm-maintenancewindowtask.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Region" : String,
  "S3Bucket" : String,
  "S3Prefix" : String
}

```

### YAML

```yaml

  Region: String
  S3Bucket: String
  S3Prefix: String

```

## Properties

`Region`

The AWS Region where the S3 bucket is located.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Bucket`

The name of an S3 bucket where execution logs are stored.

_Required_: Yes

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Prefix`

The Amazon S3 bucket subfolder.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchOutputConfig

MaintenanceWindowAutomationParameters

All content copied from https://docs.aws.amazon.com/.
