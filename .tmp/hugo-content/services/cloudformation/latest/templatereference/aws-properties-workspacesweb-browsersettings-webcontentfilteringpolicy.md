This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::BrowserSettings WebContentFilteringPolicy

The policy that specifies which URLs end users are allowed to access or which URLs or domain categories they are restricted from accessing for enhanced security.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedUrls" : [ String, ... ],
  "BlockedCategories" : [ String, ... ],
  "BlockedUrls" : [ String, ... ]
}

```

### YAML

```yaml

  AllowedUrls:
    - String
  BlockedCategories:
    - String
  BlockedUrls:
    - String

```

## Properties

`AllowedUrls`

URLs and domains that are always accessible to end users.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockedCategories`

Categories of websites that are blocked on the end user's browsers.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockedUrls`

URLs and domains that end users cannot access.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::WorkSpacesWeb::DataProtectionSettings

All content copied from https://docs.aws.amazon.com/.
