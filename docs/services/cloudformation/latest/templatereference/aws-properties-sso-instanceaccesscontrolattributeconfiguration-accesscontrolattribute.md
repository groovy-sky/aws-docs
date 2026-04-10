This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::InstanceAccessControlAttributeConfiguration AccessControlAttribute

These are IAM Identity Center identity store attributes that you can configure for use
in attributes-based access control (ABAC). You can create permissions policies that
determine who can access your AWS resources based upon the configured
attribute values. When you enable ABAC and specify `AccessControlAttributes`,
IAM Identity Center passes the attribute values of the authenticated user into IAM for
use in policy evaluation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : AccessControlAttributeValue
}

```

### YAML

```yaml

  Key: String
  Value:
    AccessControlAttributeValue

```

## Properties

`Key`

The name of the attribute associated with your identities in your identity source. This
is used to map a specified attribute in your identity source with an attribute in IAM Identity Center.

_Required_: Yes

_Type_: String

_Pattern_: `[\p{L}\p{Z}\p{N}_.:\/=+\-@]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value used for mapping a specified attribute to an identity source.

_Required_: Yes

_Type_: [AccessControlAttributeValue](aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSO::InstanceAccessControlAttributeConfiguration

AccessControlAttributeValue

All content copied from https://docs.aws.amazon.com/.
