This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Runtime CodeConfiguration

The configuration for the source code that defines how the agent runtime code should be executed, including the code location, runtime environment, and entry point.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Code" : Code,
  "EntryPoint" : [ String, ... ],
  "Runtime" : String
}

```

### YAML

```yaml

  Code:
    Code
  EntryPoint:
    - String
  Runtime: String

```

## Properties

`Code`

The source code location and configuration details.

_Required_: Yes

_Type_: [Code](aws-properties-bedrockagentcore-runtime-code.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntryPoint`

The entry point for the code execution, specifying the function or method that should be invoked when the code runs.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Runtime`

The runtime environment for executing the code (for example, Python 3.9 or Node.js 18).

_Required_: Yes

_Type_: String

_Allowed values_: `PYTHON_3_10 | PYTHON_3_11 | PYTHON_3_12 | PYTHON_3_13 | PYTHON_3_14`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Code

ContainerConfiguration

All content copied from https://docs.aws.amazon.com/.
