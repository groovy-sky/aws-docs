This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain WindowStartTime

A custom start time for the off-peak window, in Coordinated Universal Time (UTC). The window length will always be 10 hours, so you can't
specify an end time. For example, if you specify 11:00 P.M. UTC as a start time, the end time will automatically be set to 9:00 A.M.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Hours" : Integer,
  "Minutes" : Integer
}

```

### YAML

```yaml

  Hours: Integer
  Minutes: Integer

```

## Properties

`Hours`

The start hour of the window in Coordinated Universal Time (UTC), using 24-hour time. For example, 17 refers to 5:00 P.M. UTC.
The minimum value is 0 and the maximum value is 23.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `23`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Minutes`

The start minute of the window, in UTC. The minimum value is 0 and the maximum value is 59.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `59`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VPCOptions

ZoneAwarenessConfig

All content copied from https://docs.aws.amazon.com/.
