This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::LoggerDefinition LoggerDefinitionVersion

A
logger definition version contains a list of [loggers](../userguide/aws-properties-greengrass-loggerdefinition-logger.md).

###### Note

After you create a logger definition version that contains the loggers you want to
deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

In
an CloudFormation template, `LoggerDefinitionVersion` is the property type
of the `InitialVersion` property in the [`AWS::Greengrass::LoggerDefinition`](../userguide/aws-resource-greengrass-loggerdefinition.md) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Loggers" : [ Logger, ... ]
}

```

### YAML

```yaml

  Loggers:
    - Logger

```

## Properties

`Loggers`

The loggers in this version.

_Required_: Yes

_Type_: Array of [Logger](aws-properties-greengrass-loggerdefinition-logger.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [LoggerDefinitionVersion](../../../../reference/greengrass/v1/apireference/definitions-loggerdefinitionversion.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logger

AWS::Greengrass::LoggerDefinitionVersion

All content copied from https://docs.aws.amazon.com/.
