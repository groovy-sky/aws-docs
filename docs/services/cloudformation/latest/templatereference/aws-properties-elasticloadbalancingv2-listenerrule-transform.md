This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule Transform

The `Transform` property type specifies Property description not available. for an [AWS::ElasticLoadBalancingV2::ListenerRule](aws-resource-elasticloadbalancingv2-listenerrule.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HostHeaderRewriteConfig" : RewriteConfigObject,
  "Type" : String,
  "UrlRewriteConfig" : RewriteConfigObject
}

```

### YAML

```yaml

  HostHeaderRewriteConfig:
    RewriteConfigObject
  Type: String
  UrlRewriteConfig:
    RewriteConfigObject

```

## Properties

`HostHeaderRewriteConfig`

Property description not available.

_Required_: No

_Type_: [RewriteConfigObject](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfigobject.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UrlRewriteConfig`

Property description not available.

_Required_: No

_Type_: [RewriteConfigObject](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticloadbalancingv2-listenerrule-rewriteconfigobject.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetGroupTuple

AWS::ElasticLoadBalancingV2::LoadBalancer
