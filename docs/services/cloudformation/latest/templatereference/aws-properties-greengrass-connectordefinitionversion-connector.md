This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Greengrass::ConnectorDefinitionVersion Connector

Connectors are modules that provide built-in integration with local infrastructure, device
protocols, AWS, and other cloud services. For more information, see [Integrate\
with Services and Protocols Using Greengrass Connectors](https://docs.aws.amazon.com/greengrass/v1/developerguide/connectors.html) in the _AWS IoT Greengrass Version 1 Developer Guide_.

In an
CloudFormation template, the `Connectors` property of the [`AWS::Greengrass::ConnectorDefinitionVersion`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-connectordefinitionversion.html) resource contains a list of `Connector` property types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectorArn" : String,
  "Id" : String,
  "Parameters" : Json
}

```

### YAML

```yaml

  ConnectorArn: String
  Id: String
  Parameters: Json

```

## Properties

`ConnectorArn`

The Amazon Resource Name (ARN) of the connector.

For more information about connectors provided by AWS, see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/v1/developerguide/connectors-list.html).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Id`

A descriptive or arbitrary ID for the connector. This value must be unique within the
connector definition version. Maximum length is 128 characters with pattern
`[a-zA-Z0-9:_-]+`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

The parameters or configuration that the connector uses.

For more information about connectors provided by AWS, see [Greengrass Connectors Provided by AWS](https://docs.aws.amazon.com/greengrass/v1/developerguide/connectors-list.html).

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [Connector](https://docs.aws.amazon.com/greengrass/v1/apireference/definitions-connector.html) in the _AWS IoT Greengrass Version 1 API Reference_

- [AWS IoT Greengrass Version 1 Developer Guide](https://docs.aws.amazon.com/greengrass/v1/developerguide)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Greengrass::ConnectorDefinitionVersion

AWS::Greengrass::CoreDefinition
