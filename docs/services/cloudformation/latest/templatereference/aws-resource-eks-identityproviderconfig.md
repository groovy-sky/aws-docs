This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::IdentityProviderConfig

Associates an identity provider configuration to a cluster.

If you want to authenticate identities using an identity provider, you can create an
identity provider configuration and associate it to your cluster. After configuring
authentication to your cluster you can create Kubernetes `Role` and
`ClusterRole` objects, assign permissions to them, and then bind them to
the identities using Kubernetes `RoleBinding` and `ClusterRoleBinding`
objects. For more information see [Using RBAC\
Authorization](https://kubernetes.io/docs/reference/access-authn-authz/rbac) in the Kubernetes documentation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EKS::IdentityProviderConfig",
  "Properties" : {
      "ClusterName" : String,
      "IdentityProviderConfigName" : String,
      "Oidc" : OidcIdentityProviderConfig,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::EKS::IdentityProviderConfig
Properties:
  ClusterName: String
  IdentityProviderConfigName: String
  Oidc:
    OidcIdentityProviderConfig
  Tags:
    - Tag
  Type: String

```

## Properties

`ClusterName`

The name of your cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IdentityProviderConfigName`

The name of the configuration.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Oidc`

An object representing an OpenID Connect (OIDC) identity provider configuration.

_Required_: No

_Type_: [OidcIdentityProviderConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-identityproviderconfig-oidcidentityproviderconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that assists with categorization and organization.
Each tag consists of a key and an optional value. You define both. Tags don't
propagate to any other cluster or AWS resources.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-identityproviderconfig-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the identity provider configuration. The only type available is
`oidc`.

_Required_: Yes

_Type_: String

_Allowed values_: `oidc`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myIdentityProviderConfig" }`

For the IdentityProviderConfig, Ref returns the physical resource ID of the config.
For example, `cluster-name/oidc/identity-provider-config-name`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IdentityProviderConfigArn`

The Amazon Resource Name (ARN) associated with the identity provider config.

## Remarks

_Creating an identity provider config and Fargate profile_
_resources in the same template._

If CloudFormation attempts to create both resources at the same time, resource
creation fails. If you want to create both resources in the same template, then add
the `DependsOn` property in your template, as shown in the
examples.

## Examples

### Create an identity provider config

The following example creates a an identity provider config. If you're not
creating an `EKSFargateProfile` in the same template, remove the
`"DependsOn"` line in the following example. For more
information, see [`AWS::EKS::FargateProfile`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-fargateprofile.html).

#### JSON

```json

{
  "EKSIdpConfig": {
    "DependsOn": "EKSFargateProfile",
    "Type": "AWS::EKS::IdentityProviderConfig",
    "Properties": {
      "ClusterName": "my-cluster",
      "Type": "oidc",
      "Oidc": {
        "ClientId": "kubernetes",
        "IssuerUrl": "https://example.com"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  EKSIdpConfig:
    DependsOn: EKSFargateProfile
    Type: AWS::EKS::IdentityProviderConfig
    Properties:
      ClusterName: my-cluster
      Type: oidc
      Oidc:
        ClientId: "kubernetes"
        IssuerUrl: "https://example.com"
```

## See also

- [Authenticating users for your cluster from an OpenID Connect identity\
provider](https://docs.aws.amazon.com/eks/latest/userguide/authenticate-oidc-identity-provider.html) in the _Amazon EKS User Guide_.

- [AssociateIdentityProviderConfig](https://docs.aws.amazon.com/eks/latest/APIReference/API_AssociateIdentityProviderConfig.html) in the _Amazon EKS API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

OidcIdentityProviderConfig
