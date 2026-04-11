# IamInstanceProfileAssociation

Describes an association between an IAM instance profile and an instance.

## Contents

**associationId**

The ID of the association.

Type: String

Required: No

**iamInstanceProfile**

The IAM instance profile.

Type: [IamInstanceProfile](api-iaminstanceprofile.md) object

Required: No

**instanceId**

The ID of the instance.

Type: String

Required: No

**state**

The state of the association.

Type: String

Valid Values: `associating | associated | disassociating | disassociated`

Required: No

**timestamp**

The time the IAM instance profile was associated with the instance.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/iaminstanceprofileassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/iaminstanceprofileassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/iaminstanceprofileassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IamInstanceProfile

IamInstanceProfileSpecification
