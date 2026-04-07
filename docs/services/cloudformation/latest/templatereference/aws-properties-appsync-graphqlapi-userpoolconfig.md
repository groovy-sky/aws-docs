This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::GraphQLApi UserPoolConfig

The `UserPoolConfig` property type specifies the optional authorization
configuration for using Amazon Cognito user pools with your GraphQL endpoint for an AWS AppSync GraphQL API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppIdClientRegex" : String,
  "AwsRegion" : String,
  "DefaultAction" : String,
  "UserPoolId" : String
}

```

### YAML

```yaml

  AppIdClientRegex: String
  AwsRegion: String
  DefaultAction: String
  UserPoolId: String

```

## Properties

`AppIdClientRegex`

A regular expression for validating the incoming Amazon Cognito user pool app client ID. If this value
isn't set, no filtering is applied.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`AwsRegion`

The AWS Region in which the user pool was created.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DefaultAction`

The action that you want your GraphQL API to take when a request that uses Amazon Cognito user pool authentication doesn't match the Amazon Cognito user
pool configuration.

When specifying Amazon Cognito user pools as the default authentication, you must set the
value for `DefaultAction` to `ALLOW` if specifying
`AdditionalAuthenticationProviders`.

_Required_: No

_Type_: String

_Allowed values_: `ALLOW | DENY`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`UserPoolId`

The user pool ID.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::AppSync::GraphQLSchema
