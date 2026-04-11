This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway CustomClaimValidationType

Defines the name of a custom claim field and rules for finding matches to authenticate its value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizingClaimMatchValue" : AuthorizingClaimMatchValueType,
  "InboundTokenClaimName" : String,
  "InboundTokenClaimValueType" : String
}

```

### YAML

```yaml

  AuthorizingClaimMatchValue:
    AuthorizingClaimMatchValueType
  InboundTokenClaimName: String
  InboundTokenClaimValueType: String

```

## Properties

`AuthorizingClaimMatchValue`

Defines the value or values to match for and the relationship of the match.

_Required_: Yes

_Type_: [AuthorizingClaimMatchValueType](aws-properties-bedrockagentcore-gateway-authorizingclaimmatchvaluetype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InboundTokenClaimName`

The name of the custom claim field to check.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9_.-:]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InboundTokenClaimValueType`

The data type of the claim value to check for.

- Use `STRING` if you want to find an exact match to a string you define.

- Use `STRING_ARRAY` if you want to fnd a match to at least one value in an array you define.

_Required_: Yes

_Type_: String

_Allowed values_: `STRING | STRING_ARRAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClaimMatchValueType

CustomJWTAuthorizerConfiguration

All content copied from https://docs.aws.amazon.com/.
