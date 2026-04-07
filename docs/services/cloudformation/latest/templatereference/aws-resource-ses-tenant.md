This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::Tenant

Create a tenant.

_Tenants_ are logical containers that group related SES resources together.
Each tenant can have its own set of resources like email identities, configuration sets,
and templates, along with reputation metrics and sending status. This helps isolate and manage
email sending for different customers or business units within your Amazon SES API v2 account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::Tenant",
  "Properties" : {
      "ResourceAssociations" : [ ResourceAssociation, ... ],
      "Tags" : [ Tag, ... ],
      "TenantName" : String
    }
}

```

### YAML

```yaml

Type: AWS::SES::Tenant
Properties:
  ResourceAssociations:
    - ResourceAssociation
  Tags:
    - Tag
  TenantName: String

```

## Properties

`ResourceAssociations`

The list of resources to associate with the tenant.

_Required_: No

_Type_: Array of [ResourceAssociation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-tenant-resourceassociation.html)

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of objects that define the tags (keys and values) associated with the tenant.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-tenant-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TenantName`

The name of a tenant. The name can contain up to 64 alphanumeric characters, including
letters, numbers, hyphens (-) and underscores (\_) only.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-_]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref returns the resource name.

### Fn::GetAtt

`Arn`

Returns the ARN of the tenant.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Template

ResourceAssociation
