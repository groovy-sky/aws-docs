This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Rule Actions

A list of actions to be run when the rule is triggered.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssignContactCategoryActions" : [ Json, ... ],
  "CreateCaseActions" : [ CreateCaseAction, ... ],
  "EndAssociatedTasksActions" : [ Json, ... ],
  "EventBridgeActions" : [ EventBridgeAction, ... ],
  "SendNotificationActions" : [ SendNotificationAction, ... ],
  "SubmitAutoEvaluationActions" : [ SubmitAutoEvaluationAction, ... ],
  "TaskActions" : [ TaskAction, ... ],
  "UpdateCaseActions" : [ UpdateCaseAction, ... ]
}

```

### YAML

```yaml

  AssignContactCategoryActions:
    - Json
  CreateCaseActions:
    - CreateCaseAction
  EndAssociatedTasksActions:
    - Json
  EventBridgeActions:
    - EventBridgeAction
  SendNotificationActions:
    - SendNotificationAction
  SubmitAutoEvaluationActions:
    - SubmitAutoEvaluationAction
  TaskActions:
    - TaskAction
  UpdateCaseActions:
    - UpdateCaseAction

```

## Properties

`AssignContactCategoryActions`

Information about the contact category action. The syntax can be empty, for example,
`{}`.

_Required_: No

_Type_: Array of Json

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateCaseActions`

Property description not available.

_Required_: No

_Type_: Array of [CreateCaseAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-rule-createcaseaction.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndAssociatedTasksActions`

Property description not available.

_Required_: No

_Type_: Array of Json

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventBridgeActions`

Information about the EventBridge action.

_Required_: No

_Type_: Array of [EventBridgeAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-rule-eventbridgeaction.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SendNotificationActions`

Information about the send notification action.

_Required_: No

_Type_: Array of [SendNotificationAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-rule-sendnotificationaction.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubmitAutoEvaluationActions`

Property description not available.

_Required_: No

_Type_: Array of [SubmitAutoEvaluationAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-rule-submitautoevaluationaction.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskActions`

Information about the task action. This field is required if
`TriggerEventSource` is one of the following values:
`OnZendeskTicketCreate` \| `OnZendeskTicketStatusUpdate` \|
`OnSalesforceCaseCreate`

_Required_: No

_Type_: Array of [TaskAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-rule-taskaction.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateCaseActions`

Property description not available.

_Required_: No

_Type_: Array of [UpdateCaseAction](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-rule-updatecaseaction.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Connect::Rule

CreateCaseAction
