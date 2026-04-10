This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::SecurityProfile FlowModule

A list of Flow Modules an AI Agent can invoke as a tool

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FlowModuleId" : String,
  "Type" : String
}

```

### YAML

```yaml

  FlowModuleId: String
  Type: String

```

## Properties

`FlowModuleId`

If of Flow Modules invocable as tool

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Only Type we support is MCP.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataTableAccessControlConfiguration

GranularAccessControlConfiguration

All content copied from https://docs.aws.amazon.com/.
