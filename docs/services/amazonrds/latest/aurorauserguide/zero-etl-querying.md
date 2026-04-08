# Adding data to a source Aurora DB cluster and querying it

To finish creating a zero-ETL integration that replicates data from Amazon Aurora into
Amazon Redshift, you must create a database in the target destination.

For connections with Amazon Redshift, connect to your Amazon Redshift cluster or workgroup and create a database with a
reference to your integration identifier. Then, you can add data to your source Aurora
DB cluster and see it replicated in Amazon Redshift or Amazon SageMaker.

###### Topics

- [Creating a target database](#zero-etl.create-db)

- [Adding data to the source DB cluster](#zero-etl.add-data-rds)

- [Querying your Aurora data in Amazon Redshift](#zero-etl.query-data-redshift)

- [Data type differences between Aurora and Amazon Redshift databases](#zero-etl.data-type-mapping)

- [DDL operations for Aurora PostgreSQL](#zero-etl.ddl-postgres)

## Creating a target database

Before you can start replicating data into Amazon Redshift, after you create an integration,
you must create a database in your target data warehouse. This
database must include a reference to the integration identifier. You can use the Amazon Redshift
console or the Query editor v2 to create the database.

For instructions to create a destination database, see [Create a destination database in Amazon Redshift](../../../redshift/latest/mgmt/zero-etl-using-creating-db.md#zero-etl-using.create-db).

## Adding data to the source DB cluster

After you configure your integration, you can populate the source Aurora DB cluster
with data that you want to replicate into your data warehouse.

###### Note

There are differences between data types in Amazon Aurora and the target analytics warehouse. For a
table of data type mappings, see [Data type differences between Aurora and Amazon Redshift databases](#zero-etl.data-type-mapping).

First, connect to the source DB cluster using the MySQL or
PostgreSQL client of your choice. For instructions, see [Connecting to an Amazon Aurora DB cluster](aurora-connecting.md).

Then, create a table and insert a row of sample data.

###### Important

Make sure that the table has a primary key. Otherwise, it can't be replicated to
the target data warehouse.

The pg\_dump and pg\_restore PostgreSQL utilities initially create tables without a primary key and then add it afterwards. If you're using one of these utilities, we recommend first creating a schema and then loading data in a separate command.

**MySQL**

The following example uses the [MySQL Workbench utility](https://dev.mysql.com/downloads/workbench).

```sql

CREATE DATABASE my_db;

USE my_db;

CREATE TABLE books_table (ID int NOT NULL, Title VARCHAR(50) NOT NULL, Author VARCHAR(50) NOT NULL,
Copyright INT NOT NULL, Genre VARCHAR(50) NOT NULL, PRIMARY KEY (ID));

INSERT INTO books_table VALUES (1, 'The Shining', 'Stephen King', 1977, 'Supernatural fiction');
```

**PostgreSQL**

The following example uses the `psql`
PostgreSQL interactive terminal. When connecting to the cluster, include the named
database that you specified when creating the integration.

```sql

psql -h mycluster.cluster-123456789012.us-east-2.rds.amazonaws.com -p 5432 -U username -d named_db;

named_db=> CREATE TABLE books_table (ID int NOT NULL, Title VARCHAR(50) NOT NULL, Author VARCHAR(50) NOT NULL,
Copyright INT NOT NULL, Genre VARCHAR(50) NOT NULL, PRIMARY KEY (ID));

named_db=> INSERT INTO books_table VALUES (1, 'The Shining', 'Stephen King', 1977, 'Supernatural fiction');
```

## Querying your Aurora data in Amazon Redshift

After you add data to the Aurora DB cluster, it's replicated into the destination database and is ready to be
queried.

###### To query the replicated data

1. Navigate to the Amazon Redshift console and choose **Query editor**
**v2** from the left navigation pane.

2. Connect to your cluster or workgroup and choose your destination database
    (which you created from the integration) from the dropdown menu
    ( **destination\_database** in this example). For
    instructions to create a destination database, see [Create a destination database in Amazon Redshift](../../../redshift/latest/mgmt/zero-etl-using-creating-db.md#zero-etl-using.create-db).

3. Use a SELECT statement to query your data. In this example, you can run the
    following command to select all data from the table that you created in the
    source Aurora DB cluster:

```sql

SELECT * from my_db."books_table";
```

![Run a SELECT statement within the query editor. The result is a single row of sample data that was added to the Amazon RDS database.](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/zero-etl-redshift-editor.png)

- `my_db` is the Aurora database schema name. This option is only needed for MySQL databases.

- `books_table` is the Aurora table name.

You can also query the data using the a command line client. For example:

```sql

destination_database=# select * from my_db."books_table";

 ID |       Title |        Author |   Copyright |                  Genre |  txn_seq |  txn_id
----+–------------+---------------+-------------+------------------------+----------+--------+
  1 | The Shining |  Stephen King |        1977 |   Supernatural fiction |        2 |   12192
```

###### Note

For case-sensitivity, use double quotes (" ") for schema, table, and column
names. For more information, see [enable\_case\_sensitive\_identifier](../../../redshift/latest/dg/r-enable-case-sensitive-identifier.md).

## Data type differences between Aurora and Amazon Redshift databases

The following tables show the mappings of
Aurora MySQL and Aurora PostgreSQL data types to corresponding destination data types.
_Amazon Aurora currently supports only these data types for_
_zero-ETL integrations._

If a table in your source DB cluster includes an unsupported data type, the table goes
out of sync and isn't consumable by the destination target. Streaming from the source to the
target continues, but the table with the unsupported data type isn't available. To fix
the table and make it available in the target destination, you must manually revert the breaking change
and then refresh the integration by running `ALTER DATABASE...INTEGRATION
                    REFRESH`.

###### Note

You can't refresh zero-ETL integrations with an Amazon SageMaker lakehouse. Instead, delete and
try to create the integration again.

###### Topics

- [Aurora MySQL](#zero-etl.data-type-mapping-mysql)

- [Aurora PostgreSQL](#zero-etl.data-type-mapping-postgres)

### Aurora MySQL

Aurora MySQL data typeTarget data type Description LimitationsINT INTEGERSigned four-byte integerNoneSMALLINTSMALLINTSigned two-byte integerNoneTINYINTSMALLINTSigned two-byte integerNoneMEDIUMINTINTEGERSigned four-byte integerNoneBIGINTBIGINTSigned eight-byte integerNoneINT UNSIGNEDBIGINTSigned eight-byte integerNoneTINYINT UNSIGNEDSMALLINTSigned two-byte integerNoneMEDIUMINT UNSIGNEDINTEGERSigned four-byte integerNoneBIGINT UNSIGNEDDECIMAL(20,0)Exact numeric of selectable precisionNoneDECIMAL(p,s) = NUMERIC(p,s)DECIMAL(p,s)Exact numeric of selectable precision

Precision greater than 38 and scale greater than 37 not
supported

DECIMAL(p,s) UNSIGNED = NUMERIC(p,s)
UNSIGNEDDECIMAL(p,s)Exact numeric of selectable precision

Precision greater than 38 and scale greater than 37 not
supported

FLOAT4/REALREALSingle precision floating-point numberNoneFLOAT4/REAL UNSIGNEDREALSingle precision floating-point numberNoneDOUBLE/REAL/FLOAT8DOUBLE PRECISIONDouble precision floating-point numberNoneDOUBLE/REAL/FLOAT8 UNSIGNEDDOUBLE PRECISIONDouble precision floating-point numberNoneBIT(n)VARBYTE(8)Variable-length binary valueNoneBINARY(n)VARBYTE(n)Variable-length binary valueNoneVARBINARY(n)VARBYTE(n)Variable-length binary valueNoneCHAR(n)VARCHAR(n)Variable-length string valueNoneVARCHAR(n)VARCHAR(n)Variable-length string valueNoneTEXTVARCHAR(65535)Variable-length string value up to 65,535
charactersNoneTINYTEXTVARCHAR(255)Variable-length string value up to 255
charactersNoneMEDIUMTEXTVARCHAR(65535)Variable-length string value up to 65,535
charactersNoneLONGTEXTVARCHAR(65535)Variable-length string value up to 65,535
charactersNoneENUMVARCHAR(1020)Variable-length string value up to 1,020
charactersNoneSETVARCHAR(1020)Variable-length string value up to 1,020
charactersNoneDATEDATECalendar date (year, month, day)NoneDATETIMETIMESTAMPDate and time (without time zone)NoneTIMESTAMP(p)TIMESTAMPDate and time (without time zone)NoneTIMEVARCHAR(18)Variable-length string value up to 18
charactersNoneYEARVARCHAR(4)Variable-length string value up to 4
charactersNoneJSONSUPERSemistructured data or documents as valuesNone

### Aurora PostgreSQL

Zero-ETL integrations for Aurora PostgreSQL don't support custom data types or data
types created by extensions.

Aurora PostgreSQL data typeAmazon Redshift data typeDescriptionLimitationsarraySUPERSemistructured data or documents as valuesNonebigintBIGINTSigned eight-byte integerNonebigserialBIGINTSigned eight-byte integerNonebit varying(n)VARBYTE(n)Variable-length binary value up to 16,777,216
bytesNonebit(n)VARBYTE(n)Variable-length binary value up to 16,777,216
bytesNonebit, bit varyingVARBYTE(16777216)Variable-length binary value up to 16,777,216
bytesNonebooleanBOOLEANLogical boolean (true/false)NonebyteaVARBYTE(16777216)Variable-length binary value up to 16,777,216
bytesNonechar(n)CHAR(n)Fixed-length character string value up to 65,535
bytesNonechar varying(n)VARCHAR(65535)Variable-length character string value up to
65,535 charactersNonecidBIGINT

Signed eight-byte integer

Nonecidr

VARCHAR(19)

Variable-length string value up to 19 characters

NonedateDATECalendar date (year, month, day)

Values greater than 294,276 A.D. not supported

double precisionDOUBLE PRECISIONDouble precision floating-point numbersSubnormal values not fully supported

gtsvector

VARCHAR(65535)

Variable-length string value up to 65,535 characters

Noneinet

VARCHAR(19)

Variable-length string value up to 19 characters

NoneintegerINTEGERSigned four-byte integerNone

int2vector

SUPERSemistructured data or documents as
values.NoneintervalINTERVALDuration of timeOnly INTERVAL types that specify either a year to month or a day
to second qualifier are supported.jsonSUPERSemistructured data or documents as valuesNonejsonbSUPERSemistructured data or documents as valuesNonejsonpathVARCHAR(65535)Variable-length string value up to 65,535
charactersNone

macaddr

VARCHAR(17)Variable-length string value up to 17
charactersNone

macaddr8

VARCHAR(23)Variable-length string value up to 23
charactersNonemoneyDECIMAL(20,3)Currency amountNonenameVARCHAR(64)Variable-length string value up to 64
charactersNonenumeric(p,s)DECIMAL(p,s)User-defined fixed precision value

- `NaN` values not supported

- Precision and scale must be explicitly defined and not
greater than 38 (precision) and 37 (scale)

- Negative scale not supported

oidBIGINTSigned eight-byte integerNoneoidvectorSUPERSemistructured data or documents as
values.Nonepg\_brin\_bloom\_summaryVARCHAR(65535)Variable-length string value up to 65,535
charactersNonepg\_dependenciesVARCHAR(65535)Variable-length string value up to 65,535
charactersNonepg\_lsnVARCHAR(17)Variable-length string value up to 17
charactersNonepg\_mcv\_listVARCHAR(65535)Variable-length string value up to 65,535
charactersNonepg\_ndistinctVARCHAR(65535)Variable-length string value up to 65,535
charactersNonepg\_node\_treeVARCHAR(65535)Variable-length string value up to 65,535
charactersNonepg\_snapshotVARCHAR(65535)Variable-length string value up to 65,535
charactersNonerealREALSingle precision floating-point numberSubnormal values not fully supportedrefcursorVARCHAR(65535)Variable-length string value up to 65,535
charactersNonesmallintSMALLINTSigned two-byte integerNonesmallserialSMALLINTSigned two-byte integerNoneserialINTEGERSigned four-byte integerNonetextVARCHAR(65535)Variable-length string value up to 65,535
charactersNonetidVARCHAR(23)Variable-length string value up to 23
charactersNonetime \[(p)\] without time zoneVARCHAR(19)Variable-length string value up to 19
characters`Infinity` and `-Infinity` values not
supportedtime \[(p)\] with time zoneVARCHAR(22)Variable-length string value up to 22
characters`Infinity` and `-Infinity` values not
supportedtimestamp \[(p)\] without time zoneTIMESTAMPDate and time (without time zone)

- `Infinity` and `-Infinity`
values not supported

- Values greater than `9999-12-31` not
supported

- B.C. values not supported

timestamp \[(p)\] with time zoneTIMESTAMPTZDate and time (with time zone)

- `Infinity` and `-Infinity`
values not supported

- Values greater than `9999-12-31` not
supported

- B.C. values not supported

tsqueryVARCHAR(65535)Variable-length string value up to 65,535
charactersNonetsvectorVARCHAR(65535)Variable-length string value up to 65,535
charactersNonetxid\_snapshotVARCHAR(65535)Variable-length string value up to 65,535
charactersNoneuuidVARCHAR(36)Variable-length 36 character stringNonexidBIGINTSigned eight-byte integerNonexid8DECIMAL(20, 0)Fixed precision decimalNonexmlVARCHAR(65535)Variable-length string value up to 65,535
charactersNone

## DDL operations for Aurora PostgreSQL

Amazon Redshift is derived from PostgreSQL, so it shares several features with Aurora PostgreSQL due to their common PostgreSQL architecture. Zero-ETL integrations
leverage these similarities to streamline data replication from Aurora PostgreSQL to Amazon Redshift, mapping databases by name and utilizing the shared
database, schema, and table structure.

Consider the following points when managing Aurora PostgreSQL
zero-ETL integrations:

- Isolation is managed at the database level.

- Replication occurs at the database level.

- Aurora PostgreSQL databases are mapped to Amazon Redshift databases by name, with
data flowing to the corresponding renamed Redshift database if the original is
renamed.

Despite their similarities, Amazon Redshift and Aurora PostgreSQL have important
differences. The following sections outline Amazon Redshift system responses for common DDL
operations.

###### Topics

- [Database operations](#zero-etl.ddl-postgres-database)

- [Schema operations](#zero-etl.ddl-postgres-schema)

- [Table operations](#zero-etl.ddl-postgres-table)

### Database operations

The following table shows the system responses for database DDL operations.

DDL operationRedshift system response`CREATE DATABASE`No operation`DROP DATABASE`Amazon Redshift drops all the data in the target Redshift
database.`RENAME DATABASE`Amazon Redshift drops all the data in the original target
database and resynchronize the data in the new target database. If
the new database doesn't exist, you must manually create it. For
instructions, see [Create a destination database in Amazon Redshift](../../../redshift/latest/mgmt/zero-etl-using-creating-db.md#zero-etl-using.create-db).

### Schema operations

The following table shows the system responses for schema DDL operations.

DDL operationRedshift system response`CREATE SCHEMA`No operation`DROP SCHEMA`Amazon Redshift drops the original schema.`RENAME SCHEMA`Amazon Redshift drops the original schema then resynchronizes the
data in the new schema.

### Table operations

The following table shows the system responses for table DDL operations.

DDL operationRedshift system response`CREATE TABLE`

Amazon Redshift creates the table.

Some operations cause table creation to fail, such as creating
a table without a primary key or performing declarative
partitioning. For more information, see [Limitations](zero-etl.md#zero-etl.reqs-lims) and [Troubleshooting Aurora zero-ETL integrations](zero-etl-troubleshooting.md).

`DROP TABLE`Amazon Redshift drops the table.`TRUNCATE TABLE`Amazon Redshift truncates the table.`ALTER TABLE`
( `RENAME...`)Amazon Redshift renames the table or column.`ALTER TABLE` ( `SET
                                SCHEMA`)

Amazon Redshift drops the table in the original schema and resynchronizes
the table in the new schema.

`ALTER TABLE` ( `ADD PRIMARY
                                    KEY`)Amazon Redshift adds a primary key and resynchronizes the
table.`ALTER TABLE` ( `ADD
                                COLUMN`)Amazon Redshift adds a column to the table.`ALTER TABLE` ( `DROP
                                COLUMN`)

Amazon Redshift drops the column if it's not a primary key column.
Otherwise, it resynchronizes the table.

`ALTER TABLE` ( `SET
                                    LOGGED/UNLOGGED`)If you change the table to logged, Amazon Redshift
resynchronizes the table. If you change the table to unlogged, Amazon Redshift
drops the table.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data filtering for zero-ETL integrations

Viewing and monitoring
zero-ETL integrations

All content copied from https://docs.aws.amazon.com/.
