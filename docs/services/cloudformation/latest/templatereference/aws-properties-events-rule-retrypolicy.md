This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule RetryPolicy

A `RetryPolicy` object that includes information about the retry policy
settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumEventAgeInSeconds" : Integer,
  "MaximumRetryAttempts" : Integer
}

```

### YAML

```yaml

  MaximumEventAgeInSeconds: Integer
  MaximumRetryAttempts: Integer

```

## Properties

`MaximumEventAgeInSeconds`

The maximum amount of time, in seconds, to continue to make retry attempts.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumRetryAttempts`

The maximum number of retry attempts to make before the request fails. Retry attempts
continue until either the maximum number of attempts is made or until the duration of the
`MaximumEventAgeInSeconds` is met.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `185`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftDataParameters

RunCommandParameters

All content copied from https://docs.aws.amazon.com/.
