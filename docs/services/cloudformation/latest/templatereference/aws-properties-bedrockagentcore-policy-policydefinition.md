This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Policy PolicyDefinition

The definition structure for policies. Encapsulates different policy formats.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cedar" : CedarPolicy
}

```

### YAML

```yaml

  Cedar:
    CedarPolicy

```

## Properties

`Cedar`

The Cedar policy definition.

_Required_: Yes

_Type_: [CedarPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-policy-cedarpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CedarPolicy

AWS::BedrockAgentCore::PolicyEngine
