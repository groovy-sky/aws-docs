This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::Service

Creates a service. A service is any software application that can run on instances
containers, or serverless functions within an account or virtual private cloud (VPC).

For more information, see [Services](https://docs.aws.amazon.com/vpc-lattice/latest/ug/services.html) in the
_Amazon VPC Lattice User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::Service",
  "Properties" : {
      "AuthType" : String,
      "CertificateArn" : String,
      "CustomDomainName" : String,
      "DnsEntry" : DnsEntry,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::Service
Properties:
  AuthType: String
  CertificateArn: String
  CustomDomainName: String
  DnsEntry:
    DnsEntry
  Name: String
  Tags:
    - Tag

```

## Properties

`AuthType`

The type of IAM policy.

- `NONE`: The resource does not use an IAM policy. This is the default.

- `AWS_IAM`: The resource uses an IAM policy. When this type is used, auth is enabled and an auth policy is required.

_Required_: No

_Type_: String

_Allowed values_: `NONE | AWS_IAM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateArn`

The Amazon Resource Name (ARN) of the certificate.

_Required_: No

_Type_: String

_Pattern_: `^(arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:certificate/[0-9a-z-]+)?$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomDomainName`

The custom domain name of the service.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DnsEntry`

Describes the DNS information of the service. This field is read-only.

_Required_: No

_Type_: [DnsEntry](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-vpclattice-service-dnsentry.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the service. The name must be unique within the account. The valid characters
are a-z, 0-9, and hyphens (-). You can't use a hyphen as the first or last character, or
immediately after another hyphen.

If you don't specify a name, CloudFormation generates one. However, if
you specify a name, and later want to replace the resource, you must specify a new
name.

_Required_: No

_Type_: String

_Pattern_: `^(?!svc-)(?![-])(?!.*[-]$)(?!.*[-]{2})[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the service.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-vpclattice-service-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the service.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the service.

`CreatedAt`

The date and time that the service was created, specified in ISO-8601 format.

`DnsEntry.DomainName`

The domain name of the service.

`DnsEntry.HostedZoneId`

The ID of the hosted zone.

`Id`

The ID of the service.

`LastUpdatedAt`

The date and time that the service was last updated, specified in ISO-8601 format.

`Status`

The status of the service.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WeightedTargetGroup

DnsEntry
