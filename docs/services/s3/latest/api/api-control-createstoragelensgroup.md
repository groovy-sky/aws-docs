# CreateStorageLensGroup

Creates a new S3 Storage Lens group and associates it with the specified AWS account ID. An
S3 Storage Lens group is a custom grouping of objects based on prefix, suffix, object tags,
object size, object age, or a combination of these filters. For each Storage Lens group
that you’ve created, you can also optionally add AWS resource tags. For more information
about S3 Storage Lens groups, see [Working with S3 Storage Lens\
groups](../userguide/storage-lens-groups-overview.md).

To use this operation, you must have the permission to perform the
`s3:CreateStorageLensGroup` action. If you’re trying to create a Storage Lens
group with AWS resource tags, you must also have permission to perform the
`s3:TagResource` action. For more information about the required Storage Lens
Groups permissions, see [Setting account permissions to use S3 Storage Lens groups](../userguide/storage-lens-iam-permissions.md#storage_lens_groups_permissions).

For information about Storage Lens groups errors, see [List of Amazon S3 Storage\
Lens error codes](errorresponses.md#S3LensErrorCodeList).

## Request Syntax

```nohighlight

POST /v20180820/storagelensgroup HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<CreateStorageLensGroupRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <StorageLensGroup>
      <Filter>
         <And>
            <MatchAnyPrefix>
               <Prefix>string</Prefix>
            </MatchAnyPrefix>
            <MatchAnySuffix>
               <Suffix>string</Suffix>
            </MatchAnySuffix>
            <MatchAnyTag>
               <Tag>
                  <Key>string</Key>
                  <Value>string</Value>
               </Tag>
            </MatchAnyTag>
            <MatchObjectAge>
               <DaysGreaterThan>integer</DaysGreaterThan>
               <DaysLessThan>integer</DaysLessThan>
            </MatchObjectAge>
            <MatchObjectSize>
               <BytesGreaterThan>long</BytesGreaterThan>
               <BytesLessThan>long</BytesLessThan>
            </MatchObjectSize>
         </And>
         <MatchAnyPrefix>
            <Prefix>string</Prefix>
         </MatchAnyPrefix>
         <MatchAnySuffix>
            <Suffix>string</Suffix>
         </MatchAnySuffix>
         <MatchAnyTag>
            <Tag>
               <Key>string</Key>
               <Value>string</Value>
            </Tag>
         </MatchAnyTag>
         <MatchObjectAge>
            <DaysGreaterThan>integer</DaysGreaterThan>
            <DaysLessThan>integer</DaysLessThan>
         </MatchObjectAge>
         <MatchObjectSize>
            <BytesGreaterThan>long</BytesGreaterThan>
            <BytesLessThan>long</BytesLessThan>
         </MatchObjectSize>
         <Or>
            <MatchAnyPrefix>
               <Prefix>string</Prefix>
            </MatchAnyPrefix>
            <MatchAnySuffix>
               <Suffix>string</Suffix>
            </MatchAnySuffix>
            <MatchAnyTag>
               <Tag>
                  <Key>string</Key>
                  <Value>string</Value>
               </Tag>
            </MatchAnyTag>
            <MatchObjectAge>
               <DaysGreaterThan>integer</DaysGreaterThan>
               <DaysLessThan>integer</DaysLessThan>
            </MatchObjectAge>
            <MatchObjectSize>
               <BytesGreaterThan>long</BytesGreaterThan>
               <BytesLessThan>long</BytesLessThan>
            </MatchObjectSize>
         </Or>
      </Filter>
      <Name>string</Name>
      <StorageLensGroupArn>string</StorageLensGroupArn>
   </StorageLensGroup>
   <Tags>
      <Tag>
         <Key>string</Key>
         <Value>string</Value>
      </Tag>
   </Tags>
</CreateStorageLensGroupRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[x-amz-account-id](#API_control_CreateStorageLensGroup_RequestSyntax)**

The AWS account ID that the Storage Lens group is created from and associated with.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[CreateStorageLensGroupRequest](#API_control_CreateStorageLensGroup_RequestSyntax)**

Root level tag for the CreateStorageLensGroupRequest parameters.

Required: Yes

**[StorageLensGroup](#API_control_CreateStorageLensGroup_RequestSyntax)**

The Storage Lens group configuration.

Type: [StorageLensGroup](api-control-storagelensgroup.md) data type

Required: Yes

**[Tags](#API_control_CreateStorageLensGroup_RequestSyntax)**

The AWS resource tags that you're adding to your Storage Lens group. This parameter is optional.

Type: Array of [Tag](api-control-tag.md) data types

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/createstoragelensgroup.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/createstoragelensgroup.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/createstoragelensgroup.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/createstoragelensgroup.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/createstoragelensgroup.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/createstoragelensgroup.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/createstoragelensgroup.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/createstoragelensgroup.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/createstoragelensgroup.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/createstoragelensgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateMultiRegionAccessPoint

DeleteAccessGrant

All content copied from https://docs.aws.amazon.com/.
