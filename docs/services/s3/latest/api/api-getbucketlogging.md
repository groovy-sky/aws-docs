# GetBucketLogging

###### Note

This operation is not supported for directory buckets.

Returns the logging status of a bucket and the permissions users have to view and modify that
status.

The following operations are related to `GetBucketLogging`:

- [CreateBucket](api-createbucket.md)

- [PutBucketLogging](api-putbucketlogging.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?logging HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketLogging_RequestSyntax)**

The bucket name for which to get the logging information.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketLogging_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<BucketLoggingStatus>
   <LoggingEnabled>
      <TargetBucket>string</TargetBucket>
      <TargetGrants>
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
      </TargetGrants>
      <TargetObjectKeyFormat>
         <PartitionedPrefix>
            <PartitionDateSource>string</PartitionDateSource>
         </PartitionedPrefix>
         <SimplePrefix>
         </SimplePrefix>
      </TargetObjectKeyFormat>
      <TargetPrefix>string</TargetPrefix>
   </LoggingEnabled>
</BucketLoggingStatus>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[BucketLoggingStatus](#API_GetBucketLogging_ResponseSyntax)**

Root level tag for the BucketLoggingStatus parameters.

Required: Yes

**[LoggingEnabled](#API_GetBucketLogging_ResponseSyntax)**

Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a
bucket. For more information, see [PUT Bucket logging](https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the
_Amazon S3 API Reference_.

Type: [LoggingEnabled](https://docs.aws.amazon.com/AmazonS3/latest/API/API_LoggingEnabled.html) data type

## Examples

### Sample Request

The following request returns the logging status for `amzn-s3-demo-bucket`.

```

            GET ?logging HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
            Date: Wed, 25 Nov 2009 12:00:00 GMT
            Authorization: authorization string

```

### Sample Response: Showing an enabled logging status

This example illustrates one usage of GetBucketLogging.

```

            HTTP/1.1 200 OK
            Date: Wed, 25 Nov 2009 12:00:00 GMT
            Connection: close
            Server: AmazonS3

            <?xml version="1.0" encoding="UTF-8"?>
            <BucketLoggingStatus xmlns="http://doc.s3.amazonaws.com/2006-03-01">
             <LoggingEnabled>
              <TargetBucket>amzn-s3-demo-bucket</TargetBucket>
              <TargetPrefix>mybucket-access_log-/</TargetPrefix>
                <TargetGrants>
                  <Grant>
                   <Grantee xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                    xsi:type="CanonicalUser">
                    <ID>79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be</ID>
                   </Grantee>
                   <Permission>READ</Permission>
                 </Grant>
                </TargetGrants>
            </LoggingEnabled>
            </BucketLoggingStatus>

```

### Sample Response: Showing a disabled logging status

This example illustrates one usage of GetBucketLogging.

```

         HTTP/1.1 200 OK
         Date: Wed, 25 Nov 2009 12:00:00 GMT
         Connection: close
         Server: AmazonS3

         <?xml version="1.0" encoding="UTF-8"?>

         <BucketLoggingStatus xmlns="http://doc.s3.amazonaws.com/2006-03-01">
          <!--<LoggingEnabled><TargetBucket>myLogsBucket</TargetBucket><TargetPrefix>add/this/prefix/to/my/log/files/access_log-</TargetPrefix></LoggingEnabled>-->
         </BucketLoggingStatus>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/GetBucketLogging)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/GetBucketLogging)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/GetBucketLogging)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/GetBucketLogging)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/GetBucketLogging)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/GetBucketLogging)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/GetBucketLogging)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/GetBucketLogging)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/GetBucketLogging)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/GetBucketLogging)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetBucketLocation

GetBucketMetadataConfiguration
