This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template TemplateSourceEntity

The source entity of the template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceAnalysis" : TemplateSourceAnalysis,
  "SourceTemplate" : TemplateSourceTemplate
}

```

### YAML

```yaml

  SourceAnalysis:
    TemplateSourceAnalysis
  SourceTemplate:
    TemplateSourceTemplate

```

## Properties

`SourceAnalysis`

The source analysis, if it is based on an analysis.

_Required_: No

_Type_: [TemplateSourceAnalysis](aws-properties-quicksight-template-templatesourceanalysis.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceTemplate`

The source template, if it is based on an template.

_Required_: No

_Type_: [TemplateSourceTemplate](aws-properties-quicksight-template-templatesourcetemplate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TemplateSourceAnalysis

TemplateSourceTemplate

All content copied from https://docs.aws.amazon.com/.
