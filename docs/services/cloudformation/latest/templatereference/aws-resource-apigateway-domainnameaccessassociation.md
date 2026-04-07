This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::DomainNameAccessAssociation

The `AWS::ApiGateway::DomainNameAccessAssociation` resource creates a domain name access
association between an access association source and a private custom domain name.

Use a domain name access
association to invoke a private custom domain name while isolated from the public internet.

You can only create or delete a DomainNameAccessAssociation using CloudFormation. To reject a domain name
access association, use the AWS CLI.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::DomainNameAccessAssociation",
  "Properties" : {
      "AccessAssociationSource" : String,
      "AccessAssociationSourceType" : String,
      "DomainNameArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::DomainNameAccessAssociation
Properties:
  AccessAssociationSource: String
  AccessAssociationSourceType: String
  DomainNameArn: String
  Tags:
    - Tag

```

## Properties

`AccessAssociationSource`

The identifier of the domain name access association source. For a `VPCE`, the value is the VPC endpoint ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AccessAssociationSourceType`

The type of the domain name access association source. Only `VPCE` is currently supported.

_Required_: Yes

_Type_: String

_Allowed values_: `VPCE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainNameArn`

The ARN of the domain name.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The collection of tags. Each tag element is associated with a given resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-domainnameaccessassociation-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the domain name access association ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainNameAccessAssociationArn`

The access association ARN. For example:
`arn:aws:apigateway:us-west-2:111122223333:/accessassociations/domainname/private.example.com+abcd1234/vpcesource/vpce-abcd1234efg`.

## Examples

### Domain name access association example

The following example creates a `DomainNameAccessAssociation` resource named `MyDomainNameAccessAssociation`.

#### JSON

```json

{
    "MyDomainNameAccessAssociation": {
        "Type": "AWS::ApiGateway::DomainNameAccessAssociation",
        "Properties": {
            "DomainNameArn": {
                "Fn::GetAtt": [
                    "MyDomainName",
                    "DomainNameArn"
                ]
            },
            "AccessAssociationSource": "vpce-abcd123456",
            "AccessAssociationSourceType": "VPCE"
        }
    }
}
```

#### YAML

```yaml

MyDomainNameAccessAssociation:
  Type: AWS::ApiGateway::DomainNameAccessAssociation
  Properties:
    DomainNameArn: !GetAtt MyDomainName.DomainNameArn
    AccessAssociationSource: vpce-abcd123456
    AccessAssociationSourceType: VPCE
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
