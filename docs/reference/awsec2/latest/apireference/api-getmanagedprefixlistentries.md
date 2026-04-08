# GetManagedPrefixListEntries

Gets information about the entries for a specified managed prefix list.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**MaxResults**

The maximum number of results to return with a single call.
To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 100.

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

**PrefixListId**

The ID of the prefix list.

Type: String

Required: Yes

**TargetVersion**

The version of the prefix list for which to return the entries. The default is the current version.

Type: Long

Required: No

## Response Elements

The following elements are returned by the service.

**entrySet**

Information about the prefix list entries.

Type: Array of [PrefixListEntry](api-prefixlistentry.md) objects

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example gets the entries for the specified managed prefix list.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetManagedPrefixListEntries
&PrefixListId=pl-0123123123123aabb
&AUTHPARAMS
```

#### Sample Response

```

<GetManagedPrefixListEntriesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>39a3c79f-846f-4382-a592-example</requestId>
    <entrySet>
        <item>
            <cidr>10.0.0.0/16</cidr>
            <description>vpc-a</description>
        </item>
        <item>
            <cidr>10.2.0.0/16</cidr>
            <description>NY office</description>
        </item>
    </entrySet>
</GetManagedPrefixListEntriesResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getmanagedprefixlistentries.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getmanagedprefixlistentries.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getmanagedprefixlistentries.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getmanagedprefixlistentries.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getmanagedprefixlistentries.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getmanagedprefixlistentries.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getmanagedprefixlistentries.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getmanagedprefixlistentries.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getmanagedprefixlistentries.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getmanagedprefixlistentries.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetManagedPrefixListAssociations

GetNetworkInsightsAccessScopeAnalysisFindings
