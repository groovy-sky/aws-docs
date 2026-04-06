# SAPODataSourceProperties

The properties that are applied when using SAPOData as a flow source.

## Contents

**objectPath**

The object path specified in the SAPOData flow source.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**paginationConfig**

Sets the page size for each concurrent process that transfers OData records from your SAP
instance.

Type: [SAPODataPaginationConfig](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_SAPODataPaginationConfig.html) object

Required: No

**parallelismConfig**

Sets the number of concurrent processes that transfers OData records from your SAP
instance.

Type: [SAPODataParallelismConfig](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_SAPODataParallelismConfig.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/SAPODataSourceProperties)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/SAPODataSourceProperties)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/SAPODataSourceProperties)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SAPODataParallelismConfig

ScheduledTriggerProperties
