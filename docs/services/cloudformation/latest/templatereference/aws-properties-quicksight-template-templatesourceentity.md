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

_Type_: [TemplateSourceAnalysis](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-template-templatesourceanalysis.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceTemplate`

The source template, if it is based on an template.

_Required_: No

_Type_: [TemplateSourceTemplate](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-template-templatesourcetemplate.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TemplateSourceAnalysis

TemplateSourceTemplate
