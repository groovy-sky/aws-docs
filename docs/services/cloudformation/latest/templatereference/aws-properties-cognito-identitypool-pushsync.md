This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::IdentityPool PushSync

`PushSync` is a property of the [AWS::Cognito::IdentityPool](../userguide/aws-resource-cognito-identitypool.md) resource that defines the configuration options
to be applied to an Amazon Cognito identity pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApplicationArns" : [ String, ... ],
  "RoleArn" : String
}

```

### YAML

```yaml

  ApplicationArns:
    - String
  RoleArn: String

```

## Properties

`ApplicationArns`

The ARNs of the Amazon SNS platform applications that could be used by clients.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

An IAM role configured to allow Amazon Cognito to call Amazon SNS on behalf of the
developer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CognitoStreams

Tag

All content copied from https://docs.aws.amazon.com/.
