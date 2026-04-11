This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ConnectorDefinition ConnectorDefinitionVersion

A
connector definition version contains a list of connectors.

###### Note

After you create a connector definition version that contains the connectors you want
to deploy, you must add it to your group version. For more information, see [`AWS::Greengrass::Group`](../userguide/aws-resource-greengrass-group.md).

In an CloudFormation template, `ConnectorDefinitionVersion` is the
property type of the `InitialVersion` property in the [`AWS::Greengrass::ConnectorDefinition`](../userguide/aws-resource-greengrass-connectordefinition.md) resource.

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

_Type_: Array of [Connector](aws-properties-greengrass-connectordefinition-connector.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [ConnectorDefinitionVersion](../../../../reference/greengrass/v1/apireference/definitions-connectordefinitionversion.md) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](../../../greengrass/v1/developerguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Connector

AWS::Greengrass::ConnectorDefinitionVersion

All content copied from https://docs.aws.amazon.com/.
