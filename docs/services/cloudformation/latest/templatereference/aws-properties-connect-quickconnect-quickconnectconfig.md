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

_Type_: [PhoneNumberQuickConnectConfig](aws-properties-connect-quickconnect-phonenumberquickconnectconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueConfig`

The queue configuration. This is required only if QuickConnectType is QUEUE.

_Required_: No

_Type_: [QueueQuickConnectConfig](aws-properties-connect-quickconnect-queuequickconnectconfig.md)

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

_Type_: [UserQuickConnectConfig](aws-properties-connect-quickconnect-userquickconnectconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueueQuickConnectConfig

Tag

All content copied from https://docs.aws.amazon.com/.
