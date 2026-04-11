This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Solution HpoObjective

The metric to optimize during hyperparameter optimization (HPO).

###### Note

Amazon Personalize doesn't support configuring the `hpoObjective`
at this time.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MetricName" : String,
  "MetricRegex" : String,
  "Type" : String
}

```

### YAML

```yaml

  MetricName: String
  MetricRegex: String
  Type: String

```

## Properties

`MetricName`

The name of the metric.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MetricRegex`

A regular expression for finding the metric in the training job logs.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of the metric. Valid values are `Maximize` and `Minimize`.

_Required_: No

_Type_: String

_Allowed values_: `Maximize | Minimize`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HpoConfig

HpoResourceConfig

All content copied from https://docs.aws.amazon.com/.
