This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::Policy

Creates or updates a Cedar policy and saves it in the specified policy store. You can
create either a static policy or a policy linked to a policy template.

You can directly update only static policies. To update a template-linked policy, you
must update its linked policy template instead.

- To create a static policy, in the `Definition` include a
`Static` element that includes the Cedar policy text in the
`Statement` element.

- To create a policy that is dynamically linked to a policy template, in the
`Definition` include a `Templatelinked` element that
specifies the policy template ID and the principal and resource to associate
with this policy. If the policy template is ever updated, any policies linked to
the policy template automatically use the updated template.

###### Note

- If policy validation is enabled in the policy store, then updating a
static policy causes Verified Permissions to validate the policy against the
schema in the policy store. If the updated static policy doesn't pass
validation, the operation fails and the update isn't stored.

- When you edit a static policy, You can change only certain elements of a
static policy:

- The action referenced by the policy.

- A condition clause, such as when and unless.

You can't change these elements of a static policy:

- Changing a policy from a static policy to a template-linked
policy.

- Changing the effect of a static policy from permit or forbid.

- The principal referenced by a static policy.

- The resource referenced by a static policy.

- To update a template-linked policy, you must update the template instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VerifiedPermissions::Policy",
  "Properties" : {
      "Definition" : PolicyDefinition,
      "PolicyStoreId" : String
    }
}

```

### YAML

```yaml

Type: AWS::VerifiedPermissions::Policy
Properties:
  Definition:
    PolicyDefinition
  PolicyStoreId: String

```

## Properties

`Definition`

Specifies the policy type and content to use for the new or updated policy. The
definition structure must include either a `Static` or a
`TemplateLinked` element.

_Required_: Yes

_Type_: [PolicyDefinition](aws-properties-verifiedpermissions-policy-policydefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyStoreId`

Specifies the `PolicyStoreId` of the policy store you want to store the policy
in.

To specify a policy store, use its ID or alias name. When using an alias name, prefix it with `policy-store-alias/`. For example:

- ID: `PSEXAMPLEabcdefg111111`

- Alias name: `policy-store-alias/example-policy-store`

To view aliases, use [ListPolicyStoreAliases](../../../../reference/verifiedpermissions/latest/apireference/api-listpolicystorealiases.md).

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique id of the new or updated policy. For
example:

`{ "Ref": "SPEXAMPLEabcdefg111111" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`PolicyId`

The unique ID of the new or updated policy.

`PolicyType`

The type of the policy. This is one of the following values:

- Static

- TemplateLinked

## Examples

- [Creating a static policy](#aws-resource-verifiedpermissions-policy--examples--Creating_a_static_policy)

- [Creating a template-linked policy](#aws-resource-verifiedpermissions-policy--examples--Creating_a_template-linked_policy)

### Creating a static policy

The following example creates a static policy in the specified policy store
with the specified policy statement.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation sample template for creating a static policy for Verified Permissions",
    "Parameters": {
        "PolicyStoreId": {
            "Type": "String"
        },
        "Description": {
            "Type": "String"
        },
        "Statement": {
            "Type": "String"
        }
    },
    "Resources": {
        "StaticPolicy": {
            "Type": "AWS::VerifiedPermissions::Policy",
            "Properties": {
                "PolicyStoreId": {
                    "Ref": "PolicyStoreId"
                },
                "Definition": {
                    "Static": {
                        "Description": {
                            "Ref": "Description"
                        },
                        "Statement": {
                            "Ref": "Statement"
                        }
                    }
                }
            }
        }
    },
    "Outputs": {
        "PolicyId": {
            "Value": {
                "Fn::GetAtt": [
                    "StaticPolicy",
                    "PolicyId"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: >-
  Description": "AWS CloudFormation sample template for creating a static policy for Verified Permissions."
Parameters:
  PolicyStoreId:
    Type: String
  Description:
    Type: String
  Statement:
    Type: String
Resources:
  StaticPolicy:
    Type: AWS::VerifiedPermissions::Policy
    Properties:
      PolicyStoreId: !Ref PolicyStoreId
      Definition:
        Static:
          Description: !Ref Description
          Statement: !Ref Statement
Outputs:
  PolicyId:
    Value: !GetAtt StaticPolicy.PolicyId

```

### Creating a template-linked policy

The following example creates a policy that is linked to the specified policy
template. You must specify the type and ID for the placeholders in your
template.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation sample template for creating a template-linked policy for Verified Permissions.",
    "Parameters": {
        "PolicyStoreId": {
            "Type": "String"
        },
        "PolicyTemplateId": {
            "Type": "String"
        },
        "PrincipalType": {
            "Type": "String"
        },
        "PrincipalId": {
            "Type": "String"
        },
        "ResourceType": {
            "Type": "String"
        },
        "ResourceId": {
            "Type": "String"
        }
    },
    "Resources": {
        "TemplateLinkedPolicy": {
            "Type": "AWS::VerifiedPermissions::Policy",
            "Properties": {
                "PolicyStoreId": {
                    "Ref": "PolicyStoreId"
                },
                "Definition": {
                    "TemplateLinked": {
                        "PolicyTemplateId": {
                            "Ref": "PolicyTemplateId"
                        },
                        "Principal": {
                            "EntityType": {
                                "Ref": "PrincipalType"
                            },
                            "EntityId": {
                                "Ref": "PrincipalId"
                            }
                        },
                        "Resource": {
                            "EntityType": {
                                "Ref": "ResourceType"
                            },
                            "EntityId": {
                                "Ref": "ResourceId"
                            }
                        }
                    }
                }
            }
        }
    },
    "Outputs": {
        "PolicyId": {
            "Value": {
                "Fn::GetAtt": [
                    "TemplateLinkedPolicy",
                    "PolicyId"
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: >-
  Description": "AWS CloudFormation sample template for creating a template-linked policy for Verified Permissions."
Parameters:
  PolicyStoreId:
    Type: String
  PolicyTemplateId:
    Type: String
  PrincipalType:
    Type: String
  PrincipalId:
    Type: String
  ResourceType:
    Type: String
  ResourceId:
    Type: String
Resources:
  TemplateLinkedPolicy:
    Type: AWS::VerifiedPermissions::Policy
    Properties:
      PolicyStoreId: !Ref PolicyStoreId
      Definition:
        TemplateLinked:
          PolicyTemplateId: !Ref PolicyTemplateId
          Principal:
            EntityType: !Ref PrincipalType
            EntityId: !Ref PrincipalId
          Resource:
            EntityType: !Ref ResourceType
            EntityId: !Ref ResourceId
Outputs:
  PolicyId:
    Value: !GetAtt TemplateLinkedPolicy.PolicyId

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OpenIdConnectTokenSelection

EntityIdentifier

All content copied from https://docs.aws.amazon.com/.
