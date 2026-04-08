# Supported features in Amazon Aurora by AWS Region and Aurora DB engine

Aurora MySQL- and PostgreSQL-compatible database engines support several Amazon Aurora and
Amazon RDS features and options. The support varies across specific versions of each database
engine, and across AWS Regions. To identify Aurora database engine version support and
availability for a feature in a given AWS Region, you can use the following
sections.

Some of these features are Aurora-only capabilities. For example, Aurora
Serverless, Aurora global databases, and support for integration with AWS
machine learning services aren't supported by Amazon RDS. Other features, such as
Amazon RDS Proxy, are supported by both Amazon Aurora and Amazon RDS.

###### Supported Regions and DB engines

- [Table conventions](#Concepts.Aurora_Fea_Regions_DB-eng.Feature.TableConventions)

- [Blue/Green Deployments](concepts-aurora-fea-regions-db-eng-feature-bluegreendeployments.md)

- [Aurora cluster configurations](concepts-aurora-fea-regions-db-eng-feature-storage-type.md)

- [Database activity streams](concepts-aurora-fea-regions-db-eng-feature-dbactivitystreams.md)

- [Exporting cluster data to Amazon S3](concepts-aurora-fea-regions-db-eng-feature-exportclustertos3.md)

- [Exporting snapshot data to Amazon S3](concepts-aurora-fea-regions-db-eng-feature-exportsnapshottos3.md)

- [Aurora global databases](concepts-aurora-fea-regions-db-eng-feature-globaldatabase.md)

- [IAM database authentication](concepts-aurora-fea-regions-db-eng-feature-iamdbauth.md)

- [Kerberos authentication](concepts-aurora-fea-regions-db-eng-feature-kerberosauthentication.md)

- [Aurora machine learning](concepts-aurora-fea-regions-db-eng-feature-aurora-ml.md)

- [Performance Insights](concepts-aurora-fea-regions-db-eng-feature-perfinsights.md)

- [Zero-ETL integrations](concepts-aurora-fea-regions-db-eng-feature-zero-etl.md)

- [RDS Proxy](concepts-aurora-fea-regions-db-eng-feature-rds-proxy.md)

- [Secrets Manager integration](concepts-aurora-fea-regions-db-eng-feature-secretsmanager.md)

- [Aurora Serverless v2](concepts-aurora-fea-regions-db-eng-feature-serverlessv2.md)

- [Aurora Serverless v1](concepts-aurora-fea-regions-db-eng-feature-serverlessv1.md)

- [RDS Data API](concepts-aurora-fea-regions-db-eng-feature-data-api.md)

- [Zero-downtime patching (ZDP)](concepts-aurora-fea-regions-db-eng-feature-zdp.md)

- [Aurora PostgreSQL Limitless Database](concepts-aurora-fea-regions-db-eng-feature-auroralimitless.md)

- [Engine-native features](concepts-aurora-fea-regions-db-eng-feature-enginenativefeatures.md)

## Table conventions

The tables in the feature sections use the following patterns to specify version
numbers and level of support:

- **Version x.y** – The specific version
alone is supported.

- **Version x.y and higher** – The specified
version and all higher minor versions of its major version are supported. For
example, "version 10.11 and higher" means that versions 10.11,
10.11.1, and 10.12 are supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Regions and Availability Zones

Blue/Green Deployments

All content copied from https://docs.aws.amazon.com/.
