This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Capability ArgoCdRoleMapping

A mapping between an Argo CD role and IAM Identity CenterIAM; Identity Center identities. This defines which users or groups have specific permissions in Argo CD.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Identities" : [ SsoIdentity, ... ],
  "Role" : String
}

```

### YAML

```yaml

  Identities:
    - SsoIdentity
  Role: String

```

## Properties

`Identities`

A list of IAM Identity CenterIAM; Identity Center identities (users or groups) that should be assigned this Argo CD role.

_Required_: Yes

_Type_: Array of [SsoIdentity](aws-properties-eks-capability-ssoidentity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Role`

The Argo CD role to assign. Valid values are:

- `ADMIN` – Full administrative access to Argo CD.

- `EDITOR` – Edit access to Argo CD resources.

- `VIEWER` – Read-only access to Argo CD resources.

_Required_: Yes

_Type_: String

_Allowed values_: `ADMIN | EDITOR | VIEWER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArgoCd

AwsIdc

All content copied from https://docs.aws.amazon.com/.
