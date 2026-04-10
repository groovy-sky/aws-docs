This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable SSESpecification

Represents the settings used to enable server-side encryption.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SSEEnabled" : Boolean,
  "SSEType" : String
}

```

### YAML

```yaml

  SSEEnabled: Boolean
  SSEType: String

```

## Properties

`SSEEnabled`

Indicates whether server-side encryption is performed using an AWS
managed key or an AWS owned key. If enabled (true), server-side encryption
type is set to KMS and an AWS managed key is used (AWS KMS
charges apply). If disabled (false) or not specified,server-side encryption is set to an
AWS owned key. If you choose to use KMS encryption, you can also use
customer managed KMS keys by specifying them in the
`ReplicaSpecification.SSESpecification` object. You cannot mix AWS managed and customer managed KMS keys.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SSEType`

Server-side encryption type. The only supported value is:

- `KMS` \- Server-side encryption that uses AWS Key Management Service. The
key is stored in your account and is managed by AWS KMS (AWS KMS charges apply).

_Required_: No

_Type_: String

_Allowed values_: `AES256 | KMS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourcePolicy

StreamSpecification

All content copied from https://docs.aws.amazon.com/.
