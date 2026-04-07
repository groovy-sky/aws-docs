This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::Assistant ServerSideEncryptionConfiguration

The configuration information for the customer managed key used for
encryption.

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

The customer managed key used for encryption. The customer managed key must have
a policy that allows `kms:CreateGrant` and `kms:DescribeKey`
permissions to the IAM identity using the key to invoke Wisdom. To use
Wisdom with chat, the key policy must also allow `kms:Decrypt`,
`kms:GenerateDataKey*`, and `kms:DescribeKey` permissions to
the `connect.amazonaws.com` service principal. For more information about
setting up a customer managed key for Wisdom, see [Enable Amazon Connect Wisdom\
for your instance](https://docs.aws.amazon.com/connect/latest/adminguide/enable-wisdom.html). For information about valid ID values, see [Key\
identifiers (KeyId)](https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id) in the _AWS Key Management Service Developer_
_Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Wisdom::Assistant

Tag
