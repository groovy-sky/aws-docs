This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EvaluateOnExit

Specifies an array of up to 5 conditions to be met, and an action to take
( `RETRY` or `EXIT`) if all conditions are met. If none of the
`EvaluateOnExit` conditions in a `RetryStrategy` match, then the job is
retried.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : String,
  "OnExitCode" : String,
  "OnReason" : String,
  "OnStatusReason" : String
}

```

### YAML

```yaml

  Action: String
  OnExitCode: String
  OnReason: String
  OnStatusReason: String

```

## Properties

`Action`

Specifies the action to take if all of the specified conditions
( `onStatusReason`, `onReason`, and `onExitCode`) are met. The
values aren't case sensitive.

_Required_: Yes

_Type_: String

_Allowed values_: `RETRY | EXIT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnExitCode`

Contains a glob pattern to match against the decimal representation of the
`ExitCode` returned for a job. The pattern can be up to 512 characters long. It can
contain only numbers, and can end with an asterisk (\*) so that only the start of the string needs
to be an exact match.

The string can contain up to 512 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnReason`

Contains a glob pattern to match against the `Reason` returned for a job. The
pattern can contain up to 512 characters. It can contain letters, numbers, periods (.), colons
(:), and white space (including spaces and tabs). It can optionally end with an asterisk (\*) so
that only the start of the string needs to be an exact match.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnStatusReason`

Contains a glob pattern to match against the `StatusReason` returned for a job.
The pattern can contain up to 512 characters. It can contain letters, numbers, periods (.),
colons (:), and white spaces (including spaces or tabs). It can optionally end with an asterisk (\*)
so that only the start of the string needs to be an exact match.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EphemeralStorage

FargatePlatformConfiguration

All content copied from https://docs.aws.amazon.com/.
