This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Distribution CacheSettings

`CacheSettings` is a property of the [AWS::Lightsail::Distribution](../userguide/aws-resource-lightsail-distribution.md) resource. It describes the cache settings of an
Amazon Lightsail content delivery network (CDN) distribution.

These settings apply only to your distribution’s `CacheBehaviors` that have a `Behavior` of
`cache`. This includes the `DefaultCacheBehavior`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedHTTPMethods" : String,
  "CachedHTTPMethods" : String,
  "DefaultTTL" : Integer,
  "ForwardedCookies" : CookieObject,
  "ForwardedHeaders" : HeaderObject,
  "ForwardedQueryStrings" : QueryStringObject,
  "MaximumTTL" : Integer,
  "MinimumTTL" : Integer
}

```

### YAML

```yaml

  AllowedHTTPMethods: String
  CachedHTTPMethods: String
  DefaultTTL: Integer
  ForwardedCookies:
    CookieObject
  ForwardedHeaders:
    HeaderObject
  ForwardedQueryStrings:
    QueryStringObject
  MaximumTTL: Integer
  MinimumTTL: Integer

```

## Properties

`AllowedHTTPMethods`

The HTTP methods that are processed and forwarded to the distribution's origin.

You can specify the following options:

- `GET,HEAD` \- The distribution forwards the `GET` and
`HEAD` methods.

- `GET,HEAD,OPTIONS` \- The distribution forwards the `GET`,
`HEAD`, and `OPTIONS` methods.

- `GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE` \- The distribution forwards the
`GET`, `HEAD`, `OPTIONS`, `PUT`,
`PATCH`, `POST`, and `DELETE` methods.

If you specify `GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE`, you might need to
restrict access to your distribution's origin so users can't perform operations that you
don't want them to. For example, you might not want users to have permission to delete
objects from your origin.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CachedHTTPMethods`

The HTTP method responses that are cached by your distribution.

You can specify the following options:

- `GET,HEAD` \- The distribution caches responses to the `GET` and
`HEAD` methods.

- `GET,HEAD,OPTIONS` \- The distribution caches responses to the
`GET`, `HEAD`, and `OPTIONS` methods.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultTTL`

The default amount of time that objects stay in the distribution's cache before the
distribution forwards another request to the origin to determine whether the content has
been updated.

###### Note

The value specified applies only when the origin does not add HTTP headers such as
`Cache-Control max-age`, `Cache-Control s-maxage`, and
`Expires` to objects.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForwardedCookies`

An object that describes the cookies that are forwarded to the origin. Your content is
cached based on the cookies that are forwarded.

_Required_: No

_Type_: [CookieObject](aws-properties-lightsail-distribution-cookieobject.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForwardedHeaders`

An object that describes the headers that are forwarded to the origin. Your content is
cached based on the headers that are forwarded.

_Required_: No

_Type_: [HeaderObject](aws-properties-lightsail-distribution-headerobject.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForwardedQueryStrings`

An object that describes the query strings that are forwarded to the origin. Your
content is cached based on the query strings that are forwarded.

_Required_: No

_Type_: [QueryStringObject](aws-properties-lightsail-distribution-querystringobject.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumTTL`

The maximum amount of time that objects stay in the distribution's cache before the
distribution forwards another request to the origin to determine whether the object has
been updated.

The value specified applies only when the origin adds HTTP headers such as
`Cache-Control max-age`, `Cache-Control s-maxage`, and
`Expires` to objects.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumTTL`

The minimum amount of time that objects stay in the distribution's cache before the
distribution forwards another request to the origin to determine whether the object has
been updated.

A value of `0` must be specified for `minimumTTL` if the
distribution is configured to forward all headers to the origin.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CacheBehaviorPerPath

CookieObject

All content copied from https://docs.aws.amazon.com/.
