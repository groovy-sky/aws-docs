# DB instance classes

The DB instance class determines the computation and memory capacity of an Amazon RDS
DB instance. The DB instance class that you need depends
on your processing power and memory requirements.

A DB instance class consists of both the DB instance class type and the size. For example, db.r6g is a
memory-optimized DB instance class type powered by AWS Graviton2 processors. Within the db.r6g
instance class type, db.r6g.2xlarge is a DB instance class. The size of this class is
2xlarge.

For more information about instance class pricing, see [Amazon RDS\
pricing](https://aws.amazon.com/rds/pricing).

For more information about DB instance class types, supported DB engines, supported AWS Regions, changing your DB instance class, configuring the processor for RDS for Oracle, or hardware specifications for DB instance classes, see the following sections.

###### Topics

- [DB instance class types](concepts-dbinstanceclass-types.md)

- [Supported DB engines for DB instance classes](concepts-dbinstanceclass-support.md)

- [Determining DB instance class support in AWS Regions](concepts-dbinstanceclass-regionsupport.md)

- [Changing your DB instance class](#Concepts.DBInstanceClass.Changing)

- [Configuring the processor for a DB instance class in RDS for Oracle](user-configureprocessor.md)

- [Hardware specifications for DB instance classes](concepts-dbinstanceclass-summary.md)

## Changing your DB instance class

You can change the CPU and memory available to a DB instance by changing its DB instance
class. To change the DB instance class, modify your DB instance by following the
instructions in [Modifying an Amazon RDS DB instance](overview-dbinstance-modifying.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DB instances

DB instance class types

All content copied from https://docs.aws.amazon.com/.
