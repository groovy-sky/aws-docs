# DescribeRegions

Describes the Regions that are enabled for your account, or all Regions.

For a list of the Regions supported by Amazon EC2, see [Amazon EC2 service endpoints](../../../../services/ec2/latest/devguide/ec2-endpoints.md).

For information about enabling and disabling Regions for your account, see [Specify which AWS Regions \
your account can use](../../../../services/accounts/latest/reference/manage-acct-regions.md) in the _AWS Account Management Reference Guide_.

###### Note

The order of the elements in the response, including those within nested structures,
might vary. Applications should not assume the elements appear in a particular order.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllRegions**

Indicates whether to display all Regions, including Regions that are disabled for your account.

Type: Boolean

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `endpoint` \- The endpoint of the Region (for example, `ec2.us-east-1.amazonaws.com`).

- `opt-in-status` \- The opt-in status of the Region ( `opt-in-not-required` \| `opted-in` \|
`not-opted-in`).

- `region-name` \- The name of the Region (for example, `us-east-1`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**RegionName.N**

The names of the Regions. You can specify any Regions, whether they are enabled and disabled for your account.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**regionInfo**

Information about the Regions.

Type: Array of [Region](api-region.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example displays information about all Regions enabled for your account.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeRegions
&AUTHPARAMS
```

### Example 2

This example displays information about all Regions, even the Regions that are disabled for your account.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeRegions
&AllRegions=true
&AUTHPARAMS
```

### Example 3

This example displays information about the specified Regions only.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeRegions
&RegionName.1=us-east-1
&RegionName.2=eu-west-1
&AUTHPARAMS
```

#### Sample Response

```

<DescribeRegionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <regionInfo>
      <item>
         <regionName>us-east-1</regionName>
         <regionEndpoint>ec2.us-east-1.amazonaws.com</regionEndpoint>
         <optInStatus>opt-in-not-required</optInStatus>
         <geographySet>
            <item>
               <name>United States of America</name>
            </item>
         </geographySet>
      </item>
      <item>
         <regionName>eu-west-1</regionName>
         <regionEndpoint>ec2.eu-west-1.amazonaws.com</regionEndpoint>
         <optInStatus>opt-in-not-required</optInStatus>
         <geographySet>
            <item>
               <name>Ireland</name>
            </item>
         </geographySet>
      </item>
   </regionInfo>
</DescribeRegionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeregions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeregions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeregions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeregions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeregions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeregions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeregions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeregions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeregions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeregions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribePublicIpv4Pools

DescribeReplaceRootVolumeTasks
