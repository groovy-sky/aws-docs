This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::FunctionDefinition FunctionDefinitionVersion

A function definition version contains a list of functions.

###### Note

After you create a function definition version that contains the functions you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

In an CloudFormation template, `FunctionDefinitionVersion` is the property
type of the `InitialVersion` property in the [`AWS::Greengrass::FunctionDefinition`](../userguide/aws-resource-greengrass-functiondefinition.md) resource.

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

_Type_: [DefaultConfig](aws-properties-greengrass-functiondefinition-defaultconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Functions`

The functions in this version.

_Required_: Yes

_Type_: Array of [Function](aws-properties-greengrass-functiondefinition-function.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [FunctionDefinitionVersion](../../../../reference/greengrass/v1/apireference/definitions-functiondefinitionversion.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FunctionConfiguration

ResourceAccessPolicy

All content copied from https://docs.aws.amazon.com/.
