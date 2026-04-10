This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DocDB::DBInstance

The `AWS::DocDB::DBInstance` Amazon DocumentDB (with MongoDB compatibility) resource describes a DBInstance.
For more information, see [DBInstance](../../../documentdb/latest/developerguide/api-dbinstance.md) in the _Amazon DocumentDB Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DocDB::DBInstance",
  "Properties" : {
      "AutoMinorVersionUpgrade" : Boolean,
      "AvailabilityZone" : String,
      "CACertificateIdentifier" : String,
      "CertificateRotationRestart" : Boolean,
      "DBClusterIdentifier" : String,
      "DBInstanceClass" : String,
      "DBInstanceIdentifier" : String,
      "EnablePerformanceInsights" : Boolean,
      "PreferredMaintenanceWindow" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DocDB::DBInstance
Properties:
  AutoMinorVersionUpgrade: Boolean
  AvailabilityZone: String
  CACertificateIdentifier: String
  CertificateRotationRestart: Boolean
  DBClusterIdentifier: String
  DBInstanceClass: String
  DBInstanceIdentifier: String
  EnablePerformanceInsights: Boolean
  PreferredMaintenanceWindow: String
  Tags:
    - Tag

```

## Properties

`AutoMinorVersionUpgrade`

This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB does not perform minor version upgrades regardless of the value set.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZone`

The Amazon EC2 Availability Zone that the instance is created in.

Default: A random, system-chosen Availability Zone in the endpoint's AWS Region.

Example: `us-east-1d`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CACertificateIdentifier`

The identifier of the CA certificate for this DB instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CertificateRotationRestart`

Specifies whether the DB instance is restarted when you rotate your
SSL/TLS certificate.

By default, the DB instance is restarted when you rotate your SSL/TLS certificate. The certificate
is not updated until the DB instance is restarted.

###### Important

Set this parameter only if you are _not_ using SSL/TLS to connect to the DB instance.

If you are using SSL/TLS to connect to the DB instance, see [Updating Your Amazon DocumentDB TLS \
Certificates](../../../documentdb/latest/developerguide/ca-cert-rotation.md) and
[Encrypting Data in Transit](../../../documentdb/latest/developerguide/security-encryption-ssl.md) in the _Amazon DocumentDB Developer_
_Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBClusterIdentifier`

The identifier of the cluster that the instance will belong to.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DBInstanceClass`

The compute and memory capacity of the instance; for example,
`db.m4.large`. If you change the class of an instance there
can be some interruption in the cluster's service.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBInstanceIdentifier`

The instance identifier. This parameter is stored as a lowercase string.

Constraints:

- Must contain from 1 to 63 letters, numbers, or hyphens.

- The first character must be a letter.

- Cannot end with a hyphen or contain two consecutive hyphens.

Example: `mydbinstance`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnablePerformanceInsights`

A value that indicates whether to enable Performance Insights for the DB Instance. For
more information, see [Using Amazon\
Performance Insights](../../../documentdb/latest/developerguide/performance-insights.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreferredMaintenanceWindow`

The time range each week during which system maintenance can occur, in Universal
Coordinated Time (UTC).

Format: `ddd:hh24:mi-ddd:hh24:mi`

The default is a 30-minute window selected at random from an 8-hour block of time for
each AWS Region, occurring on a random day of the week.

Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun

Constraints: Minimum 30-minute window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be assigned to the instance. You can assign up to
10 tags to an instance.

_Required_: No

_Type_: Array of [Tag](aws-properties-docdb-dbinstance-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DBInstance's name, such as `sample-cluster-instance`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Endpoint`

The connection endpoint for the instance.
For example: `sample-cluster.cluster-abcdefghijkl.us-east-1.docdb.amazonaws.com`.

`Port`

The port number on which the database accepts connections, such as `27017`.

## Examples

#### JSON

```json

{
   "Type" : "AWS::DocDB::DBInstance",
   "Properties" : {
      "AutoMinorVersionUpgrade" : true,
      "AvailabilityZone" : "us-east-1c",
      "DBClusterIdentifier" : "sample-cluster",
      "DBInstanceClass" : "db.r5.large",
      "DBInstanceIdentifier" : "sample-cluster-instance-0",
      "PreferredMaintenanceWindow" : "sat:06:54-sat:07:24",
      "Tags" : [{ "Key": "String","Value": "String" }]
   }
}
```

#### YAML

```yaml

Type: "AWS::DocDB::DBInstance"
Properties:
   AutoMinorVersionUpgrade: true
   AvailabilityZone: "us-east-1c"
   DBClusterIdentifier: "sample-cluster"
   DBInstanceClass: "db.r5.large"
   DBInstanceIdentifier: "sample-cluster-instance-0"
   PreferredMaintenanceWindow: "sat:06:54-sat:07:24"
   Tags:
      -
         Key: "String"
         Value: "String"
```

## See also

- [DBInstance](../../../documentdb/latest/developerguide/api-dbinstance.md)

- [CreateDBInstance](../../../documentdb/latest/developerguide/api-createdbinstance.md)

- [DeleteDBInstance](../../../documentdb/latest/developerguide/api-deletedbinstance.md)

- [DescribeDBInstances](../../../documentdb/latest/developerguide/api-describedbinstances.md)

- [ModifyDBInstance](../../../documentdb/latest/developerguide/api-modifydbinstance.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
