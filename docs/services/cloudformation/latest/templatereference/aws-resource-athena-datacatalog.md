This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::DataCatalog

The AWS::Athena::DataCatalog resource specifies an Amazon Athena data catalog, which
contains a name, description, type, parameters, and tags. For more information, see
[DataCatalog](../../../../reference/athena/latest/apireference/api-datacatalog.md) in the _Amazon Athena API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Athena::DataCatalog",
  "Properties" : {
      "ConnectionType" : String,
      "Description" : String,
      "Error" : String,
      "Name" : String,
      "Parameters" : {Key: Value, ...},
      "Status" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Athena::DataCatalog
Properties:
  ConnectionType: String
  Description: String
  Error: String
  Name: String
  Parameters:
    Key: Value
  Status: String
  Tags:
    - Tag
  Type: String

```

## Properties

`ConnectionType`

The type of connection for a `FEDERATED` data catalog (for example,
`REDSHIFT`, `MYSQL`, or `SQLSERVER`). For
information about individual connectors, see [Available data source\
connectors](../../../athena/latest/ug/connectors-available.md).

_Required_: No

_Type_: String

_Allowed values_: `DYNAMODB | MYSQL | POSTGRESQL | REDSHIFT | ORACLE | SYNAPSE | SQLSERVER | DB2 | OPENSEARCH | BIGQUERY | GOOGLECLOUDSTORAGE | HBASE | DOCUMENTDB | CMDB | TPCDS | TIMESTREAM | SAPHANA | SNOWFLAKE | DATALAKEGEN2 | DB2AS400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the data catalog.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Error`

Text of the error that occurred during data catalog creation or deletion.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the data catalog. The catalog name must be unique for the AWS account and can use a maximum of 128 alphanumeric, underscore, at sign,
or hyphen characters.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Parameters`

Specifies the Lambda function or functions to use for creating the data
catalog. This is a mapping whose values depend on the catalog type.

- For the `HIVE` data catalog type, use the following syntax. The
`metadata-function` parameter is required. `The
                          sdk-version` parameter is optional and defaults to the currently
supported version.

`metadata-function=lambda_arn,
                              sdk-version=version_number`

- For the `LAMBDA` data catalog type, use one of the following sets
of required parameters, but not both.

- If you have one Lambda function that processes metadata
and another for reading the actual data, use the following syntax. Both
parameters are required.

`metadata-function=lambda_arn,
                                      record-function=lambda_arn`

- If you have a composite Lambda function that processes
both metadata and data, use the following syntax to specify your Lambda function.

`function=lambda_arn`

- The `GLUE` type takes a catalog ID parameter and is required. The
`
                                      catalog_id
                                  ` is the account ID of the
AWS account to which the AWS Glue Data Catalog
belongs.

`catalog-id=catalog_id`

- The `GLUE` data catalog type also applies to the default
`AwsDataCatalog` that already exists in your account, of
which you can have only one and cannot modify.

- The `FEDERATED` data catalog type uses one of the following
parameters, but not both. Use `connection-arn` for an existing
AWS Glue connection. Use `connection-type` and
`connection-properties` to specify the configuration setting for
a new connection.

- `connection-arn:<glue_connection_arn_to_reuse>`

- `lambda-role-arn` (optional): The execution role to use for the
Lambda function. If not provided, one is created.

- `connection-type:MYSQL|REDSHIFT|....,
                                      connection-properties:"<json_string>"`

For _`<json_string>`_, use escaped
JSON text, as in the following example.

`"{\"spill_bucket\":\"my_spill\",\"spill_prefix\":\"athena-spill\",\"host\":\"abc12345.snowflakecomputing.com\",\"port\":\"1234\",\"warehouse\":\"DEV_WH\",\"database\":\"TEST\",\"schema\":\"PUBLIC\",\"SecretArn\":\"arn:aws:secretsmanager:ap-south-1:111122223333:secret:snowflake-XHb67j\"}"`

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Maximum_: `51200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the creation or deletion of the data catalog.

- The `LAMBDA`, `GLUE`, and `HIVE` data catalog
types are created synchronously. Their status is either
`CREATE_COMPLETE` or `CREATE_FAILED`.

- The `FEDERATED` data catalog type is created asynchronously.

Data catalog creation status:

- `CREATE_IN_PROGRESS`: Federated data catalog creation in
progress.

- `CREATE_COMPLETE`: Data catalog creation complete.

- `CREATE_FAILED`: Data catalog could not be created.

- `CREATE_FAILED_CLEANUP_IN_PROGRESS`: Federated data catalog
creation failed and is being removed.

- `CREATE_FAILED_CLEANUP_COMPLETE`: Federated data catalog creation
failed and was removed.

- `CREATE_FAILED_CLEANUP_FAILED`: Federated data catalog creation
failed but could not be removed.

Data catalog deletion status:

- `DELETE_IN_PROGRESS`: Federated data catalog deletion in
progress.

- `DELETE_COMPLETE`: Federated data catalog deleted.

- `DELETE_FAILED`: Federated data catalog could not be
deleted.

_Required_: No

_Type_: String

_Allowed values_: `CREATE_IN_PROGRESS | CREATE_COMPLETE | CREATE_FAILED | CREATE_FAILED_CLEANUP_IN_PROGRESS | CREATE_FAILED_CLEANUP_COMPLETE | CREATE_FAILED_CLEANUP_FAILED | DELETE_IN_PROGRESS | DELETE_COMPLETE | DELETE_FAILED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags (key-value pairs) to associate with this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-athena-datacatalog-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of data catalog: `LAMBDA` for a federated catalog,
`GLUE` for AWS Glue Catalog, or `HIVE` for an external hive
metastore.

_Required_: Yes

_Type_: String

_Allowed values_: `LAMBDA | GLUE | HIVE | FEDERATED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the data catalog.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create an Athena data catalog

The following example template creates a custom Hive data catalog in
Athena.

#### JSON

```json

{
    "Resources":{
        "MyAthenaDataCatalog":{
            "Type":"AWS::Athena::DataCatalog",
            "Properties":{
                "Name":"MyCustomDataCatalog",
                "Type":"HIVE",
                "Description":"Custom Hive Catalog Description",
                "Tags":[
                    {
                        "Key":"key1",
                        "Value":"value1"
                    },
                    {
                        "Key":"key2",
                        "Value":"value2"
                    }
                ],
                "Parameters":{
                    "metadata-function":"arn:aws:lambda:us-west-2:111122223333:function:lambdaname"
                }
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  MyAthenaDataCatalog:
    Type: AWS::Athena::DataCatalog
    Properties:
      Name: MyCustomDataCatalog
      Type: HIVE
      Description: Custom Hive Catalog Description
      Tags:
        - Key: "key1"
          Value: "value1"
        - Key: "key2"
          Value: "value2"
      Parameters:
        metadata-function: "arn:aws:lambda:us-west-2:111122223333:function:lambdaname"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
