This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Capability

An object representing a managed capability in an Amazon EKS cluster. This includes all configuration, status, and health information for the capability.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EKS::Capability",
  "Properties" : {
      "CapabilityName" : String,
      "ClusterName" : String,
      "Configuration" : CapabilityConfiguration,
      "DeletePropagationPolicy" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::EKS::Capability
Properties:
  CapabilityName: String
  ClusterName: String
  Configuration:
    CapabilityConfiguration
  DeletePropagationPolicy: String
  RoleArn: String
  Tags:
    - Tag
  Type: String

```

## Properties

`CapabilityName`

The unique name of the capability within the cluster.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterName`

The name of the Amazon EKS cluster that contains this capability.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configuration`

The configuration settings for the capability. The structure varies depending on the capability type.

_Required_: No

_Type_: [CapabilityConfiguration](aws-properties-eks-capability-capabilityconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletePropagationPolicy`

The delete propagation policy for the capability. Currently, the only supported value is `RETAIN`, which keeps all resources managed by the capability when the capability is deleted.

_Required_: Yes

_Type_: String

_Allowed values_: `RETAIN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role that the capability uses to interact with AWS services.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[a-z-]*:iam::[0-9]+:role/[a-zA-Z0-9+=,.@_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-eks-capability-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of capability. Valid values are `ACK`, `ARGOCD`, or `KRO`.

_Required_: Yes

_Type_: String

_Allowed values_: `ARGOCD | ACK | KRO`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the capability.

`Configuration.ArgoCd.AwsIdc.IdcManagedApplicationArn`

Property description not available.

`Configuration.ArgoCd.ServerUrl`

Property description not available.

`CreatedAt`

The Unix epoch timestamp in seconds for when the capability was created.

`ModifiedAt`

The Unix epoch timestamp in seconds for when the capability was last modified.

`Status`

The current status of the capability. Valid values include:

- `CREATING` – The capability is being created.

- `ACTIVE` – The capability is running and available.

- `UPDATING` – The capability is being updated.

- `DELETING` – The capability is being deleted.

- `CREATE_FAILED` – The capability creation failed.

- `UPDATE_FAILED` – The capability update failed.

- `DELETE_FAILED` – The capability deletion failed.

`Version`

The version of the capability software that is currently running.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ArgoCd

All content copied from https://docs.aws.amazon.com/.
