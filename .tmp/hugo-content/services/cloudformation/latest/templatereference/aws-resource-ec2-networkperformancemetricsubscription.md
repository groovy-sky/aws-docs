This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkPerformanceMetricSubscription

Describes Infrastructure Performance subscriptions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkPerformanceMetricSubscription",
  "Properties" : {
      "Destination" : String,
      "Metric" : String,
      "Source" : String,
      "Statistic" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkPerformanceMetricSubscription
Properties:
  Destination: String
  Metric: String
  Source: String
  Statistic: String

```

## Properties

`Destination`

The Region or Availability Zone that's the target for the subscription. For example, `eu-west-1`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Metric`

The metric used for the subscription.

_Required_: Yes

_Type_: String

_Allowed values_: `aggregate-latency`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Source`

The Region or Availability Zone that's the source for the subscription. For example, `us-east-1`.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Statistic`

The statistic used for the subscription.

_Required_: Yes

_Type_: String

_Allowed values_: `p50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `primaryIdentifier` property, which consists of
the following properties: `source`, `destination`, `metric`,
and `statistic`, with each value separated by a pipe (\|). For example,
`{ "Ref": "us-east-1|us-east-2|aggregate-latency|p50" }`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::NetworkInterfacePermission

AWS::EC2::PlacementGroup

All content copied from https://docs.aws.amazon.com/.
