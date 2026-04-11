This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::WirelessDevice SessionKeysAbpV11

Session keys for ABP v1.1.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppSKey" : String,
  "FNwkSIntKey" : String,
  "NwkSEncKey" : String,
  "SNwkSIntKey" : String
}

```

### YAML

```yaml

  AppSKey: String
  FNwkSIntKey: String
  NwkSEncKey: String
  SNwkSIntKey: String

```

## Properties

`AppSKey`

The AppSKey is a secret key, which you should handle in a similar way as you would an
application password. You can protect the AppSKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](../userguide/dynamic-references.md#dynamic-references-secretsmanager) to reference this value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FNwkSIntKey`

The FNwkSIntKey is a secret key, which you should handle in a similar way as you would
an application password. You can protect the FNwkSIntKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](../userguide/dynamic-references.md#dynamic-references-secretsmanager) to reference this value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NwkSEncKey`

The NwkSEncKey is a secret key, which you should handle in a similar way as you would an
application password. You can protect the NwkSEncKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](../userguide/dynamic-references.md#dynamic-references-secretsmanager) to reference this value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SNwkSIntKey`

The SNwkSIntKey is a secret key, which you should handle in a similar way as you would
an application password. You can protect the SNwkSIntKey value by storing it in the AWS Secrets Manager and use the [secretsmanager](../userguide/dynamic-references.md#dynamic-references-secretsmanager) to reference this value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{32}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SessionKeysAbpV10x

Tag

All content copied from https://docs.aws.amazon.com/.
