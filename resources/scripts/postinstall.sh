#!/bin/sh

id prometheus >/dev/null 2>&1
if [ $? -gt 0 ]; then
  echo "Creating local user 'prometheus'"
  useradd prometheus -c 'Prometheus exporter' --system -M -N
fi
if [ -f /etc/prometheus-c5-exporter.conf ]; then
  echo "Adjust permissions for config files"
  chmod 640 /etc/prometheus-c5-exporter.conf
fi
if [ -f /etc/init.d/prometheus-c5-exporter ]; then
  echo "Adjust permission for systemV init script"
  chmod 755 /etc/init.d/prometheus-c5-exporter
fi
