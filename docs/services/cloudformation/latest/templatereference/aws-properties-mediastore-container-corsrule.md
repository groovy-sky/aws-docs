This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaStore::Container CorsRule

A rule for a CORS policy. You can add up to 100 rules to a CORS policy. If more than
one rule applies, the service uses the first applicable rule listed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedHeaders" : [ String, ... ],
  "AllowedMethods" : [ String, ... ],
  "AllowedOrigins" : [ String, ... ],
  "ExposeHeaders" : [ String, ... ],
  "MaxAgeSeconds" : Integer
}

```

### YAML

```yaml

  AllowedHeaders:
    - String
  AllowedMethods:
    - String
  AllowedOrigins:
    - String
  ExposeHeaders:
    - String
  MaxAgeSeconds: Integer

```

## Properties

`AllowedHeaders`

Specifies which headers are allowed in a preflight `OPTIONS` request
through the `Access-Control-Request-Headers` header. Each header name that is
specified in `Access-Control-Request-Headers` must have a corresponding entry in
the rule. Only the headers that were requested are sent back.

This element can contain only one wildcard character (\*).

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedMethods`

Identifies an HTTP method that the origin that is specified in the rule is allowed to
execute.

Each CORS rule must contain at least one `AllowedMethods` and one
`AllowedOrigins` element.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedOrigins`

One or more response headers that you want users to be able to access from their
applications (for example, from a JavaScript `XMLHttpRequest` object).

Each CORS rule must have at least one `AllowedOrigins` element. The string
value can include only one wildcard character (\*), for example, http://\*.example.com.
Additionally, you can specify only one wildcard character to allow cross-origin access for
all origins.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExposeHeaders`

One or more headers in the response that you want users to be able to access from
their applications (for example, from a JavaScript `XMLHttpRequest`
object).

This element is optional for each rule.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxAgeSeconds`

The time in seconds that your browser caches the preflight response for the specified
resource.

A CORS rule can have only one `MaxAgeSeconds` element.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `2147483647`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaStore::Container

MetricPolicy

All content copied from https://docs.aws.amazon.com/.
