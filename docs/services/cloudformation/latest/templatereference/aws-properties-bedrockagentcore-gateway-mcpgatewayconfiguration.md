This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway MCPGatewayConfiguration

The configuration for a Model Context Protocol (MCP) gateway. This structure defines how the gateway implements the MCP protocol.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Instructions" : String,
  "SearchType" : String,
  "SupportedVersions" : [ String, ... ]
}

```

### YAML

```yaml

  Instructions: String
  SearchType: String
  SupportedVersions:
    - String

```

## Properties

`Instructions`

The instructions for using the Model Context Protocol gateway. These instructions provide guidance on how to interact with the gateway.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SearchType`

The search type for the Model Context Protocol gateway. This field specifies how the gateway handles search operations.

_Required_: No

_Type_: String

_Allowed values_: `SEMANTIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SupportedVersions`

The supported versions of the Model Context Protocol. This field specifies which versions of the protocol the gateway can use.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LambdaInterceptorConfiguration

WorkloadIdentityDetails

All content copied from https://docs.aws.amazon.com/.
