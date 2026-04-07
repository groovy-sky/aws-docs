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

- [Blue/Green Deployments](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.BlueGreenDeployments.html)

- [Aurora cluster configurations](concepts-aurora-fea-regions-db-eng-feature-storage-type.md)

- [Database activity streams](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.DBActivityStreams.html)

- [Exporting cluster data to Amazon S3](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.ExportClusterToS3.html)

- [Exporting snapshot data to Amazon S3](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.ExportSnapshotToS3.html)

- [Aurora global databases](concepts-aurora-fea-regions-db-eng-feature-globaldatabase.md)

- [IAM database authentication](concepts-aurora-fea-regions-db-eng-feature-iamdbauth.md)

- [Kerberos authentication](concepts-aurora-fea-regions-db-eng-feature-kerberosauthentication.md)

- [Aurora machine learning](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.Aurora_ML.html)

- [Performance Insights](concepts-aurora-fea-regions-db-eng-feature-perfinsights.md)

- [Zero-ETL integrations](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.Zero-ETL.html)

- [RDS Proxy](concepts-aurora-fea-regions-db-eng-feature-rds-proxy.md)

- [Secrets Manager integration](concepts-aurora-fea-regions-db-eng-feature-secretsmanager.md)

- [Aurora Serverless v2](concepts-aurora-fea-regions-db-eng-feature-serverlessv2.md)

- [Aurora Serverless v1](concepts-aurora-fea-regions-db-eng-feature-serverlessv1.md)

- [RDS Data API](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.Data_API.html)

- [Zero-downtime patching (ZDP)](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.ZDP.html)

- [Aurora PostgreSQL Limitless Database](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.AuroraLimitless.html)

- [Engine-native features](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.EngineNativeFeatures.html)

## Table conventions

The tables in the feature sections use the following patterns to specify version
numbers and level of support:

- **Version x.y** – The specific version
alone is supported.

- **Version x.y and higher** – The specified
version and all higher minor versions of its major version are supported. For
example, "version 10.11 and higher" means that versions 10.11,
10.11.1, and 10.12 are supported.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Regions and Availability Zones

Blue/Green Deployments
