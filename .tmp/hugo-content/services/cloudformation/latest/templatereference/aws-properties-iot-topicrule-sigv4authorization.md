This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule SigV4Authorization

For more information, see [Signature Version 4 signing process](../../../../general/latest/gr/signature-version-4.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String,
  "ServiceName" : String,
  "SigningRegion" : String
}

```

### YAML

```yaml

  RoleArn: String
  ServiceName: String
  SigningRegion: String

```

## Properties

`RoleArn`

The ARN of the signing role.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

The service name to use while signing with Sig V4.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SigningRegion`

The signing region.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Action

SnsAction

All content copied from https://docs.aws.amazon.com/.
