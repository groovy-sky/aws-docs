This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ContinuousDeploymentPolicy ContinuousDeploymentPolicyConfig

Contains the configuration for a continuous deployment policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "SingleHeaderPolicyConfig" : SingleHeaderPolicyConfig,
  "SingleWeightPolicyConfig" : SingleWeightPolicyConfig,
  "StagingDistributionDnsNames" : [ String, ... ],
  "TrafficConfig" : TrafficConfig,
  "Type" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  SingleHeaderPolicyConfig:
    SingleHeaderPolicyConfig
  SingleWeightPolicyConfig:
    SingleWeightPolicyConfig
  StagingDistributionDnsNames:
    - String
  TrafficConfig:
    TrafficConfig
  Type: String

```

## Properties

`Enabled`

A Boolean that indicates whether this continuous deployment policy is enabled (in
effect). When this value is `true`, this policy is enabled and in effect.
When this value is `false`, this policy is not enabled and has no
effect.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleHeaderPolicyConfig`

This configuration determines which HTTP requests are sent to the staging
distribution. If the HTTP request contains a header and value that matches what you
specify here, the request is sent to the staging distribution. Otherwise the request is
sent to the primary distribution.

_Required_: No

_Type_: [SingleHeaderPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleWeightPolicyConfig`

This configuration determines the percentage of HTTP requests that are sent to the
staging distribution.

_Required_: No

_Type_: [SingleWeightPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-continuousdeploymentpolicy-singleweightpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StagingDistributionDnsNames`

The CloudFront domain name of the staging distribution. For example:
`d111111abcdef8.cloudfront.net`.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficConfig`

Contains the parameters for routing production traffic from your primary to staging
distributions.

_Required_: No

_Type_: [TrafficConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of traffic configuration.

_Required_: No

_Type_: String

_Allowed values_: `SingleWeight | SingleHeader`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CloudFront::ContinuousDeploymentPolicy

SessionStickinessConfig
