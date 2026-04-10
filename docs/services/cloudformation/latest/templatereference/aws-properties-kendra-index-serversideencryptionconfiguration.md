This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Index ServerSideEncryptionConfiguration

Provides the identifier of the AWS KMS customer master key (CMK) used
to encrypt data indexed by Amazon Kendra. We suggest that you use a CMK from your
account to help secure your index. Amazon Kendra doesn't support asymmetric CMKs.

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

The identifier of the AWS KMS key. Amazon Kendra doesn't support
asymmetric keys.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Search

Tag

All content copied from https://docs.aws.amazon.com/.
