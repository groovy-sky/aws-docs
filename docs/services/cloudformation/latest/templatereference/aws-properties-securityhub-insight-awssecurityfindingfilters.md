This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::Insight AwsSecurityFindingFilters

A collection of filters that are applied to all active findings aggregated by AWS Security Hub CSPM.

You can filter by up to ten finding attributes. For each attribute, you can provide up to
20 filter values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsAccountId" : [ StringFilter, ... ],
  "AwsAccountName" : [ StringFilter, ... ],
  "CompanyName" : [ StringFilter, ... ],
  "ComplianceAssociatedStandardsId" : [ StringFilter, ... ],
  "ComplianceSecurityControlId" : [ StringFilter, ... ],
  "ComplianceSecurityControlParametersName" : [ StringFilter, ... ],
  "ComplianceSecurityControlParametersValue" : [ StringFilter, ... ],
  "ComplianceStatus" : [ StringFilter, ... ],
  "Confidence" : [ NumberFilter, ... ],
  "CreatedAt" : [ DateFilter, ... ],
  "Criticality" : [ NumberFilter, ... ],
  "Description" : [ StringFilter, ... ],
  "FindingProviderFieldsConfidence" : [ NumberFilter, ... ],
  "FindingProviderFieldsCriticality" : [ NumberFilter, ... ],
  "FindingProviderFieldsRelatedFindingsId" : [ StringFilter, ... ],
  "FindingProviderFieldsRelatedFindingsProductArn" : [ StringFilter, ... ],
  "FindingProviderFieldsSeverityLabel" : [ StringFilter, ... ],
  "FindingProviderFieldsSeverityOriginal" : [ StringFilter, ... ],
  "FindingProviderFieldsTypes" : [ StringFilter, ... ],
  "FirstObservedAt" : [ DateFilter, ... ],
  "GeneratorId" : [ StringFilter, ... ],
  "Id" : [ StringFilter, ... ],
  "Keyword" : [ KeywordFilter, ... ],
  "LastObservedAt" : [ DateFilter, ... ],
  "MalwareName" : [ StringFilter, ... ],
  "MalwarePath" : [ StringFilter, ... ],
  "MalwareState" : [ StringFilter, ... ],
  "MalwareType" : [ StringFilter, ... ],
  "NetworkDestinationDomain" : [ StringFilter, ... ],
  "NetworkDestinationIpV4" : [ IpFilter, ... ],
  "NetworkDestinationIpV6" : [ IpFilter, ... ],
  "NetworkDestinationPort" : [ NumberFilter, ... ],
  "NetworkDirection" : [ StringFilter, ... ],
  "NetworkProtocol" : [ StringFilter, ... ],
  "NetworkSourceDomain" : [ StringFilter, ... ],
  "NetworkSourceIpV4" : [ IpFilter, ... ],
  "NetworkSourceIpV6" : [ IpFilter, ... ],
  "NetworkSourceMac" : [ StringFilter, ... ],
  "NetworkSourcePort" : [ NumberFilter, ... ],
  "NoteText" : [ StringFilter, ... ],
  "NoteUpdatedAt" : [ DateFilter, ... ],
  "NoteUpdatedBy" : [ StringFilter, ... ],
  "ProcessLaunchedAt" : [ DateFilter, ... ],
  "ProcessName" : [ StringFilter, ... ],
  "ProcessParentPid" : [ NumberFilter, ... ],
  "ProcessPath" : [ StringFilter, ... ],
  "ProcessPid" : [ NumberFilter, ... ],
  "ProcessTerminatedAt" : [ DateFilter, ... ],
  "ProductArn" : [ StringFilter, ... ],
  "ProductFields" : [ MapFilter, ... ],
  "ProductName" : [ StringFilter, ... ],
  "RecommendationText" : [ StringFilter, ... ],
  "RecordState" : [ StringFilter, ... ],
  "Region" : [ StringFilter, ... ],
  "RelatedFindingsId" : [ StringFilter, ... ],
  "RelatedFindingsProductArn" : [ StringFilter, ... ],
  "ResourceApplicationArn" : [ StringFilter, ... ],
  "ResourceApplicationName" : [ StringFilter, ... ],
  "ResourceAwsEc2InstanceIamInstanceProfileArn" : [ StringFilter, ... ],
  "ResourceAwsEc2InstanceImageId" : [ StringFilter, ... ],
  "ResourceAwsEc2InstanceIpV4Addresses" : [ IpFilter, ... ],
  "ResourceAwsEc2InstanceIpV6Addresses" : [ IpFilter, ... ],
  "ResourceAwsEc2InstanceKeyName" : [ StringFilter, ... ],
  "ResourceAwsEc2InstanceLaunchedAt" : [ DateFilter, ... ],
  "ResourceAwsEc2InstanceSubnetId" : [ StringFilter, ... ],
  "ResourceAwsEc2InstanceType" : [ StringFilter, ... ],
  "ResourceAwsEc2InstanceVpcId" : [ StringFilter, ... ],
  "ResourceAwsIamAccessKeyCreatedAt" : [ DateFilter, ... ],
  "ResourceAwsIamAccessKeyPrincipalName" : [ StringFilter, ... ],
  "ResourceAwsIamAccessKeyStatus" : [ StringFilter, ... ],
  "ResourceAwsIamAccessKeyUserName" : [ StringFilter, ... ],
  "ResourceAwsIamUserUserName" : [ StringFilter, ... ],
  "ResourceAwsS3BucketOwnerId" : [ StringFilter, ... ],
  "ResourceAwsS3BucketOwnerName" : [ StringFilter, ... ],
  "ResourceContainerImageId" : [ StringFilter, ... ],
  "ResourceContainerImageName" : [ StringFilter, ... ],
  "ResourceContainerLaunchedAt" : [ DateFilter, ... ],
  "ResourceContainerName" : [ StringFilter, ... ],
  "ResourceDetailsOther" : [ MapFilter, ... ],
  "ResourceId" : [ StringFilter, ... ],
  "ResourcePartition" : [ StringFilter, ... ],
  "ResourceRegion" : [ StringFilter, ... ],
  "ResourceTags" : [ MapFilter, ... ],
  "ResourceType" : [ StringFilter, ... ],
  "Sample" : [ BooleanFilter, ... ],
  "SeverityLabel" : [ StringFilter, ... ],
  "SeverityNormalized" : [ NumberFilter, ... ],
  "SeverityProduct" : [ NumberFilter, ... ],
  "SourceUrl" : [ StringFilter, ... ],
  "ThreatIntelIndicatorCategory" : [ StringFilter, ... ],
  "ThreatIntelIndicatorLastObservedAt" : [ DateFilter, ... ],
  "ThreatIntelIndicatorSource" : [ StringFilter, ... ],
  "ThreatIntelIndicatorSourceUrl" : [ StringFilter, ... ],
  "ThreatIntelIndicatorType" : [ StringFilter, ... ],
  "ThreatIntelIndicatorValue" : [ StringFilter, ... ],
  "Title" : [ StringFilter, ... ],
  "Type" : [ StringFilter, ... ],
  "UpdatedAt" : [ DateFilter, ... ],
  "UserDefinedFields" : [ MapFilter, ... ],
  "VerificationState" : [ StringFilter, ... ],
  "VulnerabilitiesExploitAvailable" : [ StringFilter, ... ],
  "VulnerabilitiesFixAvailable" : [ StringFilter, ... ],
  "WorkflowState" : [ StringFilter, ... ],
  "WorkflowStatus" : [ StringFilter, ... ]
}

```

### YAML

```yaml

  AwsAccountId:
    - StringFilter
  AwsAccountName:
    - StringFilter
  CompanyName:
    - StringFilter
  ComplianceAssociatedStandardsId:
    - StringFilter
  ComplianceSecurityControlId:
    - StringFilter
  ComplianceSecurityControlParametersName:
    - StringFilter
  ComplianceSecurityControlParametersValue:
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
  FindingProviderFieldsConfidence:
    - NumberFilter
  FindingProviderFieldsCriticality:
    - NumberFilter
  FindingProviderFieldsRelatedFindingsId:
    - StringFilter
  FindingProviderFieldsRelatedFindingsProductArn:
    - StringFilter
  FindingProviderFieldsSeverityLabel:
    - StringFilter
  FindingProviderFieldsSeverityOriginal:
    - StringFilter
  FindingProviderFieldsTypes:
    - StringFilter
  FirstObservedAt:
    - DateFilter
  GeneratorId:
    - StringFilter
  Id:
    - StringFilter
  Keyword:
    - KeywordFilter
  LastObservedAt:
    - DateFilter
  MalwareName:
    - StringFilter
  MalwarePath:
    - StringFilter
  MalwareState:
    - StringFilter
  MalwareType:
    - StringFilter
  NetworkDestinationDomain:
    - StringFilter
  NetworkDestinationIpV4:
    - IpFilter
  NetworkDestinationIpV6:
    - IpFilter
  NetworkDestinationPort:
    - NumberFilter
  NetworkDirection:
    - StringFilter
  NetworkProtocol:
    - StringFilter
  NetworkSourceDomain:
    - StringFilter
  NetworkSourceIpV4:
    - IpFilter
  NetworkSourceIpV6:
    - IpFilter
  NetworkSourceMac:
    - StringFilter
  NetworkSourcePort:
    - NumberFilter
  NoteText:
    - StringFilter
  NoteUpdatedAt:
    - DateFilter
  NoteUpdatedBy:
    - StringFilter
  ProcessLaunchedAt:
    - DateFilter
  ProcessName:
    - StringFilter
  ProcessParentPid:
    - NumberFilter
  ProcessPath:
    - StringFilter
  ProcessPid:
    - NumberFilter
  ProcessTerminatedAt:
    - DateFilter
  ProductArn:
    - StringFilter
  ProductFields:
    - MapFilter
  ProductName:
    - StringFilter
  RecommendationText:
    - StringFilter
  RecordState:
    - StringFilter
  Region:
    - StringFilter
  RelatedFindingsId:
    - StringFilter
  RelatedFindingsProductArn:
    - StringFilter
  ResourceApplicationArn:
    - StringFilter
  ResourceApplicationName:
    - StringFilter
  ResourceAwsEc2InstanceIamInstanceProfileArn:
    - StringFilter
  ResourceAwsEc2InstanceImageId:
    - StringFilter
  ResourceAwsEc2InstanceIpV4Addresses:
    - IpFilter
  ResourceAwsEc2InstanceIpV6Addresses:
    - IpFilter
  ResourceAwsEc2InstanceKeyName:
    - StringFilter
  ResourceAwsEc2InstanceLaunchedAt:
    - DateFilter
  ResourceAwsEc2InstanceSubnetId:
    - StringFilter
  ResourceAwsEc2InstanceType:
    - StringFilter
  ResourceAwsEc2InstanceVpcId:
    - StringFilter
  ResourceAwsIamAccessKeyCreatedAt:
    - DateFilter
  ResourceAwsIamAccessKeyPrincipalName:
    - StringFilter
  ResourceAwsIamAccessKeyStatus:
    - StringFilter
  ResourceAwsIamAccessKeyUserName:
    - StringFilter
  ResourceAwsIamUserUserName:
    - StringFilter
  ResourceAwsS3BucketOwnerId:
    - StringFilter
  ResourceAwsS3BucketOwnerName:
    - StringFilter
  ResourceContainerImageId:
    - StringFilter
  ResourceContainerImageName:
    - StringFilter
  ResourceContainerLaunchedAt:
    - DateFilter
  ResourceContainerName:
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
  Sample:
    - BooleanFilter
  SeverityLabel:
    - StringFilter
  SeverityNormalized:
    - NumberFilter
  SeverityProduct:
    - NumberFilter
  SourceUrl:
    - StringFilter
  ThreatIntelIndicatorCategory:
    - StringFilter
  ThreatIntelIndicatorLastObservedAt:
    - DateFilter
  ThreatIntelIndicatorSource:
    - StringFilter
  ThreatIntelIndicatorSourceUrl:
    - StringFilter
  ThreatIntelIndicatorType:
    - StringFilter
  ThreatIntelIndicatorValue:
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
  VulnerabilitiesExploitAvailable:
    - StringFilter
  VulnerabilitiesFixAvailable:
    - StringFilter
  WorkflowState:
    - StringFilter
  WorkflowStatus:
    - StringFilter

```

## Properties

`AwsAccountId`

The AWS account ID in which a finding is generated.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsAccountName`

The name of the AWS account in which a finding is generated.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CompanyName`

The name of the findings provider (company) that owns the solution (product) that
generates findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceAssociatedStandardsId`

The unique identifier of a standard in which a control is enabled. This field consists of the resource portion of the
Amazon Resource Name (ARN) returned for a standard in the [DescribeStandards](https://docs.aws.amazon.com/securityhub/1.0/APIReference/API_DescribeStandards.html) API response.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceSecurityControlId`

The unique identifier of a control across standards. Values for this field typically consist of an
AWS service and a number, such as APIGateway.5.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceSecurityControlParametersName`

The name of a security control parameter.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceSecurityControlParametersValue`

The current value of a security control parameter.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceStatus`

Exclusive to findings that are generated as the result of a check run against a specific
rule in a supported standard, such as CIS AWS Foundations. Contains security
standard-related finding details.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Confidence`

A finding's confidence. Confidence is defined as the likelihood that a finding
accurately identifies the behavior or issue that it was intended to identify.

Confidence is scored on a 0-100 basis using a ratio scale, where 0 means zero percent
confidence and 100 means 100 percent confidence.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreatedAt`

A timestamp that indicates when the security findings provider
created the potential security issue that a finding reflects.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Criticality`

The level of importance assigned to the resources associated with the finding.

A score of 0 means that the underlying resources have no criticality, and a score of 100
is reserved for the most critical resources.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A finding's description.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingProviderFieldsConfidence`

The finding provider value for the finding confidence. Confidence is defined as the likelihood
that a finding accurately identifies the behavior or issue that it was intended to
identify.

Confidence is scored on a 0-100 basis using a ratio scale, where 0 means zero percent
confidence and 100 means 100 percent confidence.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingProviderFieldsCriticality`

The finding provider value for the level of importance assigned to the resources associated with
the findings.

A score of 0 means that the underlying resources have no criticality, and a score of 100
is reserved for the most critical resources.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingProviderFieldsRelatedFindingsId`

The finding identifier of a related finding that is identified by the finding provider.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingProviderFieldsRelatedFindingsProductArn`

The ARN of the solution that generated a related finding that is identified by the finding provider.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingProviderFieldsSeverityLabel`

The finding provider value for the severity label.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingProviderFieldsSeverityOriginal`

The finding provider's original value for the severity.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingProviderFieldsTypes`

One or more finding types that the finding provider assigned to the finding. Uses the format of `namespace/category/classifier`
that classify a finding.

Valid namespace values are: Software and Configuration Checks \| TTPs \| Effects \| Unusual
Behaviors \| Sensitive Data Identifications

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirstObservedAt`

A timestamp that indicates when the security findings provider first
observed the potential security issue that a finding captured.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GeneratorId`

The identifier for the solution-specific component (a discrete unit of logic) that
generated a finding. In various security findings providers' solutions, this generator can
be called a rule, a check, a detector, a plugin, etc.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

The security findings provider-specific identifier for a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Keyword`

This field is deprecated. A keyword for a finding.

_Required_: No

_Type_: Array of [KeywordFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-keywordfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastObservedAt`

A timestamp that indicates when the security findings provider most recently observed a change in the resource that is involved in the finding.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MalwareName`

The name of the malware that was observed.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MalwarePath`

The filesystem path of the malware that was observed.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MalwareState`

The state of the malware that was observed.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MalwareType`

The type of the malware that was observed.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkDestinationDomain`

The destination domain of network-related information about a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkDestinationIpV4`

The destination IPv4 address of network-related information about a finding.

_Required_: No

_Type_: Array of [IpFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-ipfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkDestinationIpV6`

The destination IPv6 address of network-related information about a finding.

_Required_: No

_Type_: Array of [IpFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-ipfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkDestinationPort`

The destination port of network-related information about a finding.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkDirection`

Indicates the direction of network traffic associated with a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkProtocol`

The protocol of network-related information about a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkSourceDomain`

The source domain of network-related information about a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkSourceIpV4`

The source IPv4 address of network-related information about a finding.

_Required_: No

_Type_: Array of [IpFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-ipfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkSourceIpV6`

The source IPv6 address of network-related information about a finding.

_Required_: No

_Type_: Array of [IpFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-ipfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkSourceMac`

The source media access control (MAC) address of network-related information about a
finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkSourcePort`

The source port of network-related information about a finding.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteText`

The text of a note.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteUpdatedAt`

The timestamp of when the note was updated.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoteUpdatedBy`

The principal that created a note.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessLaunchedAt`

A timestamp that identifies when the process was launched.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessName`

The name of the process.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessParentPid`

The parent process ID. This field accepts positive integers between `O` and `2147483647`.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessPath`

The path to the process executable.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessPid`

The process ID.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProcessTerminatedAt`

A timestamp that identifies when the process was terminated.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductArn`

The ARN generated by Security Hub CSPM that uniquely identifies a third-party company
(security findings provider) after this provider's product (solution that generates
findings) is registered with Security Hub CSPM.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductFields`

A data type where security findings providers can include additional solution-specific
details that aren't part of the defined `AwsSecurityFinding` format.

_Required_: No

_Type_: Array of [MapFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-mapfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProductName`

The name of the solution (product) that generates findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecommendationText`

The recommendation of what to do about the issue described in a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordState`

The updated record state for the finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The Region from which the finding was generated.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelatedFindingsId`

The solution-generated identifier for a related finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelatedFindingsProductArn`

The ARN of the solution that generated a related finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceApplicationArn`

The ARN of the application that is related to a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceApplicationName`

The name of the application that is related to a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsEc2InstanceIamInstanceProfileArn`

The IAM profile ARN of the instance.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsEc2InstanceImageId`

The Amazon Machine Image (AMI) ID of the instance.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsEc2InstanceIpV4Addresses`

The IPv4 addresses associated with the instance.

_Required_: No

_Type_: Array of [IpFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-ipfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsEc2InstanceIpV6Addresses`

The IPv6 addresses associated with the instance.

_Required_: No

_Type_: Array of [IpFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-ipfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsEc2InstanceKeyName`

The key name associated with the instance.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsEc2InstanceLaunchedAt`

The date and time the instance was launched.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsEc2InstanceSubnetId`

The identifier of the subnet that the instance was launched in.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsEc2InstanceType`

The instance type of the instance.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsEc2InstanceVpcId`

The identifier of the VPC that the instance was launched in.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsIamAccessKeyCreatedAt`

The creation date/time of the IAM access key related to a finding.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsIamAccessKeyPrincipalName`

The name of the principal that is associated with an IAM access key.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsIamAccessKeyStatus`

The status of the IAM access key related to a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsIamAccessKeyUserName`

This field is deprecated. The username associated with the IAM access key related to a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsIamUserUserName`

The name of an IAM user.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsS3BucketOwnerId`

The canonical user ID of the owner of the S3 bucket.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAwsS3BucketOwnerName`

The display name of the owner of the S3 bucket.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceContainerImageId`

The identifier of the image related to a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceContainerImageName`

The name of the image related to a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceContainerLaunchedAt`

A timestamp that identifies when the container was started.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceContainerName`

The name of the container related to a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceDetailsOther`

The details of a resource that doesn't have a specific subfield for the resource type
defined.

_Required_: No

_Type_: Array of [MapFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-mapfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

The canonical identifier for the given resource type.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcePartition`

The canonical AWS partition name that the Region is assigned to.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceRegion`

The canonical AWS external Region name where this resource is located.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

A list of AWS tags associated with a resource at the time the finding was
processed.

_Required_: No

_Type_: Array of [MapFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-mapfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

Specifies the type of the resource that details are provided for.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sample`

Indicates whether or not sample findings are included in the filter results.

_Required_: No

_Type_: Array of [BooleanFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-booleanfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SeverityLabel`

The label of a finding's severity.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SeverityNormalized`

Deprecated. The normalized severity of a finding.
Instead of providing `Normalized`, provide `Label`.

The value of `Normalized` can be an integer between `0` and `100`.

If you provide `Label` and don't provide `Normalized`, then
`Normalized` is set automatically as follows.

- `INFORMATIONAL` \- 0

- `LOW` \- 1

- `MEDIUM` \- 40

- `HIGH` \- 70

- `CRITICAL` \- 90

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SeverityProduct`

Deprecated. This attribute isn't included in findings. Instead of providing
`Product`, provide `Original`.

The native severity as defined by the AWS service or integrated partner product that
generated the finding.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-numberfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceUrl`

A URL that links to a page about the current finding in the security findings provider's
solution.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThreatIntelIndicatorCategory`

The category of a threat intelligence indicator.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThreatIntelIndicatorLastObservedAt`

A timestamp that identifies the last observation of a threat intelligence indicator.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThreatIntelIndicatorSource`

The source of the threat intelligence.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThreatIntelIndicatorSourceUrl`

The URL for more details from the source of the threat intelligence.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThreatIntelIndicatorType`

The type of a threat intelligence indicator.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThreatIntelIndicatorValue`

The value of a threat intelligence indicator.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

A finding's title.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

A finding type in the format of `namespace/category/classifier` that
classifies a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdatedAt`

A timestamp that indicates when the security findings provider last
updated the finding record.

For more information about the validation and formatting of timestamp fields in AWS Security Hub CSPM, see [Timestamps](https://docs.aws.amazon.com/securityhub/1.0/APIReference/Welcome.html#timestamps).

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-datefilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserDefinedFields`

A list of name/value string pairs associated with the finding. These are custom,
user-defined fields added to a finding.

_Required_: No

_Type_: Array of [MapFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-mapfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerificationState`

The veracity of a finding.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VulnerabilitiesExploitAvailable`

Indicates whether a software vulnerability in your environment has a known exploit. You can filter findings by this
field only if you use Security Hub CSPM and Amazon Inspector.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VulnerabilitiesFixAvailable`

Indicates whether a vulnerability is fixed in a newer version of the affected software packages. You can filter
findings by this field only if you use Security Hub CSPM and Amazon Inspector.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowState`

The workflow state of a finding.

Note that this field is deprecated. To search for a finding based on its workflow
status, use `WorkflowStatus`.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowStatus`

The status of the investigation into a finding. Allowed values are the following.

- `NEW` \- The initial state of a finding, before it is reviewed.

Security Hub CSPM also resets the workflow status from `NOTIFIED` or
`RESOLVED` to `NEW` in the following cases:

- `RecordState` changes from `ARCHIVED` to `ACTIVE`.

- `Compliance.Status` changes from `PASSED` to either `WARNING`,
`FAILED`, or `NOT_AVAILABLE`.

- `NOTIFIED` \- Indicates that the resource owner has been notified about
the security issue. Used when the initial reviewer is not the resource owner, and
needs intervention from the resource owner.

If one of the following occurs, the workflow status is changed automatically from
`NOTIFIED` to `NEW`:

- `RecordState` changes from `ARCHIVED` to
`ACTIVE`.

- `Compliance.Status` changes from `PASSED` to `FAILED`,
`WARNING`, or `NOT_AVAILABLE`.

- `SUPPRESSED` \- Indicates that you reviewed the finding and don't believe that any action is
needed.

The workflow status of a `SUPPRESSED` finding does not change if
`RecordState` changes from `ARCHIVED` to
`ACTIVE`.

- `RESOLVED` \- The finding was reviewed and remediated and is now
considered resolved.

The finding remains `RESOLVED` unless one of the following occurs:

- `RecordState` changes from `ARCHIVED` to
`ACTIVE`.

- `Compliance.Status` changes from `PASSED` to `FAILED`,
`WARNING`, or `NOT_AVAILABLE`.

In those cases, the workflow status is automatically reset to `NEW`.

For findings from controls, if `Compliance.Status` is `PASSED`,
then Security Hub CSPM automatically sets the workflow status to `RESOLVED`.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-insight-stringfilter.html)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SecurityHub::Insight

BooleanFilter
