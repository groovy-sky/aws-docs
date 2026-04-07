This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::CoreDefinition CoreDefinitionVersion

A core
definition version contains a Greengrass [core](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-coredefinition-core.html).

###### Note

After you create a core definition version that contains the core you want to deploy,
you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

In an
CloudFormation template, `CoreDefinitionVersion` is the property type
of the `InitialVersion` property in the [`AWS::Greengrass::CoreDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-coredefinition.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cores" : [ Core, ... ]
}

```

### YAML

```yaml

  Cores:
    - Core

```

## Properties

`Cores`

The Greengrass core in this version. Currently, the `Cores` property for a
core definition version can contain only one core.

_Required_: Yes

_Type_: Array of [Core](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-coredefinition-core.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [CoreDefinitionVersion](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-coredefinitionversion.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Core

AWS::Greengrass::CoreDefinitionVersion
