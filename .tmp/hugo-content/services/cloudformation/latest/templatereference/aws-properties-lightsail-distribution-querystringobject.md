This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Distribution QueryStringObject

`QueryStringObject` is a property of the [CacheSettings](../userguide/aws-properties-lightsail-distribution-cachesettings.md) property. It describes the query string parameters that an
Amazon Lightsail content delivery network (CDN) distribution to bases
caching on.

For the query strings that you specify, your distribution caches separate versions of
the specified content based on the query string values in viewer requests.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Option" : Boolean,
  "QueryStringsAllowList" : [ String, ... ]
}

```

### YAML

```yaml

  Option: Boolean
  QueryStringsAllowList:
    - String

```

## Properties

`Option`

Indicates whether the distribution forwards and caches based on query strings.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStringsAllowList`

The specific query strings that the distribution forwards to the origin.

Your distribution caches content based on the specified query strings.

If the `option` parameter is true, then your distribution forwards all query
strings, regardless of what you specify using the `QueryStringsAllowList`
parameter.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputOrigin

Tag

All content copied from https://docs.aws.amazon.com/.
