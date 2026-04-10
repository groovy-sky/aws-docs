This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::StreamingDistribution S3Origin

A complex type that contains information about the Amazon S3 bucket from which you want
CloudFront to get your media files for distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainName" : String,
  "OriginAccessIdentity" : String
}

```

### YAML

```yaml

  DomainName: String
  OriginAccessIdentity: String

```

## Properties

`DomainName`

The DNS name of the Amazon S3 origin.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginAccessIdentity`

The CloudFront origin access identity to associate with the distribution. Use an origin
access identity to configure the distribution so that end users can only access objects
in an Amazon S3 bucket through CloudFront.

If you want end users to be able to access objects using either the CloudFront URL or the
Amazon S3 URL, specify an empty `OriginAccessIdentity` element.

To delete the origin access identity from an existing distribution, update the
distribution configuration and include an empty `OriginAccessIdentity`
element.

To replace the origin access identity, update the distribution configuration and
specify the new origin access identity.

For more information, see [Using an Origin Access Identity to Restrict Access to Your Amazon S3 Content](../../../amazoncloudfront/latest/developerguide/private-content-restricting-access-to-s3.md) in
the _Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [S3Origin](../../../../reference/cloudfront/latest/apireference/api-s3origin.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging

StreamingDistributionConfig

All content copied from https://docs.aws.amazon.com/.
