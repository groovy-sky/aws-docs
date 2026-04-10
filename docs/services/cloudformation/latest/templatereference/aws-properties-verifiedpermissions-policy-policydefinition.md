This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::Policy PolicyDefinition

A structure that defines a Cedar policy. It includes the policy type, a description,
and a policy body. This is a top level data type used to create a policy.

This data type is used as a request parameter for the [CreatePolicy](../../../../reference/verifiedpermissions/latest/apireference/api-createpolicy.md) operation. This structure must always have either an
`Static` or a `TemplateLinked` element.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Static" : StaticPolicyDefinition,
  "TemplateLinked" : TemplateLinkedPolicyDefinition
}

```

### YAML

```yaml

  Static:
    StaticPolicyDefinition
  TemplateLinked:
    TemplateLinkedPolicyDefinition

```

## Properties

`Static`

A structure that describes a static policy. An static policy doesn't use a template or allow
placeholders for entities.

_Required_: No

_Type_: [StaticPolicyDefinition](aws-properties-verifiedpermissions-policy-staticpolicydefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateLinked`

A structure that describes a policy that was instantiated from a template. The
template can specify placeholders for `principal` and `resource`.
When you use [CreatePolicy](../../../../reference/verifiedpermissions/latest/apireference/api-createpolicy.md) to create a policy from a template, you specify the exact
principal and resource to use for the instantiated policy.

_Required_: No

_Type_: [TemplateLinkedPolicyDefinition](aws-properties-verifiedpermissions-policy-templatelinkedpolicydefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EntityIdentifier

StaticPolicyDefinition

All content copied from https://docs.aws.amazon.com/.
