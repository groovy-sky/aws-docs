# NATIVE\_NETWORK\_ENCRYPTION option settings

You can specify encryption requirements on both the server and the client. The DB instance can act as a client
when, for example, it uses a database link to connect to another database. You might want to avoid forcing
encryption on the server side. For example, you might not want to force all client communications to use
encryption because the server requires it. In this case, you can force encryption on the client side using the
`SQLNET.*CLIENT` options.

Amazon RDS supports the following settings for the `NATIVE_NETWORK_ENCRYPTION`
option.

###### Note

When you use commas to separate values for an option setting, don't put a space
after the comma.

Option settingValid valuesDefault valuesDescription

`SQLNET.ALLOW_WEAK_CRYPTO_CLIENTS`

`TRUE`, `FALSE`

`TRUE`

The behavior of the server when a client using a non-secure cipher attempts to connect to
the database. If `TRUE`, clients can connect even if they aren't patched with the
July 2021 PSU.

If the setting is `FALSE`, clients can connect to the database only when they
are patched with the July 2021 PSU. Before you set
`SQLNET.ALLOW_WEAK_CRYPTO_CLIENTS` to `FALSE`, make sure that the
following conditions are met:

- `SQLNET.ENCRYPTION_TYPES_SERVER` and
`SQLNET.ENCRYPTION_TYPES_CLIENT` have one matching encryption method
that is not `DES`, `3DES`, or `RC4` (all key
lengths).

- `SQLNET.CHECKSUM_TYPES_SERVER` and
`SQLNET.CHECKSUM_TYPES_CLIENT` have one matching secure checksumming
method that is not `MD5`.

- The client is patched with the July 2021 PSU. If the client isn't patched, the
client loses the connection and receives the `ORA-12269` error.

`SQLNET.ALLOW_WEAK_CRYPTO`

`TRUE`, `FALSE`

`TRUE`

The behavior of the server when a client using a non-secure cipher attempts to connect to
the database. The following ciphers are considered not secure:

- `DES` encryption method (all key lengths)

- `3DES` encryption method (all key lengths)

- `RC4` encryption method (all key lengths)

- `MD5` checksumming method

If the setting is `TRUE`, clients can connect when they use the preceding
non-secure ciphers.

If the setting is `FALSE`, the database prevents clients from connecting when
they use the preceding non-secure ciphers. Before you set
`SQLNET.ALLOW_WEAK_CRYPTO` to `FALSE`, make sure that the following
conditions are met:

- `SQLNET.ENCRYPTION_TYPES_SERVER` and
`SQLNET.ENCRYPTION_TYPES_CLIENT` have one matching encryption method
that is not `DES`, `3DES`, or `RC4` (all key
lengths).

- `SQLNET.CHECKSUM_TYPES_SERVER` and
`SQLNET.CHECKSUM_TYPES_CLIENT` have one matching secure checksumming
method that is not `MD5`.

- The client is patched with the July 2021 PSU. If the client isn't patched, the
client loses the connection and receives the `ORA-12269` error.

`SQLNET.CRYPTO_CHECKSUM_CLIENT`

`Accepted`, `Rejected`, `Requested`, `Required`

`Requested`

The data integrity behavior when a DB instance connects to the client, or a server acting
as a client. When a DB instance uses a database link, it acts as a client.

`Requested` indicates that the client doesn't require the DB instance to perform
a checksum.

`SQLNET.CRYPTO_CHECKSUM_SERVER`

`Accepted`, `Rejected`, `Requested`, `Required`

`Requested`

The data integrity behavior when a client, or a server acting as a client, connects to the
DB instance. When a DB instance uses a database link, it acts as a client.

`Requested` indicates that the DB instance doesn't require the client to perform
a checksum.

`SQLNET.CRYPTO_CHECKSUM_TYPES_CLIENT`

`SHA256`, `SHA384`, `SHA512`, `SHA1`,
`MD5`

`SHA256`, `SHA384`, `SHA512`

A list of checksum algorithms.

You can specify either one value or a comma-separated list of values. If you use a comma,
don't insert a space after the comma; otherwise, you receive an
`InvalidParameterValue` error.

This parameter and `SQLNET.CRYPTO_CHECKSUM_TYPES_SERVER ` must have a common
cipher.

`SQLNET.CRYPTO_CHECKSUM_TYPES_SERVER`

`SHA256`, `SHA384`, `SHA512`, `SHA1`,
`MD5`

`SHA256`, `SHA384`, `SHA512`, `SHA1`,
`MD5`

A list of checksum algorithms.

You can specify either one value or a comma-separated list of values. If you use a comma,
don't insert a space after the comma; otherwise, you receive an
`InvalidParameterValue` error.

This parameter and `SQLNET.CRYPTO_CHECKSUM_TYPES_CLIENT` must have a common
cipher.

`SQLNET.ENCRYPTION_CLIENT`

`Accepted`, `Rejected`, `Requested`, `Required`

`Requested`

The encryption behavior of the client when a client, or a server acting as a client,
connects to the DB instance. When a DB instance uses a database link, it acts as a
client.

`Requested` indicates that the client does not require traffic from the server
to be encrypted.

`SQLNET.ENCRYPTION_SERVER`

`Accepted`, `Rejected`, `Requested`, `Required`

`Requested`

The encryption behavior of the server when a client, or a server acting as a client,
connects to the DB instance. When a DB instance uses a database link, it acts as a
client.

`Requested` indicates that the DB instance does not require traffic from the
client to be encrypted.

`SQLNET.ENCRYPTION_TYPES_CLIENT`

`RC4_256`, `AES256`, `AES192`, `3DES168`,
`RC4_128`, `AES128`, `3DES112`, `RC4_56`,
`DES`, `RC4_40`, `DES40`

`RC4_256`, `AES256`, `AES192`, `3DES168`,
`RC4_128`, `AES128`, `3DES112`, `RC4_56`,
`DES`, `RC4_40`, `DES40`

A list of encryption algorithms used by the client. The client
attempts to decrypt the server input by trying each algorithm in
order, proceeding until an algorithm succeeds or the end of the list
is reached.

Amazon RDS uses the following default list from Oracle. RDS starts with
`RC4_256` and proceeds down the list in order. You
can change the order or limit the algorithms that the DB instance
will accept.

01. `RC4_256`: RSA RC4 (256-bit key size)

02. `AES256`: AES (256-bit key size)

03. `AES192`: AES (192-bit key size)

04. `3DES168`: 3-key Triple-DES (112-bit effective key size)

05. `RC4_128`: RSA RC4 (128-bit key size)

06. `AES128`: AES (128-bit key size)

07. `3DES112`: 2-key Triple-DES (80-bit effective key size)

08. `RC4_56`: RSA RC4 (56-bit key size)

09. `DES`: Standard DES (56-bit key size)

10. `RC4_40`: RSA RC4 (40-bit key size)

11. `DES40`: DES40 (40-bit key size)

You can specify either one value or a comma-separated list of values. If you a comma, don't
insert a space after the comma; otherwise, you receive an `InvalidParameterValue`
error.

This parameter and `SQLNET.SQLNET.ENCRYPTION_TYPES_SERVER` must have a common
cipher.

`SQLNET.ENCRYPTION_TYPES_SERVER`

`RC4_256`, `AES256`, `AES192`, `3DES168`,
`RC4_128`, `AES128`, `3DES112`, `RC4_56`,
`DES`, `RC4_40`, `DES40`

`RC4_256`, `AES256`, `AES192`, `3DES168`,
`RC4_128`, `AES128`, `3DES112`, `RC4_56`,
`DES`, `RC4_40`, `DES40`

A list of encryption algorithms used by the DB instance. The DB instance uses each
algorithm, in order, to attempt to decrypt the client input until an algorithm succeeds or
until the end of the list is reached.

Amazon RDS uses the following default list from Oracle. You can change the order or limit the
algorithms that the client will accept.

01. `RC4_256`: RSA RC4 (256-bit key size)

02. `AES256`: AES (256-bit key size)

03. `AES192`: AES (192-bit key size)

04. `3DES168`: 3-key Triple-DES (112-bit effective key size)

05. `RC4_128`: RSA RC4 (128-bit key size)

06. `AES128`: AES (128-bit key size)

07. `3DES112`: 2-key Triple-DES (80-bit effective key size)

08. `RC4_56`: RSA RC4 (56-bit key size)

09. `DES`: Standard DES (56-bit key size)

10. `RC4_40`: RSA RC4 (40-bit key size)

11. `DES40`: DES40 (40-bit key size)

You can specify either one value or a comma-separated list of values. If you a comma, don't
insert a space after the comma; otherwise, you receive an `InvalidParameterValue`
error.

This parameter and `SQLNET.SQLNET.ENCRYPTION_TYPES_SERVER` must have a common
cipher.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Native network encryption (NNE)

Adding the option

All content copied from https://docs.aws.amazon.com/.
