This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::IdentityPool CognitoStreams

`CognitoStreams` is a property of the [AWS::Cognito::IdentityPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html) resource that defines configuration options for
Amazon Cognito streams.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String,
  "StreamingStatus" : String,
  "StreamName" : String
}

```

### YAML

```yaml

  RoleArn: String
  StreamingStatus: String
  StreamName: String

```

## Properties

`RoleArn`

The Amazon Resource Name (ARN) of the role Amazon Cognito can assume to publish to the
stream. This role must grant access to Amazon Cognito (cognito-sync) to invoke
`PutRecord` on your Amazon Cognito stream.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamingStatus`

Status of the Amazon Cognito streams. Valid values are: `ENABLED` or
`DISABLED`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamName`

The name of the Amazon Cognito stream to receive updates. This stream must be in the
developer's account and in the same Region as the identity pool.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CognitoIdentityProvider

PushSync
