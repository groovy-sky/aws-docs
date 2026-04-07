This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinition FunctionDefinitionVersion

A function definition version contains a list of functions.

###### Note

After you create a function definition version that contains the functions you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

In an CloudFormation template, `FunctionDefinitionVersion` is the property
type of the `InitialVersion` property in the [`AWS::Greengrass::FunctionDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-functiondefinition.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultConfig" : DefaultConfig,
  "Functions" : [ Function, ... ]
}

```

### YAML

```yaml

  DefaultConfig:
    DefaultConfig
  Functions:
    - Function

```

## Properties

`DefaultConfig`

The default configuration that applies to all Lambda functions in the
group. Individual Lambda functions can override these settings.

_Required_: No

_Type_: [DefaultConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-functiondefinition-defaultconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Functions`

The functions in this version.

_Required_: Yes

_Type_: Array of [Function](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-functiondefinition-function.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [FunctionDefinitionVersion](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-functiondefinitionversion.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FunctionConfiguration

ResourceAccessPolicy
