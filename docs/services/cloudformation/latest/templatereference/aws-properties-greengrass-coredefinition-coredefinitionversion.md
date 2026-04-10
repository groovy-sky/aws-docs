This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::CoreDefinition CoreDefinitionVersion

A core
definition version contains a Greengrass [core](../userguide/aws-properties-greengrass-coredefinition-core.md).

###### Note

After you create a core definition version that contains the core you want to deploy,
you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

In an
CloudFormation template, `CoreDefinitionVersion` is the property type
of the `InitialVersion` property in the [`AWS::Greengrass::CoreDefinition`](../userguide/aws-resource-greengrass-coredefinition.md) resource.

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

_Type_: Array of [Core](aws-properties-greengrass-coredefinition-core.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [CoreDefinitionVersion](../../../../reference/greengrass/v1/apireference/definitions-coredefinitionversion.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Core

AWS::Greengrass::CoreDefinitionVersion

All content copied from https://docs.aws.amazon.com/.
