#!/bin/bash
if [ x"${WAIT_FOR_MYSQL}" != "x" ] && [ "${WAIT_FOR_MYSQL}" == "true" ]; then
   echo "Wait for MySQL container"
   sleep 60
else
   echo "Value is not assigned to a variable ...${WAIT_FOR_MYSQL}"
fi

sh ./faq-chat-bot-app
while :
do
	sleep 1
done
