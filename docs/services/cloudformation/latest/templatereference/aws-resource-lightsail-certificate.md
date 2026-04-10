This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Certificate

The `AWS::Lightsail::Certificate` resource specifies an SSL/TLS certificate
that you can use with a content delivery network (CDN) distribution and a container
service.

###### Note

For information about certificates that you can use with a load balancer, see [AWS::Lightsail::LoadBalancerTlsCertificate](../userguide/aws-resource-lightsail-loadbalancertlscertificate.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::Certificate",
  "Properties" : {
      "CertificateName" : String,
      "DomainName" : String,
      "SubjectAlternativeNames" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::Certificate
Properties:
  CertificateName: String
  DomainName: String
  SubjectAlternativeNames:
    - String
  Tags:
    - Tag

```

## Properties

`CertificateName`

The name of the certificate.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainName`

The domain name of the certificate.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

`SubjectAlternativeNames`

An array of strings that specify the alternate domains (such as `example.org`)
and subdomains (such as `blog.example.com`) of the certificate.

_Required_: No

_Type_: Array of String

_Update requires_: Updates are not supported.

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md)
in the _AWS CloudFormation User Guide_.

###### Note

The `Value` of `Tags` is optional for Lightsail resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-lightsail-certificate-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a unique identifier for this resource.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CertificateArn`

The Amazon Resource Name (ARN) of the certificate.

`Status`

The validation status of the certificate.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
