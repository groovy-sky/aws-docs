This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ContinuousDeploymentPolicy TrafficConfig

The traffic configuration of your continuous deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SingleHeaderConfig" : SingleHeaderConfig,
  "SingleWeightConfig" : SingleWeightConfig,
  "Type" : String
}

```

### YAML

```yaml

  SingleHeaderConfig:
    SingleHeaderConfig
  SingleWeightConfig:
    SingleWeightConfig
  Type: String

```

## Properties

`SingleHeaderConfig`

Determines which HTTP requests are sent to the staging distribution.

_Required_: No

_Type_: [SingleHeaderConfig](aws-properties-cloudfront-continuousdeploymentpolicy-singleheaderconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleWeightConfig`

Contains the percentage of traffic to send to the staging distribution.

_Required_: No

_Type_: [SingleWeightConfig](aws-properties-cloudfront-continuousdeploymentpolicy-singleweightconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of traffic configuration.

_Required_: Yes

_Type_: String

_Allowed values_: `SingleWeight | SingleHeader`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SingleWeightPolicyConfig

AWS::CloudFront::Distribution

All content copied from https://docs.aws.amazon.com/.
