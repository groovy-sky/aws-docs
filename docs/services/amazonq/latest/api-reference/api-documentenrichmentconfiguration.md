# DocumentEnrichmentConfiguration

Provides the configuration information for altering document metadata and content
during the document ingestion process.

For more information, see [Custom document\
enrichment](../business-use-dg/custom-document-enrichment.md).

## Contents

**inlineConfigurations**

Configuration information to alter document attributes or metadata fields and content
when ingesting documents into Amazon Q Business.

Type: Array of [InlineDocumentEnrichmentConfiguration](api-inlinedocumentenrichmentconfiguration.md) objects

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Required: No

**postExtractionHookConfiguration**

Provides the configuration information for invoking a Lambda function in
AWS Lambda to alter document metadata and content when ingesting
documents into Amazon Q Business.

You can configure your Lambda function using the
`PreExtractionHookConfiguration` parameter if you want to apply advanced
alterations on the original or raw documents.

If you want to apply advanced alterations on the Amazon Q Business structured documents,
you must configure your Lambda function using
`PostExtractionHookConfiguration`.

You can only invoke one Lambda function. However, this function can invoke
other functions it requires.

For more information, see [Custom document enrichment](../business-use-dg/custom-document-enrichment.md).

Type: [HookConfiguration](api-hookconfiguration.md) object

Required: No

**preExtractionHookConfiguration**

Provides the configuration information for invoking a Lambda function in
AWS Lambda to alter document metadata and content when ingesting
documents into Amazon Q Business.

You can configure your Lambda function using the
`PreExtractionHookConfiguration` parameter if you want to apply advanced
alterations on the original or raw documents.

If you want to apply advanced alterations on the Amazon Q Business structured documents,
you must configure your Lambda function using
`PostExtractionHookConfiguration`.

You can only invoke one Lambda function. However, this function can invoke
other functions it requires.

For more information, see [Custom document enrichment](../business-use-dg/custom-document-enrichment.md).

Type: [HookConfiguration](api-hookconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/documentenrichmentconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/documentenrichmentconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/documentenrichmentconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentDetails

EligibleDataSource

All content copied from https://docs.aws.amazon.com/.
