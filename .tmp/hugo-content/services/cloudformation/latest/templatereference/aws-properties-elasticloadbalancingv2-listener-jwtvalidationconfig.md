This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::Listener JwtValidationConfig

The `JwtValidationConfig` property type specifies Property description not available. for an [AWS::ElasticLoadBalancingV2::Listener](aws-resource-elasticloadbalancingv2-listener.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalClaims" : [ JwtValidationActionAdditionalClaim, ... ],
  "Issuer" : String,
  "JwksEndpoint" : String
}

```

### YAML

```yaml

  AdditionalClaims:
    - JwtValidationActionAdditionalClaim
  Issuer: String
  JwksEndpoint: String

```

## Properties

`AdditionalClaims`

Property description not available.

_Required_: No

_Type_: Array of [JwtValidationActionAdditionalClaim](aws-properties-elasticloadbalancingv2-listener-jwtvalidationactionadditionalclaim.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Issuer`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JwksEndpoint`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JwtValidationActionAdditionalClaim

ListenerAttribute

All content copied from https://docs.aws.amazon.com/.
