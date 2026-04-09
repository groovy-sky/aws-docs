# GetAccessPointPolicy

Returns the access point policy associated with the specified access point.

The following actions are related to `GetAccessPointPolicy`:

- [PutAccessPointPolicy](api-control-putaccesspointpolicy.md)

- [DeleteAccessPointPolicy](api-control-deleteaccesspointpolicy.md)

## Request Syntax

```nohighlight

GET /v20180820/accesspoint/name/policy HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetAccessPointPolicy_RequestSyntax)**

The name of the access point whose policy you want to retrieve.

For using this parameter with Amazon S3 on Outposts with the REST API, you must specify the name and the x-amz-outpost-id as well.

For using this parameter with S3 on Outposts with the AWS SDK and CLI, you must specify the ARN of the access point accessed in the format `arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/accesspoint/<my-accesspoint-name>`. For example, to access the access point `reports-ap` through Outpost `my-outpost` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap`. The value must be URL encoded.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_GetAccessPointPolicy_RequestSyntax)**

The account ID for the account that owns the specified access point.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetAccessPointPolicyResult>
   <Policy>string</Policy>
</GetAccessPointPolicyResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetAccessPointPolicyResult](#API_control_GetAccessPointPolicy_ResponseSyntax)**

Root level tag for the GetAccessPointPolicyResult parameters.

Required: Yes

**[Policy](#API_control_GetAccessPointPolicy_ResponseSyntax)**

The access point policy associated with the specified access point.

Type: String

## Examples

### Sample request

The following request returns the access point of the specified Amazon S3 on Outposts.

```HTTP

           GET /v20180820/accesspoint/example-access-point/policy  HTTP/1.1
           Host: s3-outposts.<Region>.amazonaws.com
           Date: Wed, 28 Oct 2020 22:32:00 GMT
           Authorization: authorization string
           x-amz-account-id: 123456789012
           x-amz-outpost-id: op-123456

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getaccesspointpolicy.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getaccesspointpolicy.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getaccesspointpolicy.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getaccesspointpolicy.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getaccesspointpolicy.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getaccesspointpolicy.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getaccesspointpolicy.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getaccesspointpolicy.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getaccesspointpolicy.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getaccesspointpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetAccessPointForObjectLambda

GetAccessPointPolicyForObjectLambda

All content copied from https://docs.aws.amazon.com/.
