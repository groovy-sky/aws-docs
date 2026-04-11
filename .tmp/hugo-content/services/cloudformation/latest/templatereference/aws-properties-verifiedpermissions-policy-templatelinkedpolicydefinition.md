This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::Policy TemplateLinkedPolicyDefinition

A structure that describes a policy created by instantiating a policy template.

###### Note

You can't directly update a template-linked policy. You must update the associated
policy template instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PolicyTemplateId" : String,
  "Principal" : EntityIdentifier,
  "Resource" : EntityIdentifier
}

```

### YAML

```yaml

  PolicyTemplateId: String
  Principal:
    EntityIdentifier
  Resource:
    EntityIdentifier

```

## Properties

`PolicyTemplateId`

The unique identifier of the policy template used to create this policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]*$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: Updates are not supported.

`Principal`

The principal associated with this template-linked policy. Verified Permissions substitutes this principal for the
`?principal` placeholder in the policy template when it evaluates an authorization
request.

_Required_: No

_Type_: [EntityIdentifier](aws-properties-verifiedpermissions-policy-entityidentifier.md)

_Update requires_: Updates are not supported.

`Resource`

The resource associated with this template-linked policy. Verified Permissions substitutes this resource for the
`?resource` placeholder in the policy template when it evaluates an authorization
request.

_Required_: No

_Type_: [EntityIdentifier](aws-properties-verifiedpermissions-policy-entityidentifier.md)

_Update requires_: Updates are not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StaticPolicyDefinition

AWS::VerifiedPermissions::PolicyStore

All content copied from https://docs.aws.amazon.com/.
