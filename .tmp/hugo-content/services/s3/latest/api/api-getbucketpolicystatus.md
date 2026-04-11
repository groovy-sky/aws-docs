# GetBucketPolicyStatus

###### Note

This operation is not supported for directory buckets.

Retrieves the policy status for an Amazon S3 bucket, indicating whether the bucket is public. In order to
use this operation, you must have the `s3:GetBucketPolicyStatus` permission. For more
information about Amazon S3 permissions, see [Specifying Permissions in a\
Policy](../dev/using-with-s3-actions.md).

For more information about when Amazon S3 considers a bucket public, see [The Meaning of "Public"](../dev/access-control-block-public-access.md#access-control-block-public-access-policy-status).

The following operations are related to `GetBucketPolicyStatus`:

- [Using Amazon S3 Block Public Access](../dev/access-control-block-public-access.md)

- [GetPublicAccessBlock](api-getpublicaccessblock.md)

- [PutPublicAccessBlock](api-putpublicaccessblock.md)

- [DeletePublicAccessBlock](api-deletepublicaccessblock.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?policyStatus HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketPolicyStatus_RequestSyntax)**

The name of the Amazon S3 bucket whose policy status you want to retrieve.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketPolicyStatus_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<PolicyStatus>
   <IsPublic>boolean</IsPublic>
</PolicyStatus>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[PolicyStatus](#API_GetBucketPolicyStatus_ResponseSyntax)**

Root level tag for the PolicyStatus parameters.

Required: Yes

**[IsPublic](#API_GetBucketPolicyStatus_ResponseSyntax)**

The policy status for this bucket. `TRUE` indicates that this bucket is public.
`FALSE` indicates that the bucket is not public.

Type: Boolean

## Examples

### Sample Request

The following request gets a bucket policy status.

```

            GET /<bucket-name>?policyStatus HTTP/1.1
            Host: <bucket-name>.s3.<Region>.amazonaws.com
            x-amz-date: <Thu, 15 Nov 2016 00:17:21 GMT>
            Authorization: <signatureValue>

```

### Sample Response

This example illustrates one usage of GetBucketPolicyStatus.

```

            HTTP/1.1 200 OK
            x-amz-id-2: ITnGT1y4REXAMPLEPi4hklTXouTf0hccUjo0iCPEXAMPLEutBj3M7fPGlWO2SEWp
            x-amz-request-id: 51991EXAMPLE5321
            Date: Thu, 15 Nov 2016 00:17:22 GMT
            Server: AmazonS3
            Content-Length: 0

            <PolicyStatus>
               <IsPublic>TRUE</IsPublic>
            </PolicyStatus>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/getbucketpolicystatus.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/getbucketpolicystatus.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/getbucketpolicystatus.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/getbucketpolicystatus.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/getbucketpolicystatus.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/getbucketpolicystatus.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/getbucketpolicystatus.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/getbucketpolicystatus.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/getbucketpolicystatus.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/getbucketpolicystatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketPolicy

GetBucketReplication

All content copied from https://docs.aws.amazon.com/.
