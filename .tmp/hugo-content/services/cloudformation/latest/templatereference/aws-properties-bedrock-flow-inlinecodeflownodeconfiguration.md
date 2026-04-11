This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow InlineCodeFlowNodeConfiguration

Contains configurations for an inline code node in your flow. Inline code nodes let you write and execute code directly within your flow, enabling data transformations, custom logic, and integrations without needing an external Lambda function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Code" : String,
  "Language" : String
}

```

### YAML

```yaml

  Code: String
  Language: String

```

## Properties

`Code`

The code that's executed in your inline code node. The code can access input data from previous nodes in the flow, perform operations on that data, and produce output that can be used by other nodes in your flow.

The code must be valid in the programming `language` that you specify.

_Required_: Yes

_Type_: String

_Maximum_: `5000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Language`

The programming language used by your inline code node.

The code must be valid in the programming `language` that you specify. Currently, only Python 3 ( `Python_3`) is supported.

_Required_: Yes

_Type_: String

_Allowed values_: `Python_3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GuardrailConfiguration

KnowledgeBaseFlowNodeConfiguration

All content copied from https://docs.aws.amazon.com/.
