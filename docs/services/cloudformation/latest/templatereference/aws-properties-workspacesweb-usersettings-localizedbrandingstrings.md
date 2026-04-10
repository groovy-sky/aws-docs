This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::UserSettings LocalizedBrandingStrings

Localized text strings for a specific language that customize the web portal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BrowserTabTitle" : String,
  "ContactButtonText" : String,
  "ContactLink" : String,
  "LoadingText" : String,
  "LoginButtonText" : String,
  "LoginDescription" : String,
  "LoginTitle" : String,
  "WelcomeText" : String
}

```

### YAML

```yaml

  BrowserTabTitle: String
  ContactButtonText: String
  ContactLink: String
  LoadingText: String
  LoginButtonText: String
  LoginDescription: String
  LoginTitle: String
  WelcomeText: String

```

## Properties

`BrowserTabTitle`

The text displayed in the browser tab title.

_Required_: Yes

_Type_: String

_Pattern_: ``^[^<>&'`~\\]*$``

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContactButtonText`

The text displayed on the contact button. This field is optional and defaults to "Contact us".

_Required_: No

_Type_: String

_Pattern_: ``^[^<>&'`~\\]*$``

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContactLink`

A contact link URL. The URL must start with `https://` or `mailto:`. If not provided, the contact button will be hidden from the web portal screen.

_Required_: No

_Type_: String

_Pattern_: `^(https?://|mailto:).*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadingText`

The text displayed during session loading. This field is optional and defaults to "Loading your session".

_Required_: No

_Type_: String

_Pattern_: ``^[^<>&'`~\\]*$``

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoginButtonText`

The text displayed on the login button. This field is optional and defaults to "Sign In".

_Required_: No

_Type_: String

_Pattern_: ``^[^<>&'`~\\]*$``

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoginDescription`

The description text for the login section. This field is optional and defaults to "Sign in to your session".

_Required_: No

_Type_: String

_Pattern_: ``^[^<>&'`~\\]*$``

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoginTitle`

The title text for the login section. This field is optional and defaults to "Sign In".

_Required_: No

_Type_: String

_Pattern_: ``^[^<>&'`~\\]*$``

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WelcomeText`

The welcome text displayed on the sign-in page.

_Required_: Yes

_Type_: String

_Pattern_: ``^[^<>&'`~\\]*$``

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageMetadata

Tag

All content copied from https://docs.aws.amazon.com/.
