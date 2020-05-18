#!/usr/bin/env bash
txt="10s";
for i in {1..10} ; do
    # shellcheck disable=SC2028
     echo $txt $i
    sleep 1
done
