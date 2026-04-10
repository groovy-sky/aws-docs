This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile MachineLearningDetectionConfig

The `MachineLearningDetectionConfig` property type controls confidence of the
machine learning model.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfidenceLevel" : String
}

```

### YAML

```yaml

  ConfidenceLevel: String

```

## Properties

`ConfidenceLevel`

The model confidence level.

There are three levels of confidence, `"high"`, `"medium"`, and
`"low"`.

The higher the confidence level, the lower the sensitivity, and the lower the alarm
frequency will be.

_Required_: No

_Type_: String

_Allowed values_: `LOW | MEDIUM | HIGH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BehaviorCriteria

MetricDimension

All content copied from https://docs.aws.amazon.com/.
