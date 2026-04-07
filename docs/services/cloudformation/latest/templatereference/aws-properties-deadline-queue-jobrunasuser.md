This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Queue JobRunAsUser

Identifies the user for a job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Posix" : PosixUser,
  "RunAs" : String,
  "Windows" : WindowsUser
}

```

### YAML

```yaml

  Posix:
    PosixUser
  RunAs: String
  Windows:
    WindowsUser

```

## Properties

`Posix`

The user and group that the jobs in the queue run as.

_Required_: No

_Type_: [PosixUser](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-deadline-queue-posixuser.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RunAs`

Specifies whether the job should run using the queue's system user or if the job should
run using the worker agent system user.

_Required_: Yes

_Type_: String

_Allowed values_: `QUEUE_CONFIGURED_USER | WORKER_AGENT_USER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Windows`

Identifies a Microsoft Windows user.

_Required_: No

_Type_: [WindowsUser](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-deadline-queue-windowsuser.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JobAttachmentSettings

PosixUser
