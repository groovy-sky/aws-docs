This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution HpoResourceConfig

Describes the resource configuration for hyperparameter optimization (HPO).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxNumberOfTrainingJobs" : String,
  "MaxParallelTrainingJobs" : String
}

```

### YAML

```yaml

  MaxNumberOfTrainingJobs: String
  MaxParallelTrainingJobs: String

```

## Properties

`MaxNumberOfTrainingJobs`

The maximum number of training
jobs when you create a
solution
version.
The maximum value for `maxNumberOfTrainingJobs` is
`40`.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxParallelTrainingJobs`

The maximum number of parallel training
jobs when you create a
solution
version.
The maximum value for `maxParallelTrainingJobs` is
`10`.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HpoObjective

IntegerHyperParameterRange

All content copied from https://docs.aws.amazon.com/.
