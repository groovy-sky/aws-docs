This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::DistributionTenant Customizations

Customizations for the distribution tenant. For each distribution tenant, you can specify the geographic restrictions, and the Amazon Resource Names (ARNs) for the ACM certificate and AWS WAF web ACL. These are specific values that you can override or disable from the multi-tenant distribution that was used to create the distribution tenant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Certificate" : Certificate,
  "GeoRestrictions" : GeoRestrictionCustomization,
  "WebAcl" : WebAclCustomization
}

```

### YAML

```yaml

  Certificate:
    Certificate
  GeoRestrictions:
    GeoRestrictionCustomization
  WebAcl:
    WebAclCustomization

```

## Properties

`Certificate`

The AWS Certificate Manager (ACM) certificate.

_Required_: No

_Type_: [Certificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-distributiontenant-certificate.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeoRestrictions`

The geographic restrictions.

_Required_: No

_Type_: [GeoRestrictionCustomization](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-distributiontenant-georestrictioncustomization.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebAcl`

The AWS WAF web ACL.

_Required_: No

_Type_: [WebAclCustomization](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-distributiontenant-webaclcustomization.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Certificate

DomainResult
