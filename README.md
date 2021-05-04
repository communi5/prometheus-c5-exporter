# Prometheus C5 Exporter

A prometheus exporter for C5 application processes

## About

The prometheus-c5-exporter daemon listens on a dedicated port for incoming
queries of a [Prometheus](https://prometheus.io) server and queries local C5 processes for metric data.

The metrics are collected using an internal REST based format of the supported 
C5 daemons. The reponses are parsed, adjusted in naming, labels added where required 
and written in Prometheus [text format](https://prometheus.io/docs/instrumenting/exposition_formats/#text-based-format).

Currently supported C5 processes are:

- C5 Proxy `sipproxyd`
- C5 ACD Queue `acdqueued`
- C5 Registrar `registrard`
- C5 CSTAGW `cstagwd`
- C5 Notification Server `notification-server`

additional 3rd party exporter included for Dialogic XMS metrics/licenses monitoring

These metrics are usually displayed with [Grafana](https://grafana.com). Dashboards are included 
for basic visualization.

## Installation and Configuration

To install promtheus-c5-exporter it is recommended to use the provided
packages.

After package installation a systemd unit is available for controlling the c5-exporter. A user `prometheus` will be created during installation if not yet existing.

### Configuration

Only minimal configuration is required. Querying for a particular endpoint 
can be disabled to avoid spurios warnings in the logs.

Example configuration:

```
listenAddress = ":9055"

### Query sipproxyd process
sipproxydEnabled = true
sipproxydTrunksEnabled = true

### Query acdqueued process
acdqueuedEnabled = true

### Query registard process
registrardEnabled = false

### Query cstagwd process
cstaEnabled = false

### Query notification-server process
notificationEnabled = false

### 3rd party XMS
resourceCountersEnabled = false
resourceLicensesEnabled = false

```

## Building and Packaging

To build prometheus-c5-exporter only a recent Go version (v1.15+) is required.

For a quick build use: 

    go build main.go

To build a static binary use: 

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags '-s -w' -o c5exporter main.go

### Using Visual Studio Code

Tasks for Visual Studio Code are available to simplify testing and building of the application. The "Go for Visual Studio Code" plugin is also recommended.

To run the tasks in Visual Studio use:
- Use the command palette (`Ctrl-Shift-P`) and choose "Tasks: Run Task"
- Choose a task like `Build Static` or `Build tagged release package`

A quick binary build can be run using `Ctrl-Shift-B`.

### Packaging

To create native packages for Debian/Ubuntu or Red Hat Linux a tool called [goreleaser](https://github.com/goreleaser/goreleaser) is required.

To package an offical version be sure to follow these steps:

- Create a version tag (e.g. `v1.0.3`)
- Ensure to have a clean workspace (no uncommitted files)
- Run `goreleaser release --skip-publish --rm-dist`

## Parsing of C5 Metrics

An example output for the internal C5 REST endpoint as used by sessionconsole is seen here:

```json
{
	"proxyResponseTimeStampAndState:" : "2020-01-19 11:40:01  active",
	"proxyState" : "active",
	"buildVersion:" : "Version: 6.0.2.57, compiled on Jan 15 2020, 13:06:31 built by TELES Communication Systems GmbH",
	"startupTime:" : "2020-01-19 04:01:04.503",
	"memoryUsage" : "C5 Heap Health: OK  - Mem used: 2%  - Mem used: 57MB  - Mem total: 2048MB  - Max: 3% - UpdCtr: 13198",
	"tuQueueStatus" : "OK - checked: 1830",
	"counterInfos" : [
	  "       Event counters                              absolute   curr   last",
	  "  0 TRANSPORT_MESSAGE_IN                              6502      0     72",
	  "  1 TRANSPORT_MESSAGE_OUT                             7088      0     79",
	  "253 TRANSPORT_TCP_MESSAGE_IN                             0      0      0",
	  "254 TRANSPORT_TCP_MESSAGE_OUT                            0      0      0",
	  "  2 REQUEST_METHOD_INVITE_IN                             0      0      0",
	  "  6 REQUEST_METHOD_SUBSCRIBE_IN                        334      0      5",
	  " 30 REQUEST_METHOD_NOOP_IN                            4964      0     54",
	  "  9 REQUEST_METHOD_NOTIFY_OUT                           39      0      0",
	  " 52 CALL_CONTROL_ORIG_CALL_SETUP_SUCCESS                 0      0      0",
	  " 54 CALL_CONTROL_ORIG_CALL_FAST_CONNECTED                0      0      0",
	  " 53 CALL_CONTROL_ORIG_CALL_CONNECTED                     0      0      0",
	  " 47 CALL_CONTROL_ORIG_CLIENT_ERROR                       0      0      0",
	  " 48 CALL_CONTROL_ORIG_SERVER_ERROR                       0      0      0",
	  " 49 CALL_CONTROL_ORIG_GLOBAL_ERROR                       0      0      0",
	  " 50 CALL_CONTROL_ORIG_REDIRECTION                        0      0      0",
	  " 51 CALL_CONTROL_ORIG_AUTHENTICATION_REQUIRED            0      0      0",
	  "190 OVERLOAD_PROTECTION_LIMIT_REACHED                    0      0      0",
	  "214 OVERLOAD_HEAP_WARNING_REJECTED_IN_REQUESTS           0      0      0",
	  "215 OVERLOAD_HEAP_CRITICAL_REJECTED_IN_REQUESTS          0      0      0",
	  "191 OVERLOAD_LIMIT1_REJECTED_IN_REQUESTS                 0      0      0",
	  "192 OVERLOAD_LIMIT2_REJECTED_IN_REQUESTS                 0      0      0",
	  "193 OVERLOAD_LIMIT3_REJECTED_IN_REQUESTS                 0      0      0",
	  "194 OVERLOAD_LIMIT4_REJECTED_IN_REQUESTS                 0      0      0",
	  "367 CALLS_LIMIT_REACHED                                  0      0      0",
	  "368 BT_CALLS_LIMIT_REACHED                               0      0      0",
	  "369 USER_CALLS_LIMIT_REACHED                             0      0      0",
	  " 46 CALL_CONTROL_AUTHENTICATION_ERROR                    0      0      0",
	  "227 CALL_CONTROL_IN_ACL_DENY                             0      0      0",
	  "228 CALL_CONTROL_OUT_ACL_DENY                            0      0      0",
	  "329 IP_FILTER_DENIED                                     0      0      0",
	  "330 IP_FILTER_NOT_ALLOWED                                0      0      0",
	  " 76 PRESENCE_AUTHENTICATION_ERROR                        0      0      0",
	  " 77 TRANSACTION_AND_TU_RETRY_IN                         50      0      0",
	  " 78 TRANSACTION_AND_TU_RETRY_OUT                        46      0      0",
	  " 83 TRANSACTION_AND_TU_CONN_VERIFICATION_RELEASED        0      0      0",
	  " 93 LOCATION_DNS_RESOLVER_ERROR                          0      0      0",
	  " 95 LOCATION_DNS_QUERY_TIMEOUT                           0      0      0",
	  "129 DATABASE_ERRORS                                      6      0      0",
	  "366 DATABASE_NOSQL_ERRORS                                0      0      0",
	  "144 ROUTING_ERRORS                                       0      0      0",
	  "177 SNMP_REQUESTS                                      908      0     10",
	  "178 SNMP_TRAPS                                           5      0      0",
	  "267 GENERAL_RCC_IN_COMMANDS                              3      0      0",
	  "268 GENERAL_RCC_OUT_COMMANDS                             3      0      0",
	  "350 WS_AGENT_EV_IN                                       0      0      0",
	  "351 WS_AGENT_EV_OUT                                      0      0      0",
	  "352 WS_CALL_EV                                           0      0      0",
	  "360 WS_CALL_SYNC_IN                                      0      0      0",
	  "359 WS_CALL_SYNC_OUT                                     0      0      0",
	  "362 WS_CALL_NOTIFY_IN                                    0      0      0",
	  "361 WS_CALL_NOTIFY_OUT                                   0      0      0",
	  "379 PUSH_CALL_NOTIFY                                     0      0      0",
	  "380 PUSH_CALL_NOTIFY_ERROR                               0      0      0",
	  "       Usage counters                              current    min    max   lMin   lMax   lAvg",
	  " 45 CALL_CONTROL_ACTIVE_CALLS                           0      0      0      0      0      0",
	  "309 BT_ACTIVE_CALLS                                     0      0      0      0      0      0",
	  " 75 PRESENCE_ACTIVE_SUBSCRIPTIONS                       6      6      6      6      6      6",
	  " 82 TRANSACTION_AND_TU_ACTIVE_SESSIONS                  0      0      0      0      0      0",
	  "189 TRANSACTION_AND_TU_ACTIVE_UA_SESSIONS               0      0      0      0      0      0",
	  " 81 TRANSACTION_AND_TU_ACTIVE_TRANSACTION_USERS         0      0      0      0      2      0",
	  "322 TRANSACTION_AND_TU_ACTIVE_INVITE_SERVER             0      0      0      0      0      0",
	  "233 TRANSPORT_TCP_ACTIVE_IN_CONNECTION                  0      0      0      0      0      0",
	  "234 TRANSPORT_TCP_ACTIVE_TRUSTED_IN_CONNECTION          0      0      0      0      0      0",
	  "235 TRANSPORT_TCP_ACTIVE_OUT_CONNECTION                 0      0      0      0      0      0",
	  "236 TRANSPORT_TCP_ACTIVE_TRUSTED_OUT_CONNECTION         0      0      0      0      0      0",
	  "264 GENERAL_RCC_ACTIVE_CONNECTIONS                      0      0      0      0      0      0",
	  "349 WS_CONNECTIONS                                      6      6      6      6      6      6",
	  [
		" 84 TRANSACTION_AND_TU_TU_MANAGER_QUEUE_SIZE          0      0      0      0      0      0",
		"                                                      0      0      0      0      0      0",
		"                                                      0      0      0      0      1      0",
		"                                                      0      0      0      0      0      0",
		"                                                      0      0      0      0      1      0"
	  ]
	]
}
```

This will be parsed and automatic naming will be applied. For a running

Example response for a prometheus query to `http://<host>:9055/metrics`:

```
registrard_audit_ua_session_released_total 0
registrard_cass_err_conn_tmo_total{idx="0"} 0
registrard_cass_err_pending_requ_tmo_total{idx="0"} 0
registrard_cass_err_requ_tmo_total{idx="0"} 0
registrard_cluster_active_registrations_current 0
registrard_cluster_active_registrations_lastavg 0
registrard_cluster_active_registrations_lastmax 0
registrard_cluster_active_registrations_lastmin 0
registrard_connected_session_timeout_total 0
registrard_database_errors_total 0
registrard_database_nosql_errors_total 0
```
