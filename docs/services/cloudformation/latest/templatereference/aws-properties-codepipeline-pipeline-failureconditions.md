This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline FailureConditions

The configuration that specifies the result, such as rollback, to occur upon stage
failure. For more information about conditions, see [Stage conditions](../../../codepipeline/latest/userguide/stage-conditions.md)
and [How do\
stage conditions work?](../../../codepipeline/latest/userguide/concepts-how-it-works-conditions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Conditions" : [ Condition, ... ],
  "Result" : String,
  "RetryConfiguration" : RetryConfiguration
}

```

### YAML

```yaml

  Conditions:
    - Condition
  Result: String
  RetryConfiguration:
    RetryConfiguration

```

## Properties

`Conditions`

The conditions that are configured as failure conditions. For more information about
conditions, see [Stage conditions](../../../codepipeline/latest/userguide/stage-conditions.md)
and [How do\
stage conditions work?](../../../codepipeline/latest/userguide/concepts-how-it-works-conditions.md).

_Required_: No

_Type_: Array of [Condition](aws-properties-codepipeline-pipeline-condition.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Result`

The specified result for when the failure conditions are met, such as rolling back the
stage.

_Required_: No

_Type_: String

_Allowed values_: `ROLLBACK | RETRY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetryConfiguration`

The retry configuration specifies automatic retry for a failed stage, along with the
configured retry mode.

_Required_: No

_Type_: [RetryConfiguration](aws-properties-codepipeline-pipeline-retryconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentVariable

GitBranchFilterCriteria

All content copied from https://docs.aws.amazon.com/.
