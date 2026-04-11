This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Distribution CookieObject

`CookieObject` is a property of the [CacheSettings](../userguide/aws-properties-lightsail-distribution-cachesettings.md) property. It describes whether an Amazon Lightsail
content delivery network (CDN) distribution forwards cookies to the origin and, if so,
which ones.

For the cookies that you specify, your distribution caches separate versions of the
specified content based on the cookie values in viewer requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CookiesAllowList" : [ String, ... ],
  "Option" : String
}

```

### YAML

```yaml

  CookiesAllowList:
    - String
  Option: String

```

## Properties

`CookiesAllowList`

The specific cookies to forward to your distribution's origin.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Option`

Specifies which cookies to forward to the distribution's origin for a cache
behavior.

Use one of the following configurations for your distribution:

- **`all`** \- Forwards all cookies to your origin.

- **`none`** \- Doesn’t forward cookies to your origin.

- **`allow-list`** \- Forwards only the cookies that you specify using the
`CookiesAllowList` parameter.

_Required_: No

_Type_: String

_Allowed values_: `none | allow-list | all`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CacheSettings

HeaderObject

All content copied from https://docs.aws.amazon.com/.
