This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GuardDuty::PublishingDestination CFNDestinationProperties

Contains the Amazon Resource Name (ARN) of the resource that receives the published
findings, such as an S3 bucket, and the ARN of the KMS key that is used to encrypt these
published findings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationArn" : String,
  "KmsKeyArn" : String
}

```

### YAML

```yaml

  DestinationArn: String
  KmsKeyArn: String

```

## Properties

`DestinationArn`

The ARN of the resource where the findings are published.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The ARN of the KMS key to use for encryption.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::GuardDuty::PublishingDestination

TagItem
