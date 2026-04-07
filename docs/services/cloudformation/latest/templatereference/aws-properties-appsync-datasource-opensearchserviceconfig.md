This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DataSource OpenSearchServiceConfig

The `OpenSearchServiceConfig` property type specifies the
`AwsRegion` and `Endpoints` for an Amazon OpenSearch Service
domain in your account for an AWS AppSync data source.

`OpenSearchServiceConfig` is a property of the [AWS::AppSync::DataSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-datasource.html) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsRegion" : String,
  "Endpoint" : String
}

```

### YAML

```yaml

  AwsRegion: String
  Endpoint: String

```

## Properties

`AwsRegion`

The AWS Region.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

The endpoint.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LambdaConfig

RdsHttpEndpointConfig
