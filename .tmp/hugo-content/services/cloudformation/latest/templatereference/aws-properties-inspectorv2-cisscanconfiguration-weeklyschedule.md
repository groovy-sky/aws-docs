This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CisScanConfiguration WeeklySchedule

A weekly schedule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Days" : [ String, ... ],
  "StartTime" : Time
}

```

### YAML

```yaml

  Days:
    - String
  StartTime:
    Time

```

## Properties

`Days`

The weekly schedule's days.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `7`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The weekly schedule's start time.

_Required_: Yes

_Type_: [Time](aws-properties-inspectorv2-cisscanconfiguration-time.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Time

AWS::InspectorV2::CodeSecurityIntegration

All content copied from https://docs.aws.amazon.com/.
