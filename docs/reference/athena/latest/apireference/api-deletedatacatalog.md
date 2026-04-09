# DeleteDataCatalog

Deletes a data catalog.

## Request Syntax

```nohighlight

{
   "DeleteCatalogOnly": boolean,
   "Name": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DeleteCatalogOnly](#API_DeleteDataCatalog_RequestSyntax)**

Deletes the Athena Data Catalog. You can only use this with the `FEDERATED`
catalogs. You usually perform this before registering the connector with Glue Data
Catalog. After deletion, you will have to manage the Glue Connection and Lambda
function.

Type: Boolean

Required: No

**[Name](#API_DeleteDataCatalog_RequestSyntax)**

The name of the data catalog to delete.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Required: Yes

## Response Syntax

```nohighlight

{
   "DataCatalog": {
      "ConnectionType": "string",
      "Description": "string",
      "Error": "string",
      "Name": "string",
      "Parameters": {
         "string" : "string"
      },
      "Status": "string",
      "Type": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[DataCatalog](#API_DeleteDataCatalog_ResponseSyntax)**

Contains information about a data catalog in an AWS account.

###### Note

In the Athena console, data catalogs are listed as "data sources" on
the **Data sources** page under the **Data source name** column.

Type: [DataCatalog](api-datacatalog.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

Indicates a platform issue, which may be due to a transient condition or
outage.

HTTP Status Code: 500

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
required parameter may be missing or out of range.

**AthenaErrorCode**

The error code returned when the query execution failed to process, or when the
processing request for the named query failed.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/deletedatacatalog.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/deletedatacatalog.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/deletedatacatalog.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/deletedatacatalog.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/deletedatacatalog.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/deletedatacatalog.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/deletedatacatalog.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/deletedatacatalog.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/deletedatacatalog.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/deletedatacatalog.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteCapacityReservation

DeleteNamedQuery

All content copied from https://docs.aws.amazon.com/.
