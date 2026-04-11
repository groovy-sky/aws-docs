This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::AuthPolicy

Creates or updates the auth policy. The policy string in JSON must not contain newlines or
blank lines.

For more information, see [Auth policies](../../../vpc-lattice/latest/ug/auth-policies.md) in the _Amazon VPC_
_Lattice User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::AuthPolicy",
  "Properties" : {
      "Policy" : Json,
      "ResourceIdentifier" : String
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::AuthPolicy
Properties:
  Policy: Json
  ResourceIdentifier: String

```

## Properties

`Policy`

The auth policy.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceIdentifier`

The ID or ARN of the service network or service for which the policy is created.

_Required_: Yes

_Type_: String

_Pattern_: `^((((sn)|(svc))-[0-9a-z]{17})|(arn(:[a-z0-9]+([.-][a-z0-9]+)*){2}(:([a-z0-9]+([.-][a-z0-9]+)*)?){2}:((servicenetwork/sn)|(service/svc))-[0-9a-z]{17}))$`

_Minimum_: `17`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the auth policy.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`State`

The state of the auth policy. The auth policy is only active when the auth type is set to
`AWS_IAM`. If you provide a policy, then authentication and
authorization decisions are made based on this policy and the client's IAM policy. If the auth
type is `NONE`, then any auth policy you provide will remain inactive.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::VpcLattice::DomainVerification

All content copied from https://docs.aws.amazon.com/.
