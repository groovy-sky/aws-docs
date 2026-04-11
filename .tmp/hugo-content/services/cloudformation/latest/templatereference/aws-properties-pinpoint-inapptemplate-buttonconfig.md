This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::InAppTemplate ButtonConfig

Specifies the behavior of buttons that appear in an in-app message template.

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

Optional button configuration to use for in-app messages sent to Android devices. This
button configuration overrides the default button configuration.

_Required_: No

_Type_: [OverrideButtonConfiguration](aws-properties-pinpoint-inapptemplate-overridebuttonconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultConfig`

Specifies the default behavior of a button that appears in an in-app message. You can
optionally add button configurations that specifically apply to iOS, Android, or web
browser users.

_Required_: No

_Type_: [DefaultButtonConfiguration](aws-properties-pinpoint-inapptemplate-defaultbuttonconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IOS`

Optional button configuration to use for in-app messages sent to iOS devices. This
button configuration overrides the default button configuration.

_Required_: No

_Type_: [OverrideButtonConfiguration](aws-properties-pinpoint-inapptemplate-overridebuttonconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Web`

Optional button configuration to use for in-app messages sent to web applications.
This button configuration overrides the default button configuration.

_Required_: No

_Type_: [OverrideButtonConfiguration](aws-properties-pinpoint-inapptemplate-overridebuttonconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BodyConfig

DefaultButtonConfiguration

All content copied from https://docs.aws.amazon.com/.
