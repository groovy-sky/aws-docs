# DeleteAccessPoint

Deletes the specified access point.

All Amazon S3 on Outposts REST API requests for this action require an additional parameter of `x-amz-outpost-id` to be passed with the request. In addition, you must use an S3 on Outposts endpoint hostname prefix instead of `s3-control`. For an example of the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the `x-amz-outpost-id` derived by using the access point ARN, see the [Examples](api-control-deleteaccesspoint.md#API_control_DeleteAccessPoint_Examples) section.

The following actions are related to `DeleteAccessPoint`:

- [CreateAccessPoint](api-control-createaccesspoint.md)

- [GetAccessPoint](api-control-getaccesspoint.md)

- [ListAccessPoints](api-control-listaccesspoints.md)

## Request Syntax

```nohighlight

DELETE /v20180820/accesspoint/name HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_DeleteAccessPoint_RequestSyntax)**

The name of the access point you want to delete.

For using this parameter with Amazon S3 on Outposts with the REST API, you must specify the name and the x-amz-outpost-id as well.

For using this parameter with S3 on Outposts with the AWS SDK and CLI, you must specify the ARN of the access point accessed in the format `arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/accesspoint/<my-accesspoint-name>`. For example, to access the access point `reports-ap` through Outpost `my-outpost` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap`. The value must be URL encoded.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_DeleteAccessPoint_RequestSyntax)**

The AWS account ID for the account that owns the specified access point.

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

## Examples

### DeleteAccessPoint syntax for Amazon S3 on Outposts

The following request deletes the access point of the specified Outpost.

```HTTP

           DELETE  /v20180820/accesspoint/example-access-point  HTTP/1.1
           Host: s3-outposts.<Region>.amazonaws.com
           Date: Wed, 28 Oct 2020 22:32:00 GMT
           x-amz-account-id: example-account-id
           x-amz-outpost-id: op-01ac5d28a6a232904
           Authorization: authorization string

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/deleteaccesspoint.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/deleteaccesspoint.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/deleteaccesspoint.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/deleteaccesspoint.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/deleteaccesspoint.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/deleteaccesspoint.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/deleteaccesspoint.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/deleteaccesspoint.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/deleteaccesspoint.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/deleteaccesspoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteAccessGrantsLocation

DeleteAccessPointForObjectLambda

All content copied from https://docs.aws.amazon.com/.
