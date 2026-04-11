# DeleteAccessPointScope

Deletes an existing access point scope for a directory bucket.

###### Note

When you delete the scope of an access point, all prefixes and permissions are deleted.

To use this operation, you must have the permission to perform the
`s3express:DeleteAccessPointScope` action.

For information about REST API errors, see [REST error responses](errorresponses.md#RESTErrorResponses).

## Request Syntax

```nohighlight

DELETE /v20180820/accesspoint/name/scope HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_DeleteAccessPointScope_RequestSyntax)**

The name of the access point with the scope that you want to delete.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_DeleteAccessPointScope_RequestSyntax)**

The AWS account ID that owns the access point with the scope that you want to delete.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/deleteaccesspointscope.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/deleteaccesspointscope.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/deleteaccesspointscope.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/deleteaccesspointscope.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/deleteaccesspointscope.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/deleteaccesspointscope.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/deleteaccesspointscope.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/deleteaccesspointscope.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/deleteaccesspointscope.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/deleteaccesspointscope.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteAccessPointPolicyForObjectLambda

DeleteBucket

All content copied from https://docs.aws.amazon.com/.
