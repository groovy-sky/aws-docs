This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::OdbNetwork

The `AWS::ODB::OdbNetwork` resource creates an ODB network. An ODB network provides the networking foundation for Oracle Database resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ODB::OdbNetwork",
  "Properties" : {
      "AvailabilityZone" : String,
      "AvailabilityZoneId" : String,
      "BackupSubnetCidr" : String,
      "ClientSubnetCidr" : String,
      "CrossRegionS3RestoreSources" : [ String, ... ],
      "CustomDomainName" : String,
      "DefaultDnsPrefix" : String,
      "DeleteAssociatedResources" : Boolean,
      "DisplayName" : String,
      "KmsAccess" : String,
      "KmsPolicyDocument" : String,
      "S3Access" : String,
      "S3PolicyDocument" : String,
      "StsAccess" : String,
      "StsPolicyDocument" : String,
      "Tags" : [ Tag, ... ],
      "ZeroEtlAccess" : String
    }
}

```

### YAML

```yaml

Type: AWS::ODB::OdbNetwork
Properties:
  AvailabilityZone: String
  AvailabilityZoneId: String
  BackupSubnetCidr: String
  ClientSubnetCidr: String
  CrossRegionS3RestoreSources:
    - String
  CustomDomainName: String
  DefaultDnsPrefix: String
  DeleteAssociatedResources: Boolean
  DisplayName: String
  KmsAccess: String
  KmsPolicyDocument: String
  S3Access: String
  S3PolicyDocument: String
  StsAccess: String
  StsPolicyDocument: String
  Tags:
    - Tag
  ZeroEtlAccess: String

```

## Properties

`AvailabilityZone`

The Availability Zone (AZ) where the ODB network is located.

Required when creating an ODB network. Specify either AvailabilityZone or AvailabilityZoneId to define the location of the network.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AvailabilityZoneId`

The AZ ID of the AZ where the ODB network is located.

Required when creating an ODB network. Specify either AvailabilityZone or AvailabilityZoneId to define the location of the network.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BackupSubnetCidr`

The CIDR range of the backup subnet in the ODB network.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ClientSubnetCidr`

The CIDR range of the client subnet in the ODB network.

Required when creating an ODB network.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CrossRegionS3RestoreSources`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomDomainName`

The domain name for the resources in the ODB network.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultDnsPrefix`

The DNS prefix to the default DNS domain name. The default DNS domain name is oraclevcn.com.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeleteAssociatedResources`

Specifies whether to delete associated OCI networking resources along with the ODB network.

Required when creating an ODB network.

_Required_: Conditional

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The user-friendly name of the ODB network.

Required when creating an ODB network.

_Required_: Conditional

_Type_: String

_Pattern_: `^[a-zA-Z_](?!.*--)[a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsAccess`

Configuration for AWS Key Management Service (KMS) access from the ODB network.

_Required_: No

_Type_: [String](aws-properties-odb-odbnetwork-kmsaccess.md)

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsPolicyDocument`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Access`

The configuration for Amazon S3 access from the ODB network.

_Required_: No

_Type_: [String](aws-properties-odb-odbnetwork-s3access.md)

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3PolicyDocument`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StsAccess`

Configuration for AWS Security Token Service (STS) access from the ODB network.

_Required_: No

_Type_: [String](aws-properties-odb-odbnetwork-stsaccess.md)

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StsPolicyDocument`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to assign to the Odb Network.

_Required_: No

_Type_: Array of [Tag](aws-properties-odb-odbnetwork-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZeroEtlAccess`

The configuration for Zero-ETL access from the ODB network.

_Required_: No

_Type_: [String](aws-properties-odb-odbnetwork-zeroetlaccess.md)

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier of the ODB network. For example:

`{ "Ref": "myOdbNetwork" }`

For the ODB network `myOdbNetwork`, `Ref` returns the unique identifier of the ODB network.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`OciNetworkAnchorId`

The unique identifier of the OCI network anchor for the ODB network.

`OciResourceAnchorName`

The name of the OCI resource anchor that's associated with the ODB network.

`OciVcnUrl`

The URL for the VCN that's associated with the ODB network.

`OdbNetworkArn`

The Amazon Resource Name (ARN) of the ODB network.

`OdbNetworkId`

The unique identifier of the ODB network.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CrossRegionS3RestoreSourcesAccess

All content copied from https://docs.aws.amazon.com/.
