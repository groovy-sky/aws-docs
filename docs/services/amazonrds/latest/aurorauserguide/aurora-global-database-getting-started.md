# Getting started with Amazon Aurora Global Database

To get started with Aurora Global Database, first decide which Aurora DB engine you want to
use and in which AWS Regions. Only specific versions of the Aurora MySQL and Aurora PostgreSQL
database engines in certain AWS Regions support Aurora Global Database. For the complete
list, see [Supported Regions and DB engines for Aurora global databases](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.GlobalDatabase.html).

You can create an Aurora Global Database in one of the following ways:

- **Create a new Aurora Global Database with new Aurora DB clusters and Aurora DB instances**
– You can do this by following the steps in [Creating an Amazon Aurora global database](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-creating.html). After
you create the primary Aurora DB cluster, you then add the secondary AWS Region by following the
steps in [Adding an AWS Region to an Amazon Aurora global database](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-attaching.html).

- **Use an existing Aurora DB cluster that supports the Aurora Global**
**Database feature and add an AWS Region to it** – You can do this only
if your existing Aurora DB cluster uses a DB engine version that
is global-compatible.

Check whether you can choose **Add region** for
**Action** on the AWS Management Console when your Aurora DB cluster is selected. If
you can, you can use that Aurora DB cluster for your Aurora global cluster. For more
information, see [Adding an AWS Region to an Amazon Aurora global database](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-attaching.html).

Before creating an Aurora Global Database, we recommend that you understand all configuration requirements.

###### Topics

- [Configuration requirements of an Amazon Aurora global database](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.configuration.requirements.html)

- [Creating an Amazon Aurora global database](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-creating.html)

- [Adding an AWS Region to an Amazon Aurora global database](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-attaching.html)

- [Creating a headless Aurora DB cluster in a secondary Region](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database-attach.console.headless.html)

- [Creating an Amazon Aurora global database from an Aurora or Amazon RDS snapshot](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.use-snapshot.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Using Aurora Global Database

Configuration requirements
