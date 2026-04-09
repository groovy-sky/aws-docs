# Amazon Aurora DSQL cluster connectivity tools

Aurora DSQL is compatible with many third-party database drivers and ORM libraries. AWS provides two types of tools to simplify working with Aurora DSQL:

- **[Connectors](../../../../services/aurora-dsql/latest/userguide/section-connectors.md)** – Authentication plugins that extend database drivers to handle IAM token generation automatically. Use connectors when working directly with database drivers.

- **Adapters and dialects** – Extensions for specific ORM frameworks that provide IAM authentication and improved Aurora DSQL compatibility. Use adapters when working with a supported ORM framework.

## Aurora DSQL adapters and dialects

The following table shows the available adapters and dialects for Aurora DSQL.

Programming languageORM/FrameworkRepository linkJavaHibernate[https://github.com/awslabs/aurora-dsql-orms/tree/main/java/hibernate](https://github.com/awslabs/aurora-dsql-orms/tree/main/java/hibernate)PythonDjango[https://github.com/awslabs/aurora-dsql-orms/tree/main/python/django](https://github.com/awslabs/aurora-dsql-orms/tree/main/python/django)PythonSQLAlchemy[https://github.com/awslabs/aurora-dsql-orms/tree/main/python/sqlalchemy](https://github.com/awslabs/aurora-dsql-orms/tree/main/python/sqlalchemy)PythonTortoise ORM[https://github.com/awslabs/aurora-dsql-orms/tree/main/python/tortoise-orm](https://github.com/awslabs/aurora-dsql-orms/tree/main/python/tortoise-orm)

## Database driver samples

The following table shows sample code for connecting to Aurora DSQL using third-party database drivers.

Programming languageDriverSample repository linkC++libpq[https://github.com/aws-samples/aurora-dsql-samples/tree/main/cpp/libpq](https://github.com/aws-samples/aurora-dsql-samples/tree/main/cpp/libpq)C# (.NET)Npgsql[https://github.com/aws-samples/aurora-dsql-samples/tree/main/dotnet/npgsql](https://github.com/aws-samples/aurora-dsql-samples/tree/main/dotnet/npgsql)Gopgx[https://github.com/aws-samples/aurora-dsql-samples/tree/main/go/pgx](https://github.com/aws-samples/aurora-dsql-samples/tree/main/go/pgx)JavaHikariCP + pgJDBC[https://github.com/aws-samples/aurora-dsql-samples/tree/main/java/pgjdbc](https://github.com/aws-samples/aurora-dsql-samples/tree/main/java/pgjdbc)JavaScriptnode-postgres (AWS Lambda)[https://github.com/aws-samples/aurora-dsql-samples/tree/main/lambda](https://github.com/aws-samples/aurora-dsql-samples/tree/main/lambda)JavaScriptnode-postgres[https://github.com/aws-samples/aurora-dsql-samples/tree/main/javascript/node-postgres](https://github.com/aws-samples/aurora-dsql-samples/tree/main/javascript/node-postgres)JavaScriptPostgres.js[https://github.com/aws-samples/aurora-dsql-samples/tree/main/javascript/postgres-js](https://github.com/aws-samples/aurora-dsql-samples/tree/main/javascript/postgres-js)Pythonasyncpg[https://github.com/aws-samples/aurora-dsql-samples/tree/main/python/asyncpg](https://github.com/aws-samples/aurora-dsql-samples/tree/main/python/asyncpg)PythonPsycopg[https://github.com/aws-samples/aurora-dsql-samples/tree/main/python/psycopg](https://github.com/aws-samples/aurora-dsql-samples/tree/main/python/psycopg)PythonPsycopg2[https://github.com/aws-samples/aurora-dsql-samples/tree/main/python/psycopg2](https://github.com/aws-samples/aurora-dsql-samples/tree/main/python/psycopg2)Rubypg[https://github.com/aws-samples/aurora-dsql-samples/tree/main/ruby/ruby-pg](https://github.com/aws-samples/aurora-dsql-samples/tree/main/ruby/ruby-pg)RustSQLx[https://github.com/aws-samples/aurora-dsql-samples/tree/main/rust/sqlx](https://github.com/aws-samples/aurora-dsql-samples/tree/main/rust/sqlx)

## ORM and framework samples

The following table shows sample code for using third-party ORM libraries and frameworks with Aurora DSQL.

Programming languageORM/FrameworkSample repository linkJavaHibernate[https://github.com/awslabs/aurora-dsql-orms/tree/main/java/hibernate/examples/pet-clinic-app](https://github.com/awslabs/aurora-dsql-orms/tree/main/java/hibernate/examples/pet-clinic-app)JavaLiquibase[https://github.com/aws-samples/aurora-dsql-samples/tree/main/java/liquibase](https://github.com/aws-samples/aurora-dsql-samples/tree/main/java/liquibase)JavaSpring Boot[https://github.com/aws-samples/aurora-dsql-samples/tree/main/java/spring\_boot](https://github.com/aws-samples/aurora-dsql-samples/tree/main/java/spring_boot)PythonDjango[https://github.com/awslabs/aurora-dsql-orms/tree/main/python/django/examples/pet-clinic-app](https://github.com/awslabs/aurora-dsql-orms/tree/main/python/django/examples/pet-clinic-app)PythonSQLAlchemy[https://github.com/awslabs/aurora-dsql-orms/tree/main/python/sqlalchemy/examples/pet-clinic-app](https://github.com/awslabs/aurora-dsql-orms/tree/main/python/sqlalchemy/examples/pet-clinic-app)PythonTortoise ORM[https://github.com/awslabs/aurora-dsql-orms/tree/main/python/tortoise-orm/example](https://github.com/awslabs/aurora-dsql-orms/tree/main/python/tortoise-orm/example)RubyRails[https://github.com/aws-samples/aurora-dsql-samples/tree/main/ruby/rails](https://github.com/aws-samples/aurora-dsql-samples/tree/main/ruby/rails)TypeScriptPrisma[https://github.com/aws-samples/aurora-dsql-samples/tree/main/typescript/prisma](https://github.com/aws-samples/aurora-dsql-samples/tree/main/typescript/prisma)TypeScriptSequelize[https://github.com/aws-samples/aurora-dsql-samples/tree/main/typescript/sequelize](https://github.com/aws-samples/aurora-dsql-samples/tree/main/typescript/sequelize)TypeScriptTypeORM[https://github.com/aws-samples/aurora-dsql-samples/tree/main/typescript/type-orm](https://github.com/aws-samples/aurora-dsql-samples/tree/main/typescript/type-orm)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VSCode

Loading data

All content copied from https://docs.aws.amazon.com/.
