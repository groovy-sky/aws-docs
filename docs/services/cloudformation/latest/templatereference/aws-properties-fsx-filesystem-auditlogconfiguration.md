This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::FileSystem AuditLogConfiguration

The configuration that Amazon FSx for Windows File Server uses to audit and log
user accesses of files, folders, and file shares on the Amazon FSx for Windows File Server
file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuditLogDestination" : String,
  "FileAccessAuditLogLevel" : String,
  "FileShareAccessAuditLogLevel" : String
}

```

### YAML

```yaml

  AuditLogDestination: String
  FileAccessAuditLogLevel: String
  FileShareAccessAuditLogLevel: String

```

## Properties

`AuditLogDestination`

The Amazon Resource Name (ARN) for the destination of the audit logs.
The destination can be any Amazon CloudWatch Logs log group ARN or
Amazon Kinesis Data Firehose delivery stream ARN.

The name of the Amazon CloudWatch Logs log group must begin with
the `/aws/fsx` prefix. The name of the Amazon Kinesis Data
Firehose delivery stream must begin with the `aws-fsx` prefix.

The destination ARN (either CloudWatch Logs log group or Kinesis
Data Firehose delivery stream) must be in the same AWS partition,
AWS Region, and AWS account as your Amazon FSx file system.

_Required_: No

_Type_: String

_Pattern_: `^arn:[^:]{1,63}:[^:]{0,63}:[^:]{0,63}:(?:|\d{12}):[^/].{0,1023}$`

_Minimum_: `8`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileAccessAuditLogLevel`

Sets which attempt type is logged by Amazon FSx for file and folder accesses.

- `SUCCESS_ONLY` \- only successful attempts to access files
or folders are logged.

- `FAILURE_ONLY` \- only failed attempts to access files
or folders are logged.

- `SUCCESS_AND_FAILURE` \- both successful attempts and
failed attempts to access files or folders are logged.

- `DISABLED` \- access auditing of files and folders is turned off.

_Required_: Yes

_Type_: String

_Allowed values_: `DISABLED | SUCCESS_ONLY | FAILURE_ONLY | SUCCESS_AND_FAILURE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileShareAccessAuditLogLevel`

Sets which attempt type is logged by Amazon FSx for file share accesses.

- `SUCCESS_ONLY` \- only successful attempts to access file
shares are logged.

- `FAILURE_ONLY` \- only failed attempts to access file
shares are logged.

- `SUCCESS_AND_FAILURE` \- both successful attempts and
failed attempts to access file shares are logged.

- `DISABLED` \- access auditing of file shares is turned off.

_Required_: Yes

_Type_: String

_Allowed values_: `DISABLED | SUCCESS_ONLY | FAILURE_ONLY | SUCCESS_AND_FAILURE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::FSx::FileSystem

ClientConfigurations

All content copied from https://docs.aws.amazon.com/.
