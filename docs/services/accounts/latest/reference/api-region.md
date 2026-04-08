# Region

This is a structure that expresses the Region for a given account, consisting of a
name and opt-in status.

## Contents

**RegionName**

The Region code of a given Region (for example, `us-east-1`).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 50.

Required: No

**RegionOptStatus**

One of potential statuses a Region can undergo (Enabled, Enabling, Disabled,
Disabling, Enabled\_By\_Default).

Type: String

Valid Values: `ENABLED | ENABLING | DISABLING | DISABLED | ENABLED_BY_DEFAULT`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/account-2021-02-01/region.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/account-2021-02-01/region.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/account-2021-02-01/region.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContactInformation

ValidationExceptionField

All content copied from https://docs.aws.amazon.com/.
