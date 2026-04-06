This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::DataCatalog Tag

A label that you assign to a resource. Athena resources include
workgroups, data catalogs, and capacity reservations. Each tag consists of a key and an
optional value, both of which you define. For example, you can use tags to categorize
Athena resources by purpose, owner, or environment. Use a consistent set
of tag keys to make it easier to search and filter the resources in your account. For
best practices, see [Tagging\
Best Practices](https://docs.aws.amazon.com/whitepapers/latest/tagging-best-practices/tagging-best-practices.html). Tag keys can be from 1 to 128 UTF-8 Unicode characters, and
tag values can be from 0 to 256 UTF-8 Unicode characters. Tags can use letters and
numbers representable in UTF-8, and the following characters: + - = . \_ : / @. Tag keys
and values are case-sensitive. Tag keys must be unique per resource. If you specify more
than one tag, separate them by commas.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

A tag key. The tag key length is from 1 to 128 Unicode characters in UTF-8. You can
use letters and numbers representable in UTF-8, and the following characters: + - = . \_
: / @. Tag keys are case-sensitive and must be unique per resource.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A tag value. The tag value length is from 0 to 256 Unicode characters in UTF-8. You
can use letters and numbers representable in UTF-8, and the following characters: + - =
. \_ : / @. Tag values are case-sensitive.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Athena::DataCatalog

AWS::Athena::NamedQuery
