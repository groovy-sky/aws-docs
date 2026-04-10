This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::SecurityConfiguration

Use a `SecurityConfiguration` resource to configure data encryption, Kerberos authentication (available in Amazon EMR release version 5.10.0 and later), and Amazon S3 authorization for EMRFS (available in EMR 5.10.0 and later). You can re-use a security configuration for any number of clusters in your account. For more information and example security configuration JSON objects, see [Create a Security Configuration](../../../emr/latest/managementguide/emr-create-security-configuration.md) in the _Amazon EMR Management Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMR::SecurityConfiguration",
  "Properties" : {
      "Name" : String,
      "SecurityConfiguration" : Json
    }
}

```

### YAML

```yaml

Type: AWS::EMR::SecurityConfiguration
Properties:
  Name: String
  SecurityConfiguration: Json

```

## Properties

`Name`

The name of the security configuration.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10280`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityConfiguration`

The security configuration details in JSON format. For JSON parameters and examples, see
[Use Security\
Configurations to Set Up Cluster Security](../../../emr/latest/managementguide/emr-security-configurations.md) in the _Amazon EMR_
_Management Guide_.

_Required_: Yes

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns returns the security configuration name, such as `mySecurityConfiguration`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VolumeSpecification

AWS::EMR::Step

All content copied from https://docs.aws.amazon.com/.
