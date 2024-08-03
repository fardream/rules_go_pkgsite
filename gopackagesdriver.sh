#!/usr/bin/env bash

exec bazelisk run -- @rules_go//go/tools/gopackagesdriver "${@}"
