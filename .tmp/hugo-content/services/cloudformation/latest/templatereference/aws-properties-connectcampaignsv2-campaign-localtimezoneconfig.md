This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign LocalTimeZoneConfig

The configuration of timezone for recipient.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultTimeZone" : String,
  "LocalTimeZoneDetection" : [ String, ... ]
}

```

### YAML

```yaml

  DefaultTimeZone: String
  LocalTimeZoneDetection:
    - String

```

## Properties

`DefaultTimeZone`

The timezone to use for all recipients.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalTimeZoneDetection`

Detects methods for the recipient's timezone.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventTrigger

OpenHours

All content copied from https://docs.aws.amazon.com/.
