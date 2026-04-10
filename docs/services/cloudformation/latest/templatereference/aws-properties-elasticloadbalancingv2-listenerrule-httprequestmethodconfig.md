This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::ListenerRule HttpRequestMethodConfig

Information about an HTTP method condition.

HTTP defines a set of request methods, also referred to as HTTP verbs. For more
information, see the [HTTP Method\
Registry](https://www.iana.org/assignments/http-methods/http-methods.xhtml). You can also define custom HTTP methods.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Values:
    - String

```

## Properties

`Values`

The name of the request method. The maximum length is 40 characters. The allowed characters
are A-Z, hyphen (-), and underscore (\_). The comparison is case sensitive. Wildcards are not
supported; therefore, the method name must be an exact match.

If you specify multiple strings, the condition is satisfied if one of the strings matches
the HTTP request method. We recommend that you route GET and HEAD requests in the same way,
because the response to a HEAD request may be cached.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpHeaderConfig

JwtValidationActionAdditionalClaim

All content copied from https://docs.aws.amazon.com/.
