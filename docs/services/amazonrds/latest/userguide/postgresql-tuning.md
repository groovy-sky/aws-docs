# Tuning with wait events for RDS for PostgreSQL

Wait events are an important tuning tool for RDS for PostgreSQL. When you can find out why
sessions are waiting for resources and what they are doing, you're better able to
reduce bottlenecks. You can use the information in this section to find possible causes and
corrective actions. This section also discusses basic PostgreSQL tuning concepts.

The wait events in this section are specific to RDS for PostgreSQL.

###### Topics

- [Essential concepts for RDS for PostgreSQL tuning](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Tuning.concepts.html)

- [RDS for PostgreSQL wait events](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Tuning.concepts.summary.html)

- [Client:ClientRead](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.clientread.html)

- [Client:ClientWrite](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.clientwrite.html)

- [CPU](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.cpu.html)

- [IO:BufFileRead and IO:BufFileWrite](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.iobuffile.html)

- [IO:DataFileRead](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.iodatafileread.html)

- [IO:WALWrite](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.iowalwrite.html)

- [IPC:parallel wait events](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rpg-ipc-parallel.html)

- [IPC:ProcArrayGroupUpdate](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/apg-rpg-ipcprocarraygroup.html)

- [Lock:advisory](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lockadvisory.html)

- [Lock:extend](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lockextend.html)

- [Lock:Relation](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lockrelation.html)

- [Lock:transactionid](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.locktransactionid.html)

- [Lock:tuple](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.locktuple.html)

- [LWLock:BufferMapping (LWLock:buffer\_mapping)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lwl-buffer-mapping.html)

- [LWLock:BufferIO (IPC:BufferIO)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lwlockbufferio.html)

- [LWLock:buffer\_content (BufferContent)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lwlockbuffercontent.html)

- [LWLock:lock\_manager (LWLock:lockmanager)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lw-lock-manager.html)

- [LWLock:pg\_stat\_statements](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/apg-rpg-lwlockpgstat.html)

- [LWLock:SubtransSLRU (LWLock:SubtransControlLock)](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.lwlocksubtransslru.html)

- [Timeout:PgSleep](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.timeoutpgsleep.html)

- [Timeout:VacuumDelay](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/wait-event.timeoutvacuumdelay.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Working with
parameters

Essential concepts for RDS for PostgreSQL tuning
