This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinitionVersion DefaultConfig

The
default configuration that applies to all Lambda functions in the function
definition version. Individual Lambda functions can override these
settings.

In
an CloudFormation template, `DefaultConfig` is a property of the [`AWS::Greengrass::FunctionDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinitionversion.html) resource.

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

_Type_: [Execution](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-functiondefinitionversion-execution.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [FunctionDefaultConfig](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-functiondefaultconfig.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Greengrass::FunctionDefinitionVersion

Environment
