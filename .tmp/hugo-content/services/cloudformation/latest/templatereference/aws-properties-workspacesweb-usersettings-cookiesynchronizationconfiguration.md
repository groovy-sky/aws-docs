This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::UserSettings CookieSynchronizationConfiguration

The configuration that specifies which cookies should be synchronized from the end
user's local browser to the remote browser.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Allowlist" : [ CookieSpecification, ... ],
  "Blocklist" : [ CookieSpecification, ... ]
}

```

### YAML

```yaml

  Allowlist:
    - CookieSpecification
  Blocklist:
    - CookieSpecification

```

## Properties

`Allowlist`

The list of cookie specifications that are allowed to be synchronized to the remote
browser.

_Required_: Yes

_Type_: Array of [CookieSpecification](aws-properties-workspacesweb-usersettings-cookiespecification.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Blocklist`

The list of cookie specifications that are blocked from being synchronized to the remote
browser.

_Required_: No

_Type_: Array of [CookieSpecification](aws-properties-workspacesweb-usersettings-cookiespecification.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CookieSpecification

ImageMetadata

All content copied from https://docs.aws.amazon.com/.
