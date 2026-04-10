This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput MaintenanceConfiguration

The maintenance configuration settings applied to this router input.

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

Configuration settings for default maintenance scheduling.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredDayTime`

Configuration for preferred day and time maintenance settings.

_Required_: No

_Type_: [PreferredDayTimeMaintenanceConfiguration](aws-properties-mediaconnect-routerinput-preferreddaytimemaintenanceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FlowTransitEncryptionKeyConfiguration

MediaConnectFlowRouterInputConfiguration

All content copied from https://docs.aws.amazon.com/.
