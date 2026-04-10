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

_Type_: Array of [CreateCaseAction](aws-properties-connect-rule-createcaseaction.md)

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

_Type_: Array of [EventBridgeAction](aws-properties-connect-rule-eventbridgeaction.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SendNotificationActions`

Information about the send notification action.

_Required_: No

_Type_: Array of [SendNotificationAction](aws-properties-connect-rule-sendnotificationaction.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubmitAutoEvaluationActions`

Property description not available.

_Required_: No

_Type_: Array of [SubmitAutoEvaluationAction](aws-properties-connect-rule-submitautoevaluationaction.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskActions`

Information about the task action. This field is required if
`TriggerEventSource` is one of the following values:
`OnZendeskTicketCreate` \| `OnZendeskTicketStatusUpdate` \|
`OnSalesforceCaseCreate`

_Required_: No

_Type_: Array of [TaskAction](aws-properties-connect-rule-taskaction.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdateCaseActions`

Property description not available.

_Required_: No

_Type_: Array of [UpdateCaseAction](aws-properties-connect-rule-updatecaseaction.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Connect::Rule

CreateCaseAction

All content copied from https://docs.aws.amazon.com/.
