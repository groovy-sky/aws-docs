This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate Outputs

Describes the output destinations of the experiment report.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExperimentReportS3Configuration" : ExperimentReportS3Configuration
}

```

### YAML

```yaml

  ExperimentReportS3Configuration:
    ExperimentReportS3Configuration

```

## Properties

`ExperimentReportS3Configuration`

The S3 destination for the experiment report.

_Required_: Yes

_Type_: [ExperimentReportS3Configuration](aws-properties-fis-experimenttemplate-experimentreports3configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExperimentTemplateTargetFilter

S3Configuration

All content copied from https://docs.aws.amazon.com/.
