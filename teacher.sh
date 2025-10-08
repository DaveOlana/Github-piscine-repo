#!/bin/bash

interviewn= grep -H "licen" interviews/* |grep "\"" | cut -fl -d ":" | rev | cut -fl -d "-" | rev
interview="cat interviews/interview-$interviewn"
export interviewnum=$interviewn
echo $interviewnum
$interview
echo $MAIN_SUSPECT