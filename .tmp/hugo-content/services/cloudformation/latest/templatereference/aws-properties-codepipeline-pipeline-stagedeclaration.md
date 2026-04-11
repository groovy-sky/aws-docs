This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline StageDeclaration

Represents information about a stage and its definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ ActionDeclaration, ... ],
  "BeforeEntry" : BeforeEntryConditions,
  "Blockers" : [ BlockerDeclaration, ... ],
  "Name" : String,
  "OnFailure" : FailureConditions,
  "OnSuccess" : SuccessConditions
}

```

### YAML

```yaml

  Actions:
    - ActionDeclaration
  BeforeEntry:
    BeforeEntryConditions
  Blockers:
    - BlockerDeclaration
  Name: String
  OnFailure:
    FailureConditions
  OnSuccess:
    SuccessConditions

```

## Properties

`Actions`

The actions included in a stage.

_Required_: Yes

_Type_: Array of [ActionDeclaration](aws-properties-codepipeline-pipeline-actiondeclaration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BeforeEntry`

The method to use when a stage allows entry. For example, configuring this field for
conditions will allow entry to the stage when the conditions are met.

_Required_: No

_Type_: [BeforeEntryConditions](aws-properties-codepipeline-pipeline-beforeentryconditions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Blockers`

Reserved for future use.

_Required_: No

_Type_: Array of [BlockerDeclaration](aws-properties-codepipeline-pipeline-blockerdeclaration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the stage.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnFailure`

The method to use when a stage has not completed successfully. For example,
configuring this field for rollback will roll back a failed stage automatically to the
last successful pipeline execution in the stage.

_Required_: No

_Type_: [FailureConditions](aws-properties-codepipeline-pipeline-failureconditions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnSuccess`

The method to use when a stage has succeeded. For example, configuring this field for
conditions will allow the stage to succeed when the conditions are met.

_Required_: No

_Type_: [SuccessConditions](aws-properties-codepipeline-pipeline-successconditions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RuleTypeId

StageTransition

All content copied from https://docs.aws.amazon.com/.
