This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OSIS::Pipeline EncryptionAtRestOptions

Options to control how OpenSearch encrypts buffer data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyArn" : String
}

```

### YAML

```yaml

  KmsKeyArn: String

```

## Properties

`KmsKeyArn`

The ARN of the KMS key used to encrypt buffer data.
By default, data is encrypted using an AWS owned key.

_Required_: Yes

_Type_: String

_Minimum_: `7`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatchLogDestination

LogPublishingOptions
