This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget ApiGatewayToolOverride

Settings to override configurations for a tool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "Method" : String,
  "Name" : String,
  "Path" : String
}

```

### YAML

```yaml

  Description: String
  Method: String
  Name: String
  Path: String

```

## Properties

`Description`

The description of the tool. Provides information about the purpose and usage of the tool. If not provided, uses the description from the API's OpenAPI specification.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Method`

The HTTP method to expose for the specified path.

_Required_: Yes

_Type_: String

_Allowed values_: `GET | DELETE | HEAD | OPTIONS | PATCH | PUT | POST`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of tool. Identifies the tool in the Model Context Protocol.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

Resource path in the REST API (e.g., `/pets`). Must explicitly match an existing path in the REST API.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiGatewayToolFilter

ApiKeyCredentialProvider

All content copied from https://docs.aws.amazon.com/.
