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

_Type_: Array of [PolicyDefinitionRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionrule.html)

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Types`

The custom types defined within the policy definition.

_Required_: No

_Type_: Array of [PolicyDefinitionType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-automatedreasoningpolicy-policydefinitiontype.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variables`

The variables used within the policy definition.

_Required_: No

_Type_: Array of [PolicyDefinitionVariable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-automatedreasoningpolicy-policydefinitionvariable.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the policy definition.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Bedrock::AutomatedReasoningPolicy

PolicyDefinitionRule
