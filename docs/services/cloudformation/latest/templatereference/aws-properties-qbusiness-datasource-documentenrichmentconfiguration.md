This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource DocumentEnrichmentConfiguration

Provides the configuration information for altering document metadata and content
during the document ingestion process.

For more information, see [Custom document\
enrichment](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/custom-document-enrichment.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InlineConfigurations" : [ InlineDocumentEnrichmentConfiguration, ... ],
  "PostExtractionHookConfiguration" : HookConfiguration,
  "PreExtractionHookConfiguration" : HookConfiguration
}

```

### YAML

```yaml

  InlineConfigurations:
    - InlineDocumentEnrichmentConfiguration
  PostExtractionHookConfiguration:
    HookConfiguration
  PreExtractionHookConfiguration:
    HookConfiguration

```

## Properties

`InlineConfigurations`

Configuration information to alter document attributes or metadata fields and content
when ingesting documents into Amazon Q Business.

_Required_: No

_Type_: Array of [InlineDocumentEnrichmentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-datasource-inlinedocumentenrichmentconfiguration.html)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostExtractionHookConfiguration`

Configuration information for invoking a Lambda function in AWS Lambda on the structured documents with their metadata and text extracted.
You can use a Lambda function to apply advanced logic for creating,
modifying, or deleting document metadata and content. For more information, see [Using Lambda functions](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/cde-lambda-operations.html).

_Required_: No

_Type_: [HookConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-datasource-hookconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreExtractionHookConfiguration`

Configuration information for invoking a Lambda function in AWS Lambda on the original or raw documents before extracting their metadata and
text. You can use a Lambda function to apply advanced logic for creating,
modifying, or deleting document metadata and content. For more information, see [Using Lambda functions](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/cde-lambda-operations.html).

_Required_: No

_Type_: [HookConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-datasource-hookconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DocumentAttributeValue

HookConfiguration
