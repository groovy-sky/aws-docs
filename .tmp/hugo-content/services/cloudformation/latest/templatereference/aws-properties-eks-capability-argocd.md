This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Capability ArgoCd

The `ArgoCd` property type specifies Property description not available. for an [AWS::EKS::Capability](aws-resource-eks-capability.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsIdc" : AwsIdc,
  "Namespace" : String,
  "NetworkAccess" : NetworkAccess,
  "RbacRoleMappings" : [ ArgoCdRoleMapping, ... ],
  "ServerUrl" : String
}

```

### YAML

```yaml

  AwsIdc:
    AwsIdc
  Namespace: String
  NetworkAccess:
    NetworkAccess
  RbacRoleMappings:
    - ArgoCdRoleMapping
  ServerUrl: String

```

## Properties

`AwsIdc`

Property description not available.

_Required_: Yes

_Type_: [AwsIdc](aws-properties-eks-capability-awsidc.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Namespace`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkAccess`

Property description not available.

_Required_: No

_Type_: [NetworkAccess](aws-properties-eks-capability-networkaccess.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RbacRoleMappings`

Property description not available.

_Required_: No

_Type_: Array of [ArgoCdRoleMapping](aws-properties-eks-capability-argocdrolemapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerUrl`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EKS::Capability

ArgoCdRoleMapping

All content copied from https://docs.aws.amazon.com/.
