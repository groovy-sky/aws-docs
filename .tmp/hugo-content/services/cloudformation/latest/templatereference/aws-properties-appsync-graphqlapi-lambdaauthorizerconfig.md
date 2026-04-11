This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::GraphQLApi LambdaAuthorizerConfig

Configuration for AWS Lambda function authorization.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizerResultTtlInSeconds" : Integer,
  "AuthorizerUri" : String,
  "IdentityValidationExpression" : String
}

```

### YAML

```yaml

  AuthorizerResultTtlInSeconds: Integer
  AuthorizerUri: String
  IdentityValidationExpression: String

```

## Properties

`AuthorizerResultTtlInSeconds`

The number of seconds a response should be cached for. The default is 0 seconds, which
disables caching. If you don't specify a value for
`authorizerResultTtlInSeconds`, the default value is used. The maximum
value is one hour (3600 seconds). The Lambda function can override this by
returning a `ttlOverride` key in its response.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `3600`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`AuthorizerUri`

The ARN of the Lambda function to be called for authorization. This may be a standard
Lambda ARN, a version ARN ( `.../v3`) or alias ARN.

_Note_: This Lambda function must have the following resource-based
policy assigned to it. When configuring Lambda authorizers in the console, this is done
for you. To do so with the AWS CLI, run the following:

`aws lambda add-permission --function-name
                "arn:aws:lambda:us-east-2:111122223333:function:my-function" --statement-id
                "appsync" --principal appsync.amazonaws.com --action lambda:InvokeFunction`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`IdentityValidationExpression`

A regular expression for validation of tokens before the Lambda function is
called.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnhancedMetricsConfig

LogConfig

All content copied from https://docs.aws.amazon.com/.
