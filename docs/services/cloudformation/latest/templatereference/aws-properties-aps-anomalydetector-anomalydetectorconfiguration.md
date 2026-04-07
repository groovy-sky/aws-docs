This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::AnomalyDetector AnomalyDetectorConfiguration

The configuration for the anomaly detection algorithm.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RandomCutForest" : RandomCutForestConfiguration
}

```

### YAML

```yaml

  RandomCutForest:
    RandomCutForestConfiguration

```

## Properties

`RandomCutForest`

The Random Cut Forest algorithm configuration for anomaly detection.

_Required_: Yes

_Type_: [RandomCutForestConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-aps-anomalydetector-randomcutforestconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::APS::AnomalyDetector

IgnoreNearExpected
