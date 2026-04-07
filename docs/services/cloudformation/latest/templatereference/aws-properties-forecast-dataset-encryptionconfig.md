This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Forecast::Dataset EncryptionConfig

An AWS Key Management Service (KMS) key and an AWS Identity and Access Management (IAM) role that Amazon Forecast can assume to
access the key. You can specify this optional object in the
[CreateDataset](https://docs.aws.amazon.com/forecast/latest/dg/API_CreateDataset.html) and [CreatePredictor](https://docs.aws.amazon.com/forecast/latest/dg/API_CreatePredictor.html) requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyArn" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  KmsKeyArn: String
  RoleArn: String

```

## Properties

`KmsKeyArn`

The Amazon Resource Name (ARN) of the KMS key.

_Required_: No

_Type_: String

_Pattern_: `arn:aws[-a-z]*:kms:.*:key/.*`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that Amazon Forecast can assume to access the AWS KMS key.

Passing a role across AWS accounts is not allowed. If you pass a role that isn't in your
account, you get an `InvalidInputException` error.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-\_\.\/\:]+$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttributesItems

Schema
