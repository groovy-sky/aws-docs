# GetDatabase

Returns a database object for the specified database and data catalog.

## Request Syntax

```nohighlight

{
   "CatalogName": "string",
   "DatabaseName": "string",
   "WorkGroup": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CatalogName](#API_GetDatabase_RequestSyntax)**

The name of the data catalog that contains the database to return.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Required: Yes

**[DatabaseName](#API_GetDatabase_RequestSyntax)**

The name of the database to return.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**[WorkGroup](#API_GetDatabase_RequestSyntax)**

The name of the workgroup for which the metadata is being fetched. Required if
requesting an IAM Identity Center enabled AWS Glue Data Catalog.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: No

## Response Syntax

```nohighlight

{
   "Database": {
      "Description": "string",
      "Name": "string",
      "Parameters": {
         "string" : "string"
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Database](#API_GetDatabase_ResponseSyntax)**

The database returned.

Type: [Database](api-database.md) object

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

**MetadataException**

An exception that Athena received when it called a custom metastore.
Occurs if the error is not caused by user input ( `InvalidRequestException`)
or from the Athena platform ( `InternalServerException`). For
example, if a user-created Lambda function is missing permissions, the
Lambda
`4XX` exception is returned in a `MetadataException`.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/getdatabase.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/getdatabase.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/getdatabase.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/getdatabase.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/getdatabase.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/getdatabase.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/getdatabase.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/getdatabase.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/getdatabase.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/getdatabase.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetCapacityReservation

GetDataCatalog

All content copied from https://docs.aws.amazon.com/.
