# UserStackAssociationError

Describes the error that is returned when a user can’t be associated with or disassociated from a stack.

## Contents

**ErrorCode**

The error code for the error that is returned when a user can’t be associated with or disassociated from a stack.

Type: String

Valid Values: `STACK_NOT_FOUND | USER_NAME_NOT_FOUND | DIRECTORY_NOT_FOUND | INTERNAL_ERROR`

Required: No

**ErrorMessage**

The error message for the error that is returned when a user can’t be associated with or disassociated from a stack.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**UserStackAssociation**

Information about the user and associated stack.

Type: [UserStackAssociation](api-userstackassociation.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/userstackassociationerror.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/userstackassociationerror.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/userstackassociationerror.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserStackAssociation

VolumeConfig

All content copied from https://docs.aws.amazon.com/.
