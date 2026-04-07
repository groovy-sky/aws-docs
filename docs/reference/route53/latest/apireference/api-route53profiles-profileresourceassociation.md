# ProfileResourceAssociation

The association between a Route 53 Profile and resources.

## Contents

**CreationTime**

The date and time that the Profile resource association was created, in Unix time format and Coordinated Universal Time (UTC).

Type: Timestamp

Required: No

**Id**

ID of the Profile resource association.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**ModificationTime**

The date and time that the Profile resource association was modified, in Unix time format and Coordinated Universal Time (UTC).

Type: Timestamp

Required: No

**Name**

Name of the Profile resource association.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 64.

Pattern: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

Required: No

**OwnerId**

AWS account ID of the Profile resource association owner.

Type: String

Length Constraints: Minimum length of 12. Maximum length of 32.

Required: No

**ProfileId**

Profile ID of the Profile that the resources are associated with.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**ResourceArn**

The Amazon Resource Name (ARN) of the resource association.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**ResourceProperties**

If the DNS resource is a DNS Firewall rule group, this indicates the priority.

Type: String

Required: No

**ResourceType**

Resource type, such as a private hosted zone, interface VPC endpoint, or DNS Firewall rule group.

Type: String

Required: No

**Status**

Status of the Profile resource association.

Type: String

Valid Values: `COMPLETE | DELETING | UPDATING | CREATING | DELETED | FAILED`

Required: No

**StatusMessage**

Additional information about the Profile resource association.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53profiles-2018-05-10/ProfileResourceAssociation)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53profiles-2018-05-10/ProfileResourceAssociation)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53profiles-2018-05-10/ProfileResourceAssociation)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProfileAssociation

ProfileSummary
