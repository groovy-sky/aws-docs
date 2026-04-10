This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::UserSettings CookieSpecification

Specifies a single cookie or set of cookies in an end user's browser.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Domain" : String,
  "Name" : String,
  "Path" : String
}

```

### YAML

```yaml

  Domain: String
  Name: String
  Path: String

```

## Properties

`Domain`

The domain of the cookie.

_Required_: Yes

_Type_: String

_Pattern_: `^(\.?)(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)*[a-z0-9][a-z0-9-]{0,61}[a-z0-9]$`

_Minimum_: `0`

_Maximum_: `253`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the cookie.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path of the cookie.

_Required_: No

_Type_: String

_Pattern_: `^/(\S)*$`

_Minimum_: `0`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BrandingConfiguration

CookieSynchronizationConfiguration

All content copied from https://docs.aws.amazon.com/.
