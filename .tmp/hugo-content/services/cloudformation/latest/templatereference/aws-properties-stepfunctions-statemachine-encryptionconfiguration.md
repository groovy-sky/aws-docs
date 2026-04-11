This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::StepFunctions::StateMachine EncryptionConfiguration

Settings to configure server-side encryption for a state machine. By default, Step Functions provides transparent server-side encryption. With this configuration, you can specify a customer managed AWS KMS key for encryption.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsDataKeyReusePeriodSeconds" : Integer,
  "KmsKeyId" : String,
  "Type" : String
}

```

### YAML

```yaml

  KmsDataKeyReusePeriodSeconds: Integer
  KmsKeyId: String
  Type: String

```

## Properties

`KmsDataKeyReusePeriodSeconds`

Maximum duration that Step Functions will reuse data keys. When the period expires, Step Functions will call `GenerateDataKey`. Only applies to customer managed keys.

_Required_: No

_Type_: Integer

_Minimum_: `60`

_Maximum_: `900`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

An alias, alias ARN, key ID, or key ARN of a symmetric encryption AWS KMS key to encrypt data. To specify a AWS KMS key in a different AWS account, you must use the key ARN or alias ARN.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Encryption option for a state machine.

_Required_: Yes

_Type_: String

_Allowed values_: `CUSTOMER_MANAGED_KMS_KEY | AWS_OWNED_KEY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogsLogGroup

LogDestination

All content copied from https://docs.aws.amazon.com/.
