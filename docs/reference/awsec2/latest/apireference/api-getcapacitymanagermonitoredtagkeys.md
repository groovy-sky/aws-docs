---
title: "GetCapacityManagerMonitoredTagKeys"
---

# GetCapacityManagerMonitoredTagKeys

Retrieves the tag keys that are currently being monitored by EC2 Capacity Manager. Monitored tag keys are included as dimensions in capacity metric data, enabling you to group and filter metrics by tag values.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.
If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**MaxResults**

The maximum number of results to return in a single call. To retrieve the remaining results, make another call with the returned `NextToken` value. If not specified, up to 1000 results are returned.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**NextToken**

The token for the next page of results. Use the value returned from a previous call to retrieve additional results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**capacityManagerTagKeySet**

The list of tag keys being monitored by Capacity Manager, including their current status and metadata.

Type: Array of [CapacityManagerMonitoredTagKey](api-capacitymanagermonitoredtagkey.md) objects

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

This example retrieves the tag keys currently being monitored by Capacity Manager.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetCapacityManagerMonitoredTagKeys
&AUTHPARAMS
```

#### Sample Response

```

<GetCapacityManagerMonitoredTagKeysResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestId>
    <capacityManagerTagKeySet>
        <item>
            <tagKey>Environment</tagKey>
            <status>activated</status>
            <statusMessage/>
            <capacityManagerProvided>false</capacityManagerProvided>
            <earliestDatapointTimestamp>2025-12-23T20:00:00+00:00</earliestDatapointTimestamp>
        </item>
        <item>
            <tagKey>aws:autoscaling:groupName</tagKey>
            <status>activated</status>
            <statusMessage/>
            <capacityManagerProvided>true</capacityManagerProvided>
            <earliestDatapointTimestamp>2025-12-23T20:00:00+00:00</earliestDatapointTimestamp>
        </item>
    </capacityManagerTagKeySet>
</GetCapacityManagerMonitoredTagKeysResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getcapacitymanagermonitoredtagkeys.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetCapacityManagerMetricDimensions

GetCapacityReservationUsage

All content copied from https://docs.aws.amazon.com/.
