# Release notes for prometheus-c5-exporter

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
