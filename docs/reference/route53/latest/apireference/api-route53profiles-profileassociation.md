# ProfileAssociation

An association between a Route 53 Profile and a VPC.

## Contents

**CreationTime**

The date and time that the Profile association was created, in Unix time format and Coordinated Universal Time (UTC).

Type: Timestamp

Required: No

**Id**

ID of the Profile association.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**ModificationTime**

The date and time that the Profile association was modified, in Unix time format and Coordinated Universal Time (UTC).

Type: Timestamp

Required: No

**Name**

Name of the Profile association.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: No

**OwnerId**

AWS account ID of the Profile association owner.

Type: String

Length Constraints: Minimum length of 12. Maximum length of 32.

Required: No

**ProfileId**

ID of the Profile.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**ResourceId**

The Amazon Resource Name (ARN) of the VPC.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**Status**

Status of the Profile association.

Type: String

Valid Values: `COMPLETE | DELETING | UPDATING | CREATING | DELETED | FAILED`

Required: No

**StatusMessage**

Additional information about the Profile association.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/route53profiles-2018-05-10/profileassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53profiles-2018-05-10/profileassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53profiles-2018-05-10/profileassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Profile

ProfileResourceAssociation
