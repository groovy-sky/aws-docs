This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ContinuousDeploymentPolicy SingleWeightConfig

This configuration determines the percentage of HTTP requests that are sent to the
staging distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SessionStickinessConfig" : SessionStickinessConfig,
  "Weight" : Number
}

```

### YAML

```yaml

  SessionStickinessConfig:
    SessionStickinessConfig
  Weight: Number

```

## Properties

`SessionStickinessConfig`

Session stickiness provides the ability to define multiple requests from a single
viewer as a single session. This prevents the potentially inconsistent experience of
sending some of a given user's requests to your staging distribution, while others are
sent to your primary distribution. Define the session duration using TTL values.

_Required_: No

_Type_: [SessionStickinessConfig](aws-properties-cloudfront-continuousdeploymentpolicy-sessionstickinessconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

The percentage of traffic to send to a staging distribution, expressed as a decimal
number between 0 and 0.15. For example, a value of 0.10 means 10% of traffic is sent to the staging distribution.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SingleHeaderPolicyConfig

SingleWeightPolicyConfig

All content copied from https://docs.aws.amazon.com/.
