# DeleteFleetSuccessItem

Describes an EC2 Fleet that was successfully deleted.

## Contents

**currentFleetState**

The current state of the EC2 Fleet.

Type: String

Valid Values: `submitted | active | deleted | failed | deleted_running | deleted_terminating | modifying`

Required: No

**fleetId**

The ID of the EC2 Fleet.

Type: String

Required: No

**previousFleetState**

The previous state of the EC2 Fleet.

Type: String

Valid Values: `submitted | active | deleted | failed | deleted_running | deleted_terminating | modifying`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletefleetsuccessitem.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletefleetsuccessitem.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletefleetsuccessitem.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteFleetErrorItem

DeleteLaunchTemplateVersionsResponseErrorItem
