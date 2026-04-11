This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource ServerSideEncryptionConfiguration

Contains the configuration for server-side encryption.

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

The Amazon Resource Name (ARN) of the AWS KMS key used to encrypt the resource.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-cn|-us-gov|-eusc|-iso(-[b-f])?)?:kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SemanticChunkingConfiguration

SharePointCrawlerConfiguration

All content copied from https://docs.aws.amazon.com/.
