This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::Api CognitoConfig

Describes an Amazon Cognito configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppIdClientRegex" : String,
  "AwsRegion" : String,
  "UserPoolId" : String
}

```

### YAML

```yaml

  AppIdClientRegex: String
  AwsRegion: String
  UserPoolId: String

```

## Properties

`AppIdClientRegex`

A regular expression for validating the incoming Amazon Cognito user pool app
client ID. If this value isn't set, no filtering is applied.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsRegion`

The AWS Region in which the user pool was created.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The user pool ID.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthProvider

DnsMap

All content copied from https://docs.aws.amazon.com/.
