This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMIncidents::ResponsePlan IncidentTemplate

The `IncidentTemplate` property type specifies details used to create an
incident when using this response plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DedupeString" : String,
  "Impact" : Integer,
  "IncidentTags" : [ Tag, ... ],
  "NotificationTargets" : [ NotificationTargetItem, ... ],
  "Summary" : String,
  "Title" : String
}

```

### YAML

```yaml

  DedupeString:
    String
  Impact: Integer
  IncidentTags:
    - Tag
  NotificationTargets:
    - NotificationTargetItem
  Summary: String
  Title: String

```

## Properties

`DedupeString`

Used to create only one incident record for an incident.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Impact`

Defines the impact to the customers. Providing an impact overwrites the impact
provided by a response plan.

###### Possible impacts:

- `1` \- Critical impact, this typically relates to full application
failure that impacts many to all customers.

- `2` \- High impact, partial application failure with impact to many
customers.

- `3` \- Medium impact, the application is providing reduced service
to customers.

- `4` \- Low impact, customer might aren't impacted by the problem
yet.

- `5` \- No impact, customers aren't currently impacted but urgent
action is needed to avoid impact.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncidentTags`

Tags to assign to the template. When the `StartIncident` API action is
called, Incident Manager assigns the tags specified in the template to the
incident.

_Required_: No

_Type_: Array of [Tag](aws-properties-ssmincidents-responseplan-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationTargets`

The Amazon Simple Notification Service (Amazon SNS) targets that Amazon Q Developer in chat applications uses to notify the chat channel of updates
to an incident. You can also make updates to the incident through the chat channel using
the Amazon SNS topics.

_Required_: No

_Type_: Array of [NotificationTargetItem](aws-properties-ssmincidents-responseplan-notificationtargetitem.md)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Summary`

The summary describes what has happened during the incident.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the incident is a brief and easily recognizable.

_Required_: Yes

_Type_: String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynamicSsmParameterValue

Integration

All content copied from https://docs.aws.amazon.com/.
