This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule BatchParameters

The custom parameters to be used when the target is an AWS Batch job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArrayProperties" : BatchArrayProperties,
  "JobDefinition" : String,
  "JobName" : String,
  "RetryStrategy" : BatchRetryStrategy
}

```

### YAML

```yaml

  ArrayProperties:
    BatchArrayProperties
  JobDefinition: String
  JobName: String
  RetryStrategy:
    BatchRetryStrategy

```

## Properties

`ArrayProperties`

The array properties for the submitted job, such as the size of the array. The array size
can be between 2 and 10,000. If you specify array properties for a job, it becomes an array
job. This parameter is used only if the target is an AWS Batch job.

_Required_: No

_Type_: [BatchArrayProperties](aws-properties-events-rule-batcharrayproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobDefinition`

The ARN or name of the job definition to use if the event target is an AWS Batch job. This job definition must already exist.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobName`

The name to use for this execution of the job, if the target is an AWS Batch
job.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryStrategy`

The retry strategy to use for failed jobs, if the target is an AWS Batch job.
The retry strategy is the number of times to retry the failed job execution. Valid values are
1–10. When you specify a retry strategy here, it overrides the retry strategy defined in the
job definition.

_Required_: No

_Type_: [BatchRetryStrategy](aws-properties-events-rule-batchretrystrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Set the BatchParameters parameters

The following example sets the `BatchParameters` parameter for a job that has 950 objects and will be tried 8 times.

#### JSON

```json

"BatchParameters": {
  "ArrayProperties": {
      "Size": 950
   },
  "JobDefinition": "NewBatchJobDefinition",
  "JobName": "Job1",
  "RetryStrategy": {
      "Attempts": 8
   }
}
```

#### YAML

```yaml

BatchParameters:
  ArrayProperties:
    Size: 950
  JobDefinition: NewBatchJobDefinition
  JobName: Job1
  RetryStrategy:
    Attempts: 8
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchArrayProperties

BatchRetryStrategy

All content copied from https://docs.aws.amazon.com/.
