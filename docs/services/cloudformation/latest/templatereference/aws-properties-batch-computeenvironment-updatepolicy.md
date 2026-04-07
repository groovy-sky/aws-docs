This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ComputeEnvironment UpdatePolicy

Specifies the infrastructure update policy for the Amazon EC2 compute environment. For more information
about infrastructure updates, see [Updating compute environments](https://docs.aws.amazon.com/batch/latest/userguide/updating-compute-environments.html) in the
_AWS Batch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JobExecutionTimeoutMinutes" : Integer,
  "TerminateJobsOnUpdate" : Boolean
}

```

### YAML

```yaml

  JobExecutionTimeoutMinutes: Integer
  TerminateJobsOnUpdate: Boolean

```

## Properties

`JobExecutionTimeoutMinutes`

Specifies the job timeout (in minutes) when the compute environment infrastructure is
updated. The default value is 30.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TerminateJobsOnUpdate`

Specifies whether jobs are automatically terminated when the compute environment
infrastructure is updated. The default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LaunchTemplateSpecificationOverride

AWS::Batch::ConsumableResource
