This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::ComputationModel AnomalyDetectionComputationModelConfiguration

Contains the configuration for anomaly detection computation models.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputProperties" : String,
  "ResultProperty" : String
}

```

### YAML

```yaml

  InputProperties: String
  ResultProperty: String

```

## Properties

`InputProperties`

The list of input properties for the anomaly detection model.

_Required_: Yes

_Type_: String

_Pattern_: `^\$\{[a-z][a-z0-9_]*\}$`

_Minimum_: `4`

_Maximum_: `67`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResultProperty`

The property where the anomaly detection results will be stored.

_Required_: Yes

_Type_: String

_Pattern_: `^\$\{[a-z][a-z0-9_]*\}$`

_Minimum_: `4`

_Maximum_: `67`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTSiteWise::ComputationModel

AssetModelPropertyBindingValue

All content copied from https://docs.aws.amazon.com/.
