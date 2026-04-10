This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Workteam CognitoMemberDefinition

Identifies a Amazon Cognito user group. A user group can be used in on or more work
teams.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CognitoClientId" : String,
  "CognitoUserGroup" : String,
  "CognitoUserPool" : String
}

```

### YAML

```yaml

  CognitoClientId: String
  CognitoUserGroup: String
  CognitoUserPool: String

```

## Properties

`CognitoClientId`

An identifier for an application client. You must create the app client ID using Amazon Cognito.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CognitoUserGroup`

An identifier for a user group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CognitoUserPool`

An identifier for a user pool. The user pool must be in the same region as the service that you are
calling.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::Workteam

MemberDefinition

All content copied from https://docs.aws.amazon.com/.
