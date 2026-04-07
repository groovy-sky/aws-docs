This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan Workflow

Represents a workflow in a Region switch plan. A workflow defines a sequence of steps to execute during a Region switch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Steps" : [ Step, ... ],
  "WorkflowDescription" : String,
  "WorkflowTargetAction" : String,
  "WorkflowTargetRegion" : String
}

```

### YAML

```yaml

  Steps:
    - Step
  WorkflowDescription: String
  WorkflowTargetAction: String
  WorkflowTargetRegion: String

```

## Properties

`Steps`

The steps that make up the workflow.

_Required_: No

_Type_: Array of [Step](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-arcregionswitch-plan-step.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowDescription`

The description of the workflow.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowTargetAction`

The action that the workflow performs. Valid values include `activate` and `deactivate`.

_Required_: Yes

_Type_: String

_Allowed values_: `activate | deactivate | postRecovery`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowTargetRegion`

The AWS Region that the workflow targets.

_Required_: No

_Type_: String

_Pattern_: `^[a-z]{2}-[a-z-]+-\d+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TriggerCondition

Next
