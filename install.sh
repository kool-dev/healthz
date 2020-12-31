#!/usr/bin/env bash

set -euo pipefail

echo -e "Hello, we are gonna install the \033[33mlatest stable\033[39m version of healthz tool"

DEFAULT_DOWNLOAD_URL="https://github.com/kool-dev/healthz/releases/latest/download"
if [ -z "${DOWNLOAD_URL:-}" ]; then
	DOWNLOAD_URL=$DEFAULT_DOWNLOAD_URL
fi

DEFAULT_BIN="/usr/local/bin/healthz"
if [ -z "${BIN_PATH:-}" ]; then
	BIN_PATH=$DEFAULT_BIN
fi

command_exists() {
	command -v "$@" > /dev/null 2>&1
}

is_darwin() {
	case "$(uname -s)" in
	*darwin* ) true ;;
	*Darwin* ) true ;;
	* ) false;;
	esac
}

do_install () {
	ARCH=$(uname -m)
	PLAT="linux"

	if is_darwin; then
		echo "MacOS is not supported. If you need to, you can build from source."
		exit 1
	fi

	if [ "$ARCH" == "x86_64" ]; then
		ARCH="amd64"
	fi

	echo "Downloading latest binary (healthz-$PLAT-$ARCH)..."

	# TODO: fallback to wget if no curl available
	rm -f /tmp/healthz_binary
	curl -fsSL "$DOWNLOAD_URL/healthz-$PLAT-$ARCH" -o /tmp/healthz_binary

	# check for running kool process which would prevent
	# replacing existing version under Linux.
	if command_exists healthz && ! is_darwin; then
		running=$(ps aux | grep healthz | grep -v grep | wc -l | awk '{ print $1 }')
		if [ "$running" != "0" ]; then
			echo -e "\033[31;31mThere is a healthz process still running. You might need to stop them before we replace the current binary.\033[0m"
		fi
	fi

	echo -e "Moving healthz binary to $BIN_PATH..."
	if [ -w $(dirname $BIN_PATH) ]; then
		mv -f /tmp/healthz_binary $BIN_PATH
		chmod +x $BIN_PATH
	else
		echo "(requires sudo)"
		sudo mv -f /tmp/healthz_binary $BIN_PATH
		sudo chmod +x $BIN_PATH
	fi

	start_success="\033[0;32m"
	end_success="\033[0m"
	builtin echo -e "${start_success}$(healthz -v) installed successfully.${end_success}"
}

do_install
