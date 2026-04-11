This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DomainName

The `AWS::AppSync::DomainName` resource creates a
`DomainNameConfig` object to configure a custom domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::DomainName",
  "Properties" : {
      "CertificateArn" : String,
      "Description" : String,
      "DomainName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::DomainName
Properties:
  CertificateArn: String
  Description: String
  DomainName: String
  Tags:
    - Tag

```

## Properties

`CertificateArn`

The Amazon Resource Name (ARN) of the certificate. This will be an AWS Certificate Manager certificate.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z-]*:acm:[a-z0-9-]*:\d{12}:certificate/[0-9A-Za-z_/-]*$`

_Minimum_: `3`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The decription for your domain name.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The domain name.

_Required_: Yes

_Type_: String

_Pattern_: `^(\*[a-z\d-]*\.)?([a-z\d-]+\.)+[a-z\d-]+$`

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A set of tags (key-value pairs) for this domain name.

_Required_: No

_Type_: Array of [Tag](aws-properties-appsync-domainname-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::AppSync::DomainName` resource to
the intrinsic `Ref` function, the function returns the domain name.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`AppSyncDomainName`

The domain name provided by AWS AppSync.

`DomainName`

The domain name.

`DomainNameArn`

The Amazon resource name (ARN) of the domain name.

`HostedZoneId`

The ID of your Amazon Route 53 hosted zone.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RelationalDatabaseConfig

Tag

All content copied from https://docs.aws.amazon.com/.
