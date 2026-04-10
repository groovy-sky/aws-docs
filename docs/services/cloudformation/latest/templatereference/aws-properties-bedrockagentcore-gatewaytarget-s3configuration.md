This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget S3Configuration

The Amazon S3 configuration for a gateway. This structure defines how the gateway accesses files in Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketOwnerAccountId" : String,
  "Uri" : String
}

```

### YAML

```yaml

  BucketOwnerAccountId: String
  Uri: String

```

## Properties

`BucketOwnerAccountId`

The account ID of the Amazon S3 bucket owner. This ID is used for cross-account access to the bucket.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Uri`

The URI of the Amazon S3 object. This URI specifies the location of the object in Amazon S3.

_Required_: No

_Type_: String

_Pattern_: `^s3://.{1,2043}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuthCredentialProvider

SchemaDefinition

All content copied from https://docs.aws.amazon.com/.
