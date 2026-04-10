This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ResponseHeadersPolicy AccessControlAllowMethods

A list of HTTP methods that CloudFront includes as values for the
`Access-Control-Allow-Methods` HTTP response header.

For more information about the `Access-Control-Allow-Methods` HTTP response
header, see [Access-Control-Allow-Methods](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods) in the MDN Web Docs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Items" : [ String, ... ]
}

```

### YAML

```yaml

  Items:
    - String

```

## Properties

`Items`

The list of HTTP methods. Valid values are:

- `GET`

- `DELETE`

- `HEAD`

- `OPTIONS`

- `PATCH`

- `POST`

- `PUT`

- `ALL`

`ALL` is a special value that includes all of the listed HTTP
methods.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessControlAllowHeaders

AccessControlAllowOrigins

All content copied from https://docs.aws.amazon.com/.
