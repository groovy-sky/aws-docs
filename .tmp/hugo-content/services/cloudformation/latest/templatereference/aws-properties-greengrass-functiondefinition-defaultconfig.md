This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinition DefaultConfig

The default
configuration that applies to all Lambda functions in the function definition
version. Individual Lambda functions can override these settings.

In an
CloudFormation template, `DefaultConfig` is a property of the [`FunctionDefinitionVersion`](../userguide/aws-properties-greengrass-functiondefinition-functiondefinitionversion.md) property type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Execution" : Execution
}

```

### YAML

```yaml

  Execution:
    Execution

```

## Properties

`Execution`

Configuration settings for the Lambda execution environment on the AWS IoT Greengrass core.

_Required_: Yes

_Type_: [Execution](aws-properties-greengrass-functiondefinition-execution.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [FunctionDefaultConfig](../../../../reference/greengrass/v1/apireference/definitions-functiondefaultconfig.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Greengrass::FunctionDefinition

Environment

All content copied from https://docs.aws.amazon.com/.
