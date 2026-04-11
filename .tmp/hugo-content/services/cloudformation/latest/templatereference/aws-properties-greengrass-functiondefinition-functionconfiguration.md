This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinition FunctionConfiguration

The
group-specific configuration settings for a Lambda function. These settings
configure the function's behavior in the Greengrass group. For more information, see [Controlling Execution of Greengrass Lambda Functions by Using\
Group-Specific Configuration](../../../greengrass/v1/developerguide/lambda-group-config.md) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In
an CloudFormation template, `FunctionConfiguration` is a property of the
[`Function`](../userguide/aws-properties-greengrass-functiondefinition-function.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EncodingType" : String,
  "Environment" : Environment,
  "ExecArgs" : String,
  "Executable" : String,
  "MemorySize" : Integer,
  "Pinned" : Boolean,
  "Timeout" : Integer
}

```

### YAML

```yaml

  EncodingType: String
  Environment:
    Environment
  ExecArgs: String
  Executable: String
  MemorySize: Integer
  Pinned: Boolean
  Timeout: Integer

```

## Properties

`EncodingType`

The expected encoding type of the input payload for the function. Valid values are
`json` (default) and `binary`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Environment`

The environment configuration of the function.

_Required_: No

_Type_: [Environment](aws-properties-greengrass-functiondefinition-environment.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecArgs`

The execution arguments.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Executable`

The name of the function executable.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemorySize`

The memory size (in KB) required by the function.

###### Note

This property applies only to Lambda functions that run in a Greengrass
container.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Pinned`

Indicates whether the function is pinned (or _long-lived_). Pinned
functions start when the core starts and process all requests in the same container. The
default value is false.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Timeout`

The allowed execution time (in seconds) after which the function should terminate. For
pinned functions, this timeout applies for each request.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [FunctionConfiguration](../../../../reference/greengrass/v1/apireference/definitions-functionconfiguration.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Function

FunctionDefinitionVersion

All content copied from https://docs.aws.amazon.com/.
