This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::WebExperience CustomizationConfiguration

Contains the configuration information to customize the logo, font, and color of an Amazon Q Business web experience with individual files for each property or a CSS file for them all.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomCSSUrl" : String,
  "FaviconUrl" : String,
  "FontUrl" : String,
  "LogoUrl" : String
}

```

### YAML

```yaml

  CustomCSSUrl: String
  FaviconUrl: String
  FontUrl: String
  LogoUrl: String

```

## Properties

`CustomCSSUrl`

Provides the URL where the custom CSS file is hosted for an Amazon Q web experience.

_Required_: No

_Type_: String

_Pattern_: `^(https?://[a-zA-Z0-9-_.+%/]+\.css)?$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FaviconUrl`

Provides the URL where the custom favicon file is hosted for an Amazon Q web experience.

_Required_: No

_Type_: String

_Pattern_: `^(https?://[a-zA-Z0-9-_.+%/]+\.(svg|ico))?$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FontUrl`

Provides the URL where the custom font file is hosted for an Amazon Q web experience.

_Required_: No

_Type_: String

_Pattern_: `^(https?://[a-zA-Z0-9-_.+%/]+\.(ttf|woff|woff2|otf))?$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogoUrl`

Provides the URL where the custom logo file is hosted for an Amazon Q web experience.

_Required_: No

_Type_: String

_Pattern_: `^(https?://[a-zA-Z0-9-_.+%/]+\.(svg|png))?$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BrowserExtensionConfiguration

IdentityProviderConfiguration
