This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::Filter FilterCriteria

Details on the criteria used to define the filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsAccountId" : [ StringFilter, ... ],
  "CodeVulnerabilityDetectorName" : [ StringFilter, ... ],
  "CodeVulnerabilityDetectorTags" : [ StringFilter, ... ],
  "CodeVulnerabilityFilePath" : [ StringFilter, ... ],
  "ComponentId" : [ StringFilter, ... ],
  "ComponentType" : [ StringFilter, ... ],
  "Ec2InstanceImageId" : [ StringFilter, ... ],
  "Ec2InstanceSubnetId" : [ StringFilter, ... ],
  "Ec2InstanceVpcId" : [ StringFilter, ... ],
  "EcrImageArchitecture" : [ StringFilter, ... ],
  "EcrImageHash" : [ StringFilter, ... ],
  "EcrImagePushedAt" : [ DateFilter, ... ],
  "EcrImageRegistry" : [ StringFilter, ... ],
  "EcrImageRepositoryName" : [ StringFilter, ... ],
  "EcrImageTags" : [ StringFilter, ... ],
  "EpssScore" : [ NumberFilter, ... ],
  "ExploitAvailable" : [ StringFilter, ... ],
  "FindingArn" : [ StringFilter, ... ],
  "FindingStatus" : [ StringFilter, ... ],
  "FindingType" : [ StringFilter, ... ],
  "FirstObservedAt" : [ DateFilter, ... ],
  "FixAvailable" : [ StringFilter, ... ],
  "InspectorScore" : [ NumberFilter, ... ],
  "LambdaFunctionExecutionRoleArn" : [ StringFilter, ... ],
  "LambdaFunctionLastModifiedAt" : [ DateFilter, ... ],
  "LambdaFunctionLayers" : [ StringFilter, ... ],
  "LambdaFunctionName" : [ StringFilter, ... ],
  "LambdaFunctionRuntime" : [ StringFilter, ... ],
  "LastObservedAt" : [ DateFilter, ... ],
  "NetworkProtocol" : [ StringFilter, ... ],
  "PortRange" : [ PortRangeFilter, ... ],
  "RelatedVulnerabilities" : [ StringFilter, ... ],
  "ResourceId" : [ StringFilter, ... ],
  "ResourceTags" : [ MapFilter, ... ],
  "ResourceType" : [ StringFilter, ... ],
  "Severity" : [ StringFilter, ... ],
  "Title" : [ StringFilter, ... ],
  "UpdatedAt" : [ DateFilter, ... ],
  "VendorSeverity" : [ StringFilter, ... ],
  "VulnerabilityId" : [ StringFilter, ... ],
  "VulnerabilitySource" : [ StringFilter, ... ],
  "VulnerablePackages" : [ PackageFilter, ... ]
}

```

### YAML

```yaml

  AwsAccountId:
    - StringFilter
  CodeVulnerabilityDetectorName:
    - StringFilter
  CodeVulnerabilityDetectorTags:
    - StringFilter
  CodeVulnerabilityFilePath:
    - StringFilter
  ComponentId:
    - StringFilter
  ComponentType:
    - StringFilter
  Ec2InstanceImageId:
    - StringFilter
  Ec2InstanceSubnetId:
    - StringFilter
  Ec2InstanceVpcId:
    - StringFilter
  EcrImageArchitecture:
    - StringFilter
  EcrImageHash:
    - StringFilter
  EcrImagePushedAt:
    - DateFilter
  EcrImageRegistry:
    - StringFilter
  EcrImageRepositoryName:
    - StringFilter
  EcrImageTags:
    - StringFilter
  EpssScore:
    - NumberFilter
  ExploitAvailable:
    - StringFilter
  FindingArn:
    - StringFilter
  FindingStatus:
    - StringFilter
  FindingType:
    - StringFilter
  FirstObservedAt:
    - DateFilter
  FixAvailable:
    - StringFilter
  InspectorScore:
    - NumberFilter
  LambdaFunctionExecutionRoleArn:
    - StringFilter
  LambdaFunctionLastModifiedAt:
    - DateFilter
  LambdaFunctionLayers:
    - StringFilter
  LambdaFunctionName:
    - StringFilter
  LambdaFunctionRuntime:
    - StringFilter
  LastObservedAt:
    - DateFilter
  NetworkProtocol:
    - StringFilter
  PortRange:
    - PortRangeFilter
  RelatedVulnerabilities:
    - StringFilter
  ResourceId:
    - StringFilter
  ResourceTags:
    - MapFilter
  ResourceType:
    - StringFilter
  Severity:
    - StringFilter
  Title:
    - StringFilter
  UpdatedAt:
    - DateFilter
  VendorSeverity:
    - StringFilter
  VulnerabilityId:
    - StringFilter
  VulnerabilitySource:
    - StringFilter
  VulnerablePackages:
    - PackageFilter

```

## Properties

`AwsAccountId`

Details of the AWS account IDs used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeVulnerabilityDetectorName`

Property description not available.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeVulnerabilityDetectorTags`

Property description not available.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeVulnerabilityFilePath`

Property description not available.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentId`

Details of the component IDs used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentType`

Details of the component types used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2InstanceImageId`

Details of the Amazon EC2 instance image IDs used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2InstanceSubnetId`

Details of the Amazon EC2 instance subnet IDs used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ec2InstanceVpcId`

Details of the Amazon EC2 instance VPC IDs used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcrImageArchitecture`

Details of the Amazon ECR image architecture types used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcrImageHash`

Details of the Amazon ECR image hashes used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcrImagePushedAt`

Details on the Amazon ECR image push date and time used to filter findings.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-datefilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcrImageRegistry`

Details on the Amazon ECR registry used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcrImageRepositoryName`

Details on the name of the Amazon ECR repository used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EcrImageTags`

The tags attached to the Amazon ECR container image.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EpssScore`

Property description not available.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-numberfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExploitAvailable`

Property description not available.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingArn`

Details on the finding ARNs used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingStatus`

Details on the finding status types used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FindingType`

Details on the finding types used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirstObservedAt`

Details on the date and time a finding was first seen used to filter findings.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-datefilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FixAvailable`

Property description not available.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InspectorScore`

The Amazon Inspector score to filter on.

_Required_: No

_Type_: Array of [NumberFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-numberfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionExecutionRoleArn`

Property description not available.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionLastModifiedAt`

Property description not available.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-datefilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionLayers`

Property description not available.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionName`

Property description not available.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionRuntime`

Property description not available.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastObservedAt`

Details on the date and time a finding was last seen used to filter findings.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-datefilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkProtocol`

Details on network protocol used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PortRange`

Details on the port ranges used to filter findings.

_Required_: No

_Type_: Array of [PortRangeFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-portrangefilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelatedVulnerabilities`

Details on the related vulnerabilities used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceId`

Details on the resource IDs used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTags`

Details on the resource tags used to filter findings.

_Required_: No

_Type_: Array of [MapFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-mapfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

Details on the resource types used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Severity`

Details on the severity used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

Details on the finding title used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdatedAt`

Details on the date and time a finding was last updated at used to filter
findings.

_Required_: No

_Type_: Array of [DateFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-datefilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VendorSeverity`

Details on the vendor severity used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VulnerabilityId`

Details on the vulnerability ID used to filter findings.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VulnerabilitySource`

Details on the vulnerability score to filter findings by.

_Required_: No

_Type_: Array of [StringFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-stringfilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VulnerablePackages`

Details on the vulnerable packages used to filter findings.

_Required_: No

_Type_: Array of [PackageFilter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspectorv2-filter-packagefilter.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DateFilter

MapFilter
