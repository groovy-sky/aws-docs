This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinitionVersion Function

A function
is a Lambda function that's referenced from an AWS IoT Greengrass group. The
function is deployed to a Greengrass core where it runs locally. For more information, see
[Run Lambda Functions on the AWS IoT Greengrass Core](../../../greengrass/v1/developerguide/lambda-functions.md) in the
_AWS IoT Greengrass Version 1 Developer Guide_.

In an
CloudFormation template, the `Functions` property of the [`AWS::Greengrass::FunctionDefinitionVersion`](../userguide/aws-resource-greengrass-functiondefinitionversion.md) resource contains a list of `Function` property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FunctionArn" : String,
  "FunctionConfiguration" : FunctionConfiguration,
  "Id" : String
}

```

### YAML

```yaml

  FunctionArn: String
  FunctionConfiguration:
    FunctionConfiguration
  Id: String

```

## Properties

`FunctionArn`

The Amazon Resource Name (ARN) of the alias (recommended) or version of the referenced
Lambda function.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FunctionConfiguration`

The group-specific settings of the Lambda function. These settings
configure the function's behavior in the Greengrass group.

_Required_: Yes

_Type_: [FunctionConfiguration](aws-properties-greengrass-functiondefinitionversion-functionconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Id`

A descriptive or arbitrary ID for the function. This value must be unique within the
function definition version. Maximum length is 128 characters with pattern
`[a-zA-Z0-9:_-]+`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Function](../../../../reference/greengrass/v1/apireference/definitions-function.md)
in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Execution

FunctionConfiguration

All content copied from https://docs.aws.amazon.com/.
