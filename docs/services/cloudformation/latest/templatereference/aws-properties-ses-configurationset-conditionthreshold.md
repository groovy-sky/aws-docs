This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSet ConditionThreshold

Specifies the threshold conditions for the automatic validation settings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionThresholdEnabled" : String,
  "OverallConfidenceThreshold" : OverallConfidenceThreshold
}

```

### YAML

```yaml

  ConditionThresholdEnabled: String
  OverallConfidenceThreshold:
    OverallConfidenceThreshold

```

## Properties

`ConditionThresholdEnabled`

Enables or disables the automatic validation feature. When enabled, SES validates recipients to calculate a confidence score regarding their validity.

_Required_: Yes

_Type_: String

_Pattern_: `ENABLED|DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OverallConfidenceThreshold`

The validation settings that define the minimum score required for SES to allow an email to be sent.

_Required_: No

_Type_: [OverallConfidenceThreshold](aws-properties-ses-configurationset-overallconfidencethreshold.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchivingOptions

DashboardOptions

All content copied from https://docs.aws.amazon.com/.
