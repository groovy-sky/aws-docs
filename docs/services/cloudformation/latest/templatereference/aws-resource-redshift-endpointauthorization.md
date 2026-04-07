This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Redshift::EndpointAuthorization

Describes an endpoint authorization for authorizing Redshift-managed VPC endpoint access to a cluster across AWS accounts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Redshift::EndpointAuthorization",
  "Properties" : {
      "Account" : String,
      "ClusterIdentifier" : String,
      "Force" : Boolean,
      "VpcIds" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Redshift::EndpointAuthorization
Properties:
  Account: String
  ClusterIdentifier: String
  Force: Boolean
  VpcIds:
    - String

```

## Properties

`Account`

The AWS account ID of either the cluster owner (grantor) or grantee.
If `Grantee` parameter is true, then the `Account` value is of the grantor.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClusterIdentifier`

The cluster identifier.

_Required_: Yes

_Type_: String

_Pattern_: `^(?=^[a-z][a-z0-9]*(-[a-z0-9]+)*$).{1,63}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Force`

Indicates whether to force the revoke action.
If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcIds`

The virtual private cloud (VPC) identifiers to grant access to.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AllowedAllVPCs`

Indicates whether all VPCs in the grantee account are allowed access to the cluster.

`AllowedVPCs`

The VPCs allowed access to the cluster.

`AuthorizeTime`

The time (UTC) when the authorization was created.

`ClusterStatus`

The status of the cluster.

`EndpointCount`

The number of Redshift-managed VPC endpoints created for the authorization.

`Grantee`

The AWS account ID of the grantee of the cluster.

`Grantor`

The AWS account ID of the cluster owner.

`Status`

The status of the authorization action.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcSecurityGroup

AWS::Redshift::EventSubscription
