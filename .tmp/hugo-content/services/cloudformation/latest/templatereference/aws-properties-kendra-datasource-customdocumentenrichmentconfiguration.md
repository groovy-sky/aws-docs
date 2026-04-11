This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource CustomDocumentEnrichmentConfiguration

Provides the configuration information for altering document metadata and content
during the document ingestion process.

For more information, see [Customizing document metadata\
during the ingestion process](../../../kendra/latest/dg/custom-document-enrichment.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InlineConfigurations" : [ InlineCustomDocumentEnrichmentConfiguration, ... ],
  "PostExtractionHookConfiguration" : HookConfiguration,
  "PreExtractionHookConfiguration" : HookConfiguration,
  "RoleArn" : String
}

```

### YAML

```yaml

  InlineConfigurations:
    - InlineCustomDocumentEnrichmentConfiguration
  PostExtractionHookConfiguration:
    HookConfiguration
  PreExtractionHookConfiguration:
    HookConfiguration
  RoleArn: String

```

## Properties

`InlineConfigurations`

Configuration information to alter document attributes or metadata fields and content
when ingesting documents into Amazon Kendra.

_Required_: No

_Type_: Array of [InlineCustomDocumentEnrichmentConfiguration](aws-properties-kendra-datasource-inlinecustomdocumentenrichmentconfiguration.md)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostExtractionHookConfiguration`

Configuration information for invoking a Lambda function in AWS Lambda on
the structured documents with their metadata and text extracted. You can use a Lambda
function to apply advanced logic for creating, modifying, or deleting document metadata
and content. For more information, see [Advanced data manipulation](../../../kendra/latest/dg/custom-document-enrichment.md#advanced-data-manipulation).

_Required_: No

_Type_: [HookConfiguration](aws-properties-kendra-datasource-hookconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreExtractionHookConfiguration`

Configuration information for invoking a Lambda function in AWS Lambda on
the original or raw documents before extracting their metadata and text. You can use a
Lambda function to apply advanced logic for creating, modifying, or deleting document
metadata and content. For more information, see [Advanced data manipulation](../../../kendra/latest/dg/custom-document-enrichment.md#advanced-data-manipulation).

_Required_: No

_Type_: [HookConfiguration](aws-properties-kendra-datasource-hookconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of an IAM role with permission to run
`PreExtractionHookConfiguration` and
`PostExtractionHookConfiguration` for altering document metadata and
content during the document ingestion process. For more information, see [an IAM roles for Amazon Kendra](../../../kendra/latest/dg/iam-roles.md).

_Required_: No

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectionConfiguration

DatabaseConfiguration

All content copied from https://docs.aws.amazon.com/.
