This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Distribution HeaderObject

`HeaderObject` is a property of the [CacheSettings](../userguide/aws-properties-lightsail-distribution-cachesettings.md) property. It describes the request headers used by your distribution, which caches your content based on the request headers.

For the headers that you specify, your distribution caches separate versions of the
specified content based on the header values in viewer requests. For example, suppose that
viewer requests for logo.jpg contain a custom product header that
has a value of either acme or apex. Also, suppose that you configure your
distribution to cache your content based on values in the product header. Your
distribution forwards the product header to the origin and caches the response
from the origin once for each header value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeadersAllowList" : [ String, ... ],
  "Option" : String
}

```

### YAML

```yaml

  HeadersAllowList:
    - String
  Option: String

```

## Properties

`HeadersAllowList`

The specific headers to forward to your distribution's origin.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Option`

The headers that you want your distribution to forward to your origin. Your distribution caches your content based
on these headers.

Use one of the following configurations for your distribution:

- **`all`** \- Forwards all headers to your origin..

- **`none`** \- Forwards only the default headers.

- **`allow-list`** \- Forwards only the headers that you specify using the
`HeadersAllowList` parameter.

_Required_: No

_Type_: String

_Allowed values_: `none | allow-list | all`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CookieObject

InputOrigin

All content copied from https://docs.aws.amazon.com/.
