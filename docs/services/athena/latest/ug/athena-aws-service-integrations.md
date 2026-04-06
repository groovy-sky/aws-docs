# AWS service integrations with Athena

You can use Athena to query data from the AWS services listed in this section. To see the
Regions that each service supports, see [Regions and\
endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html) in the _Amazon Web Services General Reference_.

###### AWS services integrated with Athena

- [CloudFormation](#integ-ate-cfn)

- [Amazon CloudFront](#integ-ate-cf)

- [AWS CloudTrail](#integ-ate-ct)

- [Amazon\
DataZone](#integ-ate-dz)

- [Elastic Load Balancing](#integ-ate-eb)

- [Amazon EMR\
Studio](#integ-ate-emr-studio)

- [AWS Glue Data Catalog](#integ-ate-gc)

- [AWS Identity and Access Management\
(IAM)](#integ-ate-iam)

- [Amazon Quick](#integ-ate-qs)

- [Amazon S3\
Inventory](#integ-ate-s3)

- [AWS Step Functions](#integ-ate-sf)

- [AWS Systems Manager\
Inventory](#integ-ate-sys)

- [Amazon Virtual Private Cloud](#integ-ate-vpc)

For information about each integration, see the following sections.

**CloudFormation**

_Capacity reservation_

Reference topic: [AWS::Athena::CapacityReservation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-capacityreservation.html) in the
_AWS CloudFormation User Guide_

Specifies a capacity reservation with the provided name and number
of requested data processing units. For more information, see [Manage query processing capacity](https://docs.aws.amazon.com/athena/latest/ug/capacity-management.html) in the
_Amazon Athena User Guide_ and [CreateCapacityReservation](https://docs.aws.amazon.com/athena/latest/APIReference/API_CreateCapacityReservation.html) in the _Amazon Athena API Reference_.

_Data catalog_

Reference topic: [AWS::Athena::DataCatalog](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html) in the
_AWS CloudFormation User Guide_

Specify an Athena data catalog, including a name, description,
type, parameters, and tags. For more information, see [Understanding tables, databases, and data catalogs in Athena](https://docs.aws.amazon.com/athena/latest/ug/understanding-tables-databases-and-the-data-catalog.html) in the _Amazon Athena User Guide_ and [CreateDataCatalog](https://docs.aws.amazon.com/athena/latest/APIReference/API_CreateDataCatalog.html) in the _Amazon_
_Athena API Reference_.

_Named query_

Reference topic: [AWS::Athena::NamedQuery](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-namedquery.html) in the
_AWS CloudFormation User Guide_

Specify named queries with CloudFormation and run them in Athena. Named
queries allow you to map a query name to a query and then run it as
a saved query from the Athena console. For more information, see
[Use saved queries](https://docs.aws.amazon.com/athena/latest/ug/saved-queries.html) in
the _Amazon Athena User Guide_ and [CreateNamedQuery](https://docs.aws.amazon.com/athena/latest/APIReference/API_CreateNamedQuery.html) in the
_Amazon Athena API Reference_.

_Prepared statement_

Reference topic: [AWS::Athena::PreparedStatement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-preparedstatement.html) in the
_AWS CloudFormation User Guide_

Specifies a prepared statement for use with SQL queries in Athena.
A prepared statement contains parameter placeholders whose values
are supplied at execution time. For more information, see [Use parameterized queries](https://docs.aws.amazon.com/athena/latest/ug/querying-with-prepared-statements.html) in the
_Amazon Athena User Guide_ and [CreatePreparedStatement](https://docs.aws.amazon.com/athena/latest/APIReference/API_CreatePreparedStatement.html) in the
_Amazon Athena API Reference_.

_Workgroup_

Reference topic: [AWS::Athena::WorkGroup](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-workgroup.html) in the _CloudFormation User_
_Guide_

Specify Athena workgroups using AWS CloudFormation. Use Athena workgroups to
isolate queries for you or your group from other queries in the same
account. For more information, see [Use workgroups to control query access and costs](https://docs.aws.amazon.com/athena/latest/ug/workgroups-manage-queries-control-costs.html)
in the _Amazon Athena User Guide_ and [CreateWorkGroup](https://docs.aws.amazon.com/athena/latest/APIReference/API_CreateWorkGroup.html) in the _Amazon Athena API_
_Reference_.

**Amazon CloudFront**

Reference topic: [Query Amazon CloudFront logs](https://docs.aws.amazon.com/athena/latest/ug/cloudfront-logs.html)

Use Athena to query Amazon CloudFront logs. For more information about using CloudFront, see
the [Amazon CloudFront Developer Guide](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide).

**AWS CloudTrail**

Reference topic: [Query AWS CloudTrail logs](https://docs.aws.amazon.com/athena/latest/ug/cloudtrail-logs.html)

Using Athena with CloudTrail logs is a powerful way to enhance your analysis of AWS
service activity. For example, you can use queries to identify trends and
further isolate activity by attribute, such as source IP address or user. You
can create tables for querying logs directly from the CloudTrail console, and use
those tables to run queries in Athena. For more information, see [Use the CloudTrail console to create an Athena table for CloudTrail logs](https://docs.aws.amazon.com/athena/latest/ug/create-cloudtrail-table-ct.html).

**Amazon DataZone**

Reference topic: [Use Amazon DataZone in Athena](https://docs.aws.amazon.com/athena/latest/ug/datazone-using.html)

Use [Amazon DataZone](https://aws.amazon.com/datazone) to share, search,
and discover data at scale across organizational boundaries. DataZone
simplifies your experience across AWS analytics services like Athena, AWS Glue,
and AWS Lake Formation. If you have large amounts of data in different data sources, you
can use Amazon DataZone to build business use case based groupings of people, data and
tools.

In Athena, you can use the query editor to access and query DataZone
environments. For more information, see [Use Amazon DataZone in Athena](https://docs.aws.amazon.com/athena/latest/ug/datazone-using.html).

**Elastic Load Balancing**

Reference topic: [Query Application Load Balancer logs](https://docs.aws.amazon.com/athena/latest/ug/application-load-balancer-logs.html)

Querying Application Load Balancer logs allows you to see the source of
traffic, latency, and bytes transferred to and from Elastic Load Balancing instances and
backend applications. For more information, see [Query Application Load Balancer logs](https://docs.aws.amazon.com/athena/latest/ug/application-load-balancer-logs.html).

Reference topic: [Query Classic Load Balancer logs](https://docs.aws.amazon.com/athena/latest/ug/elasticloadbalancer-classic-logs.html)

Query Classic Load Balancer logs to analyze and understand traffic patterns to
and from Elastic Load Balancing instances and backend applications. You can see the source of
traffic, latency, and bytes transferred. For more information, see [Query Classic Load Balancer logs](https://docs.aws.amazon.com/athena/latest/ug/elasticloadbalancer-classic-logs.html).

**Amazon EMR Studio**

Reference topic: [Use the Amazon Athena\
SQL editor in EMR Studio](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-athena.html)

You can use Athena in an EMR Studio to develop and run interactive queries.
This makes it possible for you to use EMR Studio for SQL analytics on Athena
from the same Amazon EMR interface that you use for your Spark, Scala, and other
workloads. With the Athena integration in EMR Studio, you can perform the
following tasks:

- Perform Athena SQL queries

- View query results

- View query history

- View saved queries

- Perform parameterized queries

- View databases, tables, and views for a data catalog

The following Athena features are not available in Amazon EMR Studio:

- Admin features like creating or updating Athena workgroups, data
sources, or capacity reservations

- Athena for Spark or Spark notebooks

- DataZone integration

- Step Functions

EMR Studio integration with Athena is available in all AWS Regions where
EMR Studio and Athena are available. For more information about using Athena
in EMR Studio, see [Use the Amazon Athena\
SQL editor in EMR Studio](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-studio-athena.html) in the
_Amazon EMR Management Guide_.

**AWS Glue Data Catalog**

Reference topic: [Use AWS Glue Data Catalog to connect to your data](https://docs.aws.amazon.com/athena/latest/ug/data-sources-glue.html)

Athena integrates with the AWS Glue Data Catalog, which offers a persistent metadata
store for your data in Amazon S3. This allows you to create tables and query data in
Athena based on a central metadata store available throughout your Amazon Web Services
account and integrated with the ETL and data discovery features of AWS Glue.
For more information, see [Use AWS Glue Data Catalog to connect to your data](https://docs.aws.amazon.com/athena/latest/ug/data-sources-glue.html) and [What is AWS Glue](https://docs.aws.amazon.com/glue/latest/dg/what-is-glue.html) in the
_AWS Glue Developer Guide_.

**AWS Identity and Access Management (IAM)**

Reference topic: [Actions for\
Amazon Athena](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonathena.html)

You can use Athena API actions in IAM permission policies. For more
information, see [Actions for\
Amazon Athena](https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonathena.html) and [Identity and access management in Athena](https://docs.aws.amazon.com/athena/latest/ug/security-iam-athena.html).

**Amazon Quick**

Reference topic: [Connect to Amazon Athena with ODBC and JDBC drivers](https://docs.aws.amazon.com/athena/latest/ug/athena-bi-tools-jdbc-odbc.html)

Athena integrates with Amazon Quick for easy data visualization. You can use Athena
to generate reports or to explore data with business intelligence tools or SQL
clients connected with a JDBC or an ODBC driver. For more information about
Quick, see [What\
is Amazon Quick](https://docs.aws.amazon.com/quicksight/latest/user/welcome.html) in the _Amazon Quick User Guide_. For
information about using JDBC and ODBC drivers with Athena, see [Connecting to Amazon Athena with ODBC and\
JDBC Drivers](https://docs.aws.amazon.com/athena/latest/ug/athena-bi-tools-jdbc-odbc.html).

**Amazon S3 Inventory**

Reference topic: [Querying inventory with Athena](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html#storage-inventory-athena-query) in the _Amazon Simple Storage Service User Guide_

You can use Amazon Athena to query Amazon S3 inventory using standard SQL. You can use
Amazon S3 inventory to audit and report on the replication and encryption status of
your objects for business, compliance, and regulatory needs. For more
information, see [Amazon S3 inventory](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) in
the _Amazon Simple Storage Service User Guide_.

**AWS Step Functions**

Reference topic: [Call Athena with\
Step Functions](https://docs.aws.amazon.com/step-functions/latest/dg/connect-athena.html) in the _AWS Step Functions Developer Guide_

Call Athena with AWS Step Functions. AWS Step Functions can control select AWS services
directly using the [Amazon States Language](https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html). You can use Step Functions with Athena to start and stop query
execution, get query results, run ad-hoc or scheduled data queries, and retrieve
results from data lakes in Amazon S3. The Step Functions role must have permissions to use
Athena. For more information, see the [AWS Step Functions Developer Guide](https://docs.aws.amazon.com/step-functions/latest/dg).

###### Video: Orchestrate Amazon Athena Queries using AWS Step Functions

The following video demonstrates how to use Amazon Athena and AWS Step Functions to run
a regularly scheduled Athena query and generate a corresponding
report.

For an example that uses Step Functions and Amazon EventBridge to orchestrate AWS Glue DataBrew, Athena,
and Amazon Quick, see [Orchestrating an AWS Glue DataBrew job and Amazon Athena query with AWS Step Functions](https://aws.amazon.com/blogs/big-data/orchestrating-an-aws-glue-databrew-job-and-amazon-athena-query-with-aws-step-functions)
in the AWS Big Data Blog.

**AWS Systems Manager Inventory**

Reference topic: [Querying inventory data from multiple regions and accounts](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-inventory-query.html) in the
_AWS Systems Manager User Guide_

AWS Systems Manager Inventory integrates with Amazon Athena to help you query inventory data
from multiple AWS Regions and accounts. For more information, see the
[AWS Systems Manager User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide).

**Amazon Virtual Private Cloud**

Reference topic: [Query Amazon VPC flow logs](https://docs.aws.amazon.com/athena/latest/ug/vpc-flow-logs.html)

Amazon Virtual Private Cloud flow logs capture information about the IP traffic going to and from
network interfaces in a VPC. Query the logs in Athena to investigate network
traffic patterns and identify threats and risks across your Amazon VPC network. For
more information about Amazon VPC, see the [Amazon VPC User Guide](https://docs.aws.amazon.com/vpc/latest/userguide).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Set up access

Use Athena SQL
