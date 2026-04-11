This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ContinuousDeploymentPolicy SingleHeaderPolicyConfig

Defines a single header policy for a CloudFront distribution.

###### Note

This property is legacy. We recommend that you use [TrafficConfig](../userguide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.md) and specify the [SingleHeaderConfig](../userguide/aws-properties-cloudfront-continuousdeploymentpolicy-trafficconfig.md#cfn-cloudfront-continuousdeploymentpolicy-trafficconfig-singleheaderconfig) property instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Header" : String,
  "Value" : String
}

```

### YAML

```yaml

  Header: String
  Value: String

```

## Properties

`Header`

The name of the HTTP header that CloudFront uses to configure for the single header policy.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Specifies the value to assign to the header for a single header policy.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1783`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SingleHeaderConfig

SingleWeightConfig

All content copied from https://docs.aws.amazon.com/.
