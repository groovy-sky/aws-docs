This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AmazonMQ::Broker EncryptionOptions

Encryption options for the broker.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "UseAwsOwnedKey" : Boolean
}

```

### YAML

```yaml

  KmsKeyId: String
  UseAwsOwnedKey: Boolean

```

## Properties

`KmsKeyId`

The customer master key (CMK) to use for the A AWS KMS (KMS).
This key is used to encrypt your data at rest. If not provided, Amazon MQ will use a
default CMK to encrypt your data.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UseAwsOwnedKey`

Enables the use of an AWS owned CMK using AWS KMS (KMS). Set to `true` by default, if no value is provided, for example,
for RabbitMQ brokers.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConfigurationId

LdapServerMetadata
