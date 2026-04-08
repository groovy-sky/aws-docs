# EnableHttpEndpoint

Enables the HTTP endpoint for the DB cluster. By default, the HTTP endpoint
isn't enabled.

When enabled, this endpoint provides a connectionless web service API (RDS Data API)
for running SQL queries on the Aurora DB cluster. You can also query your database from inside the RDS console
with the RDS query editor.

For more information, see [Using RDS Data API](../../../../services/amazonrds/latest/aurorauserguide/data-api.md) in the
_Amazon Aurora User Guide_.

###### Note

This operation applies only to Aurora Serverless v2 and provisioned DB clusters. To enable the HTTP endpoint for Aurora Serverless v1 DB clusters,
use the `EnableHttpEndpoint` parameter of the `ModifyDBCluster` operation.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ResourceArn**

The Amazon Resource Name (ARN) of the DB cluster.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**HttpEndpointEnabled**

Indicates whether the HTTP endpoint is enabled or disabled for the DB cluster.

Type: Boolean

**ResourceArn**

The ARN of the DB cluster.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidResourceStateFault**

The operation can't be performed because another operation is in progress.

HTTP Status Code: 400

**ResourceNotFoundFault**

The specified resource ID was not found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/enablehttpendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/enablehttpendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/enablehttpendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/enablehttpendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/enablehttpendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/enablehttpendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/enablehttpendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/enablehttpendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/enablehttpendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/enablehttpendpoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DownloadDBLogFilePortion

FailoverDBCluster

All content copied from https://docs.aws.amazon.com/.
