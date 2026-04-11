This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::GatewayTarget ApiGatewayToolFilter

Specifies which operations from an API Gateway REST API are exposed as tools. Tool names and descriptions are derived from the operationId and description fields in the API's exported OpenAPI specification.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FilterPath" : String,
  "Methods" : [ String, ... ]
}

```

### YAML

```yaml

  FilterPath: String
  Methods:
    - String

```

## Properties

`FilterPath`

Resource path to match in the REST API. Supports exact paths (for example, `/pets`) or wildcard paths (for example, `/pets/*` to match all paths under `/pets`). Must match existing paths in the REST API.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Methods`

The methods to filter for.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiGatewayToolConfiguration

ApiGatewayToolOverride

All content copied from https://docs.aws.amazon.com/.
