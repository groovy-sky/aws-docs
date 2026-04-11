This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate AnalysisTemplateArtifacts

The analysis template artifacts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalArtifacts" : [ AnalysisTemplateArtifact, ... ],
  "EntryPoint" : AnalysisTemplateArtifact,
  "RoleArn" : String
}

```

### YAML

```yaml

  AdditionalArtifacts:
    - AnalysisTemplateArtifact
  EntryPoint:
    AnalysisTemplateArtifact
  RoleArn: String

```

## Properties

`AdditionalArtifacts`

Additional artifacts for the analysis template.

_Required_: No

_Type_: Array of [AnalysisTemplateArtifact](aws-properties-cleanrooms-analysistemplate-analysistemplateartifact.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntryPoint`

The entry point for the analysis template artifacts.

_Required_: Yes

_Type_: [AnalysisTemplateArtifact](aws-properties-cleanrooms-analysistemplate-analysistemplateartifact.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The role ARN for the analysis template artifacts.

_Required_: Yes

_Type_: String

_Minimum_: `32`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisTemplateArtifactMetadata

ColumnClassificationDetails

All content copied from https://docs.aws.amazon.com/.
