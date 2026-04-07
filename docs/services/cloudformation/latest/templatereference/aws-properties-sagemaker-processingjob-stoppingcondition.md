This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ProcessingJob StoppingCondition

Configures conditions under which the processing job should be stopped, such as how long the processing job has
been running. After the condition is met, the processing job is stopped.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxRuntimeInSeconds" : Integer
}

```

### YAML

```yaml

  MaxRuntimeInSeconds: Integer

```

## Properties

`MaxRuntimeInSeconds`

The maximum length of time, in seconds, that a training or compilation job can run
before it is stopped.

For compilation jobs, if the job does not complete during this time, a
`TimeOut` error is generated. We recommend starting with 900 seconds and
increasing as necessary based on your model.

For all other jobs, if the job does not complete during this time, SageMaker ends the job.
When `RetryStrategy` is specified in the job request,
`MaxRuntimeInSeconds` specifies the maximum time for all of the attempts
in total, not each individual attempt. The default value is 1 day. The maximum value is
28 days.

The maximum time that a `TrainingJob` can run in total, including any time
spent publishing metrics or archiving and uploading models after it has been stopped, is
30 days.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `777600`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3Output

Tag
