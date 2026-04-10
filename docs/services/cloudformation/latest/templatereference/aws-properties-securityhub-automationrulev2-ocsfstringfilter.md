This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 OcsfStringFilter

Enables filtering of security findings based on string field values in OCSF.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldName" : String,
  "Filter" : StringFilter
}

```

### YAML

```yaml

  FieldName: String
  Filter:
    StringFilter

```

## Properties

`FieldName`

The name of the field.

_Required_: Yes

_Type_: String

_Allowed values_: `activity_name | cloud.account.name | cloud.account.uid | cloud.provider | cloud.region | compliance.assessments.category | compliance.assessments.name | compliance.control | compliance.status | compliance.standards | finding_info.desc | finding_info.src_url | finding_info.title | finding_info.types | finding_info.uid | finding_info.related_events.uid | finding_info.related_events.product.uid | finding_info.related_events.title | metadata.product.feature.uid | metadata.product.name | metadata.product.uid | metadata.product.vendor_name | remediation.desc | remediation.references | resources.cloud_partition | resources.name | resources.region | resources.type | resources.uid | vulnerabilities.fix_coverage | class_name | vendor_attributes.severity`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

Enables filtering of security findings based on string field values in OCSF.

_Required_: Yes

_Type_: [StringFilter](aws-properties-securityhub-automationrulev2-stringfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OcsfNumberFilter

StringFilter

All content copied from https://docs.aws.amazon.com/.
