#!/bin/bash
set -e

FEISHU_CUSTOMERBOT_WEBHOOK=${FEISHU_CUSTOMERBOT_WEBHOOK:-}
FEISHU_CUSTOMERBOT_SECRET=${FEISHU_CUSTOMERBOT_SECRET:-}
MSGTYPE=${MSGTYPE:-"interactive"}
TITLE=${TITLE:-}
CONTENT=${CONTENT:-}

echo "## Sending message ##################"

feishu -webhook "${FEISHU_CUSTOMERBOT_WEBHOOK}" -secret "${FEISHU_CUSTOMERBOT_SECRET}" -msgtype "${MSGTYPE}" -title "${TITLE}" -content "${CONTENT}"

echo "## Done. ##################"
