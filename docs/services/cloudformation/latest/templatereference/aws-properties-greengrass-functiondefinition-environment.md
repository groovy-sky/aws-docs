This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinition Environment

The
environment configuration for a Lambda function on the AWS IoT Greengrass
core.

In an CloudFormation template, `Environment` is a property of the [`FunctionConfiguration`](../userguide/aws-properties-greengrass-functiondefinition-functionconfiguration.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessSysfs" : Boolean,
  "Execution" : Execution,
  "ResourceAccessPolicies" : [ ResourceAccessPolicy, ... ],
  "Variables" : Json
}

```

### YAML

```yaml

  AccessSysfs: Boolean
  Execution:
    Execution
  ResourceAccessPolicies:
    - ResourceAccessPolicy
  Variables: Json

```

## Properties

`AccessSysfs`

Indicates whether the function is allowed to access the `/sys` directory on
the core device, which allows the read device information from `/sys`.

###### Note

This property applies only to Lambda functions that run in a Greengrass
container.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Execution`

Settings for the Lambda execution environment in AWS IoT Greengrass.

_Required_: No

_Type_: [Execution](aws-properties-greengrass-functiondefinition-execution.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceAccessPolicies`

A list of the [resources](../userguide/aws-properties-greengrass-resourcedefinitionversion-resourceinstance.md) in the group that the function can access, with the corresponding
read-only or read-write permissions. The maximum is 10 resources.

###### Note

This property applies only for Lambda functions that run in a
Greengrass container.

_Required_: No

_Type_: Array of [ResourceAccessPolicy](aws-properties-greengrass-functiondefinition-resourceaccesspolicy.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Variables`

Environment variables for the Lambda function.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [FunctionConfigurationEnvironment](../../../../reference/greengrass/v1/apireference/definitions-functionconfigurationenvironment.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultConfig

Execution

All content copied from https://docs.aws.amazon.com/.
