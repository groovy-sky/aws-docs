# Supported features in Amazon RDS by AWS Region and DB engine

Support for Amazon RDS features and options varies across AWS Regions and specific versions of
each DB engine. To identify RDS DB engine version support and availability in a given
AWS Region, you can use the following sections.

Amazon RDS features are different from engine-native features and options.
For more information on engine-native features and options, see
[Engine-native features.](concepts-rds-fea-regions-db-eng-feature-enginenativefeatures.md)

###### Supported Regions and DB engines

- [Table conventions](#Concepts.RDS_Fea_Regions_DB-eng.Feature.TableConventions)

- [Feature quick reference](#Concepts.RDS_Fea_Regions_DB-eng.Feature.QuickReferenceTable)

- [Blue/Green Deployments](concepts-rds-fea-regions-db-eng-feature-bluegreendeployments.md)

- [Cross-Region automated backups](concepts-rds-fea-regions-db-eng-feature-crossregionautomatedbackups.md)

- [Cross-Region read replicas](concepts-rds-fea-regions-db-eng-feature-crossregionreadreplicas.md)

- [Database\
activity streams](concepts-rds-fea-regions-db-eng-feature-dbactivitystreams.md)

- [Dual-stack\
mode](concepts-rds-fea-regions-db-eng-feature-dualstackmode.md)

- [Export\
snapshots to S3](concepts-rds-fea-regions-db-eng-feature-exportsnapshottos3.md)

- [IAM database authentication](concepts-rds-fea-regions-db-eng-feature-iamdatabaseauthentication.md)

- [Kerberos authentication](concepts-rds-fea-regions-db-eng-feature-kerberosauthentication.md)

- [Multi-AZ DB clusters](concepts-rds-fea-regions-db-eng-feature-multiazdbclusters.md)

- [Performance Insights](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md)

- [RDS Custom](concepts-rds-fea-regions-db-eng-feature-rdscustom.md)

- [Amazon RDS Proxy](concepts-rds-fea-regions-db-eng-feature-rdsproxy.md)

- [Secrets\
Manager integration](concepts-rds-fea-regions-db-eng-feature-secretsmanager.md)

- [Zero-ETL integrations](concepts-rds-fea-regions-db-eng-feature-zeroetl.md)

- [Engine-native features](concepts-rds-fea-regions-db-eng-feature-enginenativefeatures.md)

## Table conventions

The tables in the feature sections use these patterns to specify version numbers and level of availability:

- **Version x.y** – The specific version alone is available.

- **Version x.y and higher** – The specified version and all higher minor versions of its major version are supported. For example, "version 10.11 and higher" means that
versions 10.11, 10.11.1, and 10.12 are available.

## Feature quick reference

The following quick reference table lists each feature and available RDS DB engine. Region
and specific version availability appears in the later feature sections.

FeatureRDS for Db2RDS for MariaDBRDS for MySQLRDS for OracleRDS for PostgreSQLRDS for SQL ServerBlue/Green DeploymentsNot available[Available](concepts-rds-fea-regions-db-eng-feature-bluegreendeployments.md)[Available](concepts-rds-fea-regions-db-eng-feature-bluegreendeployments.md)Not available[Available](concepts-rds-fea-regions-db-eng-feature-bluegreendeployments.md)Not availableCross-Region automated backups[Available](concepts-rds-fea-regions-db-eng-feature-crossregionautomatedbackups.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionAutomatedBackups.db2)[Available](concepts-rds-fea-regions-db-eng-feature-crossregionautomatedbackups.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionAutomatedBackups.mdb)[Available](concepts-rds-fea-regions-db-eng-feature-crossregionautomatedbackups.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionAutomatedBackups.my)[Available](concepts-rds-fea-regions-db-eng-feature-crossregionautomatedbackups.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionAutomatedBackups.ora)[Available](concepts-rds-fea-regions-db-eng-feature-crossregionautomatedbackups.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionAutomatedBackups.pg)[Available](concepts-rds-fea-regions-db-eng-feature-crossregionautomatedbackups.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionAutomatedBackups.sq)Cross-Region read replicasNot available[Available](concepts-rds-fea-regions-db-eng-feature-crossregionreadreplicas.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionReadReplicas.mdb)[Available](concepts-rds-fea-regions-db-eng-feature-crossregionreadreplicas.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionReadReplicas.my)[Available](concepts-rds-fea-regions-db-eng-feature-crossregionreadreplicas.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionReadReplicas.ora)[Available](concepts-rds-fea-regions-db-eng-feature-crossregionreadreplicas.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionReadReplicas.pg)[Available](concepts-rds-fea-regions-db-eng-feature-crossregionreadreplicas.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.CrossRegionReadReplicas.sq)Database activity streamsNot availableNot availableNot available[Available](concepts-rds-fea-regions-db-eng-feature-dbactivitystreams.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.DBActivityStreams.ora)Not available[Available](concepts-rds-fea-regions-db-eng-feature-dbactivitystreams.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.DBActivityStreams.SQLServer)Dual-stack modeNot available[Available](concepts-rds-fea-regions-db-eng-feature-dualstackmode.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.DualStackMode.mdb)[Available](concepts-rds-fea-regions-db-eng-feature-dualstackmode.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.DualStackMode.my)[Available](concepts-rds-fea-regions-db-eng-feature-dualstackmode.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.DualStackMode.ora)[Available](concepts-rds-fea-regions-db-eng-feature-dualstackmode.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.DualStackMode.pg)[Available](concepts-rds-fea-regions-db-eng-feature-dualstackmode.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.DualStackMode.sq)Export Snapshot to Amazon S3Not available[Available](concepts-rds-fea-regions-db-eng-feature-exportsnapshottos3.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.ExportSnapshotToS3.mdb)[Available](concepts-rds-fea-regions-db-eng-feature-exportsnapshottos3.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.ExportSnapshotToS3.my)Not available[Available](concepts-rds-fea-regions-db-eng-feature-exportsnapshottos3.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.ExportSnapshotToS3.pg)Not availableAWS Identity and Access Management (IAM) database authenticationNot available[Available](concepts-rds-fea-regions-db-eng-feature-iamdatabaseauthentication.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.IamDatabaseAuthentication.mdb)[Available](concepts-rds-fea-regions-db-eng-feature-iamdatabaseauthentication.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.IamDatabaseAuthentication.my)Not available[Available](concepts-rds-fea-regions-db-eng-feature-iamdatabaseauthentication.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.IamDatabaseAuthentication.pg)Not availableKerberos authentication[Available](concepts-rds-fea-regions-db-eng-feature-kerberosauthentication.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.my)Not available[Available](concepts-rds-fea-regions-db-eng-feature-kerberosauthentication.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.my)[Available](concepts-rds-fea-regions-db-eng-feature-kerberosauthentication.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.ora)[Available](concepts-rds-fea-regions-db-eng-feature-kerberosauthentication.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.pg)[Available](concepts-rds-fea-regions-db-eng-feature-kerberosauthentication.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.KerberosAuthentication.sq)Multi-AZ DB clustersNot availableNot available[Available](concepts-rds-fea-regions-db-eng-feature-multiazdbclusters.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.MultiAZDBClusters.my)Not available[Available](concepts-rds-fea-regions-db-eng-feature-multiazdbclusters.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.MultiAZDBClusters.pg)Not availablePerformance InsightsNot available[Available](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md)[Available](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md)[Available](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md)[Available](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md)[Available](concepts-rds-fea-regions-db-eng-feature-performanceinsights.md)RDS CustomNot availableNot availableNot available[Available](concepts-rds-fea-regions-db-eng-feature-rdscustom.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.RDSCustom.ora)Not available[Available](concepts-rds-fea-regions-db-eng-feature-rdscustom.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.RDSCustom.sq)RDS ProxyNot available[Available](concepts-rds-fea-regions-db-eng-feature-rdsproxy.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.RDS_Proxy.mdb)[Available](concepts-rds-fea-regions-db-eng-feature-rdsproxy.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.RDS_Proxy.my)Not available[Available](concepts-rds-fea-regions-db-eng-feature-rdsproxy.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.RDS_Proxy.pg)[Available](concepts-rds-fea-regions-db-eng-feature-rdsproxy.md#Concepts.RDS_Fea_Regions_DB-eng.Feature.RDS_Proxy.sq)Secrets Manager integration[Available](concepts-rds-fea-regions-db-eng-feature-secretsmanager.md)[Available](concepts-rds-fea-regions-db-eng-feature-secretsmanager.md)[Available](concepts-rds-fea-regions-db-eng-feature-secretsmanager.md)[Available](concepts-rds-fea-regions-db-eng-feature-secretsmanager.md)[Available](concepts-rds-fea-regions-db-eng-feature-secretsmanager.md)[Available](concepts-rds-fea-regions-db-eng-feature-secretsmanager.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Regions, Availability Zones, and Local Zones

Blue/Green Deployments

All content copied from https://docs.aws.amazon.com/.
