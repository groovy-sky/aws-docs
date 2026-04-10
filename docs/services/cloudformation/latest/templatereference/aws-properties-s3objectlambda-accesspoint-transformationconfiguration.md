This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3ObjectLambda::AccessPoint TransformationConfiguration

A configuration used when creating an Object Lambda Access Point transformation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Actions" : [ String, ... ],
  "ContentTransformation" : ContentTransformation
}

```

### YAML

```yaml

  Actions:
    - String
  ContentTransformation:
    ContentTransformation

```

## Properties

`Actions`

A container for the action of an Object Lambda Access Point configuration. Valid
inputs are `GetObject`, `HeadObject`, `ListObject`, and
`ListObjectV2`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentTransformation`

A container for the content transformation of an Object Lambda Access Point
configuration. Can include the FunctionArn and FunctionPayload. For more information,
see [AwsLambdaTransformation](../../../s3/latest/api/api-control-awslambdatransformation.md) in the _Amazon S3 API_
_Reference_.

_Required_: Yes

_Type_: [ContentTransformation](aws-properties-s3objectlambda-accesspoint-contenttransformation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublicAccessBlockConfiguration

AWS::S3ObjectLambda::AccessPointPolicy

All content copied from https://docs.aws.amazon.com/.
