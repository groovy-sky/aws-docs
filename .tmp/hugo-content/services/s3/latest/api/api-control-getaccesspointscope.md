# GetAccessPointScope

Returns the access point scope for a directory bucket.

To use this operation, you must have the permission to perform the
`s3express:GetAccessPointScope` action.

For information about REST API errors, see [REST error responses](errorresponses.md#RESTErrorResponses).

## Request Syntax

```nohighlight

GET /v20180820/accesspoint/name/scope HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetAccessPointScope_RequestSyntax)**

The name of the access point with the scope you want to retrieve.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_GetAccessPointScope_RequestSyntax)**

The AWS account ID that owns the access point with the scope that you want to retrieve.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetAccessPointScopeResult>
   <Scope>
      <Permissions>
         <Permission>string</Permission>
      </Permissions>
      <Prefixes>
         <Prefix>string</Prefix>
      </Prefixes>
   </Scope>
</GetAccessPointScopeResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetAccessPointScopeResult](#API_control_GetAccessPointScope_ResponseSyntax)**

Root level tag for the GetAccessPointScopeResult parameters.

Required: Yes

**[Scope](#API_control_GetAccessPointScope_ResponseSyntax)**

The contents of the access point scope.

Type: [Scope](api-control-scope.md) data type

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getaccesspointscope.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getaccesspointscope.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getaccesspointscope.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getaccesspointscope.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getaccesspointscope.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getaccesspointscope.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getaccesspointscope.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getaccesspointscope.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getaccesspointscope.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getaccesspointscope.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetAccessPointPolicyStatusForObjectLambda

GetBucket

All content copied from https://docs.aws.amazon.com/.
