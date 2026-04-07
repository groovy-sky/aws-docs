This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate ExperimentTemplateExperimentReportConfiguration

Describes the report configuration for the experiment template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSources" : DataSources,
  "Outputs" : Outputs,
  "PostExperimentDuration" : String,
  "PreExperimentDuration" : String
}

```

### YAML

```yaml

  DataSources:
    DataSources
  Outputs:
    Outputs
  PostExperimentDuration: String
  PreExperimentDuration: String

```

## Properties

`DataSources`

The data sources for the experiment report.

_Required_: No

_Type_: [DataSources](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fis-experimenttemplate-datasources.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Outputs`

The output destinations of the experiment report.

_Required_: Yes

_Type_: [Outputs](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fis-experimenttemplate-outputs.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostExperimentDuration`

The duration after the experiment end time for the data sources to include in the report.

_Required_: No

_Type_: String

_Pattern_: `[\S]+`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreExperimentDuration`

The duration before the experiment start time for the data sources to include in the report.

_Required_: No

_Type_: String

_Pattern_: `[\S]+`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ExperimentTemplateExperimentOptions

ExperimentTemplateLogConfiguration
