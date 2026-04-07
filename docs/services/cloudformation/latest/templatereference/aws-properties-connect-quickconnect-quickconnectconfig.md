This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::QuickConnect QuickConnectConfig

Contains configuration settings for a quick connect.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PhoneConfig" : PhoneNumberQuickConnectConfig,
  "QueueConfig" : QueueQuickConnectConfig,
  "QuickConnectType" : String,
  "UserConfig" : UserQuickConnectConfig
}

```

### YAML

```yaml

  PhoneConfig:
    PhoneNumberQuickConnectConfig
  QueueConfig:
    QueueQuickConnectConfig
  QuickConnectType: String
  UserConfig:
    UserQuickConnectConfig

```

## Properties

`PhoneConfig`

The phone configuration. This is required only if QuickConnectType is PHONE\_NUMBER.

_Required_: No

_Type_: [PhoneNumberQuickConnectConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-quickconnect-phonenumberquickconnectconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueConfig`

The queue configuration. This is required only if QuickConnectType is QUEUE.

_Required_: No

_Type_: [QueueQuickConnectConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-quickconnect-queuequickconnectconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuickConnectType`

The type of quick connect. In the Amazon Connect console, when you create a
quick connect, you are prompted to assign one of the following types: Agent (USER),
External (PHONE\_NUMBER), or Queue (QUEUE).

_Required_: Yes

_Type_: String

_Allowed values_: `PHONE_NUMBER | QUEUE | USER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserConfig`

The user configuration. This is required only if QuickConnectType is USER.

_Required_: No

_Type_: [UserQuickConnectConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-quickconnect-userquickconnectconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueueQuickConnectConfig

Tag
