This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinition Execution

Configuration
settings for the Lambda execution environment on the AWS IoT Greengrass
core.

In an CloudFormation template, `Execution` is a property of the [`DefaultConfig`](../userguide/aws-properties-greengrass-functiondefinition-defaultconfig.md) property type for a function definition version and the [`Environment`](../userguide/aws-properties-greengrass-functiondefinition-environment.md) property type for a function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsolationMode" : String,
  "RunAs" : RunAs
}

```

### YAML

```yaml

  IsolationMode: String
  RunAs:
    RunAs

```

## Properties

`IsolationMode`

The containerization that the Lambda function runs in. Valid values are
`GreengrassContainer` or `NoContainer`. Typically, this is
`GreengrassContainer`. For more information, see [Containerization](../../../greengrass/v1/developerguide/lambda-group-config.md#lambda-function-containerization) in the _AWS IoT Greengrass Version 1 Developer Guide_.

- When set on the [`DefaultConfig`](../userguide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.md) property of a function definition version, this setting is used as the
default containerization for all Lambda functions in the function
definition version.

- When set on the [`Environment`](../userguide/aws-properties-greengrass-functiondefinitionversion-environment.md) property of a function, this setting applies to the individual function and
overrides the default. Omit this value to run the function with the default
containerization.

###### Note

We recommend that you run in a Greengrass container unless your business case
requires that you run without containerization.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RunAs`

The user and group permissions used to run the Lambda function. Typically,
this is the ggc\_user and ggc\_group. For more information, see [Run as](../../../greengrass/v1/developerguide/lambda-group-config.md#lambda-access-identity.html) in the _AWS IoT Greengrass Version 1 Developer Guide_.

- When set on the [`DefaultConfig`](../userguide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.md) property of a function definition version, this setting is used as the
default access identity for all Lambda functions in the function
definition version.

- When set on the [`Environment`](../userguide/aws-properties-greengrass-functiondefinitionversion-environment.md) property of a function, this setting applies to the individual function and
overrides the default. You can override the user, group, or both. Omit this value to
run the function with the default permissions.

###### Important

Running as the root user increases risks to your data and device. Do not run as root
(UID/GID=0) unless your business case requires it. For more information and
requirements, see [Running a Lambda Function as Root](../../../greengrass/v1/developerguide/lambda-group-config.md#lambda-running-as-root).

_Required_: No

_Type_: [RunAs](aws-properties-greengrass-functiondefinition-runas.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [FunctionExecutionConfig](../../../../reference/greengrass/v1/apireference/definitions-functionexecutionconfig.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Environment

Function

All content copied from https://docs.aws.amazon.com/.
