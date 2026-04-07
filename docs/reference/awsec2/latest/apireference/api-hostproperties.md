# HostProperties

Describes the properties of a Dedicated Host.

## Contents

**cores**

The number of cores on the Dedicated Host.

Type: Integer

Required: No

**instanceFamily**

The instance family supported by the Dedicated Host. For example,
`m5`.

Type: String

Required: No

**instanceType**

The instance type supported by the Dedicated Host. For example, `m5.large`.
If the host supports multiple instance types, no **instanceType** is returned.

Type: String

Required: No

**sockets**

The number of sockets on the Dedicated Host.

Type: Integer

Required: No

**totalVCpus**

The total number of vCPUs on the Dedicated Host.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/hostproperties.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/hostproperties.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/hostproperties.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

HostOffering

HostReservation
