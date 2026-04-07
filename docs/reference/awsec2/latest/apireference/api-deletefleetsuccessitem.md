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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DeleteFleetSuccessItem)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DeleteFleetSuccessItem)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DeleteFleetSuccessItem)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteFleetErrorItem

DeleteLaunchTemplateVersionsResponseErrorItem
