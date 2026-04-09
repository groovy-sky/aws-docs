# Understand service quotas for Athena for Spark

_Service quotas_, also known as _limits_, are the
maximum number of service resources or operations that your AWS account can use. For more
information about the service quotas for other AWS services that you can use with
Amazon Athena for Spark, see [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the
_Amazon Web Services General Reference_.

###### Note

The default values are the initial quotas set by AWS, which are separate from the actual applied quota value and maximum possible service quota.
New AWS accounts might have initial lower quotas that can increase over time.
Amazon Athena for Apache Spark monitors account usage within each AWS Region, and then
automatically increases the quotas based on your usage. If your requirements exceed the
stated limits, contact customer support.

The following table lists the service quotas for Amazon Athena for Apache Spark.

NameDefaultAdjustableVersionDescriptionApache Spark DPU concurrency160NoPySpark Version 3The maximum number of data processing units (DPUs) that you can consume
concurrently for Apache Spark calculations for a single account in the
current AWS Region. A DPU is a relative measure of processing power that
consists of 4 vCPUs of compute capacity and 16 GB of memory.Apache Spark session DPU concurrency60NoPySpark Version 3The maximum number of DPUs you can consume concurrently for an Apache
Spark calculation within a session.On-Demand DPUs4NoApache Spark Version 3.5The maximum number of data processing units (DPUs) that you can consume concurrently for Apache Spark interactive sessions in the current AWS Region.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Cross-account catalog access

Athena Spark APIs

All content copied from https://docs.aws.amazon.com/.
