This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule BatchRetryStrategy

The retry strategy to use for failed jobs, if the target is an AWS Batch job.
If you specify a retry strategy here, it overrides the retry strategy defined in the job
definition.

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

The number of times to attempt to retry, if the job fails. Valid values are 1–10.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Set the BatchRetryStrategy property

The following example sets the `BatchRetryStrategy` property to 8 retries.

#### JSON

```json

"RetryStrategy": {
  "Attempts": 8
}
```

#### YAML

```yaml

RetryStrategy
  Attempts: 8
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchParameters

CapacityProviderStrategyItem

All content copied from https://docs.aws.amazon.com/.
