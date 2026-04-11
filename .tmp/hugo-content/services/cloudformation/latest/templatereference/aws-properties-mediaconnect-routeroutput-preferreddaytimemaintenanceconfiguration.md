This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput PreferredDayTimeMaintenanceConfiguration

Configuration for preferred day and time maintenance settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Day" : String,
  "Time" : String
}

```

### YAML

```yaml

  Day: String
  Time: String

```

## Properties

`Day`

The preferred day for maintenance operations.

_Required_: Yes

_Type_: String

_Allowed values_: `MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY | SUNDAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Time`

The preferred time for maintenance operations.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaLiveTransitEncryptionKeyConfiguration

RistRouterOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
