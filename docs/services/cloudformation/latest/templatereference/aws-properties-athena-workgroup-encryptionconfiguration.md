This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::WorkGroup EncryptionConfiguration

If query results are encrypted in Amazon S3, indicates the encryption option used (for
example, `SSE_KMS` or `CSE_KMS`) and key information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncryptionOption" : String,
  "KmsKey" : String
}

```

### YAML

```yaml

  EncryptionOption: String
  KmsKey: String

```

## Properties

`EncryptionOption`

Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys
( `SSE_S3`), server-side encryption with KMS-managed keys
( `SSE_KMS`), or client-side encryption with KMS-managed keys
( `CSE_KMS`) is used.

If a query runs in a workgroup and the workgroup overrides client-side settings, then
the workgroup's setting for encryption is used. It specifies whether query results must
be encrypted, for all queries that run in this workgroup.

_Required_: Yes

_Type_: String

_Allowed values_: `SSE_S3 | SSE_KMS | CSE_KMS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKey`

For `SSE_KMS` and `CSE_KMS`, this is the KMS key ARN or
ID.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomerContentEncryptionConfiguration

EngineConfiguration
