This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaStore::Container MetricPolicy

The metric policy that is associated with the container. A metric policy allows AWS Elemental MediaStore to send metrics to Amazon CloudWatch. In the policy, you must indicate whether you want MediaStore to send container-level metrics. You can also include rules to define groups of objects that you want MediaStore to send object-level metrics for.

To view examples of how to construct a metric policy for your use case, see [Example Metric Policies](https://docs.aws.amazon.com/mediastore/latest/ug/policies-metric-examples.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerLevelMetrics" : String,
  "MetricPolicyRules" : [ MetricPolicyRule, ... ]
}

```

### YAML

```yaml

  ContainerLevelMetrics: String
  MetricPolicyRules:
    - MetricPolicyRule

```

## Properties

`ContainerLevelMetrics`

A setting to enable or disable metrics at the container level.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricPolicyRules`

A parameter that holds an array of rules that enable metrics at the object level. This parameter is optional, but if you choose to include it, you must also include at least one rule. By default, you can include up to five rules. You can also [request a quota increase](https://console.aws.amazon.com/servicequotas/home?region=us-east-1) to allow up to 300 rules per policy.

_Required_: No

_Type_: Array of [MetricPolicyRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediastore-container-metricpolicyrule.html)

_Minimum_: `1`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CorsRule

MetricPolicyRule
