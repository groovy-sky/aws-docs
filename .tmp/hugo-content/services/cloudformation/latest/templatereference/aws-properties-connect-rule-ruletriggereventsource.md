This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Rule RuleTriggerEventSource

The name of the event source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventSourceName" : String,
  "IntegrationAssociationArn" : String
}

```

### YAML

```yaml

  EventSourceName: String
  IntegrationAssociationArn: String

```

## Properties

`EventSourceName`

The name of the event source.

_Required_: Yes

_Type_: String

_Allowed values_: `OnContactEvaluationSubmit | OnPostCallAnalysisAvailable | OnRealTimeCallAnalysisAvailable | OnRealTimeChatAnalysisAvailable | OnPostChatAnalysisAvailable | OnZendeskTicketCreate | OnZendeskTicketStatusUpdate | OnSalesforceCaseCreate | OnMetricDataUpdate | OnCaseCreate | OnCaseUpdate`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IntegrationAssociationArn`

The Amazon Resource Name (ARN) of the integration association.
`IntegrationAssociationArn` is required if
`TriggerEventSource` is one of the following values:
`OnZendeskTicketCreate` \| `OnZendeskTicketStatusUpdate` \|
`OnSalesforceCaseCreate`

_Required_: No

_Type_: String

_Pattern_: `^$|arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/integration-association/[-a-zA-Z0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Reference

SendNotificationAction

All content copied from https://docs.aws.amazon.com/.
