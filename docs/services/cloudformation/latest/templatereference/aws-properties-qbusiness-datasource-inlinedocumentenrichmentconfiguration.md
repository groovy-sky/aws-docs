This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource InlineDocumentEnrichmentConfiguration

Provides the configuration information for applying basic logic to alter document
metadata and content when ingesting documents into Amazon Q Business.

To apply advanced logic, to go beyond what you can do with basic logic, see [`HookConfiguration`](../../../amazonq/latest/api-reference/api-hookconfiguration.md).

For more information, see [Custom document enrichment](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/custom-document-enrichment.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Condition" : DocumentAttributeCondition,
  "DocumentContentOperator" : String,
  "Target" : DocumentAttributeTarget
}

```

### YAML

```yaml

  Condition:
    DocumentAttributeCondition
  DocumentContentOperator: String
  Target:
    DocumentAttributeTarget

```

## Properties

`Condition`

Configuration of the condition used for the target document attribute or metadata
field when ingesting documents into Amazon Q Business.

_Required_: No

_Type_: [DocumentAttributeCondition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-datasource-documentattributecondition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentContentOperator`

`TRUE` to delete content if the condition used for the target attribute is
met.

_Required_: No

_Type_: String

_Allowed values_: `DELETE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

Configuration of the target document attribute or metadata field when ingesting
documents into Amazon Q Business. You can also include a value.

_Required_: No

_Type_: [DocumentAttributeTarget](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-datasource-documentattributetarget.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ImageExtractionConfiguration

MediaExtractionConfiguration
