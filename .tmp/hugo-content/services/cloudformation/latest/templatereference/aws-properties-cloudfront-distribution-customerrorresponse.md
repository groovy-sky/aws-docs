This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution CustomErrorResponse

A complex type that controls:

- Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom
error messages before returning the response to the viewer.

- How long CloudFront caches HTTP status codes in the 4xx and 5xx range.

For more information about custom error pages, see [Customizing\
Error Responses](../../../amazoncloudfront/latest/developerguide/custom-error-pages.md) in the _Amazon CloudFront Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorCachingMinTTL" : Number,
  "ErrorCode" : Integer,
  "ResponseCode" : Integer,
  "ResponsePagePath" : String
}

```

### YAML

```yaml

  ErrorCachingMinTTL: Number
  ErrorCode: Integer
  ResponseCode: Integer
  ResponsePagePath: String

```

## Properties

`ErrorCachingMinTTL`

The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status
code specified in `ErrorCode`. When this time period has elapsed, CloudFront
queries your origin to see whether the problem that caused the error has been resolved
and the requested object is now available.

For more information, see [Customizing\
Error Responses](../../../amazoncloudfront/latest/developerguide/custom-error-pages.md) in the _Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorCode`

The HTTP status code for which you want to specify a custom error page and/or a
caching duration.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseCode`

The HTTP status code that you want CloudFront to return to the viewer along with the custom
error page. There are a variety of reasons that you might want CloudFront to return a status
code different from the status code that your origin returned to CloudFront, for
example:

- Some Internet devices (some firewalls and corporate proxies, for example)
intercept HTTP 4xx and 5xx and prevent the response from being returned to the
viewer. If you substitute `200`, the response typically won't be
intercepted.

- If you don't care about distinguishing among different client errors or server
errors, you can specify `400` or `500` as the
`ResponseCode` for all 4xx or 5xx errors.

- You might want to return a `200` status code (OK) and static
website so your customers don't know that your website is down.

If you specify a value for `ResponseCode`, you must also specify a value
for `ResponsePagePath`.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponsePagePath`

The path to the custom error page that you want CloudFront to return to a viewer when your
origin returns the HTTP status code specified by `ErrorCode`, for example,
`/4xx-errors/403-forbidden.html`. If you want to store your objects and
your custom error pages in different locations, your distribution must include a cache
behavior for which the following is true:

- The value of `PathPattern` matches the path to your custom error
messages. For example, suppose you saved custom error pages for 4xx errors in an
Amazon S3 bucket in a directory named `/4xx-errors`. Your distribution
must include a cache behavior for which the path pattern routes requests for
your custom error pages to that location, for example,
`/4xx-errors/*`.

- The value of `TargetOriginId` specifies the value of the
`ID` element for the origin that contains your custom error
pages.

If you specify a value for `ResponsePagePath`, you must also specify a
value for `ResponseCode`.

We recommend that you store custom error pages in an Amazon S3 bucket. If you store custom
error pages on an HTTP server and the server starts to return 5xx errors, CloudFront can't get
the files that you want to return to viewers because the origin server is
unavailable.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [CustomErrorResponse](../../../../reference/cloudfront/latest/apireference/api-customerrorresponse.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Cookies

CustomOriginConfig

All content copied from https://docs.aws.amazon.com/.
