This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate AnalysisSource

The structure that defines the body of the analysis template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Artifacts" : AnalysisTemplateArtifacts,
  "Text" : String
}

```

### YAML

```yaml

  Artifacts:
    AnalysisTemplateArtifacts
  Text: String

```

## Properties

`Artifacts`

The artifacts of the analysis source.

_Required_: No

_Type_: [AnalysisTemplateArtifacts](aws-properties-cleanrooms-analysistemplate-analysistemplateartifacts.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Text`

The query text.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `90000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisSchema

AnalysisSourceMetadata

All content copied from https://docs.aws.amazon.com/.
