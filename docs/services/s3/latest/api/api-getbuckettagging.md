# GetBucketTagging

###### Note

This operation is not supported for directory buckets.

Returns the tag set associated with the general purpose bucket.

To use this operation, you must have permission to perform the `s3:GetBucketTagging`
action. By default, the bucket owner has this permission and can grant this permission to others.

`GetBucketTagging` has the following special error:

- Error code: `NoSuchTagSet`

- Description: There is no tag set associated with the bucket.

The following operations are related to `GetBucketTagging`:

- [PutBucketTagging](https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html)

- [DeleteBucketTagging](api-deletebuckettagging.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?tagging HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketTagging_RequestSyntax)**

The name of the bucket for which to get the tagging information.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketTagging_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<Tagging>
   <TagSet>
      <Tag>
         <Key>string</Key>
         <Value>string</Value>
      </Tag>
   </TagSet>
</Tagging>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[Tagging](#API_GetBucketTagging_ResponseSyntax)**

Root level tag for the Tagging parameters.

Required: Yes

**[TagSet](#API_GetBucketTagging_ResponseSyntax)**

Contains the tag set.

Type: Array of [Tag](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Tag.html) data types

## Examples

### Sample Request

The following request returns the tag set of the specified bucket.

```

            GET ?tagging HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
            Date: Wed, 28 Oct 2009 22:32:00 GMT
            Authorization: authorization string

```

### Sample Response

Delete the metric configuration with a specified ID, which disables the CloudWatch metrics with
the `ExampleMetrics` value for the `FilterId` dimension.

```

         HTTP/1.1 200 OK
         Date: Wed, 25 Nov 2009 12:00:00 GMT
         Connection: close
         Server: AmazonS3

         <Tagging>
           <TagSet>
              <Tag>
                <Key>Project</Key>
               <Value>Project One</Value>
              </Tag>
              <Tag>
                <Key>User</Key>
                <Value>jsmith</Value>
              </Tag>
           </TagSet>
         </Tagging>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/GetBucketTagging)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/GetBucketTagging)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/GetBucketTagging)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/GetBucketTagging)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/GetBucketTagging)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/GetBucketTagging)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/GetBucketTagging)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/GetBucketTagging)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/GetBucketTagging)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/GetBucketTagging)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetBucketRequestPayment

GetBucketVersioning
