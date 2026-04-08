# DescribeMacHosts

Describes the specified EC2 Mac Dedicated Host or all of your EC2 Mac Dedicated Hosts.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Filter.N**

The filters.

- `availability-zone` \- The Availability Zone of the EC2 Mac Dedicated Host.

- `instance-type` \- The instance type size that the EC2 Mac Dedicated Host is
configured to support.

Type: Array of [Filter](api-filter.md) objects

Required: No

**HostId.N**

The IDs of the EC2 Mac Dedicated Hosts.

Type: Array of strings

Required: No

**MaxResults**

The maximum number of results to return for the request in a single page. The remaining results can be seen by sending another request with the returned `nextToken` value. This value can be between 5 and 500. If `maxResults` is given a larger value than 500, you receive an error.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 500.

Required: No

**NextToken**

The token to use to retrieve the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**macHostSet**

Information about the EC2 Mac Dedicated Hosts.

Type: Array of [MacHost](api-machost.md) objects

**nextToken**

The token to use to retrieve the next page of results.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describemachosts.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describemachosts.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describemachosts.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describemachosts.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describemachosts.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describemachosts.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describemachosts.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describemachosts.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describemachosts.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describemachosts.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeLockedSnapshots

DescribeMacModificationTasks
