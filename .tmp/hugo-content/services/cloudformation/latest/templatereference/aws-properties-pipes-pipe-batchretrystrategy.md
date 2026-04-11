This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe BatchRetryStrategy

The retry strategy that's associated with a job. For more information, see [Automated job retries](../../../batch/latest/userguide/job-retries.md) in the _AWS Batch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attempts" : Integer
}

```

### YAML

```yaml

  Attempts: Integer

```

## Properties

`Attempts`

The number of times to move a job to the `RUNNABLE` status. If the value of
`attempts` is greater than one, the job is retried on failure the same number
of attempts as the value.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchResourceRequirement

CapacityProviderStrategyItem

All content copied from https://docs.aws.amazon.com/.
