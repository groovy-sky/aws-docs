This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Workflow

Allows you to create a workflow with specified steps and step details the workflow
invokes after file transfer completes. After creating a workflow, you can associate the
workflow created with any transfer servers by specifying the
`workflow-details` field in `CreateServer` and
`UpdateServer` operations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Transfer::Workflow",
  "Properties" : {
      "Description" : String,
      "OnExceptionSteps" : [ WorkflowStep, ... ],
      "Steps" : [ WorkflowStep, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Transfer::Workflow
Properties:
  Description: String
  OnExceptionSteps:
    - WorkflowStep
  Steps:
    - WorkflowStep
  Tags:
    - Tag

```

## Properties

`Description`

Specifies the text description for the workflow.

_Required_: No

_Type_: String

_Pattern_: `^[\w\- ]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnExceptionSteps`

Specifies the steps (actions) to take if errors are encountered during execution of the workflow.

_Required_: No

_Type_: Array of [WorkflowStep](aws-properties-transfer-workflow-workflowstep.md)

_Maximum_: `8`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Steps`

Specifies the details for the steps that are in the specified workflow.

_Required_: Yes

_Type_: Array of [WorkflowStep](aws-properties-transfer-workflow-workflowstep.md)

_Maximum_: `8`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key-value pairs that can be used to group and search for workflows. Tags are metadata
attached to workflows for any purpose.

_Required_: No

_Type_: Array of [Tag](aws-properties-transfer-workflow-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`WorkflowId`

A unique identifier for a workflow.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebAppUnits

CopyStepDetails

All content copied from https://docs.aws.amazon.com/.
