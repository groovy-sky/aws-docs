This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::LoggerDefinition LoggerDefinitionVersion

A
logger definition version contains a list of [loggers](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-loggerdefinition-logger.html).

###### Note

After you create a logger definition version that contains the loggers you want to
deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

In
an CloudFormation template, `LoggerDefinitionVersion` is the property type
of the `InitialVersion` property in the [`AWS::Greengrass::LoggerDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-loggerdefinition.html) resource.

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

_Type_: Array of [Logger](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-loggerdefinition-logger.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [LoggerDefinitionVersion](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-loggerdefinitionversion.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Logger

AWS::Greengrass::LoggerDefinitionVersion
