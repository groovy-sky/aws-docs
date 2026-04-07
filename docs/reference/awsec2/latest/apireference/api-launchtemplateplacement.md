# LaunchTemplatePlacement

Describes the placement of an instance.

## Contents

**affinity**

The affinity setting for the instance on the Dedicated Host.

Type: String

Required: No

**availabilityZone**

The Availability Zone of the instance.

Type: String

Required: No

**availabilityZoneId**

The ID of the Availability Zone of the instance.

Type: String

Required: No

**groupId**

The Group ID of the placement group. You must specify the Placement Group **Group ID** to launch an instance in a shared placement
group.

Type: String

Required: No

**groupName**

The name of the placement group for the instance.

Type: String

Required: No

**hostId**

The ID of the Dedicated Host for the instance.

Type: String

Required: No

**hostResourceGroupArn**

The ARN of the host resource group in which to launch the instances.

Type: String

Required: No

**partitionNumber**

The number of the partition the instance should launch in. Valid only if the placement
group strategy is set to `partition`.

Type: Integer

Required: No

**spreadDomain**

Reserved for future use.

Type: String

Required: No

**tenancy**

The tenancy of the instance. An instance with a tenancy of `dedicated` runs
on single-tenant hardware.

Type: String

Valid Values: `default | dedicated | host`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateplacement.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateplacement.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateplacement.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateOverrides

LaunchTemplatePlacementRequest
