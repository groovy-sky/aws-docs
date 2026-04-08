# DeleteDBProxyEndpoint

Deletes a `DBProxyEndpoint`. Doing so removes the ability to access the DB proxy using the
endpoint that you defined. The endpoint that you delete might have provided capabilities such as read/write
or read-only operations, or using a different VPC than the DB proxy's default VPC.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBProxyEndpointName**

The name of the DB proxy endpoint to delete.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 63.

Pattern: `[a-zA-Z](?:-?[a-zA-Z0-9]+)*`

Required: Yes

## Response Elements

The following element is returned by the service.

**DBProxyEndpoint**

The data structure representing the details of the DB proxy endpoint that you delete.

Type: [DBProxyEndpoint](api-dbproxyendpoint.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBProxyEndpointNotFoundFault**

The DB proxy endpoint doesn't exist.

HTTP Status Code: 404

**InvalidDBProxyEndpointStateFault**

You can't perform this operation while the DB proxy endpoint is in a particular state.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/deletedbproxyendpoint.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/deletedbproxyendpoint.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/deletedbproxyendpoint.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/deletedbproxyendpoint.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/deletedbproxyendpoint.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/deletedbproxyendpoint.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/deletedbproxyendpoint.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/deletedbproxyendpoint.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/deletedbproxyendpoint.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/deletedbproxyendpoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteDBProxy

DeleteDBSecurityGroup

All content copied from https://docs.aws.amazon.com/.
