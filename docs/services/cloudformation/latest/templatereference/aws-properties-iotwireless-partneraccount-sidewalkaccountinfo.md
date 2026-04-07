This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTWireless::PartnerAccount SidewalkAccountInfo

Information about a Sidewalk account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppServerPrivateKey" : String
}

```

### YAML

```yaml

  AppServerPrivateKey: String

```

## Properties

`AppServerPrivateKey`

The Sidewalk application server private key. The application server private key is a
secret key, which you should handle in a similar way as you would an application password.
You can protect the application server private key by storing the value in the AWS Secrets Manager and use the [secretsmanager](../userguide/dynamic-references.md#dynamic-references-secretsmanager) to reference this value.

_Required_: Yes

_Type_: String

_Pattern_: `[a-fA-F0-9]{64}`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IoTWireless::PartnerAccount

SidewalkAccountInfoWithFingerprint
