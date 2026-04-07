# CancelSpotFleetRequestsSuccessItem

Describes a Spot Fleet request that was successfully canceled.

## Contents

**currentSpotFleetRequestState**

The current state of the Spot Fleet request.

Type: String

Valid Values: `submitted | active | cancelled | failed | cancelled_running | cancelled_terminating | modifying`

Required: No

**previousSpotFleetRequestState**

The previous state of the Spot Fleet request.

Type: String

Valid Values: `submitted | active | cancelled | failed | cancelled_running | cancelled_terminating | modifying`

Required: No

**spotFleetRequestId**

The ID of the Spot Fleet request.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/cancelspotfleetrequestssuccessitem.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/cancelspotfleetrequestssuccessitem.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/cancelspotfleetrequestssuccessitem.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CancelSpotFleetRequestsErrorItem

CapacityAllocation
