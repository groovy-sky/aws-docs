This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow TriggerConfig

The trigger settings that determine how and when Amazon AppFlow runs the specified
flow.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TriggerProperties" : ScheduledTriggerProperties,
  "TriggerType" : String
}

```

### YAML

```yaml

  TriggerProperties:
    ScheduledTriggerProperties
  TriggerType: String

```

## Properties

`TriggerProperties`

Specifies the configuration details of a schedule-triggered flow as defined by the user.
Currently, these settings only apply to the `Scheduled` trigger type.

_Required_: No

_Type_: [ScheduledTriggerProperties](aws-properties-appflow-flow-scheduledtriggerproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TriggerType`

Specifies the type of flow trigger. This can be `OnDemand`,
`Scheduled`, or `Event`.

_Required_: Yes

_Type_: String

_Allowed values_: `Scheduled | Event | OnDemand`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [TriggerConfig](../../../../reference/appflow/1-0/apireference/api-triggerconfig.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrendmicroSourceProperties

UpsolverDestinationProperties

All content copied from https://docs.aws.amazon.com/.
