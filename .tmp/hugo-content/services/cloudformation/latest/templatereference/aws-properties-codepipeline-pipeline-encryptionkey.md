This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline EncryptionKey

Represents information about the key used to encrypt data in the artifact store, such
as an AWS Key Management Service (AWS KMS) key.

`EncryptionKey` is a property of the [ArtifactStore](../userguide/aws-properties-codepipeline-pipeline-artifactstore.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Type" : String
}

```

### YAML

```yaml

  Id: String
  Type: String

```

## Properties

`Id`

The ID used to identify the key. For an AWS KMS key, you can use the
key ID, the key ARN, or the alias ARN.

###### Note

Aliases are recognized only in the account that created the AWS KMS
key. For cross-account actions, you can only use the key ID or key ARN to identify
the key. Cross-account actions involve using the role from the other account
(AccountB), so specifying the key ID will use the key from the other account
(AccountB).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of encryption key, such as an AWS KMS key. When creating or
updating a pipeline, the value must be set to 'KMS'.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Condition

EnvironmentVariable

All content copied from https://docs.aws.amazon.com/.
