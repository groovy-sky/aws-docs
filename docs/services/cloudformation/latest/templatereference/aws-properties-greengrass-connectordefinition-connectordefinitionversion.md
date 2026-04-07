This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ConnectorDefinition ConnectorDefinitionVersion

A
connector definition version contains a list of connectors.

###### Note

After you create a connector definition version that contains the connectors you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-group.html).

In an CloudFormation template, `ConnectorDefinitionVersion` is the
property type of the `InitialVersion` property in the [`AWS::Greengrass::ConnectorDefinition`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinition.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Connectors" : [ Connector, ... ]
}

```

### YAML

```yaml

  Connectors:
    - Connector

```

## Properties

`Connectors`

The connectors in this version. Only one instance of a given connector can be added to a
connector definition version at a time.

_Required_: Yes

_Type_: Array of [Connector](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrass-connectordefinition-connector.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [ConnectorDefinitionVersion](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-connectordefinitionversion.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Connector

AWS::Greengrass::ConnectorDefinitionVersion
