---
title: "DescribeCapacityManagerDataExports"
---

# DescribeCapacityManagerDataExports

Describes one or more Capacity Manager data export configurations. Returns information about export settings, delivery status, and recent export activity.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CapacityManagerDataExportId.N**

The IDs of the data export configurations to describe. If not specified, all export configurations are returned.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.
If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters to narrow the results. Supported filters include export status, creation date, and S3 bucket name.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return in a single call. If not specified, up to 1000 results are returned.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results. Use this value in a subsequent call to retrieve additional results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**capacityManagerDataExportSet**

Information about the data export configurations, including export settings, delivery status, and recent activity.

Type: Array of [CapacityManagerDataExportResponse](api-capacitymanagerdataexportresponse.md) objects

**nextToken**

The token to use to retrieve the next page of results. This value is null when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes a specific data export configuration.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeCapacityManagerDataExports
&CapacityManagerDataExportId.1=cmde-0abcd1234EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<DescribeCapacityManagerDataExportsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestId>
    <capacityManagerDataExportSet>
        <item>
            <capacityManagerDataExportId>cmde-0abcd1234EXAMPLE</capacityManagerDataExportId>
            <s3BucketName>my-capacity-exports-bucket</s3BucketName>
            <s3BucketPrefix>capacity-data/</s3BucketPrefix>
            <schedule>hourly</schedule>
            <outputFormat>parquet</outputFormat>
            <createTime>2024-01-15T10:30:00.000Z</createTime>
            <latestDeliveryStatus>delivered</latestDeliveryStatus>
            <latestDeliveryTime>2024-01-16T02:00:00.000Z</latestDeliveryTime>
            <latestDeliveryS3LocationUri>s3://my-capacity-exports-bucket/capacity-data/20240116.parquet</latestDeliveryS3LocationUri>
        </item>
    </capacityManagerDataExportSet>
</DescribeCapacityManagerDataExportsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describecapacitymanagerdataexports.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describecapacitymanagerdataexports.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describecapacitymanagerdataexports.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describecapacitymanagerdataexports.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describecapacitymanagerdataexports.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describecapacitymanagerdataexports.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describecapacitymanagerdataexports.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describecapacitymanagerdataexports.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describecapacitymanagerdataexports.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describecapacitymanagerdataexports.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeCapacityBlockStatus

DescribeCapacityReservationBillingRequests

All content copied from https://docs.aws.amazon.com/.
