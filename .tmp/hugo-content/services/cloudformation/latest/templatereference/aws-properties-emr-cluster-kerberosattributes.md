This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster KerberosAttributes

`KerberosAttributes` is a property of the `AWS::EMR::Cluster` resource. `KerberosAttributes` define the cluster-specific Kerberos configuration when Kerberos authentication is enabled using a security configuration. The cluster-specific configuration must be compatible with the security configuration. For more information see [Use Kerberos Authentication](../../../emr/latest/managementguide/emr-kerberos.md) in the _EMR Management Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ADDomainJoinPassword" : String,
  "ADDomainJoinUser" : String,
  "CrossRealmTrustPrincipalPassword" : String,
  "KdcAdminPassword" : String,
  "Realm" : String
}

```

### YAML

```yaml

  ADDomainJoinPassword: String
  ADDomainJoinUser: String
  CrossRealmTrustPrincipalPassword: String
  KdcAdminPassword: String
  Realm: String

```

## Properties

`ADDomainJoinPassword`

The Active Directory password for `ADDomainJoinUser`.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ADDomainJoinUser`

Required only when establishing a cross-realm trust with an Active Directory domain. A
user with sufficient privileges to join resources to the domain.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CrossRealmTrustPrincipalPassword`

Required only when establishing a cross-realm trust with a KDC in a different realm. The
cross-realm principal password, which must be identical across realms.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KdcAdminPassword`

The password used within the cluster for the kadmin service on the cluster-dedicated
KDC, which maintains Kerberos principals, password policies, and keytabs for the
cluster.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Realm`

The name of the Kerberos realm to which all nodes in a cluster belong. For example,
`EC2.INTERNAL`.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JobFlowInstancesConfig

KeyValue

All content copied from https://docs.aws.amazon.com/.
