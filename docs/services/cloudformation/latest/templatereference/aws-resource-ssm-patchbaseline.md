This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::PatchBaseline

The `AWS::SSM::PatchBaseline` resource defines the basic information for an
AWS Systems Manager patch baseline. A patch baseline defines which patches are
approved for installation on your instances.

For more information, see [CreatePatchBaseline](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_CreatePatchBaseline.html) in the _AWS Systems Manager API_
_Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSM::PatchBaseline",
  "Properties" : {
      "ApprovalRules" : RuleGroup,
      "ApprovedPatches" : [ String, ... ],
      "ApprovedPatchesComplianceLevel" : String,
      "ApprovedPatchesEnableNonSecurity" : Boolean,
      "AvailableSecurityUpdatesComplianceStatus" : String,
      "DefaultBaseline" : Boolean,
      "Description" : String,
      "GlobalFilters" : PatchFilterGroup,
      "Name" : String,
      "OperatingSystem" : String,
      "PatchGroups" : [ String, ... ],
      "RejectedPatches" : [ String, ... ],
      "RejectedPatchesAction" : String,
      "Sources" : [ PatchSource, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SSM::PatchBaseline
Properties:
  ApprovalRules:
    RuleGroup
  ApprovedPatches:
    - String
  ApprovedPatchesComplianceLevel: String
  ApprovedPatchesEnableNonSecurity: Boolean
  AvailableSecurityUpdatesComplianceStatus: String
  DefaultBaseline: Boolean
  Description: String
  GlobalFilters:
    PatchFilterGroup
  Name: String
  OperatingSystem: String
  PatchGroups:
    - String
  RejectedPatches:
    - String
  RejectedPatchesAction: String
  Sources:
    - PatchSource
  Tags:
    - Tag

```

## Properties

`ApprovalRules`

A set of rules used to include patches in the baseline.

_Required_: No

_Type_: [RuleGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-patchbaseline-rulegroup.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApprovedPatches`

A list of explicitly approved patches for the baseline.

For information about accepted formats for lists of approved patches and rejected patches,
see [Package\
name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the _AWS Systems Manager User Guide_.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `100 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApprovedPatchesComplianceLevel`

Defines the compliance level for approved patches. When an approved patch is reported as
missing, this value describes the severity of the compliance violation. The default value is
`UNSPECIFIED`.

_Required_: No

_Type_: String

_Allowed values_: `CRITICAL | HIGH | MEDIUM | LOW | INFORMATIONAL | UNSPECIFIED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApprovedPatchesEnableNonSecurity`

Indicates whether the list of approved patches includes non-security updates that should be
applied to the managed nodes. The default value is `false`. Applies to Linux managed
nodes only.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailableSecurityUpdatesComplianceStatus`

Indicates the status you want to assign to security patches that are available but not
approved because they don't meet the installation criteria specified in the patch
baseline.

Example scenario: Security patches that you might want installed can be skipped if you have
specified a long period to wait after a patch is released before installation. If an update to
the patch is released during your specified waiting period, the waiting period for installing the
patch starts over. If the waiting period is too long, multiple versions of the patch could be
released but never installed.

Supported for Windows Server managed nodes only.

_Required_: No

_Type_: String

_Allowed values_: `NON_COMPLIANT | COMPLIANT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultBaseline`

Indicates whether this is the default baseline. AWS Systems Manager supports creating multiple default
patch baselines. For example, you can create a default patch baseline for each operating
system.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the patch baseline.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalFilters`

A set of global filters used to include patches in the baseline.

###### Important

The `GlobalFilters` parameter can be configured only by using the AWS CLI or an AWS SDK. It can't be configured from the Patch Manager
console, and its value isn't displayed in the console.

_Required_: No

_Type_: [PatchFilterGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-patchbaseline-patchfiltergroup.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the patch baseline.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.]{3,128}$`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperatingSystem`

Defines the operating system the patch baseline applies to. The default value is
`WINDOWS`.

_Required_: No

_Type_: String

_Allowed values_: `WINDOWS | AMAZON_LINUX | AMAZON_LINUX_2 | AMAZON_LINUX_2022 | AMAZON_LINUX_2023 | UBUNTU | REDHAT_ENTERPRISE_LINUX | SUSE | CENTOS | ORACLE_LINUX | DEBIAN | MACOS | RASPBIAN | ROCKY_LINUX | ALMA_LINUX`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PatchGroups`

The name of the patch group to be registered with the patch baseline.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RejectedPatches`

A list of explicitly rejected patches for the baseline.

For information about accepted formats for lists of approved patches and rejected patches,
see [Package\
name formats for approved and rejected patch lists](https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html) in the _AWS Systems Manager User Guide_.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `100 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RejectedPatchesAction`

The action for Patch Manager to take on patches included in the
`RejectedPackages` list.

ALLOW\_AS\_DEPENDENCY

**Linux and macOS**: A package in the rejected patches list
is installed only if it is a dependency of another package. It is considered compliant with
the patch baseline, and its status is reported as `INSTALLED_OTHER`. This is the
default action if no option is specified.

**Windows Server**: Windows Server doesn't support the
concept of package dependencies. If a package in the rejected patches list and already
installed on the node, its status is reported as `INSTALLED_OTHER`. Any package not
already installed on the node is skipped. This is the default action if no option is
specified.

BLOCK

**All OSs**: Packages in the rejected patches list, and
packages that include them as dependencies, aren't installed by Patch Manager under any
circumstances.

State value assignment for patch compliance:

- If a package was installed before it was added to the rejected patches list, or is
installed outside of Patch Manager afterward, it's considered noncompliant with the patch
baseline and its status is reported as `INSTALLED_REJECTED`.

- If an update attempts to install a dependency package that is now rejected by the
baseline, when previous versions of the package were not rejected, the package being updated
is reported as `MISSING` for `SCAN` operations and as
`FAILED` for `INSTALL` operations.

_Required_: No

_Type_: String

_Allowed values_: `ALLOW_AS_DEPENDENCY | BLOCK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sources`

Information about the patches to use to update the managed nodes, including target operating
systems and source repositories. Applies to Linux managed nodes only.

_Required_: No

_Type_: Array of [PatchSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-patchbaseline-patchsource.html)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Optional metadata that you assign to a resource. Tags enable you to categorize a
resource in different ways, such as by purpose, owner, or environment. For example, you
might want to tag a patch baseline to identify the severity level of patches it
specifies and the operating system family it applies to.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ssm-patchbaseline-tag.html)

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the patch baseline ID, such as
`pb-abcde1234567890yz`.

###### Note

The ID of the default patch baseline provided by AWS is an Amazon
Resource Name (ARN), for example
`arn:aws:ssm:us-west-2:123456789012:patchbaseline/abcde1234567890yz`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

## Examples

### Create a patch baseline

The following example creates a Systems Manager patch baseline that
approves patches for Windows Server 2019 instances seven days after they are
released by Microsoft. The patch baseline also approves patches for Active
Directory seven days after they are released by Microsoft.

```language_owl_wvr_qlb

{
    "Resources": {
        "myPatchBaseline": {
            "Type": "AWS::SSM::PatchBaseline",
            "Properties": {
                "Name": "myPatchBaseline",
                "Description": "Baseline containing all updates approved for Windows instances",
                "OperatingSystem": "WINDOWS",
                "PatchGroups": [
                    "myPatchGroup"
                ],
                "ApprovalRules": {
                    "PatchRules": [
                        {
                            "PatchFilterGroup": {
                                "PatchFilters": [
                                    {
                                        "Values": [
                                            "Critical",
                                            "Important",
                                            "Moderate"
                                        ],
                                        "Key": "MSRC_SEVERITY"
                                    },
                                    {
                                        "Values": [
                                            "SecurityUpdates",
                                            "CriticalUpdates"
                                        ],
                                        "Key": "CLASSIFICATION"
                                    },
                                    {
                                        "Values": [
                                            "WindowsServer2019"
                                        ],
                                        "Key": "PRODUCT"
                                    }
                                ]
                            },
                            "ApproveAfterDays": 7,
                            "ComplianceLevel": "CRITICAL"
                        },
                        {
                            "PatchFilterGroup": {
                                "PatchFilters": [
                                    {
                                        "Values": [
                                            "Critical",
                                            "Important",
                                            "Moderate"
                                        ],
                                        "Key": "MSRC_SEVERITY"
                                    },
                                    {
                                        "Values": [
                                            "*"
                                        ],
                                        "Key": "CLASSIFICATION"
                                    },
                                    {
                                        "Values": [
                                            "APPLICATION"
                                        ],
                                        "Key": "PATCH_SET"
                                    },
                                    {
                                        "Values": [
                                            "Active Directory Rights Management Services Client 2.0"
                                        ],
                                        "Key": "PRODUCT"
                                    },
                                    {
                                        "Values": [
                                            "Active Directory"
                                        ],
                                        "Key": "PRODUCT_FAMILY"
                                    }
                                ]
                            },
                            "ApproveAfterDays": 7,
                            "ComplianceLevel": "CRITICAL"
                        }
                    ]
                }
            }
        }
    }
}
```

#### YAML

```yaml

---
Resources:
  myPatchBaseline:
    Type: AWS::SSM::PatchBaseline
    Properties:
      Name: myPatchBaseline
      Description: Baseline containing all updates approved for Windows instances
      OperatingSystem: WINDOWS
      PatchGroups:
      - myPatchGroup
      ApprovalRules:
        PatchRules:
        - PatchFilterGroup:
            PatchFilters:
            - Values:
              - Critical
              - Important
              - Moderate
              Key: MSRC_SEVERITY
            - Values:
              - SecurityUpdates
              - CriticalUpdates
              Key: CLASSIFICATION
            - Values:
              - WindowsServer2019
              Key: PRODUCT
          ApproveAfterDays: 7
          ComplianceLevel: CRITICAL
        - PatchFilterGroup:
            PatchFilters:
            - Values:
              - Critical
              - Important
              - Moderate
              Key: MSRC_SEVERITY
            - Values:
              - "*"
              Key: CLASSIFICATION
            - Values:
              - APPLICATION
              Key: PATCH_SET
            - Values:
              - Active Directory Rights Management Services Client 2.0
              Key: PRODUCT
            - Values:
              - Active Directory
              Key: PRODUCT_FAMILY
          ApproveAfterDays: 7
          ComplianceLevel: CRITICAL
```

## See also

- [CreatePatchBaseline](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_CreatePatchBaseline.html) in the _AWS Systems Manager_
_API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SSM::Parameter

PatchFilter
