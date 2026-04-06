# GetBucketAcl

###### Note

This operation is not supported for directory buckets.

This implementation of the `GET` action uses the `acl` subresource to return
the access control list (ACL) of a bucket. To use `GET` to return the ACL of the bucket, you
must have the `READ_ACP` access to the bucket. If `READ_ACP` permission is granted
to the anonymous user, you can return the ACL of the bucket without using an authorization
header.

When you use this API operation with an access point, provide the alias of the access point in place of the bucket name.

When you use this API operation with an Object Lambda access point, provide the alias of the Object Lambda access point in place of the bucket name.
If the Object Lambda access point alias in a request is not valid, the error code `InvalidAccessPointAliasError` is returned.
For more information about `InvalidAccessPointAliasError`, see [List of\
Error Codes](https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList).

###### Note

If your bucket uses the bucket owner enforced setting for S3 Object Ownership, requests to read
ACLs are still supported and return the `bucket-owner-full-control` ACL with the owner
being the account that created the bucket. For more information, see [Controlling object ownership and\
disabling ACLs](../userguide/about-object-ownership.md) in the _Amazon S3 User Guide_.

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

The following operations are related to `GetBucketAcl`:

- [ListObjects](https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html)

## Request Syntax

```nohighlight

GET /?acl HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketAcl_RequestSyntax)**

Specifies the S3 bucket whose ACL is being requested.

When you use this API operation with an access point, provide the alias of the access point in place of the bucket name.

When you use this API operation with an Object Lambda access point, provide the alias of the Object Lambda access point in place of the bucket name.
If the Object Lambda access point alias in a request is not valid, the error code `InvalidAccessPointAliasError` is returned.
For more information about `InvalidAccessPointAliasError`, see [List of\
Error Codes](https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList).

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketAcl_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<AccessControlPolicy>
   <Owner>
      <DisplayName>string</DisplayName>
      <ID>string</ID>
   </Owner>
   <AccessControlList>
      <Grant>
         <Grantee>
            <DisplayName>string</DisplayName>
            <EmailAddress>string</EmailAddress>
            <ID>string</ID>
            <xsi:type>string</xsi:type>
            <URI>string</URI>
         </Grantee>
         <Permission>string</Permission>
      </Grant>
   </AccessControlList>
</AccessControlPolicy>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[AccessControlPolicy](#API_GetBucketAcl_ResponseSyntax)**

Root level tag for the AccessControlPolicy parameters.

Required: Yes

**[Grants](#API_GetBucketAcl_ResponseSyntax)**

A list of grants.

Type: Array of [Grant](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Grant.html) data types

**[Owner](#API_GetBucketAcl_ResponseSyntax)**

Container for the bucket owner's ID.

Type: [Owner](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Owner.html) data type

## Examples

### Sample Request

The following request returns the ACL of the specified bucket.

```HTTP

GET ?acl HTTP/1.1
Host: bucket.s3.<Region>.amazonaws.com
Date: Wed, 28 Oct 2009 22:32:00 GMT
Authorization: authorization string

```

### Sample Response

```xml

HTTP/1.1 200 OK
x-amz-id-2: eftixk72aD6Ap51TnqcoF8eFidJG9Z/2mkiDFu8yU9AS1ed4OpIszj7UDNEHGran
x-amz-request-id: 318BC8BC148832E5
Date: Wed, 28 Oct 2009 22:32:00 GMT
Last-Modified: Sun, 1 Jan 2006 12:00:00 GMT
Content-Length: 124
Content-Type: text/plain
Connection: close
Server: AmazonS3
<AccessControlPolicy>
  <Owner>
    <ID>75aa57f09aa0c8caeab4f8c24e99d10f8e7faeebf76c078efc7c6caea54ba06a</ID>
  </Owner>
  <AccessControlList>
    <Grant>
      <Grantee xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
			xsi:type="CanonicalUser">
        <ID>75aa57f09aa0c8caeab4f8c24e99d10f8e7faeebf76c078efc7c6caea54ba06a</ID>
      </Grantee>
      <Permission>FULL_CONTROL</Permission>
    </Grant>
  </AccessControlList>
</AccessControlPolicy>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/GetBucketAcl)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/GetBucketAcl)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/GetBucketAcl)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/GetBucketAcl)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/GetBucketAcl)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/GetBucketAcl)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/GetBucketAcl)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/GetBucketAcl)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/GetBucketAcl)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/GetBucketAcl)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetBucketAccelerateConfiguration

GetBucketAnalyticsConfiguration
