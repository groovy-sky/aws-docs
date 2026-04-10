This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetRedshiftDataParameters

These are custom parameters to be used when the target is a Amazon Redshift cluster to invoke the
Amazon Redshift Data API BatchExecuteStatement.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Database" : String,
  "DbUser" : String,
  "SecretManagerArn" : String,
  "Sqls" : [ String, ... ],
  "StatementName" : String,
  "WithEvent" : Boolean
}

```

### YAML

```yaml

  Database: String
  DbUser: String
  SecretManagerArn: String
  Sqls:
    - String
  StatementName: String
  WithEvent: Boolean

```

## Properties

`Database`

The name of the database. Required when authenticating using temporary
credentials.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DbUser`

The database user name. Required when authenticating using temporary credentials.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretManagerArn`

The name or ARN of the secret that enables access to the database. Required when
authenticating using Secrets Manager.

_Required_: No

_Type_: String

_Pattern_: `^(^arn:aws([a-z]|\-)*:secretsmanager:([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}):(\d{12}):secret:.+)|(\$(\.[\w/_-]+(\[(\d+|\*)\])*)*)$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sqls`

The SQL statement text to run.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `100000 | 40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatementName`

The name of the SQL statement. You can name the SQL statement when you create it to
identify the query.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WithEvent`

Indicates whether to send an event back to EventBridge after the SQL statement
runs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeTargetParameters

PipeTargetSageMakerPipelineParameters

All content copied from https://docs.aws.amazon.com/.
