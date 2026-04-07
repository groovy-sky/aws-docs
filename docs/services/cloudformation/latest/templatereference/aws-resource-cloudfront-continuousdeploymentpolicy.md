This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ContinuousDeploymentPolicy

Creates a continuous deployment policy that routes a subset of production traffic
from a primary distribution to a staging distribution.

After you create and update a staging distribution, you can use a continuous
deployment policy to incrementally move traffic to the staging distribution. This enables
you to test changes to a distribution's configuration before moving all of your
production traffic to the new configuration.

For more information, see [Using\
CloudFront continuous deployment to safely test CDN configuration changes](../../../amazoncloudfront/latest/developerguide/continuous-deployment.md)
in the _Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::ContinuousDeploymentPolicy",
  "Properties" : {
      "ContinuousDeploymentPolicyConfig" : ContinuousDeploymentPolicyConfig
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::ContinuousDeploymentPolicy
Properties:
  ContinuousDeploymentPolicyConfig:
    ContinuousDeploymentPolicyConfig

```

## Properties

`ContinuousDeploymentPolicyConfig`

Contains the configuration for a continuous deployment policy.

_Required_: Yes

_Type_: [ContinuousDeploymentPolicyConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-continuousdeploymentpolicy-continuousdeploymentpolicyconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier for the continuous deployment
policy.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The identifier of the cotinuous deployment policy.

`LastModifiedTime`

The date and time when the continuous deployment policy was last modified.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ContinuousDeploymentPolicyConfig
