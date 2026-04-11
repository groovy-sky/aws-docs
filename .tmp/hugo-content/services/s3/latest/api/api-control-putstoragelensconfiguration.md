# PutStorageLensConfiguration

###### Note

This operation is not supported by directory buckets.

Puts an Amazon S3 Storage Lens configuration. For more information about S3 Storage Lens, see [Working with\
Amazon S3 Storage Lens](../dev/storage-lens.md) in the _Amazon S3 User Guide_. For a complete list of S3 Storage Lens metrics, see [S3 Storage Lens metrics glossary](../userguide/storage-lens-metrics-glossary.md) in the _Amazon S3 User Guide_.

###### Note

To use this action, you must have permission to perform the
`s3:PutStorageLensConfiguration` action. For more information, see [Setting permissions to use Amazon S3 Storage Lens](../dev/storage-lens-iam-permissions.md) in the
_Amazon S3 User Guide_.

## Request Syntax

```nohighlight

PUT /v20180820/storagelens/storagelensid HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<PutStorageLensConfigurationRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <StorageLensConfiguration>
      <AccountLevel>
         <ActivityMetrics>
            <IsEnabled>boolean</IsEnabled>
         </ActivityMetrics>
         <AdvancedCostOptimizationMetrics>
            <IsEnabled>boolean</IsEnabled>
         </AdvancedCostOptimizationMetrics>
         <AdvancedDataProtectionMetrics>
            <IsEnabled>boolean</IsEnabled>
         </AdvancedDataProtectionMetrics>
         <AdvancedPerformanceMetrics>
            <IsEnabled>boolean</IsEnabled>
         </AdvancedPerformanceMetrics>
         <BucketLevel>
            <ActivityMetrics>
               <IsEnabled>boolean</IsEnabled>
            </ActivityMetrics>
            <AdvancedCostOptimizationMetrics>
               <IsEnabled>boolean</IsEnabled>
            </AdvancedCostOptimizationMetrics>
            <AdvancedDataProtectionMetrics>
               <IsEnabled>boolean</IsEnabled>
            </AdvancedDataProtectionMetrics>
            <AdvancedPerformanceMetrics>
               <IsEnabled>boolean</IsEnabled>
            </AdvancedPerformanceMetrics>
            <DetailedStatusCodesMetrics>
               <IsEnabled>boolean</IsEnabled>
            </DetailedStatusCodesMetrics>
            <PrefixLevel>
               <StorageMetrics>
                  <IsEnabled>boolean</IsEnabled>
                  <SelectionCriteria>
                     <Delimiter>string</Delimiter>
                     <MaxDepth>integer</MaxDepth>
                     <MinStorageBytesPercentage>double</MinStorageBytesPercentage>
                  </SelectionCriteria>
               </StorageMetrics>
            </PrefixLevel>
         </BucketLevel>
         <DetailedStatusCodesMetrics>
            <IsEnabled>boolean</IsEnabled>
         </DetailedStatusCodesMetrics>
         <StorageLensGroupLevel>
            <SelectionCriteria>
               <Exclude>
                  <Arn>string</Arn>
               </Exclude>
               <Include>
                  <Arn>string</Arn>
               </Include>
            </SelectionCriteria>
         </StorageLensGroupLevel>
      </AccountLevel>
      <AwsOrg>
         <Arn>string</Arn>
      </AwsOrg>
      <DataExport>
         <CloudWatchMetrics>
            <IsEnabled>boolean</IsEnabled>
         </CloudWatchMetrics>
         <S3BucketDestination>
            <AccountId>string</AccountId>
            <Arn>string</Arn>
            <Encryption>
               <SSE-KMS>
                  <KeyId>string</KeyId>
               </SSE-KMS>
               <SSE-S3>
               </SSE-S3>
            </Encryption>
            <Format>string</Format>
            <OutputSchemaVersion>string</OutputSchemaVersion>
            <Prefix>string</Prefix>
         </S3BucketDestination>
         <StorageLensTableDestination>
            <Encryption>
               <SSE-KMS>
                  <KeyId>string</KeyId>
               </SSE-KMS>
               <SSE-S3>
               </SSE-S3>
            </Encryption>
            <IsEnabled>boolean</IsEnabled>
         </StorageLensTableDestination>
      </DataExport>
      <Exclude>
         <Buckets>
            <Arn>string</Arn>
         </Buckets>
         <Regions>
            <Region>string</Region>
         </Regions>
      </Exclude>
      <ExpandedPrefixesDataExport>
         <S3BucketDestination>
            <AccountId>string</AccountId>
            <Arn>string</Arn>
            <Encryption>
               <SSE-KMS>
                  <KeyId>string</KeyId>
               </SSE-KMS>
               <SSE-S3>
               </SSE-S3>
            </Encryption>
            <Format>string</Format>
            <OutputSchemaVersion>string</OutputSchemaVersion>
            <Prefix>string</Prefix>
         </S3BucketDestination>
         <StorageLensTableDestination>
            <Encryption>
               <SSE-KMS>
                  <KeyId>string</KeyId>
               </SSE-KMS>
               <SSE-S3>
               </SSE-S3>
            </Encryption>
            <IsEnabled>boolean</IsEnabled>
         </StorageLensTableDestination>
      </ExpandedPrefixesDataExport>
      <Id>string</Id>
      <Include>
         <Buckets>
            <Arn>string</Arn>
         </Buckets>
         <Regions>
            <Region>string</Region>
         </Regions>
      </Include>
      <IsEnabled>boolean</IsEnabled>
      <PrefixDelimiter>string</PrefixDelimiter>
      <StorageLensArn>string</StorageLensArn>
   </StorageLensConfiguration>
   <Tags>
      <Tag>
         <Key>string</Key>
         <Value>string</Value>
      </Tag>
   </Tags>
</PutStorageLensConfigurationRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[storagelensid](#API_control_PutStorageLensConfiguration_RequestSyntax)**

The ID of the S3 Storage Lens configuration.

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[a-zA-Z0-9\-\_\.]+`

Required: Yes

**[x-amz-account-id](#API_control_PutStorageLensConfiguration_RequestSyntax)**

The account ID of the requester.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[PutStorageLensConfigurationRequest](#API_control_PutStorageLensConfiguration_RequestSyntax)**

Root level tag for the PutStorageLensConfigurationRequest parameters.

Required: Yes

**[StorageLensConfiguration](#API_control_PutStorageLensConfiguration_RequestSyntax)**

The S3 Storage Lens configuration.

Type: [StorageLensConfiguration](api-control-storagelensconfiguration.md) data type

Required: Yes

**[Tags](#API_control_PutStorageLensConfiguration_RequestSyntax)**

The tag set of the S3 Storage Lens configuration.

###### Note

You can set up to a maximum of 50 tags.

Type: Array of [StorageLensTag](api-control-storagelenstag.md) data types

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/putstoragelensconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/putstoragelensconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/putstoragelensconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/putstoragelensconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/putstoragelensconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/putstoragelensconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/putstoragelensconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/putstoragelensconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/putstoragelensconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/putstoragelensconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutPublicAccessBlock

PutStorageLensConfigurationTagging

All content copied from https://docs.aws.amazon.com/.
