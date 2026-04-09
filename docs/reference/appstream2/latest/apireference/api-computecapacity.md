# ComputeCapacity

Describes the capacity for a fleet.

## Contents

**DesiredInstances**

The desired number of streaming instances.

Type: Integer

Required: No

**DesiredSessions**

The desired number of user sessions for a multi-session fleet. This is not allowed for single-session fleets.

When you create a fleet, you must set either the DesiredSessions or DesiredInstances attribute, based on the type of fleet you create. You can’t define both attributes or leave both attributes blank.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/computecapacity.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/computecapacity.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/computecapacity.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CertificateBasedAuthProperties

ComputeCapacityStatus

All content copied from https://docs.aws.amazon.com/.
