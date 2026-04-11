# AlternateContact

A structure that contains the details of an alternate contact associated with an AWS
account

## Contents

**AlternateContactType**

The type of alternate contact.

Type: String

Valid Values: `BILLING | OPERATIONS | SECURITY`

Required: No

**EmailAddress**

The email address associated with this alternate contact.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 254.

Pattern: `[\s]*[\w+=.#|!&-]+@[\w.-]+\.[\w]+[\s]*`

Required: No

**Name**

The name associated with this alternate contact.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**PhoneNumber**

The phone number associated with this alternate contact.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 25.

Pattern: `[\s0-9()+-]+`

Required: No

**Title**

The title associated with this alternate contact.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/account-2021-02-01/alternatecontact.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/account-2021-02-01/alternatecontact.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/account-2021-02-01/alternatecontact.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

ContactInformation

All content copied from https://docs.aws.amazon.com/.
