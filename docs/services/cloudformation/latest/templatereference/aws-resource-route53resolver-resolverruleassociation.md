This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::ResolverRuleAssociation

In the response to an
[AssociateResolverRule](../../../../reference/route53/latest/apireference/api-route53resolver-associateresolverrule.md),
[DisassociateResolverRule](../../../../reference/route53/latest/apireference/api-route53resolver-disassociateresolverrule.md),
or
[ListResolverRuleAssociations](../../../../reference/route53/latest/apireference/api-route53resolver-listresolverruleassociations.md)
request, provides information about an association between a resolver rule and a VPC. The association determines which
DNS queries that originate in the VPC are forwarded to your network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::ResolverRuleAssociation",
  "Properties" : {
      "Name" : String,
      "ResolverRuleId" : String,
      "VPCId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::ResolverRuleAssociation
Properties:
  Name: String
  ResolverRuleId: String
  VPCId: String

```

## Properties

`Name`

The name of an association between a Resolver rule and a VPC.

The name can be up to 64 characters long and can contain letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (\_), and spaces. The name cannot consist of only numbers.

_Required_: No

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResolverRuleId`

The ID of the Resolver rule that you associated with the VPC that is specified by `VPCId`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VPCId`

The ID of the VPC that you associated the Resolver rule with.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ResolverRuleAssociationId`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Name`

The name of an association between a resolver rule and a VPC, such as `test.example.com in beta VPC`.

`ResolverRuleAssociationId`

The ID of the resolver rule association that you want to get information about,
such as `rslvr-rrassoc-97242eaf88example`.

`ResolverRuleId`

The ID of the resolver rule that you associated with the VPC that is specified by `VPCId`,
such as `rslvr-rr-5328a0899example`.

`VPCId`

The ID of the VPC that you associated the resolver rule with, such as `vpc-03cf94c75cexample`.

## See also

- [ResolverRuleAssociation](../../../../reference/route53/latest/apireference/api-route53resolver-resolverruleassociation.md)
in the _Amazon Route 53 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetAddress

Next

All content copied from https://docs.aws.amazon.com/.
