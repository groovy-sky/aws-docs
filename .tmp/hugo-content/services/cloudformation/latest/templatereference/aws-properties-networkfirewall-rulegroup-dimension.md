This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup Dimension

The value to use in an Amazon CloudWatch custom metric dimension. This is used in the
`PublishMetrics` custom action. A CloudWatch custom metric dimension is a name/value pair that's
part of the identity of a metric.

AWS Network Firewall sets the dimension name to `CustomAction` and you provide the
dimension value.

For more information about CloudWatch custom metric dimensions, see
[Publishing Custom Metrics](../../../amazoncloudwatch/latest/monitoring/publishingmetrics.md#usingDimensions) in the [Amazon CloudWatch User\
Guide](../../../amazoncloudwatch/latest/monitoring/whatiscloudwatch.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Value" : String
}

```

### YAML

```yaml

  Value: String

```

## Properties

`Value`

The value to use in the custom metric dimension.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_ ]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomAction

Header

All content copied from https://docs.aws.amazon.com/.
