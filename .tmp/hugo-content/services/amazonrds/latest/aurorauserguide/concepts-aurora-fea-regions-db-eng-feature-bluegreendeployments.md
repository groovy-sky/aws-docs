# Supported Regions and Aurora DB engines for Blue/Green Deployments

A blue/green deployment copies a production database environment in a separate,
synchronized staging environment. By using Amazon RDS Blue/Green Deployments, you can make
changes to the database in the staging environment without affecting the production
environment. For example, you can upgrade the major or minor DB engine version, change
database parameters, or make schema changes in the staging environment. When you are
ready, you can promote the staging environment to be the new production database
environment. For more information, see [Using Amazon Aurora Blue/Green Deployments for database updates](blue-green-deployments.md).

## Blue/Green Deployments with Aurora MySQL

The Blue/Green Deployments feature is available for all versions of Aurora MySQL in
all AWS Regions, including Aurora MySQL clusters configured as Aurora Global Database.

## Blue/Green Deployments with Aurora PostgreSQL

The following Regions and engine versions are available for
Blue/Green Deployments with Aurora PostgreSQL, including Aurora PostgreSQL clusters configured as Aurora Global Database.

RegionAurora PostgreSQL 17Aurora PostgreSQL 16Aurora PostgreSQL 15Aurora PostgreSQL 14Aurora PostgreSQL 13Aurora PostgreSQL 12Aurora PostgreSQL 11

All AWS Regions

Version 17.4 and higherVersion 16.1 and higherVersion 15.4 and higherVersion 14.9 and higherVersion 13.12 and higherVersion 12.16 and higherVersion 11.21 and higher

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Supported Aurora
features by Region and engine

Aurora cluster configurations

All content copied from https://docs.aws.amazon.com/.
