# UserSetting

Describes an action and whether the action is enabled or disabled for users during their streaming sessions.

## Contents

**Action**

The action that is enabled or disabled.

Type: String

Valid Values: `CLIPBOARD_COPY_FROM_LOCAL_DEVICE | CLIPBOARD_COPY_TO_LOCAL_DEVICE | FILE_UPLOAD | FILE_DOWNLOAD | PRINTING_TO_LOCAL_DEVICE | DOMAIN_PASSWORD_SIGNIN | DOMAIN_SMART_CARD_SIGNIN | AUTO_TIME_ZONE_REDIRECTION`

Required: Yes

**Permission**

Indicates whether the action is enabled or disabled.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: Yes

**MaximumLength**

Specifies the number of characters that can be copied by end users from the local device to the remote session, and to the local device from the remote session.

This can be specified only for the `CLIPBOARD_COPY_FROM_LOCAL_DEVICE` and `CLIPBOARD_COPY_TO_LOCAL_DEVICE` actions.

This defaults to 20,971,520 (20 MB) when unspecified and the permission is `ENABLED`. This can't be specified when the permission is `DISABLED`.

The value can be between 1 and 20,971,520 (20 MB).

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/usersetting.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/usersetting.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/usersetting.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

User

UserStackAssociation

All content copied from https://docs.aws.amazon.com/.
