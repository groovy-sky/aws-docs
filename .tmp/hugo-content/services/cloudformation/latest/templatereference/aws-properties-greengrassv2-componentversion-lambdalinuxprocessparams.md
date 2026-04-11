This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion LambdaLinuxProcessParams

Contains parameters for a Linux process that contains an AWS Lambda
function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerParams" : LambdaContainerParams,
  "IsolationMode" : String
}

```

### YAML

```yaml

  ContainerParams:
    LambdaContainerParams
  IsolationMode: String

```

## Properties

`ContainerParams`

The parameters for the container in which the Lambda function runs.

_Required_: No

_Type_: [LambdaContainerParams](aws-properties-greengrassv2-componentversion-lambdacontainerparams.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IsolationMode`

The isolation mode for the process that contains the Lambda function. The
process can run in an isolated runtime environment inside the AWS IoT Greengrass
container, or as a regular process outside any container.

Default: `GreengrassContainer`

_Required_: No

_Type_: String

_Allowed values_: `GreengrassContainer | NoContainer`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaFunctionRecipeSource

LambdaVolumeMount

All content copied from https://docs.aws.amazon.com/.
