This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution DefaultCacheBehavior

A complex type that describes the default cache behavior if you don't specify a
`CacheBehavior` element or if request URLs don't match any of the values
of `PathPattern` in `CacheBehavior` elements. You must create
exactly one default cache behavior.

###### Important

If your minimum TTL is greater than 0, CloudFront will cache content for at least the duration specified in the cache policy's minimum TTL, even if the `Cache-Control: no-cache`, `no-store`, or `private` directives are present in the origin headers.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedMethods" : [ String, ... ],
  "CachedMethods" : [ String, ... ],
  "CachePolicyId" : String,
  "Compress" : Boolean,
  "DefaultTTL" : Number,
  "FieldLevelEncryptionId" : String,
  "ForwardedValues" : ForwardedValues,
  "FunctionAssociations" : [ FunctionAssociation, ... ],
  "GrpcConfig" : GrpcConfig,
  "LambdaFunctionAssociations" : [ LambdaFunctionAssociation, ... ],
  "MaxTTL" : Number,
  "MinTTL" : Number,
  "OriginRequestPolicyId" : String,
  "RealtimeLogConfigArn" : String,
  "ResponseHeadersPolicyId" : String,
  "SmoothStreaming" : Boolean,
  "TargetOriginId" : String,
  "TrustedKeyGroups" : [ String, ... ],
  "TrustedSigners" : [ String, ... ],
  "ViewerProtocolPolicy" : String
}

```

### YAML

```yaml

  AllowedMethods:
    - String
  CachedMethods:
    - String
  CachePolicyId: String
  Compress: Boolean
  DefaultTTL: Number
  FieldLevelEncryptionId: String
  ForwardedValues:
    ForwardedValues
  FunctionAssociations:
    - FunctionAssociation
  GrpcConfig:
    GrpcConfig
  LambdaFunctionAssociations:
    - LambdaFunctionAssociation
  MaxTTL: Number
  MinTTL: Number
  OriginRequestPolicyId: String
  RealtimeLogConfigArn: String
  ResponseHeadersPolicyId: String
  SmoothStreaming: Boolean
  TargetOriginId: String
  TrustedKeyGroups:
    - String
  TrustedSigners:
    - String
  ViewerProtocolPolicy: String

```

## Properties

`AllowedMethods`

A complex type that controls which HTTP methods CloudFront processes and forwards to your
Amazon S3 bucket or your custom origin. There are three choices:

- CloudFront forwards only `GET` and `HEAD` requests.

- CloudFront forwards only `GET`, `HEAD`, and
`OPTIONS` requests.

- CloudFront forwards `GET, HEAD, OPTIONS, PUT, PATCH, POST`, and
`DELETE` requests.

If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or
to your custom origin so users can't perform operations that you don't want them to. For
example, you might not want users to have permissions to delete objects from your
origin.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CachedMethods`

A complex type that controls whether CloudFront caches the response to requests using the
specified HTTP methods. There are two choices:

- CloudFront caches responses to `GET` and `HEAD`
requests.

- CloudFront caches responses to `GET`, `HEAD`, and
`OPTIONS` requests.

If you pick the second choice for your Amazon S3 Origin, you may need to forward
Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for
the responses to be cached correctly.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CachePolicyId`

The unique identifier of the cache policy that is attached to the default cache
behavior. For more information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) or [Using the managed cache policies](../../../amazoncloudfront/latest/developerguide/using-managed-cache-policies.md) in the
_Amazon CloudFront Developer Guide_.

A `DefaultCacheBehavior` must include either a `CachePolicyId`
or `ForwardedValues`. We recommend that you use a
`CachePolicyId`.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Compress`

Whether you want CloudFront to automatically compress certain files for this cache behavior.
If so, specify `true`; if not, specify `false`. For more
information, see [Serving\
Compressed Files](../../../amazoncloudfront/latest/developerguide/servingcompressedfiles.md) in the _Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultTTL`

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

This field is deprecated. We recommend that you use the `DefaultTTL` field
in a cache policy instead of this field. For more information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) or [Using the managed cache policies](../../../amazoncloudfront/latest/developerguide/using-managed-cache-policies.md) in the
_Amazon CloudFront Developer Guide_.

The default amount of time that you want objects to stay in CloudFront caches before CloudFront
forwards another request to your origin to determine whether the object has been
updated. The value that you specify applies only when your origin does not add HTTP
headers such as `Cache-Control max-age`, `Cache-Control s-maxage`,
and `Expires` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](../../../amazoncloudfront/latest/developerguide/expiration.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldLevelEncryptionId`

The value of `ID` for the field-level encryption configuration that you
want CloudFront to use for encrypting specific fields of data for the default cache
behavior.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForwardedValues`

This field is deprecated. We recommend that you use a cache policy or an origin
request policy instead of this field. For more information, see [Working with policies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/working-with-policies.html) in the
_Amazon CloudFront Developer Guide_.

If you want to include values in the cache key, use a cache policy. For more
information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) or [Using the managed cache policies](../../../amazoncloudfront/latest/developerguide/using-managed-cache-policies.md) in the
_Amazon CloudFront Developer Guide_.

If you want to send values to the origin but not include them in the cache key, use an
origin request policy. For more information, see [Creating origin request policies](../../../amazoncloudfront/latest/developerguide/controlling-origin-requests.md#origin-request-create-origin-request-policy) or [Using the managed origin request policies](../../../amazoncloudfront/latest/developerguide/using-managed-origin-request-policies.md) in the
_Amazon CloudFront Developer Guide_.

A `DefaultCacheBehavior` must include either a `CachePolicyId`
or `ForwardedValues`. We recommend that you use a
`CachePolicyId`.

A complex type that specifies how CloudFront handles query strings, cookies, and HTTP
headers.

_Required_: Conditional

_Type_: [ForwardedValues](aws-properties-cloudfront-distribution-forwardedvalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionAssociations`

A list of CloudFront functions that are associated with this cache behavior. Your functions
must be published to the `LIVE` stage to associate them with a cache
behavior.

_Required_: No

_Type_: Array of [FunctionAssociation](aws-properties-cloudfront-distribution-functionassociation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GrpcConfig`

The gRPC configuration for your cache behavior.

_Required_: No

_Type_: [GrpcConfig](aws-properties-cloudfront-distribution-grpcconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionAssociations`

A complex type that contains zero or more Lambda@Edge function associations for a
cache behavior.

_Required_: No

_Type_: Array of [LambdaFunctionAssociation](aws-properties-cloudfront-distribution-lambdafunctionassociation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxTTL`

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

This field is deprecated. We recommend that you use the `MaxTTL` field in a
cache policy instead of this field. For more information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) or [Using the managed cache policies](../../../amazoncloudfront/latest/developerguide/using-managed-cache-policies.md) in the
_Amazon CloudFront Developer Guide_.

The maximum amount of time that you want objects to stay in CloudFront caches before CloudFront
forwards another request to your origin to determine whether the object has been
updated. The value that you specify applies only when your origin adds HTTP headers such
as `Cache-Control max-age`, `Cache-Control s-maxage`, and
`Expires` to objects. For more information, see [Managing How Long Content Stays in an Edge Cache (Expiration)](../../../amazoncloudfront/latest/developerguide/expiration.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinTTL`

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

This field is deprecated. We recommend that you use the `MinTTL` field in a
cache policy instead of this field. For more information, see [Creating cache policies](../../../amazoncloudfront/latest/developerguide/controlling-the-cache-key.md#cache-key-create-cache-policy) or [Using the managed cache policies](../../../amazoncloudfront/latest/developerguide/using-managed-cache-policies.md) in the
_Amazon CloudFront Developer Guide_.

The minimum amount of time that you want objects to stay in CloudFront caches before CloudFront
forwards another request to your origin to determine whether the object has been
updated. For more information, see [Managing How Long\
Content Stays in an Edge Cache (Expiration)](../../../amazoncloudfront/latest/developerguide/expiration.md) in the
_Amazon CloudFront Developer Guide_.

You must specify `0` for `MinTTL` if you configure CloudFront to
forward all headers to your origin (under `Headers`, if you specify
`1` for `Quantity` and `*` for
`Name`).

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginRequestPolicyId`

The unique identifier of the origin request policy that is attached to the default
cache behavior. For more information, see [Creating origin request policies](../../../amazoncloudfront/latest/developerguide/controlling-origin-requests.md#origin-request-create-origin-request-policy) or [Using the managed origin request policies](../../../amazoncloudfront/latest/developerguide/using-managed-origin-request-policies.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RealtimeLogConfigArn`

The Amazon Resource Name (ARN) of the real-time log configuration that is attached to
this cache behavior. For more information, see [Real-time logs](../../../amazoncloudfront/latest/developerguide/real-time-logs.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResponseHeadersPolicyId`

The identifier for a response headers policy.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SmoothStreaming`

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

Indicates whether you want to distribute media files in the Microsoft Smooth Streaming
format using the origin that is associated with this cache behavior. If so, specify
`true`; if not, specify `false`. If you specify
`true` for `SmoothStreaming`, you can still distribute other
content using this cache behavior if the content matches the value of
`PathPattern`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetOriginId`

The value of `ID` for the origin that you want CloudFront to route requests to
when they use the default cache behavior.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustedKeyGroups`

A list of key groups that CloudFront can use to validate signed URLs or signed
cookies.

When a cache behavior contains trusted key groups, CloudFront requires signed URLs or signed
cookies for all requests that match the cache behavior. The URLs or cookies must be
signed with a private key whose corresponding public key is in the key group. The signed
URL or cookie contains information about which public key CloudFront should use to verify the
signature. For more information, see [Serving private content](../../../amazoncloudfront/latest/developerguide/privatecontent.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrustedSigners`

###### Important

We recommend using `TrustedKeyGroups` instead of
`TrustedSigners`.

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

A list of AWS account IDs whose public keys CloudFront can use to validate signed URLs or
signed cookies.

When a cache behavior contains trusted signers, CloudFront requires signed URLs or signed
cookies for all requests that match the cache behavior. The URLs or cookies must be
signed with the private key of a CloudFront key pair in a trusted signer's AWS account. The
signed URL or cookie contains information about which public key CloudFront should use to
verify the signature. For more information, see [Serving private content](../../../amazoncloudfront/latest/developerguide/privatecontent.md) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewerProtocolPolicy`

The protocol that viewers can use to access the files in the origin specified by
`TargetOriginId` when a request matches the path pattern in
`PathPattern`. You can specify the following options:

- `allow-all`: Viewers can use HTTP or HTTPS.

- `redirect-to-https`: If a viewer submits an HTTP request, CloudFront
returns an HTTP status code of 301 (Moved Permanently) to the viewer along with
the HTTPS URL. The viewer then resubmits the request using the new URL.

- `https-only`: If a viewer sends an HTTP request, CloudFront returns an
HTTP status code of 403 (Forbidden).

For more information about requiring the HTTPS protocol, see [Requiring HTTPS Between Viewers and CloudFront](../../../amazoncloudfront/latest/developerguide/using-https-viewers-to-cloudfront.md) in the
_Amazon CloudFront Developer Guide_.

###### Note

The only way to guarantee that viewers retrieve an object that was fetched from
the origin using HTTPS is never to use any other protocol to fetch the object. If
you have recently changed from HTTP to HTTPS, we recommend that you clear your
objects' cache because cached objects are protocol agnostic. That means that an edge
location will return an object from the cache regardless of whether the current
request protocol matches the protocol used previously. For more information, see
[Managing Cache\
Expiration](../../../amazoncloudfront/latest/developerguide/expiration.md) in the _Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `allow-all | https-only | redirect-to-https`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DefaultCacheBehavior](https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_DefaultCacheBehavior.html) in the _Amazon CloudFront API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomOriginConfig

Definition
