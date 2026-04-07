# RestoreManagedPrefixListVersion

Restores the entries from a previous version of a managed prefix list to a new version of the prefix list.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CurrentVersion**

The current version number for the prefix list.

Type: Long

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**PrefixListId**

The ID of the prefix list.

Type: String

Required: Yes

**PreviousVersion**

The version to restore.

Type: Long

Required: Yes

## Response Elements

The following elements are returned by the service.

**prefixList**

Information about the prefix list.

Type: [ManagedPrefixList](api-managedprefixlist.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example restores the entries from version 1 of the specified prefix
list.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=RestoreManagedPrefixListVersion
&PrefixListId=pl-0123123123123aabb
&CurrentVersion=3
&PreviousVersion=1
&AUTHPARAMS
```

#### Sample Response

```

<RestoreManagedPrefixListVersionResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>aeb3faff-8938-41a0-9747-example</requestId>
    <prefixList>
        <addressFamily>IPv4</addressFamily>
        <maxEntries>10</maxEntries>
        <ownerId>123456789012</ownerId>
        <prefixListArn>arn:aws:ec2:us-east-1:123456789012:prefix-list/pl-0123123123123aabb</prefixListArn>
        <prefixListId>pl-0123123123123aabb</prefixListId>
        <prefixListName>tgw-attachments</prefixListName>
        <state>restore-in-progress</state>
        <version>3</version>
    </prefixList>
</RestoreManagedPrefixListVersionResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/RestoreManagedPrefixListVersion)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/RestoreManagedPrefixListVersion)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/restoremanagedprefixlistversion.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/restoremanagedprefixlistversion.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/restoremanagedprefixlistversion.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/restoremanagedprefixlistversion.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/restoremanagedprefixlistversion.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/restoremanagedprefixlistversion.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/RestoreManagedPrefixListVersion)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/restoremanagedprefixlistversion.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RestoreImageFromRecycleBin

RestoreSnapshotFromRecycleBin
