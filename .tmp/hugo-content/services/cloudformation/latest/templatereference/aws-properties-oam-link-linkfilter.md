This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Oam::Link LinkFilter

When used in `MetricConfiguration` this field specifies which metric namespaces are to be shared with the monitoring account

When used in `LogGroupConfiguration` this field specifies which log groups are to share their log events with the monitoring account. Use the term `LogGroupName` and one or
more of the following operands.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filter" : String
}

```

### YAML

```yaml

  Filter: String

```

## Properties

`Filter`

When used in `MetricConfiguration` this field specifies which metric namespaces are to be shared with the monitoring account

When used in `LogGroupConfiguration` this field specifies which log groups are to share their log events with the monitoring account. Use the term `LogGroupName` and one or
more of the following operands.

Use single quotation marks (') around log group names and metric namespaces.

The matching of log group names and metric namespaces is case sensitive.
Each filter has a limit of five conditional operands. Conditional operands are `AND` and `OR`.

- `=` and `!=`

- `AND`

- `OR`

- `LIKE` and `NOT LIKE`. These can be used only as prefix searches. Include a `%` at the end
of the string that you want to search for and include.

- `IN` and `NOT IN`, using parentheses `( )`

Examples:

- `Namespace NOT LIKE 'AWS/%'` includes only namespaces that don't start with `AWS/`, such as custom namespaces.

- `Namespace IN ('AWS/EC2', 'AWS/ELB', 'AWS/S3')` includes only the metrics in the EC2, Elastic Load Balancing, and Amazon S3 namespaces.

- `Namespace = 'AWS/EC2' OR Namespace NOT LIKE 'AWS/%'` includes only the EC2 namespace and your custom namespaces.

- `LogGroupName IN ('This-Log-Group', 'Other-Log-Group')` includes only the log groups with names `This-Log-Group` and
`Other-Log-Group`.

- `LogGroupName NOT IN ('Private-Log-Group', 'Private-Log-Group-2')` includes all log groups except the log groups with names `Private-Log-Group` and
`Private-Log-Group-2`.

- `LogGroupName LIKE 'aws/lambda/%' OR LogGroupName LIKE 'AWSLogs%'` includes all log groups that have names that start with `aws/lambda/` or
`AWSLogs`.

###### Note

If you are updating a link that uses filters, you can specify `*` as the only value for the
`filter` parameter to delete the filter and share all log groups with the monitoring account.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LinkConfiguration

AWS::Oam::Sink

All content copied from https://docs.aws.amazon.com/.
