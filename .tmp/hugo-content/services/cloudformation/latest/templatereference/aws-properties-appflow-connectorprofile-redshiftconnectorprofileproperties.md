This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile RedshiftConnectorProfileProperties

The connector-specific profile properties when using Amazon Redshift.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String,
  "ClusterIdentifier" : String,
  "DataApiRoleArn" : String,
  "DatabaseName" : String,
  "DatabaseUrl" : String,
  "IsRedshiftServerless" : Boolean,
  "RoleArn" : String,
  "WorkgroupName" : String
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String
  ClusterIdentifier: String
  DataApiRoleArn: String
  DatabaseName: String
  DatabaseUrl: String
  IsRedshiftServerless: Boolean
  RoleArn: String
  WorkgroupName: String

```

## Properties

`BucketName`

A name for the associated Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

The object key for the destination bucket in which Amazon AppFlow places the files.

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterIdentifier`

The unique ID that's assigned to an Amazon Redshift cluster.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataApiRoleArn`

The Amazon Resource Name (ARN) of an IAM role that permits Amazon AppFlow to access your Amazon Redshift database through the Data API. For more
information, and for the polices that you attach to this role, see [Allow Amazon AppFlow to access Amazon Redshift databases with the Data\
API](../../../appflow/latest/userguide/security-iam-service-role-policies.md#access-redshift).

_Required_: No

_Type_: String

_Pattern_: `arn:aws:iam:.*:[0-9]+:.*`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The name of an Amazon Redshift database.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseUrl`

The JDBC URL of the Amazon Redshift cluster.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsRedshiftServerless`

Indicates whether the connector profile defines a connection to an Amazon Redshift
Serverless data warehouse.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of IAM role that grants Amazon Redshift
read-only access to Amazon S3. For more information, and for the polices that you
attach to this role, see [Allow Amazon Redshift to access your Amazon AppFlow data in Amazon S3](../../../appflow/latest/userguide/security-iam-service-role-policies.md#redshift-access-s3).

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws:iam:.*:[0-9]+:.*`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkgroupName`

The name of an Amazon Redshift workgroup.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [RedshiftConnectorProfileProperties](../../../../reference/appflow/1-0/apireference/api-redshiftconnectorprofileproperties.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftConnectorProfileCredentials

SalesforceConnectorProfileCredentials

All content copied from https://docs.aws.amazon.com/.
