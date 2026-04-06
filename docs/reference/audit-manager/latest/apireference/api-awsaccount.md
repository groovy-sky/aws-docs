# AWSAccount

The wrapper of AWS account details, such as account ID or email address.

## Contents

**emailAddress**

The email address that's associated with the AWS account.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 320.

Pattern: `^.*@.*$`

Required: No

**id**

The identifier for the AWS account.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^[0-9]{12}$`

Required: No

**name**

The name of the AWS account.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Pattern: `^[\u0020-\u007E]+$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/auditmanager-2017-07-25/AWSAccount)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/auditmanager-2017-07-25/AWSAccount)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/auditmanager-2017-07-25/AWSAccount)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssessmentReportsDestination

AWSService
