This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::InstanceAccessControlAttributeConfiguration

Enables the attribute-based access control (ABAC) feature for the specified IAM Identity Center instance. You can also specify new attributes to add to your ABAC
configuration during the enabling process. For more information about ABAC, see [Attribute-Based\
Access Control](../../../singlesignon/latest/userguide/abac.md) in the _IAM Identity Center User Guide_.

###### Note

The `InstanceAccessControlAttributeConfiguration` property has been
deprecated but is still supported for backwards compatibility purposes. We recommend
that you use the `AccessControlAttributes` property instead.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSO::InstanceAccessControlAttributeConfiguration",
  "Properties" : {
      "AccessControlAttributes" : [ AccessControlAttribute, ... ],
      "InstanceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::SSO::InstanceAccessControlAttributeConfiguration
Properties:
  AccessControlAttributes:
    - AccessControlAttribute
  InstanceArn: String

```

## Properties

`AccessControlAttributes`

Lists the attributes that are configured for ABAC in the specified IAM Identity Center
instance.

_Required_: No

_Type_: Array of [AccessControlAttribute](aws-properties-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattribute.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The ARN of the IAM Identity Center instance under which the operation will be
executed.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[a-z]{1,5}){0,3}:sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

Specifies the IAM Identity Center identity store attributes to add to your ABAC
configuration. When using an external identity provider as an identity source, you can pass
attributes through the SAML assertion. Doing so provides an alternative to configuring
attributes from the IAM Identity Center identity store. If a SAML assertion passes any of
these attributes, IAM Identity Center will replace the attribute value with the value from
the IAM Identity Center identity store.

## Examples

### Enabling and configuring attributes used for access control in IAM Identity Center

The following example enables ABAC in IAM Identity Center and creates a new
attribute key `CostCenter` that is mapped to the Value
`“${path:enterprise.costCenter}”` which is coming from your identity
source.

#### JSON

```json

{
    "Resources": {
        "ABAC": {
            "Type": "AWS::SSO::InstanceAccessControlAttributeConfiguration",
            "Properties": {
                "InstanceArn": "arn:aws:sso:::instance/ssoins-instanceId",
                "AccessControlAttributes": [
                    {
                        "Key": "CostCenter",
                        "Value": {
                            "Source": [
                                "${path:enterprise.costCenter}"
                            ]
                        }
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  ABAC:
    Type: 'AWS::SSO::InstanceAccessControlAttributeConfiguration'
    Properties:
      InstanceArn: 'arn:aws:sso:::instance/ssoins-instanceId'
      AccessControlAttributes:
        - Key: CostCenter
          Value:
            Source:
              - '${path:enterprise.costCenter}'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AccessControlAttribute

All content copied from https://docs.aws.amazon.com/.
