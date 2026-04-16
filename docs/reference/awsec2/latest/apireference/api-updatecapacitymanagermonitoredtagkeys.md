---
title: "UpdateCapacityManagerMonitoredTagKeys"
---

# UpdateCapacityManagerMonitoredTagKeys

Activates or deactivates tag keys for monitoring by EC2 Capacity Manager. Activated tag keys are included as dimensions in capacity metric data, enabling you to group and filter metrics by tag values.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ActivateTagKey.N**

The tag keys to activate for monitoring. Once activated, these tag keys will be included as dimensions in capacity metric data.

Type: Array of strings

Required: No

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.

Type: String

Required: No

**DeactivateTagKey.N**

The tag keys to deactivate. Deactivated tag keys will no longer be included as dimensions in capacity metric data.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.
If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**capacityManagerTagKeySet**

The list of tag keys affected by the update, including their current status and metadata.

Type: Array of [CapacityManagerMonitoredTagKey](api-capacitymanagermonitoredtagkey.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example activates the `Environment` and `CostCenter` tag keys for monitoring, and deactivates the `Team` tag key.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=UpdateCapacityManagerMonitoredTagKeys
&ActivateTagKey.1=Environment
&ActivateTagKey.2=CostCenter
&DeactivateTagKey.1=Team
&AUTHPARAMS
```

#### Sample Response

```

<UpdateCapacityManagerMonitoredTagKeysResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestId>
    <capacityManagerTagKeySet>
        <item>
            <tagKey>Environment</tagKey>
            <status>activating</status>
            <statusMessage/>
            <capacityManagerProvided>false</capacityManagerProvided>
        </item>
        <item>
            <tagKey>CostCenter</tagKey>
            <status>activating</status>
            <statusMessage/>
            <capacityManagerProvided>false</capacityManagerProvided>
        </item>
        <item>
            <tagKey>Team</tagKey>
            <status>deactivating</status>
            <statusMessage/>
            <capacityManagerProvided>false</capacityManagerProvided>
        </item>
    </capacityManagerTagKeySet>
</UpdateCapacityManagerMonitoredTagKeysResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/updatecapacitymanagermonitoredtagkeys.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UnmonitorInstances

UpdateCapacityManagerOrganizationsAccess

All content copied from https://docs.aws.amazon.com/.
