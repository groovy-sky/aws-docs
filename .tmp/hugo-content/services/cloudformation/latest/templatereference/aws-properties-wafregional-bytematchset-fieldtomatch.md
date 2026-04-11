This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::ByteMatchSet FieldToMatch

###### Note

AWS WAF Classic support will end on September 30, 2025.

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

Specifies where in a web request to look for `TargetString`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Data" : String,
  "Type" : String
}

```

### YAML

```yaml

  Data: String
  Type: String

```

## Properties

`Data`

When the value of `Type` is `HEADER`, enter the name of the header that you want AWS WAF to search,
for example, `User-Agent` or `Referer`. The name of the header is not case sensitive.

When the value of `Type` is `SINGLE_QUERY_ARG`, enter the name of the parameter that you want AWS WAF to search,
for example, `UserName` or `SalesRegion`. The parameter name is not case sensitive.

If the value of `Type` is any other value, omit `Data`.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The part of the web request that you want AWS WAF to search for a specified string. Parts of a request that you can search include the following:

- `HEADER`: A specified request header, for example, the value of the `User-Agent` or `Referer` header.
If you choose `HEADER` for the type, specify the name of the header in `Data`.

- `METHOD`: The HTTP method, which indicated the type of operation that the request is asking the origin to perform.

- `QUERY_STRING`: A query string, which is the part of a URL that appears after a `?` character, if any.

- `URI`: The part of a web request that identifies a resource, for example, `/images/daily-ad.jpg`.

- `BODY`: The part of a request that contains any additional data that you want to send to your web server
as the HTTP request body, such as data from a form. The request body immediately follows the request headers.
Note that only the first `8192` bytes of the request body are forwarded to AWS WAF for inspection.
To allow or block requests based on the length of the body, you can create a size constraint set.

- `SINGLE_QUERY_ARG`: The parameter in the query string that you will inspect, such as _UserName_ or _SalesRegion_. The maximum length for `SINGLE_QUERY_ARG` is 30 characters.

- `ALL_QUERY_ARGS`: Similar to `SINGLE_QUERY_ARG`, but rather than inspecting a single parameter, AWS WAF will inspect all parameters within the query for the value or regex pattern that you specify in
`TargetString`.

_Required_: Yes

_Type_: String

_Allowed values_: `URI | QUERY_STRING | HEADER | METHOD | BODY | SINGLE_QUERY_ARG | ALL_QUERY_ARGS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ByteMatchTuple

AWS::WAFRegional::GeoMatchSet

All content copied from https://docs.aws.amazon.com/.
