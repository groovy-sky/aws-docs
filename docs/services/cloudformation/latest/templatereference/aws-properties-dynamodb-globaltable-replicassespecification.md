This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable ReplicaSSESpecification

Allows you to specify a KMS key identifier to be used for server-side encryption. The
key can be specified via ARN, key ID, or alias. The key must be created in the same region
as the replica.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KMSMasterKeyId" : String
}

```

### YAML

```yaml

  KMSMasterKeyId: String

```

## Properties

`KMSMasterKeyId`

The AWS KMS key that should be used for the AWS KMS encryption.
To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN.
Note that you should only provide this parameter if the key is different from the
default DynamoDB key `alias/aws/dynamodb`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaSpecification

ReplicaStreamSpecification

All content copied from https://docs.aws.amazon.com/.
