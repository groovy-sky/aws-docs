# CloudWatch Logs Insights language query syntax

This section provides details about the Logs Insights QL. The query syntax
supports different functions and operations that include but aren't limited to
general functions, arithmetic and comparison operations, and regular
expressions.

###### Important

To avoid incurring excessive charges by running large queries, keep in
mind the following best practices:

- Select only the necessary log groups for each query.

- Always specify the narrowest possible time range for your
queries.

- When you use the console to run queries, cancel all your queries
before you close the CloudWatch Logs Insights console page. Otherwise, queries
continue to run until completion.

- When you add a CloudWatch Logs Insights widget to a dashboard, ensure that the
dashboard is not refreshing at a high frequency, because each
refresh starts a new query.

To create queries that contain multiple commands, separate the commands with
the pipe character ( **\|**).

To create queries that contain comments, set off the comments with the hash
character ( **#**).

###### Note

CloudWatch Logs Insights automatically discovers fields for different log types and
generates fields that start with the **@** character. For
more information about these fields, see [Supported logs and discovered fields](https://docs.aws.amazon.com/en_us/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData-discoverable-fields.html) in the _Amazon CloudWatch_
_User Guide_.

The following table briefly describes each command. Following this table is a
more comprehensive description of each command, with examples.

###### Note

All Logs Insights QL query commands are supported on log groups in the
Standard log class. Log groups in the Infrequent Access log class support
all Logs Insights QL query commands except `pattern`,
`diff`, and `unmask`.

**`**
**anomaly`**

Identifies unusual patterns in your log data using
machine learning.

**`**
**display`**

Displays a specific field or fields in query results.

**`**
**fields`**

Displays specific fields in query results and supports
functions and operations you can use to modify field values
and create new fields to use in your query.

**`**
**filter`**

Filters the query to return only the log events that
match one or more conditions.

**`**
**filterIndex`**

Forces a query to attempt to scan only the log groups
that are both indexed on the field mentioned in a field
index and also contain a value for the that field index.
This reduces scanned volume by attempting to scan only log
events from these log groups that contain the value
specified in the query for this field index.

This command is not supported for log groups in the
Infrequent Access log class.

**`**
**pattern`**

Automatically clusters your log data into patterns. A
pattern is shared text structure that recurs among your log
fields. CloudWatch Logs Insights provides ways for you to analyze the
patterns found in your log events. For more information, see
[Pattern analysis](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData_Patterns.html).

**`**
**diff`**

Compares the log events found in your requested time
period with the log events from a previous time period of
equal length, so that you can look for trends and find out
if certain log events are new.

**`**
**parse`**

Extracts data from a log field to create an extracted
field that you can process in your query.
**`parse`** supports
both glob mode using wildcards, and regular expressions.

**`**
**sort`**

Displays the returned log events in ascending
( `asc`) or descending ( `desc`)
order.

**`**
**SOURCE`**

Including `SOURCE` in a query is a useful
way to specify a large amount of log groups based on log
group name prefix, account identifiers, and log group class
to include in a query. This command is supported only when
you create a query in the AWS CLI or programmatically, not in
the CloudWatch console.

**`**
**stats`**

Calculate aggregate statistics using values in the log
fields.

**`**
**limit`**

Specifies a maximum number of log events that you want
your query to return. Useful with
**`sort`** to return
"top 20" or "most recent 20" results.

**`**
**dedup`**

Removes duplicate results based on specific values in
fields that you specify.

**`**
**unmask`**

Displays all the content of a log event that has some
content masked because of a data protection policy. For more
information about data protection in log groups, see [Help protect sensitive log data with masking](mask-sensitive-log-data.md).

**`unnest`**

Flattens a list taken as input to produce multiple
records with a single record for each element in the list.

**`**
**lookup`**

Enriches log events with data from a lookup table
by matching field values. Use lookup tables to add
reference data such as user details, application names,
or product information to your query results.

**[Other operations and\**
**functions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-operations-functions.html)**

CloudWatch Logs Insights also supports many comparison, arithmetic,
datetime, numeric, string, IP address, and general functions
and operations.

The following sections provide more details about the CloudWatch Logs Insights query
commands.

###### Topics

- [Logs Insights QL commands supported in log classes](cwl-analyzelogdata-classes.md)

- [anomaly](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Anomaly.html)

- [display](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Display.html)

- [fields](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Fields.html)

- [filter](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Filter.html)

- [filterIndex](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-FilterIndex.html)

- [SOURCE](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Source.html)

- [pattern](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Pattern.html)

- [diff](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Diff.html)

- [parse](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Parse.html)

- [sort](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Sort.html)

- [stats](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Stats.html)

- [limit](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Limit.html)

- [dedup](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Dedup.html)

- [unmask](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Unmask.html)

- [unnest](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Unnest.html)

- [lookup](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Lookup.html)

- [Boolean, comparison, numeric, datetime, and other functions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-operations-functions.html)

- [Fields that contain special characters](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-Guidelines.html)

- [Use aliases and comments in queries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax-alias.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CloudWatch Logs Insights query language (Logs Insights QL)

Logs Insights QL commands supported in log classes
