This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway AuthorizingClaimMatchValueType

Defines the value or values to match for and the relationship of the match.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClaimMatchOperator" : String,
  "ClaimMatchValue" : ClaimMatchValueType
}

```

### YAML

```yaml

  ClaimMatchOperator: String
  ClaimMatchValue:
    ClaimMatchValueType

```

## Properties

`ClaimMatchOperator`

Defines the relationship between the claim field value and the value or values you're matching for.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | CONTAINS | CONTAINS_ANY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClaimMatchValue`

The value or values to match for.

_Required_: Yes

_Type_: [ClaimMatchValueType](aws-properties-bedrockagentcore-gateway-claimmatchvaluetype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizerConfiguration

ClaimMatchValueType

All content copied from https://docs.aws.amazon.com/.
