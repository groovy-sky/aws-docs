This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::AutomatedReasoningPolicy PolicyDefinition

The complete policy definition containing rules, variables, and types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rules" : [ PolicyDefinitionRule, ... ],
  "Types" : [ PolicyDefinitionType, ... ],
  "Variables" : [ PolicyDefinitionVariable, ... ],
  "Version" : String
}

```

### YAML

```yaml

  Rules:
    - PolicyDefinitionRule
  Types:
    - PolicyDefinitionType
  Variables:
    - PolicyDefinitionVariable
  Version: String

```

## Properties

`Rules`

The collection of rules that define the policy logic.

_Required_: No

_Type_: Array of [PolicyDefinitionRule](aws-properties-bedrock-automatedreasoningpolicy-policydefinitionrule.md)

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Types`

The custom types defined within the policy definition.

_Required_: No

_Type_: Array of [PolicyDefinitionType](aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variables`

The variables used within the policy definition.

_Required_: No

_Type_: Array of [PolicyDefinitionVariable](aws-properties-bedrock-automatedreasoningpolicy-policydefinitionvariable.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the policy definition.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Bedrock::AutomatedReasoningPolicy

PolicyDefinitionRule

All content copied from https://docs.aws.amazon.com/.
