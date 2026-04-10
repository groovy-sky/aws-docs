This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::MaintenanceWindowTask MaintenanceWindowRunCommandParameters

The `MaintenanceWindowRunCommandParameters` property type specifies the
parameters for a `RUN_COMMAND` task type for a maintenance window task in
AWS Systems Manager. This means that these parameters are the same as those for
the `SendCommand` API call. For more information about
`SendCommand` parameters, see [SendCommand](../../../../reference/systems-manager/latest/apireference/api-sendcommand.md)
in the _AWS Systems Manager API Reference_.

For information about available parameters in SSM Command documents, you can view the
content of the document itself in the Systems Manager console. For information, see
[Viewing SSM\
command document content](../../../systems-manager/latest/userguide/viewing-ssm-document-content.md) in the _AWS Systems Manager User_
_Guide_.

`MaintenanceWindowRunCommandParameters` is a property of the [TaskInvocationParameters](../userguide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudWatchOutputConfig" : CloudWatchOutputConfig,
  "Comment" : String,
  "DocumentHash" : String,
  "DocumentHashType" : String,
  "DocumentVersion" : String,
  "NotificationConfig" : NotificationConfig,
  "OutputS3BucketName" : String,
  "OutputS3KeyPrefix" : String,
  "Parameters" : Json,
  "ServiceRoleArn" : String,
  "TimeoutSeconds" : Integer
}

```

### YAML

```yaml

  CloudWatchOutputConfig:
    CloudWatchOutputConfig
  Comment: String
  DocumentHash: String
  DocumentHashType: String
  DocumentVersion: String
  NotificationConfig:
    NotificationConfig
  OutputS3BucketName: String
  OutputS3KeyPrefix: String
  Parameters: Json
  ServiceRoleArn: String
  TimeoutSeconds: Integer

```

## Properties

`CloudWatchOutputConfig`

Configuration options for sending command output to Amazon CloudWatch Logs.

_Required_: No

_Type_: [CloudWatchOutputConfig](aws-properties-ssm-maintenancewindowtask-cloudwatchoutputconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Comment`

Information about the command or commands to run.

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentHash`

The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes
have been deprecated.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentHashType`

The SHA-256 or SHA-1 hash type. SHA-1 hashes are deprecated.

_Required_: No

_Type_: String

_Allowed values_: `Sha256 | Sha1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentVersion`

The AWS Systems Manager document (SSM document) version to use in the request. You can specify
`$DEFAULT`, `$LATEST`, or a specific version number. If you run commands
by using the AWS CLI, then you must escape the first two options by using a backslash. If you
specify a version number, then you don't need to use the backslash. For example:

`--document-version "\$DEFAULT"`

`--document-version "\$LATEST"`

`--document-version "3"`

_Required_: No

_Type_: String

_Pattern_: `([$]LATEST|[$]DEFAULT|^[1-9][0-9]*$)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationConfig`

Configurations for sending notifications about command status changes on a per-managed node
basis.

_Required_: No

_Type_: [NotificationConfig](aws-properties-ssm-maintenancewindowtask-notificationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputS3BucketName`

The name of the Amazon Simple Storage Service (Amazon S3) bucket.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputS3KeyPrefix`

The S3 bucket subfolder.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

The parameters for the `RUN_COMMAND` task execution.

The supported parameters are the same as those for the `SendCommand` API
call. For more information, see [SendCommand](../../../../reference/systems-manager/latest/apireference/api-sendcommand.md)
in the _AWS Systems Manager API Reference_.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceRoleArn`

The Amazon Resource Name (ARN) of the IAM service role for
AWS Systems Manager to assume when running a maintenance window task. If you do not specify a
service role ARN, Systems Manager uses a service-linked role in your account. If no
appropriate service-linked role for Systems Manager exists in your account, it is created when
you run `RegisterTaskWithMaintenanceWindow`.

However, for an improved security posture, we strongly recommend creating a custom
policy and custom service role for running your maintenance window tasks. The policy
can be crafted to provide only the permissions needed for your particular
maintenance window tasks. For more information, see [Setting up Maintenance Windows](../../../systems-manager/latest/userguide/sysman-maintenance-permissions.md) in the in the
_AWS Systems Manager User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutSeconds`

If this time is reached and the command hasn't already started running, it doesn't
run.

_Required_: No

_Type_: Integer

_Minimum_: `30`

_Maximum_: `2592000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MaintenanceWindowLambdaParameters

MaintenanceWindowStepFunctionsParameters

All content copied from https://docs.aws.amazon.com/.
