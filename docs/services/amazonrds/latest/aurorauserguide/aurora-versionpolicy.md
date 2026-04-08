# Amazon Aurora versions

With Amazon Aurora, you can choose the [supported relational database engine](#Aurora.VersionPolicy.SupportedEngines) that best fits your application requirements
while maintaining compatibility with the underlying engines. Aurora reuses the database engine
code from supported engines, so you can leverage existing skills, tools, and libraries for those
engines. When you create a cluster, you [specify the Amazon Aurora database engine version](#Aurora.VersionPolicy.ChoosingVersion) that you want to use. The version that
you choose determines compatibility and available features.

This documentation explains the common points and differences between Amazon Aurora and the
corresponding database engines. With this information, you can determine the software version
that you should select and how to verify the available features and bug fixes in each version.
You can also use this reference to determine an appropriate upgrade cadence and plan for your
upgrade.

## Supported database engines for Amazon Aurora database clusters

The following relational databases are available on Amazon Aurora. Aurora reuses code and
maintains compatibility with the underlying DB engines. However, Aurora has unique version
numbers, release cycles, and timelines for version deprecation. Each new Aurora version comes
with release notes that list the new features, fixes, and other changes and enhancements that
apply to each version.

Aurora databaseUser guideAvailable versionsRelease notesAmazon Aurora MySQL-Compatible Edition[Working with Amazon Aurora MySQL](aurora-auroramysql.md)[Database engine updates for Amazon Aurora MySQL](auroramysql-updates.md)[Release Notes for Aurora MySQL](../auroramysqlreleasenotes/welcome.md)Amazon Aurora PostgreSQL-Compatible Edition[Working with Amazon Aurora PostgreSQL](aurora-aurorapostgresql.md)[Database engine updates for Amazon Aurora PostgreSQL](aurorapostgresql-updates.md)[Release Notes for Aurora PostgreSQL](../aurorapostgresqlreleasenotes/welcome.md)

## Specifying the Amazon Aurora database version for your database cluster

When you create a new DB cluster with the **Create database** operation
in the AWS Management Console or the AWS CLI, or with the `CreateDBCluster` API operation, you can
specify any currently available Aurora database version (major or minor).

To learn how to create Aurora clusters, see [Creating an Amazon Aurora DB cluster](aurora-createinstance.md). To learn how to change the version of an existing
Aurora cluster, see [Modifying an Amazon Aurora DB cluster](aurora-modifying.md).

###### Note

Not every Aurora database version is available in every AWS Region. To learn more about
Regions and the available versions in each AWS Region, see [Regions and Availability Zones](concepts-regionsandavailabilityzones.md) and [Supported Regions and DB engines for Aurora global databases](concepts-aurora-fea-regions-db-eng-feature-globaldatabase.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora DB clusters

Aurora versioning

All content copied from https://docs.aws.amazon.com/.
