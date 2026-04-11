This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetBatchJobParameters

The parameters for using an AWS Batch job as a target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArrayProperties" : BatchArrayProperties,
  "ContainerOverrides" : BatchContainerOverrides,
  "DependsOn" : [ BatchJobDependency, ... ],
  "JobDefinition" : String,
  "JobName" : String,
  "Parameters" : {Key: Value, ...},
  "RetryStrategy" : BatchRetryStrategy
}

```

### YAML

```yaml

  ArrayProperties:
    BatchArrayProperties
  ContainerOverrides:
    BatchContainerOverrides
  DependsOn:
    - BatchJobDependency
  JobDefinition: String
  JobName: String
  Parameters:
    Key: Value
  RetryStrategy:
    BatchRetryStrategy

```

## Properties

`ArrayProperties`

The array properties for the submitted job, such as the size of the array. The array size can be between 2 and 10,000.
If you specify array properties for a job, it becomes an array job. This parameter is used only if the target is an AWS Batch job.

_Required_: No

_Type_: [BatchArrayProperties](aws-properties-pipes-pipe-batcharrayproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerOverrides`

The overrides that are sent to a container.

_Required_: No

_Type_: [BatchContainerOverrides](aws-properties-pipes-pipe-batchcontaineroverrides.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DependsOn`

A list of dependencies for the job. A job can depend upon a maximum of 20 jobs. You can
specify a `SEQUENTIAL` type dependency without specifying a job ID for array
jobs so that each child array job completes sequentially, starting at index 0. You can also
specify an `N_TO_N` type dependency with a job ID for array jobs. In that case,
each index child of this job must wait for the corresponding index child of each dependency
to complete before it can begin.

_Required_: No

_Type_: Array of [BatchJobDependency](aws-properties-pipes-pipe-batchjobdependency.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobDefinition`

The job definition used by this job. This value can be one of `name`,
`name:revision`, or the Amazon Resource Name (ARN) for the job definition. If
name is specified without a revision then the latest active revision is used.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobName`

The name of the job. It can be up to 128 letters long. The first character must be
alphanumeric, can contain uppercase and lowercase letters, numbers, hyphens (-), and
underscores (\_).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

Additional parameters passed to the job that replace parameter substitution placeholders
that are set in the job definition. Parameters are specified as a key and value pair
mapping. Parameters included here override any corresponding parameter defaults from the
job definition.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryStrategy`

The retry strategy to use for failed jobs. When a retry strategy is specified here, it
overrides the retry strategy defined in the job definition.

_Required_: No

_Type_: [BatchRetryStrategy](aws-properties-pipes-pipe-batchretrystrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeSourceSqsQueueParameters

PipeTargetCloudWatchLogsParameters

All content copied from https://docs.aws.amazon.com/.
