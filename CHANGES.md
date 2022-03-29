# Release notes for prometheus-c5-exporter

## v1.1.4 (2022-03-29)
Fixes:

- Fixed outdated metrics being reported

## v1.1.3 (2021-12-01)
Fixes:

- Fixed deadlock

## v1.1.2 (2021-11-16)

Fixes:

- Prevent multiple _info_ metrics (#3)

Features:

- Add datacenter & componentGroup labels

## v1.1.1 (2021-05-27)

Fixes:

- Add workaround for invalid cstagwd CASS_ERR subevents (#1)
- Rename XMS config and metrics to use `xms` prefix (#2)
- Fix XMS http error on response due to keepalive
- Move XMS user/pwd to configuration
- Improve startup logging of config
- Update dependency to metrics-1.17.2

Breaking changes:

- Configuration for XMS has been renamed to `Xms` prefix, see `prometheus-c5-exporter.conf.example`
- XMS metrics renamed from `resourcelicenses_` to `xms_license_` and `resourcecounters_` to `xms_counter_`, adjust your grafana dashboards
