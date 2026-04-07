# RDS for PostgreSQL wait events

The following table lists the wait events for RDS for PostgreSQL that most commonly
indicate performance problems, and summarizes the most common causes and corrective
actions..

Wait eventDefinition

[Client:ClientRead](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.clientread.html)

This event occurs when RDS for PostgreSQL is waiting to receive data
from the client.

[Client:ClientWrite](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.clientwrite.html)

This event occurs when RDS for PostgreSQL is waiting to write data to
the client.

[CPU](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.cpu.html)

This event occurs when a thread is active in CPU or is waiting for
CPU.

[IO:BufFileRead and IO:BufFileWrite](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.iobuffile.html)

These events occur when RDS for PostgreSQL creates temporary
files.

[IO:DataFileRead](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.iodatafileread.html)

This event occurs when a connection waits on a backend process to
read a required page from storage because the page isn't available
in shared memory.

[IO:WALWrite](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.iowalwrite.html)This event occurs when RDS for PostgreSQL is waiting for the write-ahead
log (WAL) buffers to be written to a WAL file.

[IPC:parallel wait events](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rpg-ipc-parallel.html)

These wait events indicate that a session is waiting for
inter-process communication related to parallel query execution
operations.

[IPC:ProcArrayGroupUpdate](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/apg-rpg-ipcprocarraygroup.html)

This event occurs when a session is waiting for the group leader
to update the transaction status at the end of the
transaction.

[Lock:advisory](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lockadvisory.html)

This event occurs when a PostgreSQL application uses a lock to
coordinate activity across multiple sessions.

[Lock:extend](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lockextend.html)

This event occurs when a backend process is waiting to lock a
relation to extend it while another process has a lock on that
relation for the same purpose.

[Lock:Relation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lockrelation.html)

This event occurs when a query is waiting to acquire a lock on a
table or view that's currently locked by another transaction.

[Lock:transactionid](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.locktransactionid.html)

This event occurs when a transaction is waiting for a row-level
lock.

[Lock:tuple](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.locktuple.html)

This event occurs when a backend process is waiting to acquire a
lock on a tuple.

[LWLock:BufferMapping (LWLock:buffer\_mapping)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lwl-buffer-mapping.html)

This event occurs when a session is waiting to associate a data
block with a buffer in the shared buffer pool.

[LWLock:BufferIO (IPC:BufferIO)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lwlockbufferio.html)

This event occurs when RDS for PostgreSQL is waiting for other
processes to finish their input/output (I/O) operations when
concurrently trying to access a page.

[LWLock:buffer\_content (BufferContent)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lwlockbuffercontent.html)

This event occurs when a session is waiting to read or write a
data page in memory while another session has that page locked for
writing.

[LWLock:lock\_manager (LWLock:lockmanager)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lw-lock-manager.html)

This event occurs when the RDS for PostgreSQL engine maintains the
shared lock's memory area to allocate, check, and deallocate a lock
when a fast path lock isn't possible.

[LWLock:SubtransSLRU (LWLock:SubtransControlLock)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lwlocksubtransslru.html)

This event occurs when a process is waiting to access the simple
least-recently used (SLRU) cache for a subtransaction.

[LWLock:pg\_stat\_statements](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/apg-rpg-lwlockpgstat.html)

This event occurs when the `pg_stat_statements`
extension takes an exclusive lock on the hash table that tracks SQL
statements.

[Timeout:PgSleep](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.timeoutpgsleep.html)

This event occurs when a server process has called the
`pg_sleep` function and is waiting for the sleep
timeout to expire.

[Timeout:VacuumDelay](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.timeoutvacuumdelay.html)

This event indicates that the vacuum process is sleeping because the
estimated cost limit has been reached.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RDS for PostgreSQL processes

Client:ClientRead
