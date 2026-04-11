This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe BatchJobDependency

An object that represents an AWS Batch job dependency.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JobId" : String,
  "Type" : String
}

```

### YAML

```yaml

  JobId: String
  Type: String

```

## Properties

`JobId`

The job ID of the AWS Batch job that's associated with this
dependency.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the job dependency.

_Required_: No

_Type_: String

_Allowed values_: `N_TO_N | SEQUENTIAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchEnvironmentVariable

BatchResourceRequirement

All content copied from https://docs.aws.amazon.com/.
