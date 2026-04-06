# CustomizationConfiguration

Contains the configuration information to customize the logo, font, and color of an Amazon Q Business web experience with individual files for each property or a CSS file for them all.

## Contents

**customCSSUrl**

Provides the URL where the custom CSS file is hosted for an Amazon Q web experience.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `(https?://[a-zA-Z0-9-_.+%/]+\.css)?`

Required: No

**faviconUrl**

Provides the URL where the custom favicon file is hosted for an Amazon Q web experience.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `(https?://[a-zA-Z0-9-_.+%/]+\.(svg|ico))?`

Required: No

**fontUrl**

Provides the URL where the custom font file is hosted for an Amazon Q web experience.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `(https?://[a-zA-Z0-9-_.+%/]+\.(ttf|woff|woff2|otf))?`

Required: No

**logoUrl**

Provides the URL where the custom logo file is hosted for an Amazon Q web experience.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `(https?://[a-zA-Z0-9-_.+%/]+\.(svg|png))?`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/CustomizationConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/CustomizationConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/CustomizationConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreatorModeConfiguration

CustomPluginConfiguration
