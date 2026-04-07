This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRule AutomationRulesFindingFilters

The criteria that determine which findings a rule applies to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsAccountId" : [ StringFilter, ... ],
  "CompanyName" : [ StringFilter, ... ],
  "ComplianceAssociatedStandardsId" : [ StringFilter, ... ],
  "ComplianceSecurityControlId" : [ StringFilter, ... ],
  "ComplianceStatus" : [ StringFilter, ... ],
  "Confidence" : [ NumberFilter, ... ],
  "CreatedAt" : [ DateFilter, ... ],
  "Criticality" : [ NumberFilter, ... ],
  "Description" : [ StringFilter, ... ],
  "FirstObservedAt" : [ DateFilter, ... ],
  "GeneratorId" : [ StringFilter, ... ],
  "Id" : [ StringFilter, ... ],
  "LastObservedAt" : [ DateFilter, ... ],
  "NoteText" : [ StringFilter, ... ],
  "NoteUpdatedAt" : [ DateFilter, ... ],
  "NoteUpdatedBy" : [ StringFilter, ... ],
  "ProductArn" : [ StringFilter, ... ],
  "ProductName" : [ StringFilter, ... ],
  "RecordState" : [ StringFilter, ... ],
  "RelatedFindingsId" : [ StringFilter, ... ],
  "RelatedFindingsProductArn" : [ StringFilter, ... ],
  "ResourceDetailsOther" : [ MapFilter, ... ],
  "ResourceId" : [ StringFilter, ... ],
  "ResourcePartition" : [ StringFilter, ... ],
  "ResourceRegion" : [ StringFilter, ... ],
  "ResourceTags" : [ MapFilter, ... ],
  "ResourceType" : [ StringFilter, ... ],
  "SeverityLabel" : [ StringFilter, ... ],
  "SourceUrl" : [ StringFilter, ... ],
  "Title" : [ StringFilter, ... ],
  "Type" : [ StringFilter, ... ],
  "UpdatedAt" : [ DateFilter, ... ],
  "UserDefinedFields" : [ MapFilter, ... ],
  "VerificationState" : [ StringFilter, ... ],
  "WorkflowStatus" : [ StringFilter, ... ]
}

```

### YAML

```yaml

  AwsAccountId:
    - StringFilter
  CompanyName:
    - StringFilter
  ComplianceAssociatedStandardsId:
    - StringFilter
  ComplianceSecurityControlId:
    - StringFilter
  ComplianceStatus:
    - StringFilter
  Confidence:
    - NumberFilter
  CreatedAt:
    - DateFilter
  Criticality:
    - NumberFilter
  Description:
    - StringFilter
  FirstObservedAt:
    - DateFilter
  GeneratorId:
    - StringFilter
  Id:
    - StringFilter
  LastObservedAt:
    - DateFilter
  NoteText:
    - StringFilter
  NoteUpdatedAt:
    - DateFilter
  NoteUpdatedBy:
    - StringFilter
  ProductArn:
    - StringFilter
  ProductName:
    - StringFilter
  RecordState:
    - StringFilter
  RelatedFindingsId:
    - StringFilter
  RelatedFindingsProductArn:
    - StringFilter
  ResourceDetailsOther:
    - MapFilter
  ResourceId:
    - StringFilter
  ResourcePartition:
    - StringFilter
  ResourceRegion:
    - StringFilter
  ResourceTags:
    - MapFilter
  ResourceType:
    - StringFilter
  SeverityLabel:
    - StringFilter
  SourceUrl:
    - StringFilter
  Title:
    - StringFilter
  Type:
    - StringFilter
  UpdatedAt:
    - DateFilter
  UserDefinedFields:
    - MapFilter
  VerificationState:
    - StringFilter
  WorkflowStatus:
    - StringFilter

```

## Properties

`AwsAccountId`

The AWS account ID in which a finding was generated.

Array Members: Minimum number of 1 item. Maximum number of 100 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CompanyName`

The name of the company for the product that generated the finding.
For control-based findings, the company is AWS.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceAssociatedStandardsId`

The unique identifier of a standard in which a control is enabled. This field consists of the resource portion of
the Amazon Resource Name (ARN) returned for a standard in the [DescribeStandards](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API response.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceSecurityControlId`

The security control ID for which a finding was generated. Security control IDs are the same across standards.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceStatus`

The result of a security check. This field is only used for findings generated
from controls.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Confidence`

The likelihood that a finding accurately identifies the behavior or issue that it was
intended to identify. `Confidence` is scored on a 0–100 basis using a ratio
scale. A value of `0` means 0 percent confidence, and a value of
`100` means 100 percent confidence. For example, a data exfiltration
detection based on a statistical deviation of network traffic has low confidence because an
actual exfiltration hasn't been verified. For more information, see [Confidence](https://docs.aws.amazon.com/securityhub/latest/userguide/asff-top-level-attributes.html#asff-confidence) in the _AWS Security Hub CSPM User Guide_.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreatedAt`

A timestamp that indicates when this finding record was created.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Criticality`

The level of importance that is assigned to the resources that are associated with a
finding. `Criticality` is scored on a 0–100 basis, using a ratio scale that supports
only full integers. A score of `0` means that the underlying resources have no
criticality, and a score of `100` is reserved for the most critical resources. For
more information, see [Criticality](https://docs.aws.amazon.com/securityhub/latest/userguide/asff-top-level-attributes.html#asff-criticality) in the _AWS Security Hub CSPM User Guide_.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A finding's description.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirstObservedAt`

A timestamp that indicates when the potential security issue captured by a
finding was first observed by the security findings product.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeneratorId`

The identifier for the solution-specific component that
generated a finding.

Array Members: Minimum number of 1 item. Maximum number of 100 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The product-specific identifier for a finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastObservedAt`

A timestamp that indicates when the security findings provider most recently observed a change in the resource that is involved in the finding.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteText`

The text of a user-defined note that's added to a finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteUpdatedAt`

The timestamp of when the note was updated.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteUpdatedBy`

The principal that created a note.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductArn`

The Amazon Resource Name (ARN) for a third-party product that generated a finding in
Security Hub CSPM.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductName`

Provides the name of the product that generated the finding. For
control-based findings, the product name is Security Hub CSPM.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordState`

Provides the current state of a finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelatedFindingsId`

The product-generated identifier for a related finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelatedFindingsProductArn`

The ARN for the product that generated a related finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceDetailsOther`

Custom fields and values about the resource that a finding pertains to.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [MapFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-mapfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

The identifier for the given resource type. For AWS resources that are identified by
Amazon Resource Names (ARNs), this is the ARN. For AWS resources that lack ARNs,
this is the identifier as defined by the AWS service that created the resource.
For non-AWS resources, this is a unique identifier that is associated with the
resource.

Array Members: Minimum number of 1 item. Maximum number of 100 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcePartition`

The partition in which the resource that the finding pertains to is located.
A partition is a group of AWS Regions. Each AWS account is scoped to one partition.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceRegion`

The AWS Region where the resource that a finding pertains to is located.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

A list of AWS tags associated with a resource at the time the finding was processed.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [MapFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-mapfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

A finding's title.

Array Members: Minimum number of 1 item. Maximum number of 100 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SeverityLabel`

The severity value of the finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceUrl`

Provides a URL that links to a page about the current finding in the finding product.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

A finding's title.

Array Members: Minimum number of 1 item. Maximum number of 100 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

One or more finding types in the format of namespace/category/classifier that classify
a finding. For a list of namespaces, classifiers, and categories, see [Types\
taxonomy for ASFF](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format-type-taxonomy.html) in the _AWS Security Hub CSPM User Guide_.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdatedAt`

A timestamp that indicates when the finding record was most recently updated.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserDefinedFields`

A list of user-defined name and value string pairs added to a finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [MapFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-mapfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerificationState`

Provides the veracity of a finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowStatus`

Provides information about the status of the investigation into a finding.

Array Members: Minimum number of 1 item. Maximum number of 20 items.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-automationrule-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AutomationRulesFindingFieldsUpdate

DateFilter
