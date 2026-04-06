# GetBucketOwnershipControls

###### Note

This operation is not supported for directory buckets.

Retrieves `OwnershipControls` for an Amazon S3 bucket. To use this operation, you must have
the `s3:GetBucketOwnershipControls` permission. For more information about Amazon S3 permissions,
see [Specifying\
permissions in a policy](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html).

###### Note

A bucket doesn't have `OwnershipControls` settings in the following cases:

- The bucket was created before the `BucketOwnerEnforced` ownership setting was
introduced and you've never explicitly applied this value

- You've manually deleted the bucket ownership control value using the
`DeleteBucketOwnershipControls` API operation.

By default, Amazon S3 sets `OwnershipControls` for all newly created buckets.

For information about Amazon S3 Object Ownership, see [Using Object Ownership](../userguide/about-object-ownership.md).

The following operations are related to `GetBucketOwnershipControls`:

- [PutBucketOwnershipControls](api-putbucketownershipcontrols.md)

- [DeleteBucketOwnershipControls](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketOwnershipControls.html)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?ownershipControls HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketOwnershipControls_RequestSyntax)**

The name of the Amazon S3 bucket whose `OwnershipControls` you want to retrieve.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketOwnershipControls_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<OwnershipControls>
   <Rule>
      <ObjectOwnership>string</ObjectOwnership>
   </Rule>
   ...
</OwnershipControls>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[OwnershipControls](#API_GetBucketOwnershipControls_ResponseSyntax)**

Root level tag for the OwnershipControls parameters.

Required: Yes

**[Rule](#API_GetBucketOwnershipControls_ResponseSyntax)**

The container element for an ownership control rule.

Type: Array of [OwnershipControlsRule](https://docs.aws.amazon.com/AmazonS3/latest/API/API_OwnershipControlsRule.html) data types

## Examples

### Sample GetBucketOwnershipControls Request for BucketOwnerEnforced

This example illustrates one usage of GetBucketOwnershipControls.

```

          GET /amzn-s3-demo-bucket?/ownershipControls HTTP/1.1
          Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
          Date: Mon, 29 Nov 2021 00:17:22 GMT
          Authorization: signatureValue;

```

### Sample GetBucketOwnershipControls Response

This example illustrates one usage of GetBucketOwnershipControls.

```

          HTTP/1.1 200 OK
          x-amz-id-2: Adphn7MaAHDEg9mh5JmcTN8mzyVX0JhIztSiQNaqTxnXXcYi4uiZbYdwWC3JXmh/XXVUUQwO4Vs=
          x-amz-request-id: 252631E05F84A415
          Date: Mon, 29 Nov 2021 00:17:22 GMT
          Server: AmazonS3
          Content-Length: 194

          <OwnershipControls xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
            <Rule>
              <ObjectOwnership>BucketOwnerEnforced</ObjectOwnership>
            </Rule>
          </OwnershipControls>

```

### Sample GetBucketOwnershipControls Request for BucketOwnerPreferred

This example illustrates one usage of GetBucketOwnershipControls.

```

          GET /amzn-s3-demo-bucket?/ownershipControls HTTP/1.1
          Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
          Date: Thu, 18 Jun 2017 00:17:22 GMT
          Authorization: signatureValue;

```

### Sample GetBucketOwnershipControls Response

This example illustrates one usage of GetBucketOwnershipControls.

```

          HTTP/1.1 200 OK
          x-amz-id-2: Adphn7MaAHDEg9mh5JmcTN8mzyVX0JhIztSiQNaqTxnXXcYi4uiZbYdwWC3JXmh/XXVUUQwO4Vs=
          x-amz-request-id: 252631E05F84A415
          Date: Thu, 18 Jun 2020 00:17:22 GMT
          Server: AmazonS3
          Content-Length: 194

          <OwnershipControls xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
            <Rule>
              <ObjectOwnership>BucketOwnerPreferred</ObjectOwnership>
            </Rule>
          </OwnershipControls>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/GetBucketOwnershipControls)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/GetBucketOwnershipControls)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/GetBucketOwnershipControls)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/GetBucketOwnershipControls)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/GetBucketOwnershipControls)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/GetBucketOwnershipControls)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/GetBucketOwnershipControls)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/GetBucketOwnershipControls)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/GetBucketOwnershipControls)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/GetBucketOwnershipControls)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetBucketNotificationConfiguration

GetBucketPolicy
