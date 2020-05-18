#!/usr/bin/env bash
txt="50s";
for i in {1..50} ; do
     # shellcheck disable=SC2028
     echo $txt $i
    sleep 1
done
