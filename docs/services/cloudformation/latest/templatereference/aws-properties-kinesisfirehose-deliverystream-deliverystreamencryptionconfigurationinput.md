This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream DeliveryStreamEncryptionConfigurationInput

Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side
Encryption (SSE).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyARN" : String,
  "KeyType" : String
}

```

### YAML

```yaml

  KeyARN: String
  KeyType: String

```

## Properties

`KeyARN`

If you set `KeyType` to `CUSTOMER_MANAGED_CMK`, you must specify
the Amazon Resource Name (ARN) of the CMK. If you set `KeyType` to `AWS_OWNED_CMK`, Firehose uses a service-account CMK.

_Required_: No

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyType`

Indicates the type of customer master key (CMK) to use for encryption. The default
setting is `AWS_OWNED_CMK`. For more information about CMKs, see [Customer\
Master Keys (CMKs)](../../../kms/latest/developerguide/concepts.md#master_keys).

You can use a CMK of type CUSTOMER\_MANAGED\_CMK to encrypt up to 500 delivery
streams.

###### Important

To encrypt your delivery stream, use symmetric CMKs. Kinesis Data Firehose doesn't
support asymmetric CMKs. For information about symmetric and asymmetric CMKs, see [About\
Symmetric and Asymmetric CMKs](../../../kms/latest/developerguide/symm-asymm-concepts.md) in the AWS Key Management
Service developer guide.

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_OWNED_CMK | CUSTOMER_MANAGED_CMK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataFormatConversionConfiguration

Deserializer

All content copied from https://docs.aws.amazon.com/.
