# Application

Summary information for an Amazon Q Business application.

## Contents

**applicationId**

The identifier for the Amazon Q Business application.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9-]{35}`

Required: No

**createdAt**

The Unix timestamp when the Amazon Q Business application was created.

Type: Timestamp

Required: No

**displayName**

The name of the Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Pattern: `[a-zA-Z0-9][a-zA-Z0-9_-]*`

Required: No

**identityType**

The authentication type being used by a Amazon Q Business application.

Type: String

Valid Values: `AWS_IAM_IDP_SAML | AWS_IAM_IDP_OIDC | AWS_IAM_IDC | AWS_QUICKSIGHT_IDP | ANONYMOUS`

Required: No

**quickSightConfiguration**

The Amazon Quick configuration for an Amazon Q Business application that uses Quick as the identity provider.

Type: [QuickSightConfiguration](api-quicksightconfiguration.md) object

Required: No

**status**

The status of the Amazon Q Business application. The application is ready to use when the
status is `ACTIVE`.

Type: String

Valid Values: `CREATING | ACTIVE | DELETING | FAILED | UPDATING`

Required: No

**updatedAt**

The Unix timestamp when the Amazon Q Business application was last updated.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/application.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/application.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/application.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

APISchema

AppliedAttachmentsConfiguration

All content copied from https://docs.aws.amazon.com/.
