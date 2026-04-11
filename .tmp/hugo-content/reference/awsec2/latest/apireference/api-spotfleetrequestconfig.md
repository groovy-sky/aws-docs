# SpotFleetRequestConfig

Describes a Spot Fleet request.

## Contents

**activityStatus**

The progress of the Spot Fleet request.
If there is an error, the status is `error`.
After all requests are placed, the status is `pending_fulfillment`.
If the size of the fleet is equal to or greater than its target capacity, the status is `fulfilled`.
If the size of the fleet is decreased, the status is `pending_termination`
while Spot Instances are terminating.

Type: String

Valid Values: `error | pending_fulfillment | pending_termination | fulfilled`

Required: No

**createTime**

The creation date and time of the request.

Type: Timestamp

Required: No

**spotFleetRequestConfig**

The configuration of the Spot Fleet request.

Type: [SpotFleetRequestConfigData](api-spotfleetrequestconfigdata.md) object

Required: No

**spotFleetRequestId**

The ID of the Spot Fleet request.

Type: String

Required: No

**spotFleetRequestState**

The state of the Spot Fleet request.

Type: String

Valid Values: `submitted | active | cancelled | failed | cancelled_running | cancelled_terminating | modifying`

Required: No

**TagSet.N**

The tags for a Spot Fleet resource.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/spotfleetrequestconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/spotfleetrequestconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/spotfleetrequestconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SpotFleetMonitoring

SpotFleetRequestConfigData
