#!/bin/bash
read -p "Please input your commit: "
git add .
git commit -m "$REPLY"
git push