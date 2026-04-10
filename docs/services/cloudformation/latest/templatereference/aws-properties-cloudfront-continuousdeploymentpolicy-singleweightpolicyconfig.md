This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ContinuousDeploymentPolicy SingleWeightPolicyConfig

Configure a policy that CloudFront uses to route requests to different origins or use different cache settings, based on the weight assigned to each option.

###### Note

This property is legacy. We recommend that you use [TrafficConfig](../userguide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.md) and specify the [SingleWeightConfig](../userguide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.md#cfn-cloudfront-continuousdeploymentpolicy-trafficconfig-singleweightconfig) property instead.

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

Enable session stickiness for the associated origin or cache settings.

_Required_: No

_Type_: [SessionStickinessConfig](aws-properties-cloudfront-continuousdeploymentpolicy-sessionstickinessconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Weight`

The percentage of requests that CloudFront will use to send to an associated origin or cache settings.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SingleWeightConfig

TrafficConfig

All content copied from https://docs.aws.amazon.com/.
