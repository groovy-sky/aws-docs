This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution S3OriginConfig

A complex type that contains information about the Amazon S3 origin. If the origin is a
custom origin or an S3 bucket that is configured as a website endpoint, use the
`CustomOriginConfig` element instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OriginAccessIdentity" : String,
  "OriginReadTimeout" : Integer
}

```

### YAML

```yaml

  OriginAccessIdentity: String
  OriginReadTimeout: Integer

```

## Properties

`OriginAccessIdentity`

###### Note

If you're using origin access control (OAC) instead of origin access identity,
specify an empty `OriginAccessIdentity` element. For more information,
see [Restricting access to an AWS](../../../amazoncloudfront/latest/developerguide/private-content-restricting-access-to-origin.md) in the
_Amazon CloudFront Developer Guide_.

The CloudFront origin access identity to associate with the origin. Use an origin access
identity to configure the origin so that viewers can _only_ access
objects in an Amazon S3 bucket through CloudFront. The format of the value is:

`origin-access-identity/cloudfront/ID-of-origin-access-identity`

The `
                            ID-of-origin-access-identity
                        ` is the value that CloudFront
returned in the `ID` element when you created the origin access
identity.

If you want viewers to be able to access objects using either the CloudFront URL or the Amazon S3
URL, specify an empty `OriginAccessIdentity` element.

To delete the origin access identity from an existing distribution, update the
distribution configuration and include an empty `OriginAccessIdentity`
element.

To replace the origin access identity, update the distribution configuration and
specify the new origin access identity.

For more information about the origin access identity, see [Serving Private\
Content through CloudFront](../../../amazoncloudfront/latest/developerguide/privatecontent.md) in the _Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginReadTimeout`

Specifies how long, in seconds, CloudFront waits for a response from the origin. This is
also known as the _origin response timeout_. The minimum timeout is 1
second, the maximum is 120 seconds, and the default (if you don't specify otherwise) is
30 seconds.

For more information, see [Response timeout](../../../amazoncloudfront/latest/developerguide/downloaddistvaluesorigin.md#DownloadDistValuesOriginResponseTimeout) in the
_Amazon CloudFront Developer Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [S3OriginConfig](../../../../reference/cloudfront/latest/apireference/api-s3originconfig.md) in the _Amazon CloudFront API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Restrictions

StatusCodes

All content copied from https://docs.aws.amazon.com/.
