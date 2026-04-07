# SQL statistics for Performance Insights

_SQL statistics_ are performance-related metrics about SQL queries that are collected by Performance Insights.
Performance Insights gathers statistics for each second that a query is running and for each SQL call. The SQL statistics
are an average for the selected time range.

A SQL digest is a composite of all queries having a given pattern but not necessarily having the same literal values.
The digest replaces literal values with a question mark. For example, `SELECT * FROM emp WHERE lname= ?`.
This digest might consist of the following child queries:

```sql

SELECT * FROM emp WHERE lname = 'Sanchez'
SELECT * FROM emp WHERE lname = 'Olagappan'
SELECT * FROM emp WHERE lname = 'Wu'
```

All engines support SQL statistics for digest queries.

For the region, DB engine, and instance class support information for this feature, see
[Amazon Aurora DB engine, Region, and instance class support for Performance Insights features](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.Overview.Engines.html#USER_PerfInsights.Overview.PIfeatureEngnRegSupport)

###### Topics

- [SQL statistics for Aurora MySQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.MySQL.html)

- [SQL statistics for Aurora PostgreSQL](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.PostgreSQL.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Counter metrics for Performance Insights

SQL statistics for Aurora MySQL
