This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel MaintenanceCreateSettings

The `MaintenanceCreateSettings` property type specifies Property description not available. for an [AWS::MediaLive::Channel](aws-resource-medialive-channel.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaintenanceDay" : String,
  "MaintenanceStartTime" : String
}

```

### YAML

```yaml

  MaintenanceDay: String
  MaintenanceStartTime: String

```

## Properties

`MaintenanceDay`

Choose one day of the week for maintenance. The chosen day is used for all future maintenance windows.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceStartTime`

Choose the hour that maintenance will start. The chosen time is used for all future maintenance windows.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

M3u8Settings

MediaPackageAdditionalDestinations

All content copied from https://docs.aws.amazon.com/.
