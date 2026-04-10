This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::AnomalyDetector

Anomaly detection uses the Random Cut Forest algorithm for time-series analysis. The
anomaly detector analyzes Amazon Managed Service for Prometheus metrics to identify unusual patterns
and behaviors.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::APS::AnomalyDetector",
  "Properties" : {
      "Alias" : String,
      "Configuration" : AnomalyDetectorConfiguration,
      "EvaluationIntervalInSeconds" : Integer,
      "Labels" : [ Label, ... ],
      "MissingDataAction" : MissingDataAction,
      "Tags" : [ Tag, ... ],
      "Workspace" : String
    }
}

```

### YAML

```yaml

Type: AWS::APS::AnomalyDetector
Properties:
  Alias: String
  Configuration:
    AnomalyDetectorConfiguration
  EvaluationIntervalInSeconds: Integer
  Labels:
    - Label
  MissingDataAction:
    MissingDataAction
  Tags:
    - Tag
  Workspace: String

```

## Properties

`Alias`

The user-friendly name of the anomaly detector.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configuration`

The algorithm configuration of the anomaly detector.

_Required_: Yes

_Type_: [AnomalyDetectorConfiguration](aws-properties-aps-anomalydetector-anomalydetectorconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationIntervalInSeconds`

The frequency, in seconds, at which the anomaly detector evaluates metrics.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Labels`

The Amazon Managed Service for Prometheus metric labels associated with the anomaly detector.

_Required_: No

_Type_: Array of [Label](aws-properties-aps-anomalydetector-label.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MissingDataAction`

The action taken when data is missing during evaluation.

_Required_: No

_Type_: [MissingDataAction](aws-properties-aps-anomalydetector-missingdataaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags applied to the anomaly detector.

_Required_: No

_Type_: Array of [Tag](aws-properties-aps-anomalydetector-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Workspace`

An Amazon Managed Service for Prometheus workspace is a logical and isolated Prometheus server
dedicated to ingesting, storing, and querying your Prometheus-compatible metrics.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn):aps:[a-z0-9-]+:[0-9]+:workspace/[a-zA-Z0-9-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "Id" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the anomaly detector.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Managed Service for Prometheus

AnomalyDetectorConfiguration

All content copied from https://docs.aws.amazon.com/.
