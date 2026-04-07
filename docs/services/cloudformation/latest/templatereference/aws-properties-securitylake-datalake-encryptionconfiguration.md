This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::DataLake EncryptionConfiguration

Provides encryption details of the Amazon Security Lake object. The AWS shared responsibility model applies to data protection in Amazon Security Lake. As described in this model, AWS is responsible for protecting the global infrastructure that runs all of the AWS Cloud. You are responsible for maintaining control over your content that is hosted on this infrastructure. For more details, see [Data protection](https://docs.aws.amazon.com/security-lake/latest/userguide/data-protection.html) in the Amazon Security Lake User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String
}

```

### YAML

```yaml

  KmsKeyId: String

```

## Properties

`KmsKeyId`

The ID of KMS encryption key used by Amazon Security Lake to encrypt the Security Lake object.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SecurityLake::DataLake

Expiration
