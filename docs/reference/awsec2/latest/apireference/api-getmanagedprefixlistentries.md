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

Type: Array of [PrefixListEntry](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PrefixListEntry.html) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetManagedPrefixListEntries)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetManagedPrefixListEntries)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/GetManagedPrefixListEntries)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/GetManagedPrefixListEntries)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/GetManagedPrefixListEntries)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/GetManagedPrefixListEntries)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/GetManagedPrefixListEntries)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/GetManagedPrefixListEntries)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetManagedPrefixListEntries)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/GetManagedPrefixListEntries)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetManagedPrefixListAssociations

GetNetworkInsightsAccessScopeAnalysisFindings
