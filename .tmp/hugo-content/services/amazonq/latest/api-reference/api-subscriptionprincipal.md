# SubscriptionPrincipal

A user or group in the IAM Identity Center instance connected to the Amazon Q Business
application.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**group**

The identifier of a group in the IAM Identity Center instance connected to the
Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 47.

Pattern: `([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}`

Required: No

**user**

The identifier of a user in the IAM Identity Center instance connected to the
Amazon Q Business application.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 47.

Pattern: `([0-9a-f]{10}-|)[A-Fa-f0-9]{8}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{4}-[A-Fa-f0-9]{12}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/subscriptionprincipal.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/subscriptionprincipal.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/subscriptionprincipal.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubscriptionDetails

Tag

All content copied from https://docs.aws.amazon.com/.
