This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::StorageLens SSEKMS

Specifies the use of server-side encryption using an AWS Key Management Service key (SSE-KMS)
to encrypt the delivered S3 Storage Lens metrics export file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeyId" : String
}

```

### YAML

```yaml

  KeyId: String

```

## Properties

`KeyId`

Specifies the Amazon Resource Name (ARN) of the customer managed AWS KMS key to
use for encrypting the S3 Storage Lens metrics export file. Amazon S3 only supports symmetric
encryption keys. For more information, see [Special-purpose keys](../../../kms/latest/developerguide/key-types.md) in the
_AWS Key Management Service Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SelectionCriteria

StorageLensConfiguration

All content copied from https://docs.aws.amazon.com/.
