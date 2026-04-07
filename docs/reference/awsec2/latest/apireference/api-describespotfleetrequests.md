# DescribeSpotFleetRequests

Describes your Spot Fleet requests.

Spot Fleet requests are deleted 48 hours after they are canceled and their instances
are terminated.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**MaxResults**

The maximum number of items to return for this request.
To get the next page of items, make another request with the token returned in the output.
For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Required: No

**NextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

Required: No

**SpotFleetRequestId.N**

The IDs of the Spot Fleet requests.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The token to include in another request to get the next page of items. This value is `null` when there
are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

**spotFleetRequestConfigSet**

Information about the configuration of your Spot Fleet.

Type: Array of [SpotFleetRequestConfig](api-spotfleetrequestconfig.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes all of your Spot Fleet requests.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeSpotFleetRequests
&AUTHPARAMS
```

#### Sample Response

```

<DescribeSpotFleetRequestsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>4d68a6cc-8f2e-4be1-b425-example</requestId>
    <spotFleetRequestConfigSet>
        <item>
            <spotFleetRequestId>sfr-12345678-cb31-425e-8c23-example2710</spotFleetRequestId>
            <spotFleetRequestState>cancelled</spotFleetRequestState>
            <spotFleetRequestConfig>
                <spotPrice>0.0153</spotPrice>
                <targetCapacity>20</targetCapacity>
                <iamFleetRole>arn:aws:iam::123456789011:role/spot-fleet-role</iamFleetRole>
                <launchSpecifications>
                    <item>
                        <subnetId>subnet-1a2b3c4d</subnetId>
                        <ebsOptimized>false</ebsOptimized>
                        <imageId>ami-1ecae776</imageId>
                        <instanceType>m4.xlarge</instanceType>
                    </item>
                    <item>
                        <subnetId>subnet-1a2b3c4d</subnetId>
                        <ebsOptimized>false</ebsOptimized>
                        <imageId>ami-1ecae776</imageId>
                        <instanceType>m3.medium</instanceType>
                    </item>
                </launchSpecifications>
            </spotFleetRequestConfig>
        </item>
        <item>
            <spotFleetRequestId>sfr-abcdefgh-e71f-450d-880d-examplec127</spotFleetRequestId>
            <spotFleetRequestState>active</spotFleetRequestState>
            <spotFleetRequestConfig>
                <spotPrice>0.0153</spotPrice>
                <targetCapacity>5</targetCapacity>
                <iamFleetRole>arn:aws:iam::123456789011:role/spot-fleet-role</iamFleetRole>
                <launchSpecifications>
                    <item>
                        <subnetId>subnet-abc123ab</subnetId>
                        <ebsOptimized>false</ebsOptimized>
                        <imageId>ami-1ecae776</imageId>
                        <instanceType>m4.large</instanceType>
                    </item>
                    <item>
                        <subnetId>subnet-abc123ab</subnetId>
                        <ebsOptimized>false</ebsOptimized>
                        <imageId>ami-1ecae776</imageId>
                        <instanceType>m3.medium</instanceType>
                    </item>
                </launchSpecifications>
            </spotFleetRequestConfig>
        </item>
    </spotFleetRequestConfigSet>
</DescribeSpotFleetRequestsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeSpotFleetRequests)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeSpotFleetRequests)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describespotfleetrequests.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describespotfleetrequests.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describespotfleetrequests.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describespotfleetrequests.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describespotfleetrequests.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describespotfleetrequests.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeSpotFleetRequests)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describespotfleetrequests.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeSpotFleetRequestHistory

DescribeSpotInstanceRequests
