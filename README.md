## To Those Who Wish To Download All OCW's Resources

I am a big fan of The OCW, I wish to download all OCW's resources, this repository serves that purpose.


## How to use

find a school you feel like exploring,  I prefer EE school so I would do the following steps:
1. crawling all the links from EE school
```
go run . crawl -p "https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/" -f ee.info -s "electrical-engineering-and-computer-science"
```

2. select one course I would like to learn and merge its course materials(PDF version lecture notes)

```
go run . pdf ee.info 6_172

```

## License Info
---
This software contains a package ([unipdf](https://github.com/unidoc/unipdf)) which is a commercial product and requires a license code to operate, 
Therefore the use of this software package is governed by the end-user license agreement (EULA) available at: https://unidoc.io/eula/
To obtain a Trial license code to evaluate the software, please visit https://unidoc.io/




[universe]
iops = 0
mem_limit_mb = 0
cpu_quota_percentage = 0
quota_limit_mb = 0
scsi_pr_level = 0
umask_dir = 0750
umask = 0640
bakupdir = /opt/mysql/backup/3306
mycnf = /opt/mysql/etc/3306/my.cnf
run_user = actiontech-mysql
id = mysql-2v034m
group_id = mysql-test

[mysql]
no-auto-rehash
prompt = '\\u@\\h:\\p \\R:\\m:\\s [\\d]> '
#default-character-set = utf8mb4
#tee = /data/mysql_tmp/mysql_operation.log

[mysqld]
# DO NOT MODIFY, Universe will generate this part
port = 3306
server_id = 664674382
basedir = /opt/mysql/base/5.7.25
datadir = /opt/mysql/data/3306
log_bin = /opt/mysql/log/binlog/3306/mysql-bin
tmpdir = /opt/mysql/tmp/3306
relay_log = /opt/mysql/log/relaylog/3306/mysql-relay
innodb_log_group_home_dir = /opt/mysql/log/redolog/3306
log_error = /opt/mysql/data/3306/mysql-error.log
report_host = 172.20.134.5

# BINLOG
binlog_error_action = ABORT_SERVER
binlog_format = row
binlog_rows_query_log_events = 1
log_slave_updates = 1
master_info_repository = TABLE
max_binlog_size = 250M
relay_log_info_repository = TABLE
relay_log_recovery = 1
sync_binlog = 1

# GTID #
gtid_mode = ON
enforce_gtid_consistency = 1
binlog_gtid_simple_recovery = 1

# ENGINE
default_storage_engine = InnoDB
innodb_buffer_pool_size = 1G
innodb_data_file_path = ibdata1:1G:autoextend
innodb_file_per_table = 1
innodb_flush_log_at_trx_commit=1
innodb_flush_method = O_DIRECT
innodb_io_capacity = 1000
innodb_log_buffer_size = 64M
innodb_log_file_size = 2G
innodb_log_files_in_group = 2
innodb_max_dirty_pages_pct = 60
innodb_print_all_deadlocks=1
#innodb_stats_on_metadata = 0
innodb_strict_mode = 1
#innodb_undo_logs = 128 		#Deprecated In 5.7.19
innodb_undo_tablespaces=3 	    #Deprecated In 5.7.21
innodb_max_undo_log_size=4G
innodb_undo_log_truncate=1
innodb_read_io_threads = 8
innodb_write_io_threads = 8
innodb_purge_threads = 4
innodb_buffer_pool_load_at_startup = 1
innodb_buffer_pool_dump_at_shutdown = 1
innodb_buffer_pool_dump_pct=25
innodb_sort_buffer_size = 8M
#innodb_page_cleaners = 8
innodb_buffer_pool_instances = 8
innodb_lock_wait_timeout = 10
innodb_io_capacity_max = 2000
innodb_flush_neighbors = 1
#innodb_large_prefix = 1
innodb_thread_concurrency = 4
innodb_stats_persistent_sample_pages = 64
innodb_autoinc_lock_mode = 2
innodb_online_alter_log_max_size = 1G
innodb_open_files = 4096
innodb_temp_data_file_path = ibtmp1:12M:autoextend:max:50G
innodb_rollback_segments = 128
innodb_numa_interleave = 1
innodb_monitor_enable = log_lsn_checkpoint_age,log_max_modified_age_async,log_max_modified_age_sync

# CACHE
key_buffer_size = 16M
tmp_table_size = 64M
max_heap_table_size = 64M
table_open_cache = 2000
query_cache_type = 0
query_cache_size = 0
max_connections = 2000
thread_cache_size = 200
open_files_limit = 65535
binlog_cache_size = 1M
join_buffer_size = 8M
sort_buffer_size = 2M
read_buffer_size = 8M
read_rnd_buffer_size = 8M
table_definition_cache = 2000
table_open_cache_instances = 8


# SLOW LOG
slow_query_log = 1
slow_query_log_file = /opt/mysql/data/3306/mysql-slow.log
log_slow_admin_statements = 1
log_slow_slave_statements = 1
long_query_time = 1

# SEMISYNC #
plugin_load = "rpl_semi_sync_master=semisync_master.so;rpl_semi_sync_slave=semisync_slave.so"
rpl_semi_sync_master_enabled = 1
rpl_semi_sync_slave_enabled = 0
rpl_semi_sync_master_wait_for_slave_count = 1
rpl_semi_sync_master_wait_no_slave = 0
rpl_semi_sync_master_timeout = 30000 # 30s

# CLIENT_DEPRECATE_EOF
session_track_schema = 1
session_track_state_change = 1
session_track_system_variables = '*'

# MISC
log_timestamps=SYSTEM
lower_case_table_names = 1
max_allowed_packet = 64M
read_only = 0
skip_external_locking = 1
skip_name_resolve = 1
skip_slave_start = 1
socket = /opt/mysql/data/3306/mysqld.sock
pid_file = /opt/mysql/data/3306/mysqld.pid
disabled_storage_engines = ARCHIVE,BLACKHOLE,EXAMPLE,FEDERATED,MEMORY,MERGE,NDB
log-output = TABLE,FILE
character_set_server = utf8mb4
secure_file_priv = ""
performance-schema-instrument ='wait/lock/metadata/sql/mdl=ON'
performance-schema-instrument = 'memory/% = COUNTED'
expire_logs_days = 7
max_connect_errors = 1000000
interactive_timeout = 1800
wait_timeout = 1800
log_bin_trust_function_creators = 1
performance-schema-instrument = 'stage/innodb/alter%=ON'
performance-schema-consumer-events-stages-current=ON
performance-schema-consumer-events-stages-history=ON
performance-schema-consumer-events-stages-history-long=ON

# MTS
slave-parallel-type=LOGICAL_CLOCK
slave_parallel_workers=16
slave_preserve_commit_order = ON
slave_rows_search_algorithms = 'INDEX_SCAN,HASH_SCAN'

##BaseConfig
collation_server = utf8mb4_bin
explicit_defaults_for_timestamp = 1
transaction_isolation = READ-COMMITTED

##Unused
#plugin-load-add = validate_password.so
#validate_password_policy = MEDIUM