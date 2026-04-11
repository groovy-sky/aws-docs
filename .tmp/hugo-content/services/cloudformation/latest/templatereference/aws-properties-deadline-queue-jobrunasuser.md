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

_Type_: [PosixUser](aws-properties-deadline-queue-posixuser.md)

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

_Type_: [WindowsUser](aws-properties-deadline-queue-windowsuser.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JobAttachmentSettings

PosixUser

All content copied from https://docs.aws.amazon.com/.
