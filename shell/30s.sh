#!/usr/bin/env bash
txt="30s";
for i in {1..30} ; do
     # shellcheck disable=SC2028
     echo $txt $i
    sleep 1
done
