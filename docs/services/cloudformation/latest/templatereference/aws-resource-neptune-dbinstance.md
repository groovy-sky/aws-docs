This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Neptune::DBInstance

The `AWS::Neptune::DBInstance` type creates an Amazon Neptune DB instance.

**Updating DB Instances**

You can set a deletion policy for your DB instance to control how CloudFormation handles the
instance when the stack is deleted. For Neptune DB instances, you can choose to
_retain_ the instance, to _delete_ the
instance, or to _create a snapshot_ of the instance. The default
CloudFormation behavior depends on the `DBClusterIdentifier` property:

- For `AWS::Neptune::DBInstance` resources that don't specify the
`DBClusterIdentifier` property, CloudFormation saves a snapshot of the DB
instance.

- For `AWS::Neptune::DBInstance` resources that do specify the
`DBClusterIdentifier` property, CloudFormation deletes the DB
instance.

**Deleting DB Instances**

###### Important

If a DB instance is deleted or replaced during an update, CloudFormation deletes all automated
snapshots. However, it retains manual DB snapshots. During an update that requires
replacement, you can apply a stack policy to prevent DB instances from being replaced.

When properties labeled _Update requires: Replacement_ are updated,
CloudFormation first creates a
replacement DB instance, changes references from other dependent resources to point to
the replacement DB instance, and finally deletes the old DB instance.

###### Important

We highly recommend that you take a snapshot of the database before updating the
stack. If you don't, you lose the data when CloudFormation replaces your DB instance. To
preserve your data, perform the following procedure:

1. Deactivate any applications that are using the DB instance so that there's
    no activity on the DB instance.

2. Create a snapshot of the DB instance.

3. If you want to restore your instance using a DB snapshot, modify the updated
    template with your DB instance changes and add the
    `DBSnapshotIdentifier` property with the ID of the DB snapshot
    that you want to use.

4. Update the stack.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Neptune::DBInstance",
  "Properties" : {
      "AllowMajorVersionUpgrade" : Boolean,
      "AutoMinorVersionUpgrade" : Boolean,
      "AvailabilityZone" : String,
      "DBClusterIdentifier" : String,
      "DBInstanceClass" : String,
      "DBInstanceIdentifier" : String,
      "DBParameterGroupName" : String,
      "DBSubnetGroupName" : String,
      "PreferredMaintenanceWindow" : String,
      "PubliclyAccessible" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Neptune::DBInstance
Properties:
  AllowMajorVersionUpgrade: Boolean
  AutoMinorVersionUpgrade: Boolean
  AvailabilityZone: String
  DBClusterIdentifier: String
  DBInstanceClass: String
  DBInstanceIdentifier: String
  DBParameterGroupName: String
  DBSubnetGroupName: String
  PreferredMaintenanceWindow: String
  PubliclyAccessible: Boolean
  Tags:
    - Tag

```

## Properties

`AllowMajorVersionUpgrade`

Indicates that major version upgrades are allowed. Changing this
parameter doesn't result in an outage and the change is asynchronously
applied as soon as possible. This parameter must be set to true when specifying
a value for the EngineVersion parameter that is a different major version than
the DB instance's current version.

###### Warning

When you change this parameter for an existing DB cluster, CloudFormation will replace your existing DB cluster
with a new, empty one that uses the engine version you specified.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoMinorVersionUpgrade`

Indicates that minor version patches are applied automatically.

When updating this property, some interruptions may occur.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`AvailabilityZone`

Specifies the name of the Availability Zone the DB instance is located in.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBClusterIdentifier`

If the DB instance is a member of a DB cluster, contains the name of the DB cluster that
the DB instance is a member of.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBInstanceClass`

Contains the name of the compute and memory capacity class of the DB instance.

If you update this property, some interruptions may occur.

_Required_: Yes

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DBInstanceIdentifier`

Contains a user-supplied database identifier. This identifier is the unique key that
identifies a DB instance.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBParameterGroupName`

The name of an existing DB parameter group or a reference to an
AWS::Neptune::DBParameterGroup resource created in the template.
If any of the data members of the referenced parameter
group are changed during an update, the DB instance might need to be restarted,
which causes some interruption. If the parameter group contains static parameters,
whether they were changed or not, an update triggers a reboot.

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`DBSubnetGroupName`

A DB subnet group to associate with the DB instance. If you update this value,
the new subnet group must be a subnet group in a new virtual private cloud
(VPC).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreferredMaintenanceWindow`

Specifies the weekly time range during which system maintenance can occur, in Universal
Coordinated Time (UTC).

_Required_: No

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PubliclyAccessible`

Indicates whether the DB instance is publicly accessible.

When the DB instance is publicly accessible and you connect from outside of the DB instance's virtual private
cloud (VPC), its Domain Name System (DNS) endpoint resolves to the public IP address. When you connect from within
the same VPC as the DB instance, the endpoint resolves to the private IP address. Access to the DB instance is
ultimately controlled by the security group it uses. That public access isn't permitted if the security group assigned
to the DB cluster doesn't permit it.

When the DB instance isn't publicly accessible, it is an internal DB instance with a DNS name that resolves to a
private IP address.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An arbitrary set of tags (key-value pairs) for this DB instance.

_Required_: No

_Type_: Array of [Tag](aws-properties-neptune-dbinstance-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Endpoint`

The connection endpoint for the database. For example:
`mystack-mydb-1apw1j4phylrk.cg034hpkmmjt.us-east-2.rds.amazonaws.com`.

`Port`

The port number on which the database accepts connections. For example:
`8182`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
