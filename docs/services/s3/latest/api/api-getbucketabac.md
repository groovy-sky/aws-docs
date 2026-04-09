# GetBucketAbac

Returns the attribute-based access control (ABAC) property of the general purpose bucket. If ABAC is enabled on your bucket, you can use tags on the bucket for access control. For more information, see [Enabling ABAC in general purpose buckets](../userguide/buckets-tagging-enable-abac.md).

## Request Syntax

```nohighlight

GET /?abac HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketAbac_RequestSyntax)**

The name of the general purpose bucket.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketAbac_RequestSyntax)**

The AWS account ID of the general purpose bucket's owner.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<AbacStatus>
   <Status>string</Status>
</AbacStatus>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[AbacStatus](#API_GetBucketAbac_ResponseSyntax)**

Root level tag for the AbacStatus parameters.

Required: Yes

**[Status](#API_GetBucketAbac_ResponseSyntax)**

The ABAC status of the general purpose bucket.

Type: String

Valid Values: `Enabled | Disabled`

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/getbucketabac.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/getbucketabac.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/getbucketabac.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/getbucketabac.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/getbucketabac.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/getbucketabac.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/getbucketabac.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/getbucketabac.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/getbucketabac.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/getbucketabac.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeletePublicAccessBlock

GetBucketAccelerateConfiguration

All content copied from https://docs.aws.amazon.com/.
