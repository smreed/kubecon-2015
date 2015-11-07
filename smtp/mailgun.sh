#!/bin/sh
chmod 600 /etc/postfix/sasl_passwd.db

if [ ! -f /var/spool/postfix/etc/resolv.conf ]; then
  cp /etc/resolv.conf /var/spool/postfix/etc/resolv.conf
fi

service rsyslog start
postfix start
tail -f /var/log/syslog

