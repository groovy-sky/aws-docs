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

_Type_: [PreferredDayTimeMaintenanceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-routeroutput-preferreddaytimemaintenanceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FlowTransitEncryptionKeyConfiguration

MediaConnectFlowRouterOutputConfiguration
