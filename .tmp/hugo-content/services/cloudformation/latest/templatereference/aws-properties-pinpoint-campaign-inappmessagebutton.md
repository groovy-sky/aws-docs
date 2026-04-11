This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign InAppMessageButton

Specifies the configuration of a button that appears in an in-app message.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Android" : OverrideButtonConfiguration,
  "DefaultConfig" : DefaultButtonConfiguration,
  "IOS" : OverrideButtonConfiguration,
  "Web" : OverrideButtonConfiguration
}

```

### YAML

```yaml

  Android:
    OverrideButtonConfiguration
  DefaultConfig:
    DefaultButtonConfiguration
  IOS:
    OverrideButtonConfiguration
  Web:
    OverrideButtonConfiguration

```

## Properties

`Android`

An object that defines the default behavior for a button in in-app messages
sent to Android.

_Required_: No

_Type_: [OverrideButtonConfiguration](aws-properties-pinpoint-campaign-overridebuttonconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultConfig`

An object that defines the default behavior for a button in an in-app
message.

_Required_: No

_Type_: [DefaultButtonConfiguration](aws-properties-pinpoint-campaign-defaultbuttonconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IOS`

An object that defines the default behavior for a button in in-app messages
sent to iOS devices.

_Required_: No

_Type_: [OverrideButtonConfiguration](aws-properties-pinpoint-campaign-overridebuttonconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Web`

An object that defines the default behavior for a button in in-app messages
for web applications.

_Required_: No

_Type_: [OverrideButtonConfiguration](aws-properties-pinpoint-campaign-overridebuttonconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InAppMessageBodyConfig

InAppMessageContent

All content copied from https://docs.aws.amazon.com/.
