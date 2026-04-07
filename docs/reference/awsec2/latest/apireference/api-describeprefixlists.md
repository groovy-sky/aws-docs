# DescribePrefixLists

Describes available AWS services in a prefix list format, which includes the prefix list
name and prefix list ID of the service and the IP address range for the service.

We recommend that you use [DescribeManagedPrefixLists](api-describemanagedprefixlists.md) instead.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `prefix-list-id`: The ID of a prefix list.

- `prefix-list-name`: The name of a prefix list.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return with a single call.
To retrieve the remaining results, make another call with the returned `nextToken` value.

Type: Integer

Required: No

**NextToken**

The token for the next page of results.

Type: String

Required: No

**PrefixListId.N**

One or more prefix list IDs.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**prefixListSet**

All available prefix lists.

Type: Array of [PrefixList](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PrefixList.html) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example lists all available AWS prefix lists.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribePrefixLists
&AUTHPARAMS
```

#### Sample Response

```

<DescribePrefixListsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <prefixListSet>
        <item>
            <prefixListName>com.amazonaws.us-west-2.s3</prefixListName>
            <prefixListId>pl-12345678</prefixListId>
            <cidrSet>
              <item>54.123.456.7/19</item>
            </cidrSet>
        </item>
    </prefixListSet>
    <requestId>614db4d4-ac7b-4cb6-853e-example</requestId>
</DescribePrefixListsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribePrefixLists)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribePrefixLists)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribePrefixLists)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribePrefixLists)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribePrefixLists)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribePrefixLists)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribePrefixLists)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribePrefixLists)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribePrefixLists)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribePrefixLists)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribePlacementGroups

DescribePrincipalIdFormat
