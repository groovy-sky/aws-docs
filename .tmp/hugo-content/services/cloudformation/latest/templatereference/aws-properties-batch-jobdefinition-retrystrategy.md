This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition RetryStrategy

The retry strategy that's associated with a job. For more information, see [Automated job retries](../../../batch/latest/userguide/job-retries.md) in the
_AWS Batch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attempts" : Integer,
  "EvaluateOnExit" : [ EvaluateOnExit, ... ]
}

```

### YAML

```yaml

  Attempts: Integer
  EvaluateOnExit:
    - EvaluateOnExit

```

## Properties

`Attempts`

The number of times to move a job to the `RUNNABLE` status. You can specify
between 1 and 10 attempts. If the value of `attempts` is greater than one, the job is
retried on failure the same number of attempts as the value.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluateOnExit`

Array of up to 5 objects that specify the conditions where jobs are retried or failed. If
this parameter is specified, then the `attempts` parameter must also be specified. If
none of the listed conditions match, then the job is retried.

_Required_: No

_Type_: [Array](aws-properties-batch-jobdefinition-evaluateonexit.md) of [EvaluateOnExit](aws-properties-batch-jobdefinition-evaluateonexit.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Retrying jobs

This example will retry the job attempt up to 3 times if the job status reason is
either `AGENT` or `Task failed to start`. The final rule
matches all other job failures and exits. If none of the entries in
`EvaluateOnExit` match the job failure, the job will be retried.

#### JSON

```json

{
  "Attempts": 3,
  "EvaluateOnExit": [
    {
      "Action": "RETRY",
      "OnReason": "AGENT"
    },
    {
      "Action": "RETRY",
      "OnReason": "Task failed to start"
    },
    {
      "Action": "EXIT",
      "OnReason": "*"
    }
  ]
}
```

#### YAML

```yaml

Attempts: 3
EvaluateOnExit:
  - Action: RETRY
    OnReason: AGENT
  - Action: RETRY
    OnReason: Task failed to start
  - Action: EXIT
    OnReason: '*'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceRetentionPolicy

RuntimePlatform

All content copied from https://docs.aws.amazon.com/.
