#!/bin/sh
python -u -m smtpd -c DebuggingServer -n 0.0.0.0:25 2>&1
