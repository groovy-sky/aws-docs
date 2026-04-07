This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition JobTimeout

An object that represents a job timeout configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttemptDurationSeconds" : Integer
}

```

### YAML

```yaml

  AttemptDurationSeconds: Integer

```

## Properties

`AttemptDurationSeconds`

The job timeout time (in seconds) that's measured from the job attempt's
`startedAt` timestamp. After this time passes, AWS Batch terminates your jobs if they
aren't finished. The minimum value for the timeout is 60 seconds.

For array jobs, the timeout applies to the child jobs, not to the parent array job.

For multi-node parallel (MNP) jobs, the timeout applies to the whole job, not to the
individual nodes.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ImagePullSecret

LinuxParameters
