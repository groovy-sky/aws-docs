This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::Distribution TenantConfig

###### Note

This field only supports multi-tenant distributions. You can't specify this field for standard distributions. For more information, see [Unsupported features for SaaS Manager for Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/distribution-config-options.md#unsupported-saas) in the _Amazon CloudFront Developer Guide_.

The configuration for a distribution tenant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParameterDefinitions" : [ ParameterDefinition, ... ]
}

```

### YAML

```yaml

  ParameterDefinitions:
    - ParameterDefinition

```

## Properties

`ParameterDefinitions`

The parameters that you specify for a distribution tenant.

_Required_: No

_Type_: Array of [ParameterDefinition](aws-properties-cloudfront-distribution-parameterdefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TrustStoreConfig

All content copied from https://docs.aws.amazon.com/.
