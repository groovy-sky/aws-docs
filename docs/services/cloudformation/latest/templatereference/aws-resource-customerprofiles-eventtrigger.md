This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::EventTrigger

Specifies the rules to perform an action based on customer ingested data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CustomerProfiles::EventTrigger",
  "Properties" : {
      "Description" : String,
      "DomainName" : String,
      "EventTriggerConditions" : [ EventTriggerCondition, ... ],
      "EventTriggerLimits" : EventTriggerLimits,
      "EventTriggerName" : String,
      "ObjectTypeName" : String,
      "SegmentFilter" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CustomerProfiles::EventTrigger
Properties:
  Description: String
  DomainName: String
  EventTriggerConditions:
    - EventTriggerCondition
  EventTriggerLimits:
    EventTriggerLimits
  EventTriggerName: String
  ObjectTypeName: String
  SegmentFilter: String
  Tags:
    - Tag

```

## Properties

`Description`

The description of the event trigger.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The unique name of the domain.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventTriggerConditions`

A list of conditions that determine when an event should trigger the
destination.

_Required_: Yes

_Type_: Array of [EventTriggerCondition](aws-properties-customerprofiles-eventtrigger-eventtriggercondition.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventTriggerLimits`

Defines limits controlling whether an event triggers the destination, based on
ingestion latency and the number of invocations per profile over specific time
periods.

_Required_: No

_Type_: [EventTriggerLimits](aws-properties-customerprofiles-eventtrigger-eventtriggerlimits.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventTriggerName`

The unique name of the event trigger.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ObjectTypeName`

The unique name of the object type.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z_][a-zA-Z_0-9-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentFilter`

The destination is triggered only for profiles that meet the criteria of a segment
definition.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-customerprofiles-eventtrigger-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

The timestamp of when the event trigger was created.

`LastUpdatedAt`

The timestamp of when the event trigger was most recently updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EventTriggerCondition

All content copied from https://docs.aws.amazon.com/.
