This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::EnforcedGuardrailConfiguration ModelEnforcement

Model-specific information for the enforced guardrail configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExcludedModels" : [ String, ... ],
  "IncludedModels" : [ String, ... ]
}

```

### YAML

```yaml

  ExcludedModels:
    - String
  IncludedModels:
    - String

```

## Properties

`ExcludedModels`

Models to exclude from enforcement of the guardrail.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludedModels`

Models to enforce the guardrail on.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Bedrock::EnforcedGuardrailConfiguration

SelectiveContentGuarding

All content copied from https://docs.aws.amazon.com/.
