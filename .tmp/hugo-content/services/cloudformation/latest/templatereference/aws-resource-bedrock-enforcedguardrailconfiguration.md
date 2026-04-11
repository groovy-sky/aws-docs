This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::EnforcedGuardrailConfiguration

Sets the account-level enforced guardrail configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::EnforcedGuardrailConfiguration",
  "Properties" : {
      "GuardrailIdentifier" : String,
      "GuardrailVersion" : String,
      "ModelEnforcement" : ModelEnforcement,
      "SelectiveContentGuarding" : SelectiveContentGuarding
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::EnforcedGuardrailConfiguration
Properties:
  GuardrailIdentifier: String
  GuardrailVersion: String
  ModelEnforcement:
    ModelEnforcement
  SelectiveContentGuarding:
    SelectiveContentGuarding

```

## Properties

`GuardrailIdentifier`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^(([a-z0-9]+)|(arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:guardrail/[a-z0-9]+))$`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GuardrailVersion`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[1-9][0-9]{0,7}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelEnforcement`

Model-specific information for the enforced guardrail configuration.

_Required_: No

_Type_: [ModelEnforcement](aws-properties-bedrock-enforcedguardrailconfiguration-modelenforcement.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectiveContentGuarding`

Selective content guarding controls for enforced guardrails.

_Required_: No

_Type_: [SelectiveContentGuarding](aws-properties-bedrock-enforcedguardrailconfiguration-selectivecontentguarding.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ConfigId`

Property description not available.

`CreatedAt`

Property description not available.

`CreatedBy`

Property description not available.

`GuardrailArn`

Property description not available.

`GuardrailId`

Property description not available.

`Owner`

Property description not available.

`UpdatedAt`

Property description not available.

`UpdatedBy`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebSourceConfiguration

ModelEnforcement

All content copied from https://docs.aws.amazon.com/.
