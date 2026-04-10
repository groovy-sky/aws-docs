This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::ComputeEnvironment LaunchTemplateSpecificationOverride

An object that represents a launch template to use in place of the default launch template. You must specify either the launch template ID or launch template name in the request, but not
both.

If security groups are specified using both the `securityGroupIds` parameter of
`CreateComputeEnvironment` and the launch template, the values in the
`securityGroupIds` parameter of `CreateComputeEnvironment` will be
used.

You can define up to ten (10) overrides for each compute environment.

###### Note

This object isn't applicable to jobs that are running on Fargate resources.

###### Note

To unset all override templates for a compute environment, you can pass an empty array to the [UpdateComputeEnvironment.overrides](../../../../reference/batch/latest/apireference/api-updatecomputeenvironment.md) parameter, or not include the `overrides` parameter when submitting the `UpdateComputeEnvironment` API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LaunchTemplateId" : String,
  "LaunchTemplateName" : String,
  "TargetInstanceTypes" : [ String, ... ],
  "UserdataType" : String,
  "Version" : String
}

```

### YAML

```yaml

  LaunchTemplateId: String
  LaunchTemplateName: String
  TargetInstanceTypes:
    - String
  UserdataType: String
  Version: String

```

## Properties

`LaunchTemplateId`

The ID of the launch template.

**Note:** If you specify the `launchTemplateId` you can't specify the `launchTemplateName` as well.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LaunchTemplateName`

The name of the launch template.

**Note:** If you specify the `launchTemplateName` you can't specify the `launchTemplateId` as well.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`TargetInstanceTypes`

The instance type or family that this override launch template should be applied to.

This parameter is required when defining a launch template override.

Information included in this parameter must meet the following requirements:

- Must be a valid Amazon EC2 instance type or family.

- The following AWS Batch `InstanceTypes` are not allowed: `optimal`, `default_x86_64`, and `default_arm64`.

- `targetInstanceTypes` can target only instance types and families that are included within the [`ComputeResource.instanceTypes`](../../../../reference/batch/latest/apireference/api-computeresource.md#Batch-Type-ComputeResource-instanceTypes) set. `targetInstanceTypes` doesn't need to include all of the instances from the `instanceType` set, but at least a subset. For example, if `ComputeResource.instanceTypes` includes `[m5, g5]`, `targetInstanceTypes` can include `[m5.2xlarge]` and `[m5.large]` but not `[c5.large]`.

- `targetInstanceTypes` included within the same launch template override or across launch template overrides can't overlap for the same compute environment. For example, you can't define one launch template override to target an instance family and another define an instance type within this same family.

_Required_: No

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`UserdataType`

The EKS node initialization process to use. You only need to specify this value if you are
using a custom AMI. The default value is `EKS_BOOTSTRAP_SH`. If
_imageType_ is a custom AMI based on EKS\_AL2023 or EKS\_AL2023\_NVIDIA then you
must choose `EKS_NODEADM`.

_Required_: No

_Type_: String

_Allowed values_: `EKS_BOOTSTRAP_SH | EKS_NODEADM`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Version`

The version number of the launch template,
`$Default`, or `$Latest`.

If the value is `$Default`, the default version of the launch template is used. If the value is `$Latest`, the latest version of the launch template is used.

###### Important

If the AMI ID that's used in a compute environment is from the launch template, the AMI
isn't changed when the compute environment is updated. It's only changed if the
`updateToLatestImageVersion` parameter for the compute environment is set to
`true`. During an infrastructure update, if either `$Default` or `$Latest` is specified, AWS Batch re-evaluates the launch template version, and it
might use a different version of the launch template. This is the case even if the launch
template isn't specified in the update. When updating a compute environment, changing the launch
template requires an infrastructure update of the compute environment. For more information, see
[Updating compute\
environments](../../../batch/latest/userguide/updating-compute-environments.md) in the _AWS Batch User Guide_.

Default: `$Default`

Latest: `$Latest`

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LaunchTemplateSpecification

UpdatePolicy

All content copied from https://docs.aws.amazon.com/.
