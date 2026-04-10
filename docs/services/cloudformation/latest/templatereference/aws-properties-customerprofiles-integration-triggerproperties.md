This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Integration TriggerProperties

Specifies the configuration details that control the trigger for a flow. Currently,
these settings only apply to the Scheduled trigger type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Scheduled" : ScheduledTriggerProperties
}

```

### YAML

```yaml

  Scheduled:
    ScheduledTriggerProperties

```

## Properties

`Scheduled`

Specifies the configuration details of a schedule-triggered flow that you
define.

_Required_: No

_Type_: [ScheduledTriggerProperties](aws-properties-customerprofiles-integration-scheduledtriggerproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TriggerConfig

ZendeskSourceProperties

All content copied from https://docs.aws.amazon.com/.
