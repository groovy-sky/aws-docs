# Retrieving metrics with the Performance Insights API for Aurora

When Performance Insights is turned on, the API provides visibility into instance
performance. Amazon CloudWatch Logs provides the authoritative source for vended monitoring metrics for
AWS services.

Performance Insights offers a domain-specific view of database load measured as average active sessions (AAS). This
metric appears to API consumers as a two-dimensional time-series dataset. The time dimension of the data provides DB
load data for each time point in the queried time range. Each time point decomposes overall load in relation to the
requested dimensions, such as `SQL`, `Wait-event`, `User`, or `Host`,
measured at that time point.

Amazon RDS Performance Insights monitors your Amazon Aurora cluster so that you can analyze and troubleshoot database performance. One
way to view Performance Insights data is in the AWS Management Console. Performance Insights also provides a public API so that
you can query your own data. You can use the API to do the following:

- Offload data into a database

- Add Performance Insights data to existing monitoring dashboards

- Build monitoring tools

To use the Performance Insights API, enable Performance Insights on one of your Amazon RDS DB instances. For information
about enabling Performance Insights, see [Turning Performance Insights on and off for Aurora](user-perfinsights-enabling.md). For more information about the Performance Insights API, see the
[Amazon RDS Performance Insights API\
Reference](../../../../reference/performance-insights/latest/apireference/welcome.md).

The Performance Insights API provides the following operations.

Performance Insights action

AWS CLI command

Description

[`CreatePerformanceAnalysisReport`](../../../../reference/performance-insights/latest/apireference/api-createperformanceanalysisreport.md)

[`aws pi\
							create-performance-analysis-report`](../../../cli/latest/reference/pi/createperformanceanalysisreport.md)

Creates a performance analysis report for a specific time period for the DB instance.
The result is `AnalysisReportId` which is the unique
identifier of the report.

[`DeletePerformanceAnalysisReport`](../../../../reference/performance-insights/latest/apireference/api-deleteperformanceanalysisreport.md)

[`aws pi\
							delete-performance-analysis-report`](../../../cli/latest/reference/pi/deleteperformanceanalysisreport.md)

Deletes a performance analysis report.

[`DescribeDimensionKeys`](../../../../reference/performance-insights/latest/apireference/api-describedimensionkeys.md)

[`aws pi\
									describe-dimension-keys`](../../../cli/latest/reference/pi/describe-dimension-keys.md)

Retrieves the top N dimension keys for a metric for a specific time period.

[`GetDimensionKeyDetails`](../../../../reference/performance-insights/latest/apireference/api-getdimensionkeydetails.md)

[`aws pi\
									get-dimension-key-details`](../../../cli/latest/reference/pi/get-dimension-key-details.md)

Retrieves the attributes of the specified dimension group for a DB instance or data source. For
example, if you specify a SQL ID, and if the dimension details are available,
`GetDimensionKeyDetails` retrieves the full text of the dimension
`db.sql.statement` associated with this ID. This operation is useful because
`GetResourceMetrics` and `DescribeDimensionKeys` don't support
retrieval of large SQL statement text.

[`GetPerformanceAnalysisReport`](../../../../reference/performance-insights/latest/apireference/api-getperformanceanalysisreport.md)

[`aws pi\
							get-performance-analysis-report`](../../../cli/latest/reference/pi/getperformanceanalysisreport.md)

Retrieves the report including the insights for the report. The result includes the
report status, report ID, report time details, insights, and
recommendations.

`GetResourceMetadata`

[`aws pi\
									get-resource-metadata`](../../../cli/latest/reference/pi/get-resource-metadata.md)

Retrieve the metadata for different features. For example, the metadata might indicate that a
feature is turned on or off on a specific DB instance.

[`GetResourceMetrics`](../../../../reference/performance-insights/latest/apireference/api-getresourcemetrics.md)

[`aws pi\
								get-resource-metrics`](../../../cli/latest/reference/pi/get-resource-metrics.md)

Retrieves Performance Insights metrics for a set of data sources over a time period. You can
provide specific dimension groups and dimensions, and provide aggregation and filtering criteria
for each group.

`ListAvailableResourceDimensions`

[`aws pi\
									list-available-resource-dimensions`](../../../cli/latest/reference/pi/list-available-resource-dimensions.md)

Retrieve the dimensions that can be queried for each specified metric type on a specified
instance.

`ListAvailableResourceMetrics`

[`aws pi\
									list-available-resource-metrics`](../../../cli/latest/reference/pi/list-available-resource-metrics.md)

Retrieve all available metrics of the specified metric types that can be queried for a
specified DB instance.

`ListPerformanceAnalysisReports`

[`aws pi\
							list-performance-analysis-reports`](../../../cli/latest/reference/pi/list-performance-analysis-reports.md)

Retrieves all the analysis reports available for the DB instance. The reports are listed
based on the start time of each report.

`ListTagsForResource`

[`aws pi\
							list-tags-for-resource`](../../../cli/latest/reference/pi/list-tags-for-resource.md)

Lists all the metadata tags added to the resource. The list includes the name and value
of the tag.

`TagResource`

[`aws pi\
							tag-resource`](../../../cli/latest/reference/pi/tag-resource.md)

Adds metadata tags to the Amazon RDS resource. The tag includes a name and a value.

`UntagResource`

[`aws pi\
							untag-resource`](../../../cli/latest/reference/pi/untag-resource.md)

Removes the metadata tag from the resource.

For more information about retrieving time-series metrics and AWS CLI examples for Performance Insights, see the following topics.

###### Topics

- [Retrieving time-series metrics for Performance Insights](user-perfinsights-api-timeseries.md)

- [AWS CLI examples for Performance Insights](user-perfinsights-api-examples.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Viewing Performance Insights proactive recommendations

Retrieving time-series metrics

All content copied from https://docs.aws.amazon.com/.
