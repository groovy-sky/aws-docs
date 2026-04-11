This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice OtaaV11

OTAA device object for v1.1 for create APIs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppKey" : String,
  "JoinEui" : String,
  "NwkKey" : String
}

```

### YAML

```yaml

  AppKey: String
  JoinEui: String
  NwkKey: String

```

## Properties

`AppKey`

The AppKey is a secret key, which you should handle in a similar way as you would an
application password. You can protect the AppKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](../userguide/dynamic-references.md#dynamic-references-secretsmanager) to reference this value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JoinEui`

The JoinEUI value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{16}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NwkKey`

The NwkKey is a secret key, which you should handle in a similar way as you would an
application password. You can protect the NwkKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](../userguide/dynamic-references.md#dynamic-references-secretsmanager) to reference this value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OtaaV10x

SessionKeysAbpV10x

All content copied from https://docs.aws.amazon.com/.
