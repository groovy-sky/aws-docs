This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::DistributionTenant GeoRestrictionCustomization

The customizations that you specified for the distribution tenant for geographic restrictions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Locations" : [ String, ... ],
  "RestrictionType" : String
}

```

### YAML

```yaml

  Locations:
    - String
  RestrictionType: String

```

## Properties

`Locations`

The locations for geographic restrictions.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestrictionType`

The method that you want to use to restrict distribution of your content by
country:

- `none`: No geographic restriction is enabled, meaning access to content is
not restricted by client geo location.

- `blacklist`: The `Location` elements specify the
countries in which you don't want CloudFront to distribute your content.

- `whitelist`: The `Location` elements specify the
countries in which you want CloudFront to distribute your content.

_Required_: No

_Type_: String

_Allowed values_: `blacklist | whitelist | none`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DomainResult

ManagedCertificateRequest

All content copied from https://docs.aws.amazon.com/.
