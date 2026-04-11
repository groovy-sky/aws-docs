This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource AuthorizationConfig

The `AuthorizationConfig` property type specifies the authorization type
and configuration for an AWS AppSync http data source.

`AuthorizationConfig` is a property of the [AWS AppSync DataSource HttpConfig](../userguide/aws-properties-appsync-datasource-httpconfig.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationType" : String,
  "AwsIamConfig" : AwsIamConfig
}

```

### YAML

```yaml

  AuthorizationType: String
  AwsIamConfig:
    AwsIamConfig

```

## Properties

`AuthorizationType`

The authorization type that the HTTP endpoint requires.

- **AWS\_IAM**: The authorization type is Signature Version 4
(SigV4).

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_IAM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsIamConfig`

The AWS Identity and Access Management settings.

_Required_: No

_Type_: [AwsIamConfig](aws-properties-appsync-datasource-awsiamconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppSync::DataSource

AwsIamConfig

All content copied from https://docs.aws.amazon.com/.
