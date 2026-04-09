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

Type: [SAPODataPaginationConfig](api-sapodatapaginationconfig.md) object

Required: No

**parallelismConfig**

Sets the number of concurrent processes that transfers OData records from your SAP
instance.

Type: [SAPODataParallelismConfig](api-sapodataparallelismconfig.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/sapodatasourceproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/sapodatasourceproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/sapodatasourceproperties.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SAPODataParallelismConfig

ScheduledTriggerProperties

All content copied from https://docs.aws.amazon.com/.
