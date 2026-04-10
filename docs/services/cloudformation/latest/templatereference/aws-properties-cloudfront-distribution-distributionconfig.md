This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution DistributionConfig

A distribution configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aliases" : [ String, ... ],
  "AnycastIpListId" : String,
  "CacheBehaviors" : [ CacheBehavior, ... ],
  "CNAMEs" : [ String, ... ],
  "Comment" : String,
  "ConnectionFunctionAssociation" : ConnectionFunctionAssociation,
  "ConnectionMode" : String,
  "ContinuousDeploymentPolicyId" : String,
  "CustomErrorResponses" : [ CustomErrorResponse, ... ],
  "CustomOrigin" : LegacyCustomOrigin,
  "DefaultCacheBehavior" : DefaultCacheBehavior,
  "DefaultRootObject" : String,
  "Enabled" : Boolean,
  "HttpVersion" : String,
  "IPV6Enabled" : Boolean,
  "Logging" : Logging,
  "OriginGroups" : OriginGroups,
  "Origins" : [ Origin, ... ],
  "PriceClass" : String,
  "Restrictions" : Restrictions,
  "S3Origin" : LegacyS3Origin,
  "Staging" : Boolean,
  "TenantConfig" : TenantConfig,
  "ViewerCertificate" : ViewerCertificate,
  "ViewerMtlsConfig" : ViewerMtlsConfig,
  "WebACLId" : String
}

```

### YAML

```yaml

  Aliases:
    - String
  AnycastIpListId: String
  CacheBehaviors:
    - CacheBehavior
  CNAMEs:
    - String
  Comment: String
  ConnectionFunctionAssociation:
    ConnectionFunctionAssociation
  ConnectionMode: String
  ContinuousDeploymentPolicyId: String
  CustomErrorResponses:
    - CustomErrorResponse
  CustomOrigin:
    LegacyCustomOrigin
  DefaultCacheBehavior:
    DefaultCacheBehavior
  DefaultRootObject: String
  Enabled: Boolean
  HttpVersion: String
  IPV6Enabled: Boolean
  Logging:
    Logging
  OriginGroups:
    OriginGroups
  Origins:
    - Origin
  PriceClass: String
  Restrictions:
    Restrictions
  S3Origin:
    LegacyS3Origin
  Staging: Boolean
  TenantConfig:
    TenantConfig
  ViewerCertificate:
    ViewerCertificate
  ViewerMtlsConfig:
    ViewerMtlsConfig
  WebACLId: String

```

## Properties

`Aliases`

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

A complex type that contains information about CNAMEs (alternate domain names), if
any, for this distribution.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AnycastIpListId`

###### Note

To use this field for a multi-tenant distribution, use a connection group instead. For more information, see [ConnectionGroup](../../../../reference/cloudfront/latest/apireference/api-connectiongroup.md).

ID of the Anycast static IP list that is associated with the distribution.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheBehaviors`

A complex type that contains zero or more `CacheBehavior` elements.

_Required_: No

_Type_: Array of [CacheBehavior](aws-properties-cloudfront-distribution-cachebehavior.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CNAMEs`

An alias for the CloudFront distribution's domain name.

###### Note

This property is legacy. We recommend that you use [Aliases](../userguide/aws-properties-cloudfront-distribution-distributionconfig.md#cfn-cloudfront-distribution-distributionconfig-aliases) instead.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Comment`

A comment to describe the distribution. The comment cannot be longer than
128 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionFunctionAssociation`

The distribution's connection function association.

_Required_: No

_Type_: [ConnectionFunctionAssociation](aws-properties-cloudfront-distribution-connectionfunctionassociation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionMode`

This field specifies whether the connection mode is through a standard distribution (direct) or a multi-tenant distribution with distribution tenants (tenant-only).

_Required_: No

_Type_: String

_Allowed values_: `direct | tenant-only`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContinuousDeploymentPolicyId`

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

The identifier of a continuous deployment policy. For more information, see
`CreateContinuousDeploymentPolicy`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomErrorResponses`

A complex type that controls the following:

- Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom
error messages before returning the response to the viewer.

- How long CloudFront caches HTTP status codes in the 4xx and 5xx range.

For more information about custom error pages, see [Customizing\
Error Responses](../../../amazoncloudfront/latest/developerguide/custom-error-pages.md) in the _Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Array of [CustomErrorResponse](aws-properties-cloudfront-distribution-customerrorresponse.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomOrigin`

The user-defined HTTP server that serves as the origin for content that CloudFront distributes.

###### Note

This property is legacy. We recommend that you use [Origin](../userguide/aws-properties-cloudfront-distribution-origin.md) instead.

_Required_: No

_Type_: [LegacyCustomOrigin](aws-properties-cloudfront-distribution-legacycustomorigin.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultCacheBehavior`

A complex type that describes the default cache behavior if you don't specify a
`CacheBehavior` element or if files don't match any of the values of
`PathPattern` in `CacheBehavior` elements. You must create
exactly one default cache behavior.

_Required_: Yes

_Type_: [DefaultCacheBehavior](aws-properties-cloudfront-distribution-defaultcachebehavior.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultRootObject`

When a viewer requests the root URL for your distribution, the default root object is the
object that you want CloudFront to request from your origin. For example, if your root URL is
`https://www.example.com`, you can specify CloudFront to return the
`index.html` file as the default root object. You can specify a default
root object so that viewers see a specific file or object, instead of another object in
your distribution (for example,
`https://www.example.com/product-description.html`). A default root
object avoids exposing the contents of your distribution.

You can specify the object name or a path to the object name (for example,
`index.html` or `exampleFolderName/index.html`). Your string
can't begin with a forward slash ( `/`). Only specify the object name or the
path to the object.

If you don't want to specify a default root object when you create a distribution,
include an empty `DefaultRootObject` element.

To delete the default root object from an existing distribution, update the
distribution configuration and include an empty `DefaultRootObject`
element.

To replace the default root object, update the distribution configuration and specify
the new object.

For more information about the default root object, see [Specify a default root object](../../../amazoncloudfront/latest/developerguide/defaultrootobject.md) in the _Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

From this field, you can enable or disable the selected distribution.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HttpVersion`

(Optional) Specify the HTTP version(s) that you want viewers to use to
communicate with CloudFront. The default value for new distributions is
`http1.1`.

For viewers and CloudFront to use HTTP/2, viewers must support TLSv1.2 or later, and
must support Server Name Indication (SNI).

For viewers and CloudFront to use HTTP/3, viewers must support TLSv1.3 and Server
Name Indication (SNI). CloudFront supports HTTP/3 connection migration to allow the
viewer to switch networks without losing connection. For more information
about connection migration, see [Connection Migration](https://www.rfc-editor.org/rfc/rfc9000.html) at RFC 9000. For more
information about supported TLSv1.3 ciphers, see [Supported protocols and ciphers between viewers and\
CloudFront](../../../amazoncloudfront/latest/developerguide/secure-connections-supported-viewer-protocols-ciphers.md).

_Required_: No

_Type_: String

_Allowed values_: `http1.1 | http2 | http3 | http2and3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IPV6Enabled`

###### Note

To use this field for a multi-tenant distribution, use a connection group instead. For more information, see [ConnectionGroup](../../../../reference/cloudfront/latest/apireference/api-connectiongroup.md).

If you want CloudFront to respond to IPv6 DNS requests with an IPv6 address for your
distribution, specify `true`. If you specify `false`, CloudFront
responds to IPv6 DNS requests with the DNS response code `NOERROR` and with
no IP addresses. This allows viewers to submit a second request, for an IPv4 address for
your distribution.

In general, you should enable IPv6 if you have users on IPv6 networks who want to
access your content. However, if you're using signed URLs or signed cookies to restrict
access to your content, and if you're using a custom policy that includes the
`IpAddress` parameter to restrict the IP addresses that can access your
content, don't enable IPv6. If you want to restrict access to some content by IP address
and not restrict access to other content (or restrict access but not by IP address), you
can create two distributions. For more information, see [Creating a Signed URL Using a Custom Policy](../../../amazoncloudfront/latest/developerguide/private-content-creating-signed-url-custom-policy.md) in the
_Amazon CloudFront Developer Guide_.

If you're using an Amazon Route 53 AWS Integration alias resource record set to route traffic to your CloudFront
distribution, you need to create a second alias resource record set when both of the
following are true:

- You enable IPv6 for the distribution

- You're using alternate domain names in the URLs for your objects

For more information, see [Routing\
Traffic to an Amazon CloudFront Web Distribution by Using Your Domain Name](../../../route53/latest/developerguide/routing-to-cloudfront-distribution.md) in the
_Amazon Route 53 AWS Integration Developer Guide_.

If you created a CNAME resource record set, either with Amazon Route 53 AWS Integration or with another DNS
service, you don't need to make any changes. A CNAME record will route traffic to your
distribution regardless of the IP address format of the viewer request.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Logging`

A complex type that controls whether access logs are written for the
distribution.

For more information about logging, see [Access Logs](../../../amazoncloudfront/latest/developerguide/accesslogs.md) in
the _Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: [Logging](aws-properties-cloudfront-distribution-logging.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginGroups`

A complex type that contains information about origin groups for this distribution.

Specify a value for either the `Origins` or `OriginGroups` property.

_Required_: Conditional

_Type_: [OriginGroups](aws-properties-cloudfront-distribution-origingroups.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Origins`

A complex type that contains information about origins for this distribution.

Specify a value for either the `Origins` or `OriginGroups` property.

_Required_: Conditional

_Type_: Array of [Origin](aws-properties-cloudfront-distribution-origin.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PriceClass`

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

The price class that corresponds with the maximum price that you want to pay for CloudFront
service. If you specify `PriceClass_All`, CloudFront responds to requests for your
objects from all CloudFront edge locations.

If you specify a price class other than `PriceClass_All`, CloudFront serves your
objects from the CloudFront edge location that has the lowest latency among the edge locations
in your price class. Viewers who are in or near regions that are excluded from your
specified price class may encounter slower performance.

For more information about price classes, see [Choosing the Price\
Class for a CloudFront Distribution](../../../amazoncloudfront/latest/developerguide/priceclass.md) in the _Amazon CloudFront Developer Guide_.
For information about CloudFront pricing, including how price classes (such as Price Class
100) map to CloudFront regions, see [Amazon CloudFront\
Pricing](https://aws.amazon.com/cloudfront/pricing).

_Required_: No

_Type_: String

_Allowed values_: `PriceClass_100 | PriceClass_200 | PriceClass_All | None`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Restrictions`

A complex type that identifies ways in which you want to restrict distribution of your
content.

_Required_: No

_Type_: [Restrictions](aws-properties-cloudfront-distribution-restrictions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Origin`

The origin as an Amazon S3 bucket.

###### Note

This property is legacy. We recommend that you use [Origin](../userguide/aws-properties-cloudfront-distribution-origin.md) instead.

_Required_: No

_Type_: [LegacyS3Origin](aws-properties-cloudfront-distribution-legacys3origin.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Staging`

###### Note

This field only supports standard distributions. You can't specify this field for multi-tenant distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

A Boolean that indicates whether this is a staging distribution. When this value is
`true`, this is a staging distribution. When this value is
`false`, this is not a staging distribution.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TenantConfig`

###### Note

This field only supports multi-tenant distributions. You can't specify this field for standard distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

A distribution tenant configuration.

_Required_: No

_Type_: [TenantConfig](aws-properties-cloudfront-distribution-tenantconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewerCertificate`

A complex type that determines the distribution's SSL/TLS configuration for
communicating with viewers.

_Required_: No

_Type_: [ViewerCertificate](aws-properties-cloudfront-distribution-viewercertificate.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewerMtlsConfig`

The distribution's viewer mTLS configuration.

_Required_: No

_Type_: [ViewerMtlsConfig](aws-properties-cloudfront-distribution-viewermtlsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebACLId`

###### Note

Multi-tenant distributions only support AWS WAF V2 web ACLs.

A unique identifier that specifies the AWS WAF web ACL, if any, to associate with this
distribution. To specify a web ACL created using the latest version of AWS WAF, use the
ACL ARN, for example
`arn:aws:wafv2:us-east-1:123456789012:global/webacl/ExampleWebACL/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.
To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example
`a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

AWS WAF is a web application firewall that lets you monitor the HTTP and HTTPS requests
that are forwarded to CloudFront, and lets you control access to your content. Based on
conditions that you specify, such as the IP addresses that requests originate from or
the values of query strings, CloudFront responds to requests either with the requested content
or with an HTTP 403 status code (Forbidden). You can also configure CloudFront to return a
custom error page when a request is blocked. For more information about AWS WAF, see the
[AWS WAF Developer Guide](../../../waf/latest/developerguide/what-is-aws-waf.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [DistributionConfig](../../../../reference/cloudfront/latest/apireference/api-distributionconfig.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Definition

ForwardedValues

All content copied from https://docs.aws.amazon.com/.
