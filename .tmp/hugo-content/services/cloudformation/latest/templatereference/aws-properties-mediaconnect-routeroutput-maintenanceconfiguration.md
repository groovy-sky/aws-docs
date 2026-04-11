This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput MaintenanceConfiguration

The configuration settings for maintenance operations, including preferred maintenance windows and schedules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Default" : Json,
  "PreferredDayTime" : PreferredDayTimeMaintenanceConfiguration
}

```

### YAML

```yaml

  Default: Json
  PreferredDayTime:
    PreferredDayTimeMaintenanceConfiguration

```

## Properties

`Default`

Default maintenance configuration settings.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredDayTime`

Preferred day and time maintenance configuration settings.

_Required_: No

_Type_: [PreferredDayTimeMaintenanceConfiguration](aws-properties-mediaconnect-routeroutput-preferreddaytimemaintenanceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowTransitEncryptionKeyConfiguration

MediaConnectFlowRouterOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
