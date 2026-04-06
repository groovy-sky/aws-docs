# PutStorageLensConfigurationTagging

###### Note

This operation is not supported by directory buckets.

Put or replace tags on an existing Amazon S3 Storage Lens configuration. For more information
about S3 Storage Lens, see [Assessing your storage activity and usage with Amazon S3 Storage Lens](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html) in the
_Amazon S3 User Guide_.

###### Note

To use this action, you must have permission to perform the
`s3:PutStorageLensConfigurationTagging` action. For more information, see
[Setting permissions to\
use Amazon S3 Storage Lens](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html) in the _Amazon S3 User Guide_.

## Request Syntax

```nohighlight

PUT /v20180820/storagelens/storagelensid/tagging HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<PutStorageLensConfigurationTaggingRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <Tags>
      <Tag>
         <Key>string</Key>
         <Value>string</Value>
      </Tag>
   </Tags>
</PutStorageLensConfigurationTaggingRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[storagelensid](#API_control_PutStorageLensConfigurationTagging_RequestSyntax)**

The ID of the S3 Storage Lens configuration.

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[a-zA-Z0-9\-\_\.]+`

Required: Yes

**[x-amz-account-id](#API_control_PutStorageLensConfigurationTagging_RequestSyntax)**

The account ID of the requester.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[PutStorageLensConfigurationTaggingRequest](#API_control_PutStorageLensConfigurationTagging_RequestSyntax)**

Root level tag for the PutStorageLensConfigurationTaggingRequest parameters.

Required: Yes

**[Tags](#API_control_PutStorageLensConfigurationTagging_RequestSyntax)**

The tag set of the S3 Storage Lens configuration.

###### Note

You can set up to a maximum of 50 tags.

Type: Array of [StorageLensTag](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_StorageLensTag.html) data types

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3control-2018-08-20/PutStorageLensConfigurationTagging)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3control-2018-08-20/PutStorageLensConfigurationTagging)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/PutStorageLensConfigurationTagging)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3control-2018-08-20/PutStorageLensConfigurationTagging)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/PutStorageLensConfigurationTagging)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3control-2018-08-20/PutStorageLensConfigurationTagging)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3control-2018-08-20/PutStorageLensConfigurationTagging)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3control-2018-08-20/PutStorageLensConfigurationTagging)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3control-2018-08-20/PutStorageLensConfigurationTagging)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/PutStorageLensConfigurationTagging)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PutStorageLensConfiguration

SubmitMultiRegionAccessPointRoutes
