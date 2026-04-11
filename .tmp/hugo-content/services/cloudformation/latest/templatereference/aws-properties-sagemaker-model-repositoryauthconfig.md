This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Model RepositoryAuthConfig

Specifies an authentication configuration for the private docker registry where your
model image is hosted. Specify a value for this property only if you specified
`Vpc` as the value for the `RepositoryAccessMode` field of the
`ImageConfig` object that you passed to a call to
`CreateModel` and the private Docker registry where the model image is
hosted requires authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RepositoryCredentialsProviderArn" : String
}

```

### YAML

```yaml

  RepositoryCredentialsProviderArn: String

```

## Properties

`RepositoryCredentialsProviderArn`

The Amazon Resource Name (ARN) of an AWS Lambda function that provides
credentials to authenticate to the private Docker registry where your model image is
hosted. For information about how to create an AWS Lambda function, see
[Create a Lambda function\
with the console](../../../lambda/latest/dg/getting-started-create-function.md) in the _AWS Lambda Developer_
_Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiModelConfig

S3DataSource

All content copied from https://docs.aws.amazon.com/.
