This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask CloudWatchOutputConfig

Configuration options for sending command output to Amazon CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchLogGroupName" : String,
  "CloudWatchOutputEnabled" : Boolean
}

```

### YAML

```yaml

  CloudWatchLogGroupName: String
  CloudWatchOutputEnabled: Boolean

```

## Properties

`CloudWatchLogGroupName`

The name of the CloudWatch Logs log group where you want to send command output. If you
don't specify a group name, AWS Systems Manager automatically creates a log group for you. The log group
uses the following naming format:

`aws/ssm/SystemsManagerDocumentName`

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchOutputEnabled`

Enables Systems Manager to send command output to CloudWatch Logs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSM::MaintenanceWindowTask

LoggingInfo

All content copied from https://docs.aws.amazon.com/.
